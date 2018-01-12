package main

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func MustMakeOSSClient(endpoint, accessKeyId, accessKeySecret string) *oss.Client {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		abort(err)
		client = nil
	}

	return client
}
