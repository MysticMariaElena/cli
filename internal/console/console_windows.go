// +build windows

package console

import (
	"os"

	"golang.org/x/sys/windows"
)

func SetMode() {
	// TODO pass stdout, preferably from opts.IO
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
