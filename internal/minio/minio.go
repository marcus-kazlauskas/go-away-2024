package minio

import (
	"context"
	"go-away-2024/internal/config"
	"net"
	"os"

	"github.com/gofiber/fiber/v2/log"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const PuzzleBucketName string = "tasks"
const PuzzleContentType string = "multipart/form-data"

type MinioClient struct {
	c *minio.Client
}

func NewClient(cfg *config.Config) *MinioClient {
	endpoint := net.JoinHostPort(cfg.S3.Host, cfg.S3.Port)

	m, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3.AccessKey, cfg.S3.SecretKey, ""),
		Secure: cfg.S3.SslMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("S3 client created: endpoint=%s", endpoint)
	return &MinioClient{
		c: m,
	}
}

func (m *MinioClient) UploadPuzzleInput(object *os.File) error {
	objectName := object.Name()
	objectStat, _ := object.Stat()
	object.Seek(0, 0)

	info, err := m.c.PutObject(
		context.Background(),
		PuzzleBucketName, objectName, object, objectStat.Size(),
		minio.PutObjectOptions{ContentType: PuzzleContentType},
	)
	if err != nil {
		return err
	}
	log.Infof("Successfully uploaded %s of size %d", objectName, info.Size)
	return nil
}

func (m *MinioClient) DownloadPuzzleInput(objectName string) error {
	// TODO: Logic will be added with aoc_calc package
	return nil
}
