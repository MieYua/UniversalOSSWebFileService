/*
 * Copyright (C) Mie Yua <mieyua@aliyun.com>, 2015.
 * All rights reserved.
 */

package object

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/consts"
	"errors"
	"io/ioutil"
	"log"
)

// 	Choose the object by options.
//	选择筛选Object。
/*
 *	Example:
 *	err := c.OptionObject("bucketName/test.txt", "", "")
 *	fmt.Println(err)
 *
 *	Warning:
 *	If the bucket's cors is not available or its cors hasn't been set up,
 *	response will show 403 ERROR.
 */
func (c *Client) OptionObject(opath, accessControlRequestMethod, accessControlRequestHeader, origin string) (err error) {
	cc := c.CClient

	reqStr := "/" + opath

	params := map[string]string{consts.HH_CONTENT_TYPE: "application/xml"}
	if accessControlRequestMethod != "" {
		params[consts.OH_ACCESS_CONTROL_REQUEST_METHOD] = accessControlRequestMethod
	} else {
		params[consts.OH_ACCESS_CONTROL_REQUEST_METHOD] = "PUT"
	}
	if accessControlRequestHeader != "" {
		params[consts.OH_ACCESS_CONTROL_REQUEST_HEADER] = accessControlRequestHeader
	} else {
		params[consts.OH_ACCESS_CONTROL_REQUEST_HEADER] = "x-oss-test"
	}
	if origin != "" {
		params[consts.OH_ORIGIN] = origin
	} else {
		params[consts.OH_ORIGIN] = "http://www.example.com"
	}
	resp, err := cc.DoRequest("OPTIONS", reqStr, reqStr, params, nil)
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

	//log.Println("CORS's request has passed by options.")
	return
}
