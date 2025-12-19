package minio

import (
	"context"
	"fmt"
	"go-away-2024/internal/config"
	"io"
	"net"
	"os"

	"github.com/gofiber/fiber/v2/log"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	c           *minio.Client
	bucket      string
	contentType string
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
		c:           m,
		bucket:      cfg.S3.Bucket,
		contentType: cfg.S3.ContentType,
	}
}

func NewPattern(id int64, year int32, day int32, part int32) string {
	return fmt.Sprintf("Id%dYear%dDay%dPart%d-*.txt", id, year, day, part)
}

func DataError() error {
	return fmt.Errorf("incorrect input data")
}

func (m *MinioClient) UploadPuzzleInput(name string, object *os.File) error {
	objectStat, err := object.Stat()
	if err != nil {
		return err
	}
	object.Seek(0, 0)

	info, err := m.c.PutObject(
		context.Background(),
		m.bucket, name, object, objectStat.Size(),
		minio.PutObjectOptions{ContentType: m.contentType},
	)
	if err != nil {
		return err
	}
	log.Infof("Uploaded %s of size %d", name, info.Size)
	return nil
}

func (m *MinioClient) DownloadPuzzleInput(name string, object *os.File) error {
	reader, err := m.c.GetObject(context.Background(), m.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	stat, err := reader.Stat()
	if err != nil {
		return err
	}
	if _, err := io.CopyN(object, reader, stat.Size); err != nil {
		return err
	}
	log.Infof("Downloaded %s of size %d", name, stat.Size)
	return nil
}
