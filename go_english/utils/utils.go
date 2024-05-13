package utils

import (
	"io"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func ReadIcon(filepath string) ([]byte, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func GetNewTime() string {
	time := time.Now().Format("2006-01-02 15:04:05")
	return time
}
func IntPtr(n int) uintptr {
	return uintptr(n)
}
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// ShowMessage windows下的另一种DLL方法调用
func ShowMessage(tittle, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(tittle), IntPtr(0))
	defer syscall.FreeLibrary(user32dll)
}
