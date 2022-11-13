package main

//go:generate go-bindata -fs -o=dll.go -pkg=dll -nocompress -nomemcopy ./Tasks.dll

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

func BytePtr(s []byte) uintptr {
	return uintptr(unsafe.Pointer(&s[0]))
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func Tasks(a []byte, b int) {
	Tasks_plan := syscall.MustLoadDLL("Tasks.dll")
	add := Tasks_plan.MustFindProc("qidong")
	ret, _, err := add.Call(BytePtr(a), IntPtr(b))
	if err != nil {
		if ret == 0 {
			fmt.Println("计划任务添加成功！")
		}
	}
}

func main() {
	sysType := runtime.GOOS
	if sysType == "windows" {
		var str = []byte("c:\\windows\\1.exe")
		Tasks(str, len(str))
	}

}
