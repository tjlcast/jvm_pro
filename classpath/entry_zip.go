package classpath

import "path/filepath"

/**
	表示ZIP或者JAR文件形式的类路径
 */

type ZipEntry struct {
	absPath		string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(classname string) ([]byte, Entry, error) {
	// todo
	return nil, nil, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}