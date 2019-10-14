package io

import (
	"os"
	"path/filepath"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_ufs(canonicalize0, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;")
	_ufs(getBooleanAttributes0, "getBooleanAttributes0", "(Ljava/io/File;)I")
	_ufs(getLastModifiedTime, "getLastModifiedTime", "(Ljava/io/File;)J")
}

func _ufs(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/UnixFileSystem", name, desc, method)
}

// private native String canonicalize0(String path) throws IOException;
// (Ljava/lang/String;)Ljava/lang/String;
func canonicalize0(frame *rtda.Frame) {
	pathStr := frame.GetRefVar(1)

	// todo
	path := heap.JSToGoStr(pathStr)
	path2 := filepath.Clean(path)
	if path2 != path {
		pathStr = heap.JSFromGoStr(path2)
	}

	frame.PushRef(pathStr)
}

// public native int getBooleanAttributes0(File f);
// (Ljava/io/File;)I
func getBooleanAttributes0(frame *rtda.Frame) {
	fileObj := frame.GetRefVar(1)
	path := _getPath(fileObj)

	// todo
	attributes0 := 0
	if _exists(path) {
		attributes0 |= 0x01
	}
	if _isDir(path) {
		attributes0 |= 0x04
	}

	frame.PushInt(int32(attributes0))
}

func _getPath(fileObj *heap.Object) string {
	pathStr := fileObj.GetFieldValue("path", "Ljava/lang/String;").Ref
	return heap.JSToGoStr(pathStr)
}

// http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
// exists returns whether the given file or directory exists or not
func _exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// http://stackoverflow.com/questions/8824571/golang-determining-whether-file-points-to-file-or-directory
func _isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir()
	}
	return false
}

// public native long getLastModifiedTime(File f);
//
func getLastModifiedTime(frame *rtda.Frame) {
	fileObj := frame.GetRefVar(1)

	path := _getPath(fileObj)
	t := _getLastModifiedTime(path)

	frame.PushLong(t)
}

func _getLastModifiedTime(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.ModTime().UnixNano() / 1000
	}
	return 0
}