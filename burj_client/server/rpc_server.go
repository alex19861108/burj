package server

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/alex19861108/burj/burj_client/service"

	"github.com/alex19861108/burj/burj_center/iris/common/transfer"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"golang.org/x/net/context"
)

const (
	IMAGE_ROOT_PATH   = "/tmp/data/img/"
	APK_PKG_ROOT_PATH = "/tmp/data/apk/"
)

type RpcServer struct {
	ADB service.ADB
	Cfg config.Config
}

func (s *RpcServer) Trigger(ctx context.Context, req *proto.TriggerJobRequest) (*proto.TriggerJobResponse, error) {
	imagePath := path.Join(IMAGE_ROOT_PATH, req.Job.Id)
	pkgPath := path.Join(APK_PKG_ROOT_PATH, req.Job.Id)
	// package images
	localZipPath, err := s.packImages(req.SubJob.Images, imagePath)
	if err != nil {
		s.clean([]string{imagePath})
		return &proto.TriggerJobResponse{
			Job:    req.Job,
			SubJob: req.SubJob,
		}, err
	}

	// download apk
	localApkPkgPath, err := transfer.DownloadFromHttp(req.Job.Apk.Path, pkgPath)
	if err != nil {
		s.clean([]string{imagePath, pkgPath})
		return &proto.TriggerJobResponse{
			Job:    req.Job,
			SubJob: req.SubJob,
		}, err
	}

	// do job
	err = s.ADB.Run(req.SubJob.Device.Serial, localZipPath, localApkPkgPath, req.Job.Apk.PkgName, req.Job.Apk.ClassName, req.Job.Id)
	if err != nil {
		s.clean([]string{imagePath, pkgPath})
		return &proto.TriggerJobResponse{
			Job:    req.Job,
			SubJob: req.SubJob,
		}, err
	}

	// clean tmp file
	s.clean([]string{imagePath, pkgPath})

	return &proto.TriggerJobResponse{
		SubJob: req.SubJob,
	}, nil
}

func (s *RpcServer) packImages(images []*proto.Image, dir string) (string, error) {
	for _, img := range images {
		imgFtpPath := s.Cfg.FtpServerConfig.URL + "/" + img.Path
		transfer.DownloadFromSftp(imgFtpPath, dir)
	}

	localImageZipPath := path.Join(dir, path.Base(dir)+".zip")
	err := packDirectory(dir, localImageZipPath)

	return localImageZipPath, err
}

func (s *RpcServer) clean(dirs []string) {
	for _, item := range dirs {
		os.Remove(item)
	}
}

func packDirectory(dirPath string, zipPath string) error {
	fw, err := os.Create(zipPath)
	defer fw.Close()
	if err != nil {
		panic(err)
	}

	w := zip.NewWriter(fw)
	defer w.Close()

	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		zipfile, err := os.Open(path.Join(dirPath, f.Name()))
		defer zipfile.Close()
		if err != nil {
			panic(err)
		}

		info, err := zipfile.Stat()
		if err != nil {
			panic(err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			panic(err)
		}

		header.Method = zip.Deflate

		writer, err := w.CreateHeader(header)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(writer, zipfile)
		if err != nil {
			panic(err)
		}
	}

	return err
}

func InitRpcServer(cfg config.Config) {
	listener, err := net.Listen("tcp", ":"+cfg.ServerConfig.Port.Rpc)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterJobServiceServer(s, &RpcServer{ADB: service.ADB{}, Cfg: cfg})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
