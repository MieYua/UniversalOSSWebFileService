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

//	Get all files(or a file).
//	获得所有文件（或单个文件）信息。
/*
 *	Example:
 *  All files:
 *	GET http://URL:port/v1/file
 *  or
 *  Single file:
 *  GET http://URL:port/v1/file/47
 */
func GetFiles(ctx *macaron.Context) string {
	ctx.Resp.Header().Add("Access-Control-Allow-Origin", "*")
	respj := model.SingleResponseJson{}
	fileId := ctx.ParamsInt("fileId")

	// 判断输入id
	if fileId < 0 {
		// 不合法
		respj.Meta.Code = 410
		respj.Meta.Message = "文件id参数不合法"
		log.Println("410 文件id参数不合法")
		b, _ := json.Marshal(respj)
		return string(b)
	} else if fileId == 0 {
		// 所有文件
		respjall := model.ResponseJson{}
		files, err := getAllFiles()
		if err != nil {
			respjall.Meta.Code = 503
			respjall.Meta.Message = "文件数据库读取失败 " + err.Error()
			log.Println("503 文件数据库读取失败 " + err.Error())
			b, _ := json.Marshal(respjall)
			return string(b)
		}
		for _, file := range files {
			respjall.Data = append(respjall.Data, file)
		}
		respjall.Meta.Code = 200
		respjall.Meta.Message = "成功得到所有文件信息"
		log.Println("200 成功得到所有文件信息")
		b, _ := json.Marshal(respjall)
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
		// 单个文件
		file, err := getSingleFile(fileId)
		if err != nil {
			respj.Meta.Code = 503
			respj.Meta.Message = "文件数据库读取失败 " + err.Error()
			log.Println("503 文件数据库读取失败 " + err.Error())
			b, _ := json.Marshal(respj)
			return string(b)
		}
		respj.Data = file
	}

	// 成功返回
	respj.Meta.Code = 200
	respj.Meta.Message = "成功得到文件信息"
	log.Println("200 成功得到文件信息")
	b, _ := json.Marshal(respj)
	return string(b)
}
