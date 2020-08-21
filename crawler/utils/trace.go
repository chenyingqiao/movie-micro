package utils

import (
	"os"
	"path/filepath"
	"runtime/trace"
	"strings"
)

//StartTrace 开始跟踪
func StartTrace() {
	f, err := os.Create(getCurrentDirectory() + "/trace.out")
	if err != nil {
		return
	}
	trace.Start(f)
}

//StopStrace 结束跟踪
func StopStrace() {
	trace.Stop()
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
