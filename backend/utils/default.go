package utils

import (
	"strconv"
	"sync"
	"time"
)

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var look = sync.Mutex{}

func GetTimestamp() string {
	look.Lock()
	defer look.Unlock()
	return time.Now().Format("20060102150405.000")
}
