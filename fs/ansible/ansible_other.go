// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris

package ansible

// Supported always returns false for these targets
func Supported() bool { return false }
