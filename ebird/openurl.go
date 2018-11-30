// Copyright (c) 2018, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ebird

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// OpenURL makes a best guess attempt to open a URL (typically a link) using a
// browser.
//
// TODO(kyle): move to a separate package
func OpenURL(u string) {
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", u).Start()
	case "linux":
		err = exec.Command("xdg-open", u).Start()
	case "windows":
		cmd := "url.dll,FileProtocolHandler"
		runDll32 := filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
		err = exec.Command(runDll32, cmd, u).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
}
