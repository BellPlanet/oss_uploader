package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func mustParseObjKeyAndFile(objFilePair string) (string, string) {
	p := strings.SplitN(objFilePair, ":", 2)
	objKey := p[0]
	if len(objKey) == 0 {
		abortf("invalid object key: %s", objFilePair)
		return "", ""
	}
	filePath := p[1]
	if len(filePath) == 0 {
		abortf("invalid file path: %s", objFilePair)
		return "", ""
	}

	return objKey, filePath
}

const uploadUsageDoc = `oss-uploader upload BUCKET OBJKEY:FILE...`

func uploadExecute(args []string) {
	flagSet := flag.NewFlagSet("upload", flag.ExitOnError)
	flagSet.Usage = func() { fmt.Fprintf(os.Stderr, uploadUsageDoc) }
	flagSet.Parse(args)

	if flagSet.NArg() < 1 {
		abortf("missing BUCKET")
		return
	}
	if flagSet.NArg() < 2 {
		abortf("missing OBJKEY:FILE part")
		return
	}

	bucketName := flagSet.Arg(0)
	objFilePairs := flagSet.Args()[1:]

	mustInitEnvvars()
	bucket := MustGetOSSBucket(bucketName)

	var uploadJob sync.WaitGroup
	for _, objFilePair := range objFilePairs {
		objKey, filePath := mustParseObjKeyAndFile(objFilePair)

		uploadJob.Add(1)
		go func(objKey, filePath string) {
			defer uploadJob.Done()
			err := bucket.PutObjectFromFile(objKey, filePath)
			if err == nil {
				fmt.Fprintf(os.Stderr, "uploaded file: %s to %s\n", filePath, objKey)
			} else {
				fmt.Fprintf(os.Stderr, "error: upload %s to %s failed: %+v\n", filePath, objKey, err)
			}
		}(objKey, filePath)
	}

	uploadJob.Wait()
}
