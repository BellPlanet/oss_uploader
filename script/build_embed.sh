#!/bin/bash

set -e
set -o errexit
set -o nounset
set -o pipefail

CURDIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
BUILDPATH="$(dirname "$CURDIR")"

help() {
    echo 'Usage: build_embed.sh binary_name'
}

die() {
    echo "$1"
    exit 1
}

build() {
    local binname=$1
    local platform=$2
    local full_name="${binname}-${platform}-amd64"
    local full_path="${BUILDPATH}/${full_name}"

    rm -rf "$full_path"
    echo "building ${full_path}"
    GOOS="${platform}" GOARCH=amd64 CGO_ENABLED=0 go build \
        -ldflags "-X main.OSSSettingsSet=true -X main.OSSEndpoint=${OSS_UPLOADER_ENDPOINT} -X main.OSSAccessKeyId=${OSS_UPLOADER_ACCESS_KEY_ID} -X main.OSSAccessKeySecret=${OSS_UPLOADER_ACCESS_KEY_SECRET}" \
        -o "${full_path}" \
        "./cmd/${binname}"
}

BINNAME=${1:-""}
OSS_UPLOADER_ENDPOINT=${OSS_UPLOADER_ENDPOINT:-""}
OSS_UPLOADER_ACCESS_KEY_ID=${OSS_UPLOADER_ACCESS_KEY_ID:-""}
OSS_UPLOADER_ACCESS_KEY_SECRET=${OSS_UPLOADER_ACCESS_KEY_SECRET:-""}

[[ -z $OSS_UPLOADER_ENDPOINT ]] && die 'env OSS_UPLOADER_ENDPOINT required'
[[ -z $OSS_UPLOADER_ACCESS_KEY_ID ]] && die 'env OSS_UPLOADER_ACCESS_KEY_ID required'
[[ -z $OSS_UPLOADER_ACCESS_KEY_SECRET ]] && die 'env OSS_UPLOADER_ACCESS_KEY_SECRET required'

if [[ -z $BINNAME ]]
then
    help
    exit -1
fi

case "$OSTYPE" in
    linux*)
        build "$BINNAME" linux
        ;;
    darwin*)
        build "$BINNAME" darwin
        ;;
    *)
        echo "unsupported platform $OSTYPE"
        exit -1
        ;;
esac
