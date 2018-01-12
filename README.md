# oss-uploader

[Aliyun OSS][aliyun-oss] cli uploader.

[aliyun-oss]: https://www.aliyun.com/product/oss

## Usage

### Upload

```
export OSS_UPLOADER_ENDPOINT=oss-cn-shenzhen.aliyuncs.com
export OSS_UPLOADER_ACCESS_KEY_ID=xxx
export OSS_UPLOADER_ACCESS_KEY_SECRET=supersecret
oss-uploader upload bucket myfile:myfile.zip
```

### Download

```
export OSS_UPLOADER_ENDPOINT=oss-cn-shenzhen.aliyuncs.com
export OSS_UPLOADER_ACCESS_KEY_ID=xxx
export OSS_UPLOADER_ACCESS_KEY_SECRET=supersecret
oss-uploader download bucket myfile:myfile.zip
```

## Embeded Build

To simplify usage, you can pre-embed `OSS_UPLOADER_*` settings into binary:

```
export OSS_UPLOADER_ENDPOINT=oss-cn-shenzhen.aliyuncs.com
export OSS_UPLOADER_ACCESS_KEY_ID=xxx
export OSS_UPLOADER_ACCESS_KEY_SECRET=supersecret
./script/build_embed.sh oss-uploader
```

## License

see [LICENSE](./LICENSE.md)
