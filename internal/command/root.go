package command

import (
	"github.com/khulnasoft/orchard/internal/command/context"
	"github.com/khulnasoft/orchard/internal/command/controller"
	"github.com/khulnasoft/orchard/internal/command/create"
	deletepkg "github.com/khulnasoft/orchard/internal/command/deletecmd"
	"github.com/khulnasoft/orchard/internal/command/dev"
	"github.com/khulnasoft/orchard/internal/command/get"
	"github.com/khulnasoft/orchard/internal/command/list"
	"github.com/khulnasoft/orchard/internal/command/logs"
	"github.com/khulnasoft/orchard/internal/command/pause"
	"github.com/khulnasoft/orchard/internal/command/portforward"
	"github.com/khulnasoft/orchard/internal/command/resume"
	"github.com/khulnasoft/orchard/internal/command/set"
	"github.com/khulnasoft/orchard/internal/command/ssh"
	"github.com/khulnasoft/orchard/internal/command/vnc"
	"github.com/khulnasoft/orchard/internal/command/worker"
	"github.com/khulnasoft/orchard/internal/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "orchard",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.FullVersion,
	}

	addGroupedCommands(command, "Working With Resources:",
		create.NewCommand(),
		deletepkg.NewCommand(),
		get.NewCommand(),
		list.NewCommand(),
		logs.NewCommand(),
		pause.NewCommand(),
		portforward.NewCommand(),
		resume.NewCommand(),
		set.NewCommand(),
		ssh.NewCommand(),
		vnc.NewCommand(),
	)

	addGroupedCommands(command, "Administrative Tasks:",
		context.NewCommand(),
		controller.NewCommand(),
		worker.NewCommand(),
		dev.NewCommand(),
	)

	return command
}

func addGroupedCommands(parent *cobra.Command, title string, commands ...*cobra.Command) {
	group := &cobra.Group{
		ID:    title,
		Title: title,
	}
	parent.AddGroup(group)

	for _, command := range commands {
		command.GroupID = group.ID
		parent.AddCommand(command)
	}
}
