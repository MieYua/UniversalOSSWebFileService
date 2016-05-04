/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// 	Change the setting of this bucket's logging.
//	修改Bucket的日志设置。
/*
 *	Example:
 *	err := PutBucketLogging(bucketName, targetBucket, targetPrefix)
 *	If the targetBucket is null, its default is "bucketName+'logs'".
 *	If the targetPrefix is null, its default is "MyLog-".
 */
func (c *Client) PutBucketLogging(bucketName, targetBucket, targetPrefix string) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?logging"

	lxml := types.LoggingXML{}

	if targetBucket == "" {
		lxml.LoggingEnabled.TargetBucket = bucketName + "logs"
	} else {
		lxml.LoggingEnabled.TargetBucket = targetBucket
	}
	lxml.LoggingEnabled.TargetPrefix = targetPrefix

	bs, err := xml.Marshal(lxml)
	if err != nil {
		return
	}
	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	contentType := http.DetectContentType(buffer.Bytes())
	params := map[string]string{}
	params[consts.HH_CONTENT_TYPE] = contentType

	resp, err := cc.DoRequest("PUT", reqStr, reqStr, params, buffer)

	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Println(string(body))
		return
	}

	//log.Println("The logging's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the status of this bucket's logging.
//	获得Bucket的日志设置。
/*
 *	Example:
 *	bls, err := c.GetBucketLogging(bucketName)
 */
func (c *Client) GetBucketLogging(bucketName string) (bls types.BucketLoggingStatus, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?logging"
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

	err = xml.Unmarshal(body, &bls)
	if err != nil {
		return
	}

	// log.Println("You have got the logging's setting of " + bucketName + ".")
	return
}
