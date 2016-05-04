/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// OSS Bucket设置类
type BucketSettingConfig struct {
	ACLSetting       string        `json:"ACLSetting"`
	LocationSetting  string        `json:"LocationSetting"`
	LoggingSetting   LoggingJSON   `json:"LoggingSetting,omitempty"`
	WebsiteSetting   WebsiteJSON   `json:"WebsiteSetting,omitempty"`
	RefererSetting   RefererJSON   `json:"RefererSetting,omitempty"`
	LifecycleSetting LifecycleJSON `json:"LifecycleSetting,omitempty"`
	CORSSetting      CORSJSON      `json:"CORSSetting,omitempty"`
}

// 	日志设置json数据
type LoggingJSON struct {
	TargetBucket string `json:"TargetBucket"`
	TargetPrefix string `json:"TargetPrefix"`
}

// 	默认首页及错误页设置json数据
type WebsiteJSON struct {
	IndexPage string `json:"IndexPage"`
	ErrorPage string `json:"ErrorPage"`
}

// 	白名单地址的json数据
type Referer struct {
	Referer string `json:"Referer"`
}

// 	白名单设置json数据
type RefererJSON struct {
	AllowEmptyReferer string    `json:"AllowEmptyReferer"`
	Referers          []Referer `json:"Referers"`
}

// 	生命周期规则json数据
type Rule struct {
	Id     string `json:"Id"`
	Prefix string `json:"Prefix"`
	Status string `json:"Status"`
	Days   int    `json:"Days"`
}

// 	生命周期设置json数据
type LifecycleJSON struct {
	Rules []Rule `json:"Rules"`
}

// 	跨域请求的设置规则
type CORSRule struct {
	AllowedOrigin []string `json:"AllowedOrigin"`
	AllowedMethod []string `json:"AllowedMethod"`
	AllowedHeader []string `json:"AllowedHeader"`
	ExposeHeader  []string `json:"ExposeHeader"`
	MaxAgeSeconds int      `json:"MaxAgeSeconds"`
}

// 	跨域请求设置json数据
type CORSJSON struct {
	CORSRule []CORSRule `json:"CORSRule"`
}
