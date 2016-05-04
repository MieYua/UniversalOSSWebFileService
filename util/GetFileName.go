/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package util

import (
	"strings"
)

//	GetFileName
//	Get the file name without file type.
//	文件名提取。
/*
 *	Example:
 *	fileName := GetFileName(fullName)
 */
func GetFileName(fullName string) string {
	fileName := ""
	isContains := strings.Contains(fullName, ".")
	if isContains == false {
		return fullName
	} else {
		typeDot := strings.LastIndex(fullName, ".")
		fb := []byte(fullName)
		fileName = string(fb[0:typeDot])
	}
	return fileName
}
