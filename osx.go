package main

import (
	"fmt"
	"os/exec"
)

const appleScript = `
tell application "Terminal"
    activate
    tell window 1
        -- Change font size
        set font size to 5
        
        -- Maximize window
        set miniaturized to false
        -- set position to {0, 0}
        set size to {1280, 800} -- Adjust the size as needed
    end tell
end tell
`

func main2() {
	// Execute the AppleScript command using osascript
	cmd := exec.Command("osascript", appleScript)

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running AppleScript:", err)
		return
	}

	fmt.Println("Terminal customized.")
}
