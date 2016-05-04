/*
 * Copyright (C) Mie Yua <mieyua@126.com>, 2016.
 * All rights reserved.
 */

package util

//	CompareStrings
//	Compare two strings.
//	字符串集合比较。
/*
 *	Example:
 *	isSameStrings := CompareStrings(strings1, strings2)
 */
func CompareStrings(strings1, strings2 []string) (isSameStrings bool) {
	length1 := len(strings1)
	length2 := len(strings2)
	if length1 != length2 {
		if length1 == 0 && length2 == 1 && strings2[0] == "" || length1 == 1 && length2 == 0 && strings1[0] == "" {
			isSameStrings = true
			return
		} else {
			isSameStrings = false
			return
		}
	} else {
		for i := 0; i < length1; i++ {
			if strings1[i] != strings2[i] {
				isSameStrings = false
				return
			}
		}
		isSameStrings = true
		return
	}
}
