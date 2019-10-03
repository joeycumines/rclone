// Package ansible manages the configuration and detection of ansible environments using the local filesystem, and is
// part of the implementation of ansible module support for rclone functionality
package ansible

import "errors"

type ()

var (
	unsupportedError = errors.New(`ansible support not implemented for your platform or is otherwise unavailable`)
)

// GenerateShellExec will generate a shell script that will run the provided command with the provided args, and will
// encapsulate the output in a JSON object
func GenerateShellExec() {
}
