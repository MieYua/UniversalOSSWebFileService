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

//	Delete the file by fileId.
//	删除文件（不可恢复）。
/*
 *	Example:
 *	DELETE http://URL:port/v1/file/47
 */
func DeleteFile(ctx *macaron.Context) string {
	ctx.Resp.Header().Add("Access-Control-Allow-Origin", "*")
	respj := model.SingleResponseJson{}
	fileId := ctx.ParamsInt("fileId")

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
		// 删除文件
		filePaths := make([]string, 0)
		file, err := getSingleFile(fileId)
		if err != nil {
			respj.Meta.Code = 503
			respj.Meta.Message = "文件数据库获取失败，请稍后重试。"
			b, _ := json.Marshal(respj)
			log.Println("503 " + err.Error())
			return string(b)
		}
		filePaths = append(filePaths, file.FilePath)
		O.StartFileOperation(filePaths)
		_, err = O.Delete()
		if err != nil {
			respj.Meta.Code = 503
			respj.Meta.Message = "文件删除失败，请稍后重试。"
			b, _ := json.Marshal(respj)
			log.Println("503 " + err.Error())
			return string(b)
		}
	}

	// 文件信息从数据库删除
	err := deleteFile(fileId)
	if err != nil {
		respj.Meta.Code = 503
		respj.Meta.Message = "文件上传数据库删除失败，请稍后重试。"
		b, _ := json.Marshal(respj)
		log.Println("503 " + err.Error())
		return string(b)
	}

	// 成功返回
	respj.Meta.Code = 204
	respj.Meta.Message = "文件删除成功"
	log.Println("204 文件删除成功")
	b, _ := json.Marshal(respj)
	return string(b)
}
