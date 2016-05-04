/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 	请求的Body中json数据。
type RequestJSON struct {
	FileName    string `json:"file_name,omitempty"`
	Description string `json:"description,omitempty"`
}

// 返回的json中的状态部分
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 	返回的json数据
type ResponseJson struct {
	Meta Meta          `json:"meta"`
	Data []interface{} `json:"data,omitempty"`
}

// 	返回单个数据的json数据
type SingleResponseJson struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
