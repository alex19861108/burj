package controllers

import (
	"context"
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

func (c *JobController) AnyNew() (job proto.Job, err error) {
	log.Printf("%#v\n", c.Ctx.FormValues())
	name := c.Ctx.FormValue("name")
	imageTags := c.Ctx.FormValue("image_tags")
	apkPath := c.Ctx.FormValue("apk_path")
	apkPkgName := c.Ctx.FormValue("apk_pkg_name")
	apkClassName := c.Ctx.FormValue("apk_class_name")
	//deviceLabel := c.Ctx.FormValue("device_label")

	conn, err := grpc.Dial(c.Cfg.BurjCenterConfig.RpcServer, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return job, ERROR_CONNECT_BURJ_CENTER
	}

	// get all valid node device unit
	client := proto.NewNodeServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchNodeRequest{})
	if err != nil {
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
	}

	// save job information
	job, err = c.JobService.InsertOrUpdate(job)

	// trigger subJobs
	for _, subJob := range job.SubJobs {
		conn, err := grpc.Dial(subJob.Node.Addr.Rpc, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			return job, err
		}

		client := proto.NewJobServiceClient(conn)
		response, err := client.Trigger(context.Background(), &proto.TriggerJobRequest{
			Job:    &job,
			SubJob: subJob,
		})
		if err != nil {
			return job, err
		}
		log.Printf("%#v\t%#v\n", subJob, response)
	}
	return job, nil
}

func (c *JobController) AnyDeleteBy(id string) (err error) {
	return c.JobService.DeleteByID(id)
}

func (c *JobController) AnyLog() (job proto.Job, err error) {
	jid := c.Ctx.FormValue("jid")
	subJid := c.Ctx.FormValue("sub_jid")
	result := c.Ctx.FormValue("result")

	job, err = c.JobService.GetByID(jid)
	if err != nil {
		return
	}

	for _, subJob := range job.SubJobs {
		if subJob.Id == subJid {
			subJob.Status = proto.Status_SUCCESS
			subJob.UpdateTime = time.Now().String()
			subJob.Log = result
		}
	}
	// save job information
	job, err = c.JobService.InsertOrUpdate(job)

	return job, nil
}
