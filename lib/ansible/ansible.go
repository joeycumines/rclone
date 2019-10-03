// Package ansible implements boilerplate and utilities for the implementation of ansible module support for rclone
// functionality.
//
// See https://docs.ansible.com/ansible/latest/dev_guide/developing_program_flow_modules.html#binary-modules for an
// overview of how this works.
package ansible

type (
	// Module models an Ansible module implementation
	Module interface {
		Name() string
	}
)
