package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	//多个路径
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	//当结尾是*
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	//当结尾是.jar、.JAR、.zip、.ZIP
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	//目录形式的路径
	return newDirEntry(path)
}
