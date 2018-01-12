package main

import (
	"flag"
	"fmt"
	"os"
)

const usageDoc = `Usage oss-uploader COMMAND [arg...]
oss-uploader [-help]

Commands:
  upload Upload files to OSS.
  download Download files from OSS.
`

func main() {
	help := flag.Bool("help", false, "")

	flag.Usage = func() { fmt.Fprintf(os.Stderr, usageDoc) }
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	switch flag.Arg(0) {
	case "upload":
		args := flag.Args()[1:]
		uploadExecute(args)
	case "download":
		args := flag.Args()[1:]
		downloadExecute(args)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func abort(err error) {
	fmt.Fprintf(os.Stderr, "error: %+v\n", err)
	os.Exit(-1)
}

func abortf(l string, args ...interface{}) {
	abort(fmt.Errorf(l, args...))
}
