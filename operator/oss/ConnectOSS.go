/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package oss

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss"
	"UniversalOSSWebFileService/model"
	"encoding/json"
)

//	Link to OSS Client.
//	连接OSS客户端。
/*
 *	Example:
 *	j, err := oc.ConnectOSS()
 */
func (oc *OSSClient) ConnectOSS() (j string, err error) {
	c := oss.InitiateClient(oc.EndPoint, oc.AccessKeyId, oc.AccessKeySecret)
	lambr, err := c.GetServiceInfo()
	lor := model.ConnectOSSResult{}
	if err != nil {
		lor.ConnectCondition = "Error: " + err.Error() + "."
		return
	} else {
		lor.ConnectCondition = "Link to OSS Successfully."
		lor.Owner.ID = lambr.Owner.ID
		lor.Owner.DisplayName = lambr.Owner.DisplayName
		oc.Client = c
	}
	b, _ := json.Marshal(lor)
	j = string(b)
	return
}
