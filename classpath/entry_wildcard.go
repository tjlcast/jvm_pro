package classpath

import (
	"path/filepath"
	"os"
	"strings"
)

/**
	WildcardEntry本质上也是CompositeEntry

	在 walkFn 中，根据后缀名选出JAE文件，并且返回SkipDir跳过子目录

 */

 func newWildcardEntry(path string) CompositeEntry {
 	baseDir := path[:len(path)-1] // remove "*"
 	compositeEntry := []Entry{}

 	walkFn := func(path string, info os.FileInfo, err error) error {
		 if err != nil {
		 	return err
		 }

		 if info.IsDir() && path != baseDir {
		 	return filepath.SkipDir
		 }

		 if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		 	jarEntry := newZipEntry(path)
		 	compositeEntry = append(compositeEntry, jarEntry)
		 }

		 return nil
	 }
	filepath.Walk(baseDir, walkFn)

 	return compositeEntry
 }
