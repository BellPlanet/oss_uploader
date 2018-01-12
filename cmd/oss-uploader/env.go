package main

import (
	"fmt"
	"os"
)

const (
	ENV_KEY_OSS_ENDPOINT  = "OSS_UPLOADER_ENDPOINT"
	ENV_KEY_OSS_AK_ID     = "OSS_UPLOADER_ACCESS_KEY_ID"
	ENV_KEY_OSS_AK_SECRET = "OSS_UPLOADER_ACCESS_KEY_SECRET"
)

func errEnvIsMissing(k string) error {
	return fmt.Errorf("env %s is missing", k)
}

func mustInitEnvvars() {
	for _, k := range []string{ENV_KEY_OSS_ENDPOINT, ENV_KEY_OSS_AK_ID, ENV_KEY_OSS_AK_SECRET} {
		if os.Getenv(k) == "" {
			abort(errEnvIsMissing(k))
		}
	}
}
