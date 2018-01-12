package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

const downloadUsageDoc = `oss-uploader download BUCKET OBJKEY:FILE...`

func downloadExecute(args []string) {
	flagSet := flag.NewFlagSet("download", flag.ExitOnError)
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

	var downloadJob sync.WaitGroup
	for _, objFilePair := range objFilePairs {
		objKey, filePath := mustParseObjKeyAndFile(objFilePair)

		downloadJob.Add(1)
		go func(objKey, filePath string) {
			defer downloadJob.Done()
			err := bucket.GetObjectToFile(objKey, filePath)
			if err == nil {
				fmt.Fprintf(os.Stderr, "downloaded file: %s to %s\n", objKey, filePath)
			} else {
				fmt.Fprintf(os.Stderr, "error: download %s to %s failed: %+v\n", objKey, filePath, err)
			}
		}(objKey, filePath)
	}

	downloadJob.Wait()
}
