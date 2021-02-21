/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午5:18
 */

package file

import "os"

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
