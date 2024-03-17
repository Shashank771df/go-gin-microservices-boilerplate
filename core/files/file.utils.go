package files

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type UtilsFile struct {
}

func (obj UtilsFile) GetDirectories(path string) ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, fileinfo := range files {
		if fileinfo.IsDir() {
			ret = append(ret, fileinfo)
		}
	}

	return ret, nil
}

func (obj UtilsFile) GetFiles(path string) ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, fileinfo := range files {
		if !fileinfo.IsDir() {
			ret = append(ret, fileinfo)
		}
	}

	return ret, nil
}

func (obj UtilsFile) GetDirectoriesSorted(path string) ([]os.FileInfo, error) {
	fi, err := obj.GetDirectories(path)

	if err != nil {
		return fi, err
	}

	sort.Slice(fi, func(i, j int) bool {
		return fi[i].ModTime().After(fi[j].ModTime())
	})

	return fi, nil
}

func (obj UtilsFile) GetFilesSorted(path string) ([]os.FileInfo, error) {
	fi, err := obj.GetFiles(path)

	if err != nil {
		return fi, err
	}

	sort.Slice(fi, func(i, j int) bool {
		return fi[i].ModTime().After(fi[j].ModTime())
	})

	return fi, nil
}

func (obj UtilsFile) ExtensionFile(name string) string {
	if !strings.Contains(name, ".") {
		return ""
	}

	index := strings.LastIndex(name, ".")

	return name[index:]
}

func (obj UtilsFile) CreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); err == nil {
		return errors.New("directory already exists")
	}

	return os.Mkdir(dirPath, os.ModePerm)
}

func (obj UtilsFile) CreateDirIfNotExists(dirname string) error {
	_, err := os.Stat(dirname)

	if err != nil {
		return os.Mkdir(dirname, os.ModePerm)
	}

	return nil
}

func (obj UtilsFile) CreateFile(fileName string, force bool) error {
	if _, err := os.Stat(fileName); err == nil {
		if !force {
			return errors.New("file already exists")
		}
	}

	_, err := os.Create(fileName)

	return err
}

func (obj UtilsFile) CreateFileIfNotExists(filename string) error {
	_, err := os.Stat(filename)

	if err != nil {
		f, err := os.OpenFile(filename, os.O_CREATE, 0666)

		if err == nil {
			f.Close()
		}

		return err
	}

	return nil
}

func (obj UtilsFile) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (obj UtilsFile) Check(path string) error {
	_, err := os.Stat(path)

	return err
}

func (obj UtilsFile) ReadFile(path string) []byte {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil
	}

	return data
}

func (obj UtilsFile) WriteFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}

	return err
}

func (obj UtilsFile) WriteString(path string, data string) error {
	return obj.WriteFile(path, []byte(data))
}

func (obj UtilsFile) WriteStringLn(path string, data string) error {
	return obj.WriteFile(path, []byte(data+"\n"))
}

func (obj UtilsFile) RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func (obj UtilsFile) ReadMarshal(path string, data interface{}) error {
	var buffer []byte

	if buffer = obj.ReadFile(path); buffer == nil {
		return errors.New("json file is empty")
	}

	if err := json.Unmarshal(buffer, &data); err != nil {
		return err
	}

	return nil
}

func (obj UtilsFile) WriteMarshal(path string, data interface{}) error {
	buffer, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = obj.WriteFile(path, buffer)

	if err != nil {
		return err
	}

	return nil
}

func (obj UtilsFile) FileInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}
