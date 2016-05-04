/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package controller

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/Unknwon/macaron"
	"UniversalOSSWebFileService/model"
	"encoding/json"
	"log"
)

//	Get the download address of the file.
//	获得单个文件的下载地址。
/*
 *	Example:
 *	PUT http://URL:port/v1/file/47
 *  Body(Json)
 *  {
 *		"file_name":"new file name",
 *  	"description":"new description"
 *  }
 */
func PutFile(ctx *macaron.Context) string {
	ctx.Resp.Header().Add("Access-Control-Allow-Origin", "*")
	respj := model.SingleResponseJson{}
	fileId := ctx.ParamsInt("fileId")

	// 获取Body的参数
	bytes, _ := ctx.Req.Body().Bytes()
	reqj := model.RequestJSON{}
	json.Unmarshal(bytes, &reqj)
	fileName := reqj.FileName
	description := reqj.Description

	// 判断文件id
	if fileId <= 0 {
		// 不合法
		respj.Meta.Code = 410
		respj.Meta.Message = "文件id参数不合法"
		log.Println("410 文件id参数不合法")
		b, _ := json.Marshal(respj)
		return string(b)
	} else {
		// 检查id存在性
		isExist, err := checkFileExist(fileId)
		if err != nil {
			respj.Meta.Code = 503
			respj.Meta.Message = "文件数据库读取失败 " + err.Error()
			log.Println("503 文件数据库读取失败 " + err.Error())
			b, _ := json.Marshal(respj)
			return string(b)
		}
		if isExist == false {
			respj.Meta.Code = 410
			respj.Meta.Message = "该id对应文件不存在"
			log.Println("410 该id对应文件不存在")
			b, _ := json.Marshal(respj)
			return string(b)
		}
		// 更新文件信息
		err = updateFile(fileId, fileName, description)
		if err != nil {
			respj.Meta.Code = 503
			respj.Meta.Message = "文件数据库读取失败 " + err.Error()
			log.Println("503 文件数据库读取失败 " + err.Error())
			b, _ := json.Marshal(respj)
			return string(b)
		}
	}

	// 成功返回
	respj.Meta.Code = 200
	respj.Meta.Message = "成功更新文件信息"
	log.Println("200 成功更新文件信息")
	b, _ := json.Marshal(respj)
	return string(b)
}
