/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package oss

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss"
	"UniversalOSSWebFileService/model"
	"encoding/json"
	"log"
	"os"
)

//	Check the existence and settings of buckets(xxx-file and xxx-log).
//	检查xxx-file和xxx-log是否存在和设置。
/*
 *	Example:
 *	j, err := oc.CheckOSS()
 */
func (oc *OSSClient) CheckOSS() (j string, err error) {
	configFile, err := os.Open("configuration.json")
	if err != nil {
		log.Println("打开配置文件错误：" + err.Error())
		return
	}

	config := model.Configuration{}
	decoderFile := json.NewDecoder(configFile)
	decoderFile.Decode(&config)

	fileBucketName := config.OSSConfig.OSSFileBucketName
	logBucketName := config.OSSConfig.OSSLogBucketName

	c := oss.InitiateClient(oc.EndPoint, oc.AccessKeyId, oc.AccessKeySecret)
	lambr, err := c.GetServiceInfo()
	if err != nil {
		return
	}

	// 	Check whether the buckets exist.
	isFileBucketExist := false
	isLogBucketExist := false
	cor := model.CheckOSSResult{}
	for _, v := range lambr.Buckets.Bucket {
		if v.Name == fileBucketName {
			isFileBucketExist = true
		}
		if v.Name == logBucketName {
			isLogBucketExist = true
		}
	}

	if isFileBucketExist == true {
		if isLogBucketExist == false {
			cor.CreateBucketLogging = "The bucket(" + fileBucketName + ") exist and the bucket(" + logBucketName + ") don't exist. Create the bucket(" + logBucketName + ")."
			err = c.CreateBucket(logBucketName)
			if err != nil {
				return
			}
		}
	} else {
		if isLogBucketExist == true {
			cor.CreateBucketLogging = "The bucket(" + logBucketName + ") exist and the bucket(" + fileBucketName + ") don't exist. Create the bucket(" + fileBucketName + ")."
			err = c.CreateBucket(fileBucketName)
			if err != nil {
				return
			}
		} else {
			cor.CreateBucketLogging = "The buckets(" + fileBucketName + " and " + logBucketName + ") don't exist. Create the buckets(" + fileBucketName + " and " + logBucketName + ")."
			err = c.CreateBucket(fileBucketName)
			if err != nil {
				return
			}
			err = c.CreateBucket(logBucketName)
			if err != nil {
				return
			}
		}
	}
	cor.CheckBucketExist = "The buckets(" + fileBucketName + " and " + logBucketName + ") exist."

	//	Check Buckets' settings
	isSameFile := false
	isSameLog := false

	isSameFile, err = oc.CompareSetting(fileBucketName, config.OSSConfig.OSSFileConfig)
	if err != nil {
		return
	}

	isSameLog, err = oc.CompareSetting(logBucketName, config.OSSConfig.OSSLogConfig)
	if err != nil {
		return
	}

	if isSameFile == false {
		err = oc.DefaultSetting(fileBucketName, config.OSSConfig.OSSFileConfig)
		if err != nil {
			return
		}
		if isSameLog == false {
			err = oc.DefaultSetting(""+logBucketName+"", config.OSSConfig.OSSLogConfig)
			if err != nil {
				return
			}
			cor.ChangeSettingLogging = "Default the buckets(" + fileBucketName + " and " + logBucketName + ")."
		} else {
			cor.ChangeSettingLogging = "Default the bucket(" + fileBucketName + ")."
		}
	} else {
		if isSameLog == false {
			err = oc.DefaultSetting(""+logBucketName+"", config.OSSConfig.OSSLogConfig)
			if err != nil {
				return
			}
			cor.ChangeSettingLogging = "Default the bucket(" + logBucketName + ")."
		}
	}
	cor.CheckBucketSetting = "The settings of buckets(" + fileBucketName + " and " + logBucketName + ") are ok."
	b, _ := json.Marshal(cor)
	j = string(b)
	return
}
