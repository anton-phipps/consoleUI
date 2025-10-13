/*
* This is to accept keypresses and other user interactions
 */
package interactions

import (
	"os"

	"golang.org/x/term"
)

func ReadKey() (string, error) {
	// Save the current state of the terminal
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// big enough for escape sequences
	buf := make([]byte, 3)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return "", err
	}

	// Decode common keys
	switch {
	case n == 1 && buf[0] == 3: // Ctrl+C
		return "Ctrl+C", nil
	case n == 1 && buf[0] == 13: // Enter
		return "Enter", nil
	case n == 1 && buf[0] == 27: // ESC
		return "Esc", nil
	}

	// Default: return the raw character
	return string(buf[n-1]), nil
}
