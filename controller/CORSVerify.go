/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package controller

import (
	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/Unknwon/macaron"
)

//	跨域验证：用于浏览器的跨域保护。
/*
 * Access-Control-Allow-Origin: 请求允许的源地址
 * Access-Control-Allow-Methods: 请求允许的方法
 * Access-Control-Allow-Headers: 请求允许的HTTP请求头
 */
func CORSVerify(ctx *macaron.Context) {
	ctx.Resp.Header().Add("Access-Control-Allow-Origin", "*")
	ctx.Resp.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	ctx.Resp.Header().Add("Access-Control-Allow-Headers", "accept, content-type, authorization")
}
