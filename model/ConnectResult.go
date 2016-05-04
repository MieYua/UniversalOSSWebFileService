/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 	连接服务返回json数据
type ConnectResult struct {
	ConnectResult  string   `json:"ConnectResult"`
	ConnectLogging []string `json:"ConnectLogging"`
	Owner          Owner    `json:"Owner,omitempty"`
}
