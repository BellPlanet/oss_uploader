package main

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func MustMakeOSSClient() *oss.Client {
	client, err := oss.New(OSSEndpoint, OSSAccessKeyId, OSSAccessKeySecret)
	if err != nil {
		abort(err)
		client = nil
	}

	return client
}

func MustGetOSSBucket(bucketName string) *oss.Bucket {
	client := MustMakeOSSClient()
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		abort(err)
		bucket = nil
	}

	return bucket
}
