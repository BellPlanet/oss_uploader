package main

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func MustMakeOSSClient() *oss.Client {
	client, err := oss.New(
		os.Getenv(ENV_KEY_OSS_ENDPOINT),
		os.Getenv(ENV_KEY_OSS_AK_ID),
		os.Getenv(ENV_KEY_OSS_AK_SECRET),
	)
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
