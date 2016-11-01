package tools

import (
	"runtime"
	"os"
)

func IsOSX() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}



func IsKubernetes() bool{
	if os.Getenv("KUBERNETES") == "true"  {
		return true
	}
	return false
}

func IsWercker() bool{
	if os.Getenv("WERCKER") == "true"  {
		return true
	}
	return false
}


func IsDocker() bool{
	if _, err := os.Stat("/.dockerinit"); err == nil {
		return true
	}
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
