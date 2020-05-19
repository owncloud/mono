package proto_test

import (
	"bytes"
	"context"
	"image"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/owncloud/ocis-pkg/v2/service/grpc"
	"github.com/owncloud/ocis-thumbnails/pkg/config"
	"github.com/owncloud/ocis-thumbnails/pkg/proto/v0"
	"github.com/owncloud/ocis-thumbnails/pkg/thumbnail/imgsource"
	"github.com/owncloud/ocis-thumbnails/pkg/thumbnail/storage"
	"github.com/stretchr/testify/assert"

	svc "github.com/owncloud/ocis-thumbnails/pkg/service/v0"
)

var service = grpc.Service{}

func init() {
	service = grpc.NewService(
		grpc.Namespace("com.owncloud.api"),
		grpc.Name("thumbnails"),
		grpc.Address("localhost:9992"),
	)

	cfg := config.New()
	cfg.Thumbnail.Resolutions = []string{"16x16", "32x32", "64x64", "128x128"}

	wd, _ := os.Getwd()
	fsCfg := config.FileSystemSource{
		BasePath: filepath.Join(wd, "../../../testdata/"),
	}
	err := proto.RegisterThumbnailServiceHandler(
		service.Server(),
		svc.NewService(
			svc.Config(cfg),
			svc.ThumbnailStorage(storage.NewInMemoryStorage()),
			svc.ThumbnailSource(imgsource.NewFileSystemSource(fsCfg)),
		),
	)
	if err != nil {
		log.Fatalf("could not register ThumbnailHandler: %v", err)
	}
	service.Server().Start()
}

func TestGetThumbnailInvalidImage(t *testing.T) {
	req := proto.GetRequest{
		Filepath: "invalid.png",
		Filetype: proto.GetRequest_PNG,
		Etag:     "33a64df551425fcc55e4d42a148795d9f25f89d4",
		Height:   32,
		Width:    32,
	}
	client := service.Client()
	cl := proto.NewThumbnailService("com.owncloud.api.thumbnails", client)
	_, err := cl.GetThumbnail(context.Background(), &req)

	assert.NotNil(t, err)
}

func TestGetThumbnail(t *testing.T) {
	req := proto.GetRequest{
		Filepath: "oc.png",
		Filetype: proto.GetRequest_PNG,
		Etag:     "33a64df551425fcc55e4d42a148795d9f25f89d4",
		Height:   32,
		Width:    32,
	}
	client := service.Client()
	cl := proto.NewThumbnailService("com.owncloud.api.thumbnails", client)
	rsp, err := cl.GetThumbnail(context.Background(), &req)
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
	assert.NotEmpty(t, rsp.GetThumbnail())

	img, _, _ := image.Decode(bytes.NewReader(rsp.GetThumbnail()))
	assert.Equal(t, 32, img.Bounds().Size().X)

	assert.Equal(t, "image/png", rsp.GetMimetype())
}