/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package service

import (
	"UniversalOSSWebFileService/model"
	"UniversalOSSWebFileService/operator/file"
	"UniversalOSSWebFileService/operator/oss"
	"UniversalOSSWebFileService/operator/upload"
	"encoding/json"
	"log"
	"os"
	"strings"
)

var BUCKET_NAME string
var OSS_LOCATION string

type Operation struct {
	OSSClient *oss.OSSClient
	Uploader  *upload.Uploader
	FilePack  *file.FilePack
}

//	Connect
//	Connect to OSS then check the settings of buckets.
// 	连接到OSS和数据库并检查设置。
/*
 *	Example:
 *	cr, o, err := Connect(endPoint, accessKeyId, accessKeySecret)
 */
func Connect(endPoint, accessKeyId, accessKeySecret, fileBucketName, fileBucketLocation string) (connectResult string, o *Operation, err error) {
	BUCKET_NAME = fileBucketName
	OSS_LOCATION = fileBucketLocation

	// 获得OSS客户端
	oc := oss.SetOSSClient(endPoint, accessKeyId, accessKeySecret)
	operation := Operation{}
	lor := model.ConnectOSSResult{}
	cr := model.ConnectResult{}
	j, err := oc.ConnectOSS()
	if err != nil {
		log.Println(err.Error())
		return
	}
	json.Unmarshal([]byte(j), &lor)
	cr.ConnectLogging = append(cr.ConnectLogging, "Connecting...")
	if strings.Contains(j, "RequestTimeout") {
		for i := 0; i < 5; i++ {
			cr.ConnectLogging = append(cr.ConnectLogging, "Request timeout.")
			j, err = oc.ConnectOSS()
			json.Unmarshal([]byte(j), &lor)
			cr.ConnectLogging = append(cr.ConnectLogging, "Retry to Connect...")
			if strings.Contains(j, "Connect to OSS successfully.") {
				break
			}
		}
		if strings.Contains(j, "RequestTimeout") {
			cr.ConnectLogging = append(cr.ConnectLogging, "Request timeout: Too many tries.")
			cr.ConnectResult = "Connect unsuccessfully."
			b, _ := json.Marshal(cr)
			connectResult = string(b)
			log.Println(cr.ConnectResult)
			return
		}
	}
	operation.OSSClient = oc
	cr.Owner.ID = lor.Owner.ID
	cr.Owner.DisplayName = lor.Owner.DisplayName

	// 检查OSS的bucket状态
	j, err = oc.CheckOSS()
	cor := model.CheckOSSResult{}
	json.Unmarshal([]byte(j), &cor)
	if cor.CreateBucketLogging == "" {
		cr.ConnectLogging = append(cr.ConnectLogging, cor.CheckBucketExist)
	} else {
		cr.ConnectLogging = append(cr.ConnectLogging, cor.CreateBucketLogging)
		cr.ConnectLogging = append(cr.ConnectLogging, cor.CheckBucketExist)
	}
	if cor.ChangeSettingLogging == "" {
		cr.ConnectLogging = append(cr.ConnectLogging, cor.CheckBucketSetting)
	} else {
		cr.ConnectLogging = append(cr.ConnectLogging, cor.ChangeSettingLogging)
		cr.ConnectLogging = append(cr.ConnectLogging, cor.CheckBucketSetting)
	}

	cr.ConnectResult = "Connect successfully."
	b, _ := json.Marshal(cr)
	return string(b), &operation, nil
}

//	StartFileOperation
//	Initiate to operate with a file or files of bucket.
// 	初始化xxx-file中的文件操作。
/*
 *	Example:
 *	_, o, _ := Connect(endPoint, accessKeyId, accessKeySecret, databaseUsername, databasePassword, databaseAddress, databaseName)
 *	o.StartFileOperation([]string{filePath1,filePath2,...})
 */
func (o *Operation) StartFileOperation(filePaths []string) {
	if len(filePaths) >= 0 {
		fp := file.SetFiles(o.OSSClient, BUCKET_NAME, filePaths)
		log.Println("The operations of these objects are ready.")
		o.FilePack = fp
	} else {
		log.Println("The inputs of the objects are not allowed.")
	}
}

//	Upload
//	Upload a file into bucket by POST.
// 	将文件上传至xx-file（POST上传方式）。
func (o *Operation) UploadWeb(objectPath string, file *os.File, uploaderId string) (filePath, fileType string, err error) {
	u := upload.SetUploader(o.OSSClient, BUCKET_NAME, OSS_LOCATION, objectPath, uploaderId, file)
	o.Uploader = u
	filePath, fileType, err = u.UploadFileWeb()
	return
}

//	Delete
//	Delete a file or files of bucket.
// 	完全删除xx-file中的文件。
/*
 *	Example:
 *	_, o, _ := Connect(endPoint, accessKeyId, accessKeySecret, databaseUsername, databasePassword, databaseAddress, databaseName)
 *	o.StartFileOperation([]string{filePath1,filePath2,...})
 * 	err := o.Delete()
 */
func (o *Operation) Delete() (dr string, err error) {
	fp := o.FilePack
	err = fp.DeleteFile()
	return
}
