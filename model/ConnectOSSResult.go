/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 	连接OSS返回json数据。
type ConnectOSSResult struct {
	ConnectCondition string `json:"ConnectOSSResult"`
	Owner            Owner  `json:"Owner,omitempty"`
}
