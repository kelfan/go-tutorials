package filer

import (
	"../logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	logger.ErrNotNil(err, "filer.Current can't get current path")
	logger.Info("get path success:" + dir)
	return dir
}

/**
	relative path to absolute path
 */
func Relative2abs(rel string) string {
	abs, err := filepath.Abs(rel)
	logger.ErrNotNil(err, "filer.Realtive2abs: can't covert relative path - "+rel)
	return abs
}

func CFolder(path string)  {
	CheckFolder(path, 0700)
}

func Readfile(file string) string {
	data, err := ioutil.ReadFile(file)
	logger.ErrNotNil(err, "filer.ReadFile: can't read file - " + file)
	return string(data)
}


func CheckFolder(path string, permissionBits os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, permissionBits)
		logger.ErrNotNil(err, "can't create path: "+path)
		logger.Info("filer.CheckFolder: create file success - " + path)
	} else {
		logger.Info("filer.CheckFolder: file already exist - " + path)
	}
}

//+-----+---+--------------------------+
//| rwx | 7 | Read, write and execute  |
//| rw- | 6 | Read, write              |
//| r-x | 5 | Read, and execute        |
//| r-- | 4 | Read,                    |
//| -wx | 3 | Write and execute        |
//| -w- | 2 | Write                    |
//| --x | 1 | Execute                  |
//| --- | 0 | no permissions           |
//+------------------------------------+
//
//+------------+------+-------+
//| Permission | Octal| Field |
//+------------+------+-------+
//| rwx------  | 0700 | User  |
//| ---rwx---  | 0070 | Group |
//| ------rwx  | 0007 | Other |
//+------------+------+-------+
