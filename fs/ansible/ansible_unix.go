// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package ansible

// Supported always returns true for these targets
func Supported() bool { return true }
