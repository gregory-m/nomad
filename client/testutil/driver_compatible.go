package testutil

import (
	"os/exec"
	"runtime"
	"syscall"
	"testing"
)

func ExecCompatible(t *testing.T) {
	if runtime.GOOS != "windows" && syscall.Geteuid() != 0 {
		t.Skip("Must be root on non-windows environments to run test")
	}
}

func QemuCompatible(t *testing.T) {
	if runtime.GOOS != "windows" && syscall.Geteuid() != 0 {
		t.Skip("Must be root on non-windows environments to run test")
	}
}

func RktCompatible(t *testing.T) {
	if runtime.GOOS == "windows" || syscall.Geteuid() != 0 {
		t.Skip("Must be root on non-windows environments to run test")
	}
	// else see if rkt exists
	_, err := exec.Command("rkt", "version").CombinedOutput()
	if err != nil {
		t.Skip("Must have rkt installed for rkt specific tests to run")
	}
}

func MountCompatible(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support mount")
	}

	if syscall.Geteuid() != 0 {
		t.Skip("Must be root to run test")
	}
}
