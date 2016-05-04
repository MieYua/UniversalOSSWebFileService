/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package main

import (
	. "UniversalOSSWebFileService/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"UniversalOSSWebFileService/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

const (
	TEST_URL         = "http://localhost:4000/v1"
	TEST_UPLOADER_ID = "TESTTEST-TEST-TEST-TEST-TESTTESTTEST"
	TEST_DESCRIPTION = "Test file"
	TEST_FILE_ID     = "1"
	GO_PATH          = "X://xxxxxx"
)

func TestGetSingleFile(t *testing.T) {
	Convey("获得单个文件测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/file/" + TEST_FILE_ID
		req, _ := http.NewRequest("GET", url, nil)
		fmt.Println(req)
		resp, _ := c.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(b))
	})
}

func TestGetAllFiles(t *testing.T) {
	Convey("获得所有文件测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/file"
		req, _ := http.NewRequest("GET", url, nil)
		fmt.Println(req)
		resp, _ := c.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(b))
	})
}

func TestGetSingleFileDownloadLink(t *testing.T) {
	Convey("获得单个文件下载地址测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/fileAddress/" + TEST_FILE_ID
		req, _ := http.NewRequest("GET", url, nil)
		fmt.Println(req)
		resp, _ := c.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(b))
	})
}

func TestPostFile(t *testing.T) {
	Convey("文件上传测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/file"

		buffer := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(buffer)

		bodyWriter.CreateFormField("UploaderId")
		buffer.WriteString(TEST_UPLOADER_ID)
		bodyWriter.CreateFormField("Description")
		buffer.WriteString(TEST_DESCRIPTION)

		fileWriter, _ := bodyWriter.CreateFormFile("file", "test.jpg")
		fh, _ := os.Open(GO_PATH + "/src/UniversalOSSWebFileService/test/test.jpg")
		defer fh.Close()
		io.Copy(fileWriter, fh)
		bodyWriter.Close()
		fmt.Println(url, "multipart/form-data;boundary="+bodyWriter.Boundary())
		//fmt.Println(buffer)
		resp, _ := c.Post(url, "multipart/form-data;boundary="+bodyWriter.Boundary(), buffer)
		if resp != nil {
			b, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			fmt.Println(string(b))
		}
	})
}

func TestPutFile(t *testing.T) {
	Convey("文件重命名测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/file/" + TEST_FILE_ID
		reqj := model.RequestJSON{}
		reqj.FileName = "file_rename"
		reqj.Description = "Test file new description"
		bs, _ := json.Marshal(reqj)
		buffer := new(bytes.Buffer)
		buffer.Write(bs)
		req, _ := http.NewRequest("PUT", url, buffer)
		fmt.Println(req)
		resp, _ := c.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(b))
	})
}

func TestDeleteFile(t *testing.T) {
	Convey("文件删除测试", t, func() {
		fmt.Println("")
		c := new(http.Client)
		url := TEST_URL + "/file/" + TEST_FILE_ID
		req, _ := http.NewRequest("DELETE", url, nil)
		fmt.Println(req)
		resp, _ := c.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(b))
	})
}
