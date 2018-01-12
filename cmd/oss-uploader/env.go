package main

import "os"

const (
	ENV_KEY_OSS_ENDPOINT  = "OSS_UPLOADER_ENDPOINT"
	ENV_KEY_OSS_AK_ID     = "OSS_UPLOADER_ACCESS_KEY_ID"
	ENV_KEY_OSS_AK_SECRET = "OSS_UPLOADER_ACCESS_KEY_SECRET"
)

var (
	OSSSettingsSet     string
	OSSEndpoint        string
	OSSAccessKeyId     string
	OSSAccessKeySecret string
)

func mustReadEnvvar(envkey string) string {
	value := os.Getenv(envkey)
	if value == "" {
		abortf("env %s is missing", envkey)
	}

	return value
}

func mustInitEnvvars() {
	if OSSSettingsSet != "" {
		// already set before
		return
	}

	OSSEndpoint = mustReadEnvvar(ENV_KEY_OSS_ENDPOINT)
	OSSAccessKeyId = mustReadEnvvar(ENV_KEY_OSS_AK_ID)
	OSSAccessKeySecret = mustReadEnvvar(ENV_KEY_OSS_AK_SECRET)
	OSSSettingsSet = "true"
}
