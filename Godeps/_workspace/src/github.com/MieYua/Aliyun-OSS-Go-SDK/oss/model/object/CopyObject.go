/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

// 	Copy an object of one bucket to another bucket or this one.
//	拷贝Object。
/*
 *	Example:
 *	err := c.CopyObject(pasteSrc, copySrc)
 *	copySrc: 	the file need be copied(bucket1Name/../fileName)
 *	pasteSrc:	the file src need be pasted(bucket2Name/../copy_fileName)
 *
 *	If file size is larger than 1GB, please use function UploadPartCopy
 */
func (c *Client) CopyObject(pasteSrc, copySrc string) (cor types.CopyObjectResult, err error) {
	cc := c.CClient

	if strings.HasPrefix(copySrc, "/") == false {
		copySrc = "/" + copySrc
	}
	if strings.HasPrefix(pasteSrc, "/") == false {
		pasteSrc = "/" + pasteSrc
	}
	params := map[string]string{consts.OH_COPY_OBJECT_SOURCE: copySrc}
	resp, err := cc.DoRequest("PUT", pasteSrc, pasteSrc, params, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &cor)
	if err != nil {
		return
	}

	//fmt.Println("The object(" + copySrc + ") has been copied to (" + pasteSrc + ").")
	return
}
