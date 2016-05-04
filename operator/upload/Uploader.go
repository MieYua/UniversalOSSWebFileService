/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package upload

import (
	"UniversalOSSWebFileService/operator/oss"
	"os"
)

//	Uploader: Uploader.
//	上传器。
type Uploader struct {
	Client         *oss.OSSClient
	BucketName     string
	BucketLocation string
	ObjectPath     string
	File           *os.File
	FileSize       int64
	FileStartPoint int64
	UploaderId     string
}

//	SetUploader: Set the Uploader.
//	设置上传器。
/*
 *	Example:
 *	oc := SetUploader(c, bucketName, bucketLocation, objectPath, uploaderId, file)
 */
func SetUploader(c *oss.OSSClient, bucketName, bucketLocation, objectPath, uploaderId string, file *os.File) *Uploader {
	size := int64(0)
	if file != nil {
		fi, _ := file.Stat()
		size = fi.Size()
	}

	u := Uploader{
		Client:         c,
		BucketName:     bucketName,
		BucketLocation: bucketLocation,
		ObjectPath:     objectPath,
		File:           file,
		FileSize:       size,
		FileStartPoint: 0,
		UploaderId:     uploaderId,
	}
	return &u
}
