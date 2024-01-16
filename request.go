package seniverse

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
	"time"
)

func (c *V3Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	param.Set("key", c.key)
	client.SetParams(param)

	// 发起请求
	request, err := client.Get(ctx, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.Middleware(ctx, request)
	}

	return request, err
}

func (c *V4Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置参数
	param.Set("ts", fmt.Sprintf("%d", time.Now().Unix()))
	param.Set("ttl", "600")
	param.Set("public_key", c.publicKey)

	// 签名并返回请求地址
	urlStr := c.sign(url, param)

	// 设置请求地址
	client.SetUri(urlStr)

	// 设置方式
	client.SetMethod(method)

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置格式
	client.SetContentTypeJson()

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.Middleware(ctx, request)
	}

	return request, err
}
