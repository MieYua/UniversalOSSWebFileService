/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

import (
	"time"
)

// 文件数据表
type File struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	FileName    string    `json:"file_name" xorm:"not null VARCHAR(32)"`
	FileTypeId  int       `json:"file_type_id" xorm:"not null index INT(11)"`
	FilePath    string    `json:"file_path" xorm:"not null VARCHAR(255)"`
	UploaderId  string    `json:"uploader_id" xorm:"not null index CHAR(36)"`
	UploadTime  time.Time `json:"upload_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Description string    `json:"description" xorm:"not null default '' VARCHAR(50)"`
}
