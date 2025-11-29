package minio

import (
	"go-away-2024/internal/config"
	"net"

	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

func CreateClient() {
	var err error
	endpoint := net.JoinHostPort(config.S3Cfg.Host, config.S3Cfg.Port)

	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.S3Cfg.AccessKey, config.S3Cfg.SecretKey, ""),
		Secure: config.S3Cfg.SslMode,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("S3 client created: endpoint=%s", endpoint)
}
