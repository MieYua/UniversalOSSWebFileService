/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package controller

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/Unknwon/macaron"
	"UniversalOSSWebFileService/model"
	"UniversalOSSWebFileService/util"
	"encoding/json"
	"io"
	"log"
	"os"
)

//	Upload a new file.
//	上传新文件。
/*
 *	Example:
 *	POST http://URL:port/v1/file
 *  Content-Type: multipart/form-data;boundary=2a6bd1629b6c998cf7f6b6a4793365fcc7e9f3b1f279dac932519edfd653
 *
 	--2a6bd1629b6c998cf7f6b6a4793365fcc7e9f3b1f279dac932519edfd653
	Content-Disposition: form-data; name="UploaderId"

	XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
	--2a6bd1629b6c998cf7f6b6a4793365fcc7e9f3b1f279dac932519edfd653
	Content-Disposition: form-data; name="Description"

	XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	--2a6bd1629b6c998cf7f6b6a4793365fcc7e9f3b1f279dac932519edfd653
	Content-Disposition: form-data; name="file"; filename="test.txt"
	Content-Type: application/octet-stream

	...
	--2a6bd1629b6c998cf7f6b6a4793365fcc7e9f3b1f279dac932519edfd653--
*/
func PostFile(ctx *macaron.Context) string {
	ctx.Resp.Header().Add("Access-Control-Allow-Origin", "*")
	respj := model.SingleResponseJson{}

	// 文件信息获取
	r := ctx.Req
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		respj.Meta.Code = 204
		respj.Meta.Message = "没有获得到文件信息，请检查后重试。"
		log.Println("204 " + err.Error())
		b, _ := json.Marshal(respj)
		return string(b)
	}
	defer file.Close()

	// 建立文件缓存
	fileName := header.Filename
	fileName = util.GetFileName(fileName)
	f, err := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		respj.Meta.Code = 503
		respj.Meta.Message = "新建文件缓存失败 " + err.Error()
		log.Println("503 " + err.Error())
		b, _ := json.Marshal(respj)
		return string(b)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		respj.Meta.Code = 503
		respj.Meta.Message = "文件写入失败 " + err.Error()
		log.Println("503 File's error " + err.Error())
		b, _ := json.Marshal(respj)
		return string(b)
	}
	uploaderId := r.FormValue("UploaderId")
	description := r.FormValue("Description")

	// 文件上传
	filePath, fileType, err := O.UploadWeb(f.Name(), f, uploaderId)
	if err != nil {
		respj.Meta.Code = 503
		respj.Meta.Message = "文件上传多次失败，请稍后重试。"
		b, _ := json.Marshal(respj)
		log.Println("503 " + err.Error())
		return string(b)
	}
	f.Close()
	os.Remove(header.Filename)

	// 文件信息写入数据库
	err = insertFile(fileName, filePath, fileType, uploaderId, description)
	if err != nil {
		respj.Meta.Code = 503
		respj.Meta.Message = "文件上传数据库写入失败，请稍后重试。"
		b, _ := json.Marshal(respj)
		log.Println("503 " + err.Error())
		return string(b)
	}

	// 成功返回
	respj.Meta.Code = 201
	respj.Meta.Message = "文件上传成功"
	log.Println("201 文件(" + fileName + ")上传成功")
	b, _ := json.Marshal(respj)
	return string(b)
}
