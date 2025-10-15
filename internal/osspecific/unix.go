//go:build darwin || linux || freebsd
// +build darwin linux freebsd

package osspecific

import (
	"fmt"
	"os"
	"os/exec"
)

func GetConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		return 0, 0, err
	}
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}
