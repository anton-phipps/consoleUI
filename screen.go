package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getUnixConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		return 0, 0, err
	}
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}

func getWindowsConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("mode", "con", "/status")
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Columns") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				colsStr := parts[1]
				fmt.Sscanf(colsStr, "%d", &cols)
			}
		} else if strings.Contains(line, "Lines") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				rowsStr := parts[1]
				fmt.Sscanf(rowsStr, "%d", &rows)
			}
		}
	}

	return rows, cols, nil
}

func getConsoleSize() (rows, cols int, err error) {
	if runtime.GOOS == "windows" {
		rows, cols, err = getWindowsConsoleSize()
	} else {
		rows, cols, err = getUnixConsoleSize()
	}
	return rows, cols, err
}
