/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 文件类型表
type FileType struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FileTypeName string `json:"file_type_name" xorm:"not null VARCHAR(32)"`
	Description  string `json:"description" xorm:"not null default '' VARCHAR(50)"`
}
