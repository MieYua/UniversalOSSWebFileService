/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package upload

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss"
	"UniversalOSSWebFileService/util"
	"strings"
)

//	Upload a file by POST (Web).
//	Post上传文件。
/*
 *	Example:
 *	fileRealPath, err := u.UploadFileWeb()
 */
func (u *Uploader) UploadFileWeb() (fileRealPath string, fileType string, err error) {
	filePath := u.ObjectPath
	uploaderId := u.UploaderId
	if uploaderId == "" {
		uploaderId = "default"
	}

	typeSlice := strings.LastIndex(filePath, "/")
	typeDot := strings.LastIndex(filePath, ".")
	rs := []byte(filePath)
	ls := len(rs)
	fileName := string(rs[typeSlice+1 : typeDot])
	fileType = string(rs[typeDot+1 : ls])
	fileType = strings.ToLower(fileType)
	if u.File.Name() != "" {
		fileName = u.File.Name()
	}

	filePath = uploaderId + "/" + util.RandomAlphaOrNumeric(47, true, true) + string(rs[typeDot:ls])

	uc := oss.InitiateClient(u.Client.EndPoint, u.Client.AccessKeyId, u.Client.AccessKeySecret)

	err = uc.PostObject(u.BucketName, filePath, fileName)
	if err != nil {
		for i := 0; i < 3; i++ {
			err = uc.PostObject(u.BucketName, filePath, fileName)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		return
	}
	fileRealPath = filePath
	return
}
