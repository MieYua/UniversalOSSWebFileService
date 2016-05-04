/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package oss

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss"
	ctypes "UniversalOSSWebFileService/Godeps/_workspace/src/github.com/MieYua/Aliyun-OSS-Go-SDK/oss/types"
	"UniversalOSSWebFileService/model"
	"strings"
)

//	Default the settings of buckets(xxx-file or xxx-log) by defaults.
//	将xxx-file和xxx-log的设置默认化。
/*
 *	Example:
 *	err := oc.DefaultSetting(bucketName, bsc)
 */
func (oc *OSSClient) DefaultSetting(bucketName string, bsc model.BucketSettingConfig) (err error) {
	c := oc.Client
	aclSetting := ""
	if bsc.ACLSetting == "" {
		aclSetting = "private"
	} else {
		aclSetting = bsc.ACLSetting
	}

	loggingSetting := map[string]string{}
	loggingSetting = oss.SetBucketLogging(bsc.LoggingSetting.TargetBucket, bsc.LoggingSetting.TargetPrefix)

	websiteSetting := map[string]string{}
	websiteSetting = oss.SetBucketWebsite(bsc.WebsiteSetting.IndexPage, bsc.WebsiteSetting.ErrorPage)

	refererList := []string{}
	referers := []string{}
	length := len(bsc.RefererSetting.Referers)

	for i := 0; i < length; i++ {
		refererList = append(refererList, bsc.RefererSetting.Referers[i].Referer)
	}
	referers = oss.SetBucketReferer(refererList)

	rules := []ctypes.Rule{}
	length = len(bsc.LifecycleSetting.Rules)
	for i := 0; i < length; i++ {
		rules = oss.SetBucketLifecycle(rules, bsc.LifecycleSetting.Rules[i].Id, bsc.LifecycleSetting.Rules[i].Prefix, bsc.LifecycleSetting.Rules[i].Status, bsc.LifecycleSetting.Rules[i].Days)
	}

	corsRules := []ctypes.CORSRule{}
	length = len(bsc.CORSSetting.CORSRule)

	for i := 0; i < length; i++ {
		corsRules = oss.SetBucketCORS(corsRules, bsc.CORSSetting.CORSRule[i].AllowedOrigin, bsc.CORSSetting.CORSRule[i].AllowedMethod, bsc.CORSSetting.CORSRule[i].AllowedHeader, bsc.CORSSetting.CORSRule[i].ExposeHeader, bsc.CORSSetting.CORSRule[i].MaxAgeSeconds)
	}
	if strings.Contains(bucketName, "-file") {
		err = c.SetBucket(bucketName, aclSetting, loggingSetting, websiteSetting, nil, rules, corsRules)
	} else if strings.Contains(bucketName, "-log") {
		err = c.SetBucket(bucketName, aclSetting, nil, websiteSetting, referers, rules, corsRules)
	} else {
		err = c.SetBucket(bucketName, aclSetting, loggingSetting, websiteSetting, referers, rules, corsRules)
	}
	return
}
