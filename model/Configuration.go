/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package model

// 配置类
type Configuration struct {
	ServerConfig   ServerConfig   `json:"server_config"`
	DatabaseConfig DatabaseConfig `json:"database_config"`
	OSSConfig      OSSConfig      `json:"oss_config"`
}

// 服务器配置类
type ServerConfig struct {
	Port int `json:"port"`
}

// 数据库配置类
type DatabaseConfig struct {
	MysqlAddress      string `json:"mysql_address"`
	MysqlDatabaseName string `json:"mysql_database_name"`
	MysqlUsername     string `json:"mysql_username"`
	MysqlPassword     string `json:"mysql_password"`
	MysqlTablePrefix  string `json:"mysql_table_prefix"`
}

// OSS配置类
type OSSConfig struct {
	Endpoint          string              `json:"endpoint"`
	AccessKeyId       string              `json:"access_key_id"`
	AccessKeySecret   string              `json:"access_key_secret"`
	OSSFileBucketName string              `json:"oss_file_bucket_name"`
	OSSLogBucketName  string              `json:"oss_log_bucket_name"`
	OSSFileConfig     BucketSettingConfig `json:"oss_file_config"`
	OSSLogConfig      BucketSettingConfig `json:"oss_log_config"`
}
