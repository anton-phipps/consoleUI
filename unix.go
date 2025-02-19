//go:build !windows
// +build !windows

package main

/*
#include <termios.h>
#include <unistd.h>
#include <stdio.h>
char getChar()
{
	struct termios oldt, newt;
	char ch;
	tcgetattr(STDIN_FILENO, &oldt);
	newt = oldt;
	newt.c_lflag&= ~(ICANON | ECHO);
	tcsetattr(STDIN_FILENO, TCSANOW, &newt);
	ch = getchar();
	tcsetattr(STDIN_FILENO, TCSANOW, &oldt);
}
*/
import "C"
import "fmt"

func init() {
	fmt.Println("Loading image, press any key to continue...")
	getChar()
	fmt.Printf("%sH%sJ", ESC, ESC)
}

func getChar() byte {
	return byte(C.getChar())
}
