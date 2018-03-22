package controllers

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/alex19861108/burj/burj_center/iris/config"

	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_server/iris/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"github.com/kataras/iris/mvc"
	"google.golang.org/grpc"
)

var (
	ERROR_CONNECT_BURJ_CENTER = errors.New("Cannot connect burj center")
	ERROR_NO_NODE_FOUND       = errors.New("Cannot found any node")
	ERROR_NO_TAG_FOUND        = errors.New("Cannot found any tag")
	ERROR_NO_IMAGE_FOUND      = errors.New("Cannot found any image")
)

type JobController struct {
	Cfg          config.Config
	Ctx          iris.Context
	ImageService services.ImageService
	TagService   services.TagService
	JobService   services.JobService
}

func (c *JobController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	//b.Handle("GET", "/something/{id:long}", "MyCustomHandler", nil)
}

func (c *JobController) Get() (results []proto.Job) {
	return c.JobService.GetAll()
}

func (c *JobController) GetBy(id string) (node proto.Job, err error) {
	return c.JobService.GetByID(id)
}

func (c *JobController) GetFilterBy(owner string) (results []proto.Job) {
	return c.JobService.GetByOwner(owner)
}

func (c *JobController) AnyCreate() (job proto.Job, err error) {
	log.Printf("create job, params: %#v\n", c.Ctx.FormValues())
	name := strings.TrimSpace(c.Ctx.FormValue("name"))
	imageTags := strings.TrimSpace(c.Ctx.FormValue("image_tags"))
	apkPath := strings.TrimSpace(c.Ctx.FormValue("apk_path"))
	apkPkgName := strings.TrimSpace(c.Ctx.FormValue("apk_pkg_name"))
	apkClassName := strings.TrimSpace(c.Ctx.FormValue("apk_class_name"))

	conn, err := grpc.Dial(c.Cfg.BurjCenterConfig.RpcServer, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return job, ERROR_CONNECT_BURJ_CENTER
	}

	log.Println("[job]: get image by tags")
	// get all valid node device unit
	client := proto.NewNodeServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchNodeRequest{})
	if err != nil || len(resp.NodeDeviceUnits) == 0 {
		return job, ERROR_NO_NODE_FOUND
	}

	// get all valid images by image tag
	names := strings.Split(imageTags, ",")
	tags, err := c.TagService.GetByNames(names)
	if err != nil {
		return job, ERROR_NO_TAG_FOUND
	}
	images, err := c.ImageService.GetByTags(tags)
	if err != nil || len(images) == 0 {
		return job, ERROR_NO_IMAGE_FOUND
	}

	log.Println("[job]: part images to devices")
	// part images for devices
	subJobs := c.JobService.PacketImagesToDevices(images, resp.NodeDeviceUnits)
	job = proto.Job{
		Name:    name,
		SubJobs: subJobs,
		Apk: &proto.Apk{
			Path:      apkPath,
			PkgName:   apkPkgName,
			ClassName: apkClassName,
		},
		Tags:       tags,
		Status:     proto.Status_PREPARE,
		UpdateTime: time.Now().String(),
	}

	log.Println("[job]: create job")
	// save job information
	job, err = c.JobService.InsertOrUpdate(job)

	log.Println("[job]: trigger job")
	// trigger subJobs
	go func(job proto.Job) (proto.Job, error) {
		for _, subJob := range job.SubJobs {
			conn, err := grpc.Dial(subJob.Node.Addr.Rpc, grpc.WithInsecure())
			defer conn.Close()
			if err != nil {
				return job, err
			}

			client := proto.NewJobServiceClient(conn)
			_, err = client.Trigger(context.Background(), &proto.TriggerJobRequest{
				Job:    &job,
				SubJob: subJob,
			})
			if err != nil {
				job.Status = proto.Status_ERROR
				job.UpdateTime = time.Now().String()
				for _, sj := range job.SubJobs {
					if sj.Id == subJob.Id {
						sj.Summary = err.Error()
						sj.Status = proto.Status_ERROR
						sj.UpdateTime = time.Now().String()
						break
					}
				}
				job, err = c.JobService.InsertOrUpdate(job)
				return job, err
			} else {
				if job.Status != proto.Status_ERROR && job.Status != proto.Status_SUCCESS {
					job.Status = proto.Status_RUNNING
				}
				job.UpdateTime = time.Now().String()
				for _, sj := range job.SubJobs {
					if sj.Id == subJob.Id {
						sj.Status = proto.Status_RUNNING
						sj.UpdateTime = time.Now().String()
						break
					}
				}
				job, err = c.JobService.InsertOrUpdate(job)
				return job, err
			}
		}
		return job, nil
	}(job)
	return job, nil
}

func (c *JobController) AnyDeleteBy(id string) (err error) {
	return c.JobService.DeleteByID(id)
}

func (c *JobController) GetImageTags() (tags []*proto.Tag, err error) {
	return c.TagService.GetImageTags()
}

func (c *JobController) AnyLog() (job proto.Job, err error) {
	log.Printf("receive new log: %#v\n", c.Ctx.FormValues())
	jid := c.Ctx.FormValue("jid")
	subJid := c.Ctx.FormValue("sjid")
	result := c.Ctx.FormValue("result")
	if jid == "" || subJid == "" {
		return
	}

	buf := bytes.NewBuffer(nil)
	fr, _, err := c.Ctx.FormFile("file")
	if err == nil {
		io.Copy(buf, fr)
	}
	defer fr.Close()

	job, err = c.JobService.GetByID(jid)
	if err != nil {
		return
	}

	// change sub job status
	for _, subJob := range job.SubJobs {
		if subJob.Id == subJid {
			subJob.Status = proto.Status_SUCCESS
			subJob.UpdateTime = time.Now().String()
			subJob.Summary = result
			subJob.Log = buf.String()
		}
	}

	// change job status
	if job.Status != proto.Status_SUCCESS && job.Status != proto.Status_ERROR {
		jobStatus := proto.Status_SUCCESS
		for _, subJob := range job.SubJobs {
			if subJob.Status != proto.Status_SUCCESS {
				jobStatus = proto.Status_RUNNING
			}
			if subJob.Status == proto.Status_ERROR {
				jobStatus = proto.Status_ERROR
			}
		}
		// save job information
		job.Status = jobStatus
		job, err = c.JobService.InsertOrUpdate(job)
	}

	return
}
