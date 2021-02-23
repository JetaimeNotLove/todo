/**
 * @Author: Huyantian
 * @Date: 2021/2/23 下午7:28
 */

package filedao

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"todo/pkg/file"
)

const (
	dir      = "./.todolist"
	todolist = "todolist.json"
	userList = "userlist.json"
)

func Init() {
	if !file.Exists(dir) {
		if err := os.Mkdir(dir, os.ModePerm|os.ModeDir); err != nil {
			panic(err)
		}
	}
	for _, fp := range []string{todolist, userList} {
		initFile(fp)
	}
}

func initFile(fp string) {
	if !file.Exists(filepath.Join(dir, fp)) {
		f, err := os.Create(filepath.Join(dir, fp))
		if err != nil {
			panic(f)
		}
		f.Close()
	}

}

func read(fileName string, obj interface{}) error {
	fileName = filepath.Join(dir, fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, obj)
}

func write(fileName string, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(dir, fileName), data, os.ModePerm)
}
