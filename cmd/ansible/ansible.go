package ansible

import (
	"github.com/rclone/rclone/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(root())
}

func root() (c *cobra.Command) {
	c = &cobra.Command{
		Use:   `ansible`,
		Short: `Manage Ansible support.`,
	}
	c.AddCommand(install())
	return
}

func install() (c *cobra.Command) {
	c = &cobra.Command{
		Use:   `install`,
		Short: `Create Python script wrappers for use by Ansible.`,
		Run: func(command *cobra.Command, args []string) {
		},
	}
	return
}

func module() (c *cobra.Command) {
	c = &cobra.Command{
		Use:   `module`,
		Short: `Manually run an Ansible module implementation.`,
		Run: func(command *cobra.Command, args []string) {
		},
	}
	return
}
