package minio

import (
	"go-away-2024/internal/config"
	"net"

	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(cfg *config.Config) *minio.Client {
	endpoint := net.JoinHostPort(cfg.S3.Host, cfg.S3.Port)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3.AccessKey, cfg.S3.SecretKey, ""),
		Secure: cfg.S3.SslMode,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("S3 client created: endpoint=%s", endpoint)
	return client
}
