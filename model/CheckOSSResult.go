/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 	检查OSS返回json数据
type CheckOSSResult struct {
	CheckBucketExist     string `json:"CheckBucketExist"`
	CreateBucketLogging  string `json:"CreateBucketLogging,omitempty"`
	CheckBucketSetting   string `json:"CheckBucketSetting"`
	ChangeSettingLogging string `json:"ChangeSettingLogging,omitempty"`
}
