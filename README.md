# This is a Universal Web File Service for Aliyun OSS by MieYua's SDK #
---

## Talk about the service at first ##
* <strong>The service is just a simple universal version, so there will not be any updates.</strong>
* <strong>此服务为简单通用版本，不会更新功能，想要其他内容请自行增减。</strong>
* <strong>Change the dependencies' routes to local path.</strong>
* <strong>使用godep修改依赖包，无需自行下载。</strong>
* <strong>Please read the README.md at first.</strong>
* <strong>请务必先看完文档。</strong>

---

## File Service Version ##
* <strong>Ver. 1.0 (Released on 160503)</strong>

---

## Golang Version ##
* <strong>Golang 1.5 or Higher</strong>

---

## SDK Version ##
* [SDK Link](https://github.com/MieYua/Aliyun-OSS-Go-SDK)
* <strong>Ver. 4.3 (Released on 150911)</strong>

---


## Contents ##
[1 配置文件](#configuration)  
[1.1 数据库配置文件](#config-sql)  
[1.2 全部配置文件](#config-all)  
[2 接口](#file-service-examples)  
[2.1 获取所有文件（或单个文件）信息](#get-file)  
[2.2 获取单个文件下载地址](#get-file-address)  
[2.3 简单表单上传](#post-file)  
[2.4 更新文件信息](#put-file)  
[2.5 删除文件](#delete-file)  
[3 简单文件服务历史版本](#universal-web-file-service-history)  
[4 使用包](#some-imports)

---

## Configuration ##
[Top](#contents)


### Config SQL ###
[Top](#contents)

* `xx_`一定要和[下面](#config-all)设置的前缀`prefix`一致

		CREATE TABLE `xx_file` (
		  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文件id，自增主键',
		  `file_name` varchar(32) NOT NULL COMMENT '文件名',
		  `file_type_id` int(11) unsigned NOT NULL COMMENT '文件类型id',
		  `file_path` varchar(255) NOT NULL COMMENT '文件存储位置',
		  `uploader_id` char(36) NOT NULL COMMENT '上传者id',
		  `upload_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
		  `description` varchar(20) NOT NULL DEFAULT '' COMMENT '描述，备注',
		   PRIMARY KEY (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件表';
		
		CREATE TABLE `xx_file_type` (
		  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件类型id，自增主键',
		  `file_type_name` varchar(32) NOT NULL COMMENT '文件类型名称',
		  `description` varchar(50) NOT NULL DEFAULT '' COMMENT '描述，备注',
		  PRIMARY KEY (`id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件类型表';

### Config All ###
[Top](#contents)

* OSS文件Bucket：`xx-file`；
* OSS日志Bucket：`xx-log`；
* 由于需要比较配置文件与OSS设置，先设置好`xx-file`和`xx-log`的website设置、跨域(CORS)设置和生命周期(lifecycle)设置。
* 注意不要忘了`oss_file_config`-->`LoggingSetting`的`"TargetBucket":"xx-log"`这个设置。
* 数据库(Database)设置要和上面的sql设置的`xx_`相同。

		{	
			/**
			 * -----------------------------------------
			 * 服务器配置
			 * SERVER CONFIGURATION
			 * -----------------------------------------
			 * port: 端口号（例如4000）
			 */
			"server_config":{
				"port": 4000,		// 端口号(Port)：例如4000
			},
			/**
			 * -----------------------------------------
			 * 数据库配置(MySQL)
			 * DATABASE CONFIGURATION(MySQL)
			 * -----------------------------------------
			 * mysql_address: 			数据库地址
			 * ------->  例如tcp(localhost:3306)或tcp(xxx.xxx.xxx.xxx:3306)等
			 * mysql_database_name: 	数据库名称
			 * mysql_username: 			数据库用户名
			 * mysql_password: 			数据库密码
			 * mysql_table_prefix: 		数据库内表前缀（请带下划线，如"test_"）
			 */
			"database_config":{
				"mysql_address":"tcp(localhost:3306)",
				"mysql_database_name": "xx_test",
				"mysql_username": "root",
				"mysql_password": "******",
				"mysql_table_prefix": "xx_"
			},
			/**
			 * -----------------------------------------
			 * OSS配置
			 * Aliyun OSS CONFIGURATION
			 * -----------------------------------------
			 * endpoint: bucket所在节点（不能自己乱输）
			 * ------->  北京（外）：BEIJING
			 * ------->  杭州（外）：HANGZHOU
			 * ------->  香港（外）：HONGKONG
			 * ------->  青岛（外）：QINGDAO
			 * ------->  深圳（外）：SHENZHEN
			 * ------->  美国硅谷（外）：US_WEST1
			 * ------->  北京（内）：BEIJING_IN
			 * ------->  杭州（内）：HANGZHOU_IN
			 * ------->  香港（内）：HONGKONG_IN
			 * ------->  青岛（内）：QINGDAO_IN
			 * ------->  深圳（内）：SHENZHEN_IN
			 * ------->  美国硅谷（内）：US_WEST1_IN
			 * access_key_id: 			accessKeyId
			 * access_key_secret: 		accessKeySecret
			 * oss_file_bucket_name: 	OSS存放文件的bucket名称（需自己新建，先设置好website设置、跨域(CORS)设置和生命周期(lifecycle)设置）
			 * oss_log_bucket_name: 	OSS存放日志的bucket名称（需自己新建，先设置好website设置、跨域(CORS)设置和生命周期(lifecycle)设置）
			 * oss_file_config: 		OSS存放文件的bucket规则
			 * oss_log_config: 			OSS存放日志的bucket规则
			 */
			"oss_config":{
				"endpoint":"HANGZHOU",
				"access_key_id":"******",
				"access_key_secret":"******",
				"oss_file_bucket_name":"xx-file",
				"oss_log_bucket_name":"xx-log",
				/**
				 * -----------------------------------------
				 * OSS文件Bucket配置
				 * OSS FILE BUCKET CONFIGURATION
				 * -----------------------------------------
				 * ACLSetting: 读写权限
				 * ------->  私有：private
				 * ------->  公共读：public-read
				 * ------->  公共读写：public-read-write
				 * LocationSetting: bucket所在位置（不能乱填）
				 * ------->  北京：oss-cn-beijing
				 * ------->  杭州：oss-cn-hangzhou
				 * ------->  香港：oss-cn-hongkong
				 * ------->  青岛：oss-cn-qindao
				 * ------->  深圳：oss-cn-shenzhen
				 * ------->  美国硅谷：oss-us-west1
				 * WebsiteSetting: 默认首页错误页配置
				 * LoggingSetting: 指向日志文档的bucket配置（通常就是上面的log_bucket）
				 * LifecycleSetting: 生命周期规则设置（默认例子一年后删除）
				 * CORSSetting: 跨域访问设置（默认例子）
				 */
				"oss_file_config":{
					"ACLSetting":"public-read-write",
					"LocationSetting":"oss-cn-hangzhou",
					"WebsiteSetting":{
						"IndexPage":"index.html",
						"ErrorPage":"error.html"
					},
					"LoggingSetting":{
						"TargetBucket":"xx-log",
						"TargetPrefix":""
					},
					"LifecycleSetting":{
						"Rules":[
							{
								"Id":"Delete after one year",
								"Prefix":"",
								"Status":"Enabled",
								"Days":365
							}
						]
					},
					"CORSSetting":{
						"CORSRule":[
							{
								"AllowedMethod":[
									"PUT",
									"GET",
									"POST",
									"DELETE"
								],
								"AllowedOrigin":[
									"*"
								],
								"AllowedHeader":[
									"*"
								],
								"ExposeHeader":[
									""
								],
								"MaxAgeSeconds":3600
							}
						]
					}
				},
		
				/**
				 * -----------------------------------------
				 * OSS日志Bucket配置
				 * OSS LOG BUCKET CONFIGURATION
				 * -----------------------------------------
				 * ACLSetting: 读写权限（保持私有）
				 * ------->  私有：private
				 * LocationSetting: bucket所在位置（不能乱填）
				 * ------->  北京：oss-cn-beijing
				 * ------->  杭州：oss-cn-hangzhou
				 * ------->  香港：oss-cn-hongkong
				 * ------->  青岛：oss-cn-qindao
				 * ------->  深圳：oss-cn-shenzhen
				 * ------->  美国硅谷：oss-us-west1
				 * WebsiteSetting: 默认首页错误页配置
				 * RefererSetting: 白名单配置（默认例子）
				 * LoggingSetting: 指向日志文档的bucket配置（通常就是上面的log_bucket）
				 * LifecycleSetting: 生命周期规则设置（默认例子一年后删除）
				 * CORSSetting: 跨域访问设置（默认例子）
				 */
				"oss_log_config":{
					"ACLSetting":"private",
					"LocationSetting":"oss-cn-hangzhou",
					"WebsiteSetting":{
						"IndexPage":"index.html",
						"ErrorPage":"error.html"
					},
					"RefererSetting":{
						"AllowEmptyReferer":"false",
						"Referers":[
							{
								"Referer":"http://www.aliyun.com"
							},
							{
								"Referer":"https://www.aliyun.com"
							}
						]
					},
					"LifecycleSetting":{
						"Rules":[
							{
								"Id":"Delete after one month",
								"Prefix":"",
								"Status":"Enabled",
								"Days":30
							}
						]
					},
					"CORSSetting":{
						"CORSRule":[
							{
								"AllowedMethod":[
									"PUT",
									"GET",
									"POST",
									"DELETE"
								],
								"AllowedOrigin":[
									"*"
								],
								"AllowedHeader":[
									""
								],
								"ExposeHeader":[
									""
								],
								"MaxAgeSeconds":3600
							}
						]
					}
				}
			}
		}

---

## File Service Examples ##
[Top](#contents)

### Get File ###
[Top](#contents)

*	获得所有文件信息

		GET http://URL:port/v1/file  

*	获得单个文件信息（fileId int）

		GET http://URL:port/v1/file/fileId 

### Get File Address ###
[Top](#contents)

*	获得单个文件下载地址（fileId int）
 
		GET http://URL:port/v1/fileAddress/fileId 

### Post File ###
[Top](#contents)

*	文件表单上传  

		POST http://URL:port/v1/file multipart/form-data;boundary=testboundary 
		--testboundary
		Content-Disposition: form-data; name="UploaderId"
		
		XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX 
		--testboundary
		Content-Disposition: form-data; name="Description"
	
		XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  
		--testboundary
		Content-Disposition: form-data; name="file"; filename="test.txt"
		Content-Type: application/octet-stream
		
		...
		--testboundary--

### Put File ###
[Top](#contents)

*	更新文件信息（fileId int）

		PUT http://URL:port/v1/file/fileId
		Body(Json)
		{
			"file_name":"new file name",
			"description":"new description"
		}

### Delete File ###
[Top](#contents)

*	删除文件（fileId int，不可恢复）

		DELETE http://URL:port/v1/file/fileId

---

## Universal Web File Service History ##
[Top](#contents)

### Ver. 1.0 (20160503) ###
* 原始版本（不更新）：
	* 配置数据库地址，自动生成文件表；
	* 自定Bucket，端口；
	* 简单表单上传；
	* 查询文件信息；
	* 修改文件信息；
	* 删除文件。

---


## Some Imports ##
[Top](#contents)

* 不需要下载包。
* `go get github.com/MieYua/Aliyun-OSS-Go-SDK/oss`
* `go get github.com/go-sql-driver/mysql`
* `go get github.com/go-xorm/core`
* `go get github.com/go-xorm/xorm`
* `go get github.com/Unknwon/macaron`

---
