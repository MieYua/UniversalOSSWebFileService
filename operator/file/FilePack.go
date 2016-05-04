/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package file

import (
	"UniversalOSSWebFileService/operator/oss"
)

//	FilePack: File information.
//	文件信息。
type FilePack struct {
	OSSClient  *oss.OSSClient
	BucketName string
	FilePaths  []string
}

//	SetFile: Put the file to a pack.
//	设置文件。
/*
 *	Example:
 *	fp := SetFiles(oc, bucketName, FileIds)
 */
func SetFiles(oc *oss.OSSClient, bucketName string, filePaths []string) *FilePack {
	fp := FilePack{
		OSSClient:  oc,
		BucketName: bucketName,
		FilePaths:  filePaths,
	}
	return &fp
}
