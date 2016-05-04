CREATE TABLE `hz_file` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文件id，自增主键',
  `file_name` varchar(32) NOT NULL COMMENT '文件名',
  `file_type_id` int(11) unsigned NOT NULL COMMENT '文件类型id',
  `file_path` varchar(255) NOT NULL COMMENT '文件存储位置',
  `uploader_id` char(36) NOT NULL COMMENT '上传者id',
  `upload_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述，备注',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件表';

CREATE TABLE `hz_file_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件类型id，自增主键',
  `file_type_name` varchar(32) NOT NULL COMMENT '文件类型名称',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述，备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件类型表';