package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

/**
	使用Entry表示类路径项，分别有4种实现：
	DirEntry、ZipEntry、CompositeEntry、WildcardEntry
 */

type Entry interface {
	// 查找和加载class文件，className是class文件的相对路径
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/**
	newEntry根据参数创建不同类型的entry实例
 */
func newEntry(path string) Entry {
	// 目录
	if strings.Contains(path, pathListSeparator) {
		//
		return nil
	}
	// 所有
	if strings.HasSuffix(path, "*") {
		return nil
	}
	// .jar 或则 .zip
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		// 具体文件
	}
	return nil
}
