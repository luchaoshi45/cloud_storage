package ceph

import (
	"cloud_storage/configurator"
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var cephConn *s3.S3

func init() {
	// 1. 初始化 ceph 的一些信息
	cephConfig := configurator.GetCephConfig()

	auth := aws.Auth{
		AccessKey: cephConfig.GetAttr("AccessKey"),
		SecretKey: cephConfig.GetAttr("SecretKey"),
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          cephConfig.GetAttr("Endpoint"),
		S3Endpoint:           cephConfig.GetAttr("Endpoint"),
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	// 2. 创建S3类型的连接
	cephConn = s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func GetCephBucket(bucket string) *s3.Bucket {
	return cephConn.Bucket(bucket)
}

// PutObject : 上传文件到ceph集群
func PutObject(bucket string, path string, data []byte) error {
	return GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}
