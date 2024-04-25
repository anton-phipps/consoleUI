//go:build !windows
// +build !windows

package main

import "fmt"

func init() {
	fmt.Printf("%sH%sJ", ESC, ESC)
}
