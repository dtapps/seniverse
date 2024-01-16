package seniverse

import "go.dtapp.net/golog"

type V3Client struct {
	key     string // API密钥
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

func NewV3Client(key string) (*V3Client, error) {
	return &V3Client{key: key}, nil
}

type V4Client struct {
	publicKey string
	secret    string
	gormLog   struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

func NewV4Client(publicKey string, secret string) (*V4Client, error) {
	return &V4Client{publicKey: publicKey, secret: secret}, nil
}
