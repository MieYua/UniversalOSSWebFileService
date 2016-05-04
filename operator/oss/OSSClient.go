/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package oss

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss"
)

//	OSS客户端类
type OSSClient struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	Client          *oss.Client
}

//	设置OSS客户端
/*
 *	Example:
 *	oc := SetOSSClient(endPoint, accessKeyId, accessKeySecret)
 */
func SetOSSClient(endPoint, accessKeyId, accessKeySecret string) *OSSClient {
	oc := OSSClient{
		EndPoint:        endPoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	return &oc
}
