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

// 	Change the website defaults of this bucket.
//	修改Bucket的默认首页以及错误页。
/*
 *	Example:
 *	err := PutBucketWebsite(bucketName, indexDocument, errorDocument)
 *	If the indexDocument is null, its default is "index.html".
 *	If the errorDocument is null, its default is "error.html".
 *	And the type of documents is only accepted with ".html"
 */
func (c *Client) PutBucketWebsite(bucketName, indexDocument, errorDocument string) (err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?website"

	wxml := types.WebsiteXML{}

	if indexDocument == "" {
		wxml.IndexDocument.Suffix = "index.html"
	} else {
		wxml.IndexDocument.Suffix = indexDocument
	}

	if errorDocument == "" {
		wxml.ErrorDocument.Key = "error.html"
	} else {
		wxml.ErrorDocument.Key = errorDocument
	}

	bs, err := xml.Marshal(wxml)
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

	//log.Println("The website's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the websites' default of this bucket.
//	获得Bucket的默认首页以及错误页。
/*
 *	Example:
 *	wc, err := c.GetBucketWebsite(bucketName)
 */
func (c *Client) GetBucketWebsite(bucketName string) (wc types.WebsiteConfiguration, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?website"
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

	err = xml.Unmarshal(body, &wc)
	if err != nil {
		return
	}

	// log.Println("You have got the website's setting of " + bucketName + ".")
	return
}
