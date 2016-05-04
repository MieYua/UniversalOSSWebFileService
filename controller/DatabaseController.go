/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package controller

import (
	"UniversalOSSWebFileService/model"
	"strings"
)

//	CheckFileExist: Check the file's exist.
//	检查文件是否存在。
/*
 *	Example:
 *	isExist, err := checkFileExist(fileId)
 */
func checkFileExist(fileId int) (bool, error) {
	file := model.File{}
	isExist, err := X.Id(fileId).Get(&file)
	return isExist, err
}

//	GetFile: Get a file by fileId.
//	获得单个文件详情。
/*
 *	Example:
 *	file, err := getSingleFile(fileId)
 */
func getSingleFile(fileId int) (file model.File, err error) {
	_, err = X.Id(fileId).Get(&file)
	if err != nil {
		return
	}
	return
}

//	GetFiles: Get all files in database.
//	获得所有文件详情。
/*
 *	Example:
 *	files, err := getAllFiles()
 */
func getAllFiles() ([]model.File, error) {
	files := make([]model.File, 0)
	err := X.Find(&files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

//	GetFileAddress: Get the address of the file.
//	获得文件下载地址。
/*
 *	Example:
 *	address, err := getFileAddress(fileId)
 */
func getFileAddress(fileId int) (string, error) {
	file := model.File{}
	_, err := X.Id(fileId).Get(&file)
	if err != nil {
		return "", err
	}
	address := "http://" + BUCKET_NAME + "." + OSS_LOCATION + ".aliyuncs.com/" + file.FilePath
	return address, nil
}

//	InsertFile: Insert a new file row.
//	新增文件记录。
/*
 *	Example:
 *	err := insertFile(fileName, filePath, fileType, uploaderId, description)
 */
func insertFile(fileName, filePath, fileType, uploaderId, description string) (err error) {
	file := model.File{}
	file.FileName = fileName
	file.FilePath = filePath
	fileTypeId, err := checkFileType(fileType)
	if err != nil {
		return err
	}
	file.FileTypeId = fileTypeId
	file.UploaderId = uploaderId
	file.Description = description
	_, err = X.Omit("id", "upload_time").Insert(&file)
	if err != nil {
		return err
	}
	return
}

//	UpdateFile: Update the file by fileId.
//	修改文件记录。
/*
 *	Example:
 *	err := updateFile(fileId, fileName, description)
 */
func updateFile(fileId int, fileName, description string) (err error) {
	file := model.File{}
	_, err = X.Id(fileId).Get(&file)
	if err != nil {
		return
	}
	file.FileName = fileName
	file.Description = description
	updateColumns := make([]string, 0)
	if fileName != "" {
		updateColumns = append(updateColumns, "file_name")
	}
	if description != "" {
		updateColumns = append(updateColumns, "description")
	}
	updateNumber := len(updateColumns)
	switch updateNumber {
	case 1:
		_, err = X.Id(fileId).Cols(updateColumns[0]).Update(&file)
	case 2:
		_, err = X.Id(fileId).Cols(updateColumns[0], updateColumns[1]).Update(&file)
	default:
	}
	if err != nil {
		return err
	}
	return
}

//	DeleteFile: Delete the file by fileId.
//	删除文件记录。
/*
 *	Example:
 *	err := deleteFile(fileId)
 */
func deleteFile(fileId int) (err error) {
	file := model.File{}
	_, err = X.Id(fileId).Get(&file)
	if err != nil {
		return
	}
	return
}

// CheckFileType: Check the file's type's exist.
// 检查文件类型是否存在（存在返回文件类型id，不存在新增并返回）。
/*
 *	Example:
 *	fileTypeId, err := checkFileType(fileType)
 */
func checkFileType(fileTypeName string) (fileTypeId int, err error) {
	fileTypeName = strings.ToLower(fileTypeName)
	fileType := model.FileType{}
	isExist, err := X.Where("file_type_name=?", fileTypeName).Get(&fileType)
	if err != nil {
		return 0, err
	}
	if isExist == true {
		return fileType.Id, nil
	} else {
		fileTypeNew := model.FileType{}
		fileTypeNew.FileTypeName = fileTypeName
		fileTypeNew.Description = fileTypeName + "文件"
		_, err = X.Omit("id").Insert(&fileTypeNew)
		if err != nil {
			return 0, err
		}
		_, err = X.Get(&fileTypeNew)
		if err != nil {
			return 0, err
		}
		return fileTypeNew.Id, nil
	}
	return
}
