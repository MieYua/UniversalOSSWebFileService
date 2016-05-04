/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package file

import (
	"errors"
)

//	Delete the file(s).
//	完全删除文件操作。
/*
 *	Example:
 *	dfr, err := fp.DeleteFile()
 */
func (fp *FilePack) DeleteFile() (err error) {
	c := fp.OSSClient.Client
	bucketName := fp.BucketName
	if len(fp.FilePaths) > 0 {
		err = c.DeleteMultipleObject(bucketName, fp.FilePaths)
		if err != nil {
			return
		}
	} else {
		return errors.New("没有要删除的文件")
	}
	return
}
