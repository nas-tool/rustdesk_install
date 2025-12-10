//go:build windows

package util

import (
	"os/exec"
)

func configureSysProcAttr(cmd *exec.Cmd) {
	// Windows下不需要设置Setsid，保持默认即可
}
