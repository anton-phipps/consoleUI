//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Unix Version")
	time.Sleep(time.Second * 5)
}
