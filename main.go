/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package main

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/Unknwon/macaron"
	"UniversalOSSWebFileService/controller"
	"UniversalOSSWebFileService/model"
	"encoding/json"
	"log"
	"os"
)

func main() {
	// 读取端口配置信息
	configFile, err := os.Open("configuration.json")
	if err != nil {
		log.Println("打开配置文件错误：" + err.Error())
		return
	}
	config := model.Configuration{}
	decoderFile := json.NewDecoder(configFile)
	decoderFile.Decode(&config)
	port := config.ServerConfig.Port

	//	新建马卡龙经典实例
	m := macaron.Classic()

	m.SetURLPrefix("/v1")

	//	文件操作
	m.Group("/file", func() {
		//	获得所有文件（或单个文件带fileId）信息
		m.Get("/?:fileId", controller.GetFiles)
		m.Options("/?:fileId", controller.CORSVerify)

		//	获得文件信息（获取下载地址）
		m.Get("Address/:fileId:int", controller.GetFileDownloadlink)
		m.Options("Address/:fileId:int", controller.CORSVerify)

		//	上传文件（文件存入后OSS存入随机唯一文件名）
		m.Post("/", controller.PostFile)
		m.Options("/", controller.CORSVerify)

		//	更新文件（只可重命名数据库内的文件名）
		m.Put("/:fileId:int", controller.PutFile)
		m.Options("/:fileId:int", controller.CORSVerify)

		//	删除文件（完全删除文件，不可恢复）
		m.Delete("/:fileId:int", controller.DeleteFile)
		m.Options("/:fileId:int", controller.CORSVerify)
	})

	//	运行马卡龙实例（监听端口）
	m.Run(port)
}
