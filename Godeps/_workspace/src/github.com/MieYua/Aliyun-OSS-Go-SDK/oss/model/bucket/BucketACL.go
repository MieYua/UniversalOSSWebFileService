/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package bucket

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
)

// 	Change the setting of this bucket' acl.
//	修改Bucket的ACL权限。
/*
 *	Example:
 *	err := PutBucketACL(bucketName, (consts)ACL)
 */
func (c *Client) PutBucketACL(bucketName, acl string) (err error) {
	cc := c.CClient

	params := map[string]string{consts.OH_OSS_CANNED_ACL: acl}
	reqStr := "/" + bucketName
	resp, err := cc.DoRequest("PUT", reqStr, reqStr, params, nil)
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

	//log.Println("The ACL's setting of " + bucketName + " has been changed.")
	return
}

// 	Get the setting of this bucket' acl.
//	获得Bucket的ACL权限。
/*
 *	Example:
 *	acl, err := GetBucketACL(bucketName)
 */
func (c *Client) GetBucketACL(bucketName string) (acl types.AccessControlPolicy, err error) {
	cc := c.CClient

	reqStr := "/" + bucketName + "?acl"
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

	err = xml.Unmarshal(body, &acl)
	if err != nil {
		return
	}

	//log.Println("You have got the ACL's setting of " + bucketName + ".")
	return
}
