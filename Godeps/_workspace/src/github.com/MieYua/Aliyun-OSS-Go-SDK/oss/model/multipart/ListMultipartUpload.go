/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package multipart

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// 	Give the list of the unuploaded multipart upload missions.
//	列出所有未上传完整的multipart任务列表。
/*
 *	Example:
 *	lmur, err := c.ListMultipartUpload(bucketName, map[string]string or nil)
 *	If you want to search all results, second parameter is nil.
 */
func (c *Client) ListMultipartUpload(bucketName string, params map[string]string) (lmur types.ListMultipartUploadResult, err error) {
	cc := c.CClient

	if strings.HasPrefix(bucketName, "/") == false {
		bucketName = "/" + bucketName
	}

	reqStr := bucketName + "?uploads"
	if params != nil {
		for k, v := range params {
			reqStr += "&" + k + "=" + v
		}
	}

	resp, err := cc.DoRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &lmur)
	if err != nil {
		return
	}

	//log.Println("You have got all the unuploaded parts' details of " + bucketName + ".")
	return
}
