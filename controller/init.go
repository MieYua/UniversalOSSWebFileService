/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package controller

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	_ "UniversalOSSWebFileService/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/go-xorm/core"
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/go-xorm/xorm"
	"UniversalOSSWebFileService/model"
	"UniversalOSSWebFileService/service"
	"encoding/json"
	"log"
	"os"
)

var X *xorm.Engine
var O *service.Operation
var BUCKET_NAME string
var OSS_LOCATION string

func init() {
	// 打开配置文件
	configFile, err := os.Open("configuration.json")
	if err != nil {
		log.Println("打开配置文件错误：" + err.Error())
		return
	}

	config := model.Configuration{}
	decoderFile := json.NewDecoder(configFile)
	decoderFile.Decode(&config)

	// 数据库初始化
	mysqlUsername := config.DatabaseConfig.MysqlUsername
	mysqlPassword := config.DatabaseConfig.MysqlPassword
	mysqlAddress := config.DatabaseConfig.MysqlAddress
	mysqlDatabaseName := config.DatabaseConfig.MysqlDatabaseName

	X, err = xorm.NewEngine("mysql", mysqlUsername+":"+mysqlPassword+"@"+mysqlAddress+"/"+mysqlDatabaseName+"?charset=utf8")
	if err != nil {
		log.Println("连接数据库失败：" + err.Error())
	}

	mysqlTablePrefix := config.DatabaseConfig.MysqlTablePrefix

	if mysqlTablePrefix != "" {
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, mysqlTablePrefix)
		X.SetTableMapper(tbMapper)
	}

	X.ImportFile("configurationFile.sql")

	// 文件初始化
	endPoint := getEndpoint(config.OSSConfig.Endpoint)
	accessKeyId := config.OSSConfig.AccessKeyId
	accessKeySecret := config.OSSConfig.AccessKeySecret
	BUCKET_NAME = config.OSSConfig.OSSFileBucketName
	OSS_LOCATION = config.OSSConfig.OSSFileConfig.LocationSetting
	_, O, err = service.Connect(endPoint, accessKeyId, accessKeySecret, BUCKET_NAME, OSS_LOCATION)
	if err != nil {
		log.Println("连接文件服务错误：" + err.Error())
		return
	}
}

func getEndpoint(ep string) string {
	switch ep {
	case "BEIJING":
		return consts.ENDPOINT_BEIJING
	case "HANGZHOU":
		return consts.ENDPOINT_HANGZHOU
	case "HONGKONG":
		return consts.ENDPOINT_HONGKONG
	case "QINGDAO":
		return consts.ENDPOINT_QINGDAO
	case "SHENZHEN":
		return consts.ENDPOINT_SHENZHEN
	case "US_WEST1":
		return consts.ENDPOINT_US_WEST1
	case "BEIJING_IN":
		return consts.ENDPOINT_BEIJING_INTERNAL
	case "HANGZHOU_IN":
		return consts.ENDPOINT_HANGZHOU_INTERNAL
	case "HONGKONG_IN":
		return consts.ENDPOINT_HONGKONG_INTERNAL
	case "QINGDAO_IN":
		return consts.ENDPOINT_QINGDAO_INTERNAL
	case "SHENZHEN_IN":
		return consts.ENDPOINT_SHENZHEN_INTERNAL
	case "US_WEST1_IN":
		return consts.ENDPOINT_US_WEST1_INTERNAL
	default:
	}
	return consts.ENDPOINT_HANGZHOU
}
