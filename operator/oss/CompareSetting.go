/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package oss

import (
	"UniversalOSSWebFileService/model"
	"UniversalOSSWebFileService/util"
	"strings"
)

//	Compare the settings of buckets(xxx-file or xxx-log) with defaults.
//	和xxx-file及xxx-log默认设置比较。
/*
 *	Example:
 *	isSame, err := oc.CompareSetting(c, bucketName, bsc)
 */
func (oc *OSSClient) CompareSetting(bucketName string, bsc model.BucketSettingConfig) (isSame bool, err error) {
	c := oc.Client
	acl, lc, bls, wc, rc, lfc, corsc, err := c.GetBucketSetting(bucketName)
	if err != nil {
		return
	}
	isSame = true
	if acl.AccessControlList.Grant != bsc.ACLSetting {
		isSame = false
		return
	}
	if string(lc) != bsc.LocationSetting {
		isSame = false
		return
	}
	if bls.LoggingEnabled.TargetBucket != bsc.LoggingSetting.TargetBucket || bls.LoggingEnabled.TargetPrefix != bsc.LoggingSetting.TargetPrefix {
		isSame = false
		if strings.Contains(bucketName, "-log") {
			isSame = true
		} else {
			return
		}
	}
	if wc.IndexDocument.Suffix != bsc.WebsiteSetting.IndexPage || wc.ErrorDocument.Key != bsc.WebsiteSetting.ErrorPage {
		isSame = false
		return
	}
	if rc.AllowEmptyReferer != bsc.RefererSetting.AllowEmptyReferer {
		isSame = false
		if strings.Contains(bucketName, "-file") {
			isSame = true
		} else {
			return
		}
	}
	if len(rc.RefererList.Referer) != len(bsc.RefererSetting.Referers) {
		isSame = false
		if strings.Contains(bucketName, "-file") {
			isSame = true
		} else {
			return
		}
	} else {
		length := len(rc.RefererList.Referer)
		for i := 0; i < length; i++ {
			if rc.RefererList.Referer[i] != bsc.RefererSetting.Referers[i].Referer {
				isSame = false
			}
		}
		if strings.Contains(bucketName, "-file") {
			isSame = true
		}
	}
	if len(lfc.Rule) != len(bsc.LifecycleSetting.Rules) {
		isSame = false
		return
	} else {
		length := len(lfc.Rule)
		for i := 0; i < length; i++ {
			if lfc.Rule[i].Prefix != bsc.LifecycleSetting.Rules[i].Prefix {
				isSame = false
				return
			} else if lfc.Rule[i].Status != bsc.LifecycleSetting.Rules[i].Status {
				isSame = false
				return
			} else if lfc.Rule[i].Expiration.Days != bsc.LifecycleSetting.Rules[i].Days {
				isSame = false
				return
			}
		}
	}
	if len(corsc.CORSRule) != len(bsc.CORSSetting.CORSRule) {
		isSame = false
		return
	} else {
		length := len(corsc.CORSRule)
		for i := 0; i < length; i++ {
			if !util.CompareStrings(corsc.CORSRule[i].AllowedHeader, bsc.CORSSetting.CORSRule[i].AllowedHeader) {
				isSame = false
				return
			} else if !util.CompareStrings(corsc.CORSRule[i].AllowedMethod, bsc.CORSSetting.CORSRule[i].AllowedMethod) {
				isSame = false
				return
			} else if !util.CompareStrings(corsc.CORSRule[i].AllowedOrigin, bsc.CORSSetting.CORSRule[i].AllowedOrigin) {
				isSame = false
				return
			} else if !util.CompareStrings(corsc.CORSRule[i].ExposeHeader, bsc.CORSSetting.CORSRule[i].ExposeHeader) {
				isSame = false
				return
			} else if corsc.CORSRule[i].MaxAgeSeconds != bsc.CORSSetting.CORSRule[i].MaxAgeSeconds {
				isSame = false
				return
			}
		}
	}
	return
}
