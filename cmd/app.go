// cmd documents the entire CLI surface of the application. It is the entry point for the CLI and defines the commands and flags that the CLI supports. It also contains the main function that runs the CLI application.
package cmd

import (
	"github.com/joakimen/goji/cmd/auth"
	"github.com/joakimen/goji/cmd/epic"

	"github.com/joakimen/goji/cmd/issue"

	"github.com/urfave/cli/v2"

	"github.com/joakimen/goji/pkg/config"
)

func NewApp() cli.App {
	return cli.App{
		Name:  config.CliName,
		Usage: "A small Jira CLI",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose output",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "auth",
				Usage: "Manage Jira credentials",
				Subcommands: []*cli.Command{
					{
						Name:  "login",
						Usage: "Add Jira credentials to the system keyring",
						Action: func(cCtx *cli.Context) error {
							return auth.Login()
						},
					},
					{
						Name:  "show",
						Usage: "Show stored credentials",
						Action: func(cCtx *cli.Context) error {
							return auth.Show()
						},
					},
					{
						Name:  "status",
						Usage: "Show authentication status",
						Action: func(cCtx *cli.Context) error {
							return auth.Status()
						},
					},
					{
						Name:  "logout",
						Usage: "Remove stored Jira credentials",
						Action: func(cCtx *cli.Context) error {
							return auth.Logout()
						},
					},
				},
			},
			{
				Name:  "epic",
				Usage: "Manage epics",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List epics",
						Action: func(cCtx *cli.Context) error {
							projectId := cCtx.String("project")
							jsonOutput := cCtx.Bool("json")
							all := cCtx.Bool("all")
							mine := cCtx.Bool("mine")
							return epic.List(projectId, jsonOutput, all, mine)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "project",
								Usage:   "Project key",
								Aliases: []string{"p"},
							},
							&cli.BoolFlag{
								Name:    "all",
								Usage:   "Return both resolved and unresolved epics",
								Aliases: []string{"a"},
							},
							&cli.BoolFlag{
								Name:    "json",
								Usage:   "Output as JSON",
								Aliases: []string{"j"},
							},
							&cli.BoolFlag{
								Name:    "mine",
								Usage:   "Return only epics assigned to the current user",
								Aliases: []string{"m"},
							},
						},
					},
				},
			},
			{
				Name:  "issue",
				Usage: "Manage issues",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List issues",
						Action: func(cCtx *cli.Context) error {
							projectId := cCtx.String("project")
							jsonOutput := cCtx.Bool("json")
							all := cCtx.Bool("all")
							mine := cCtx.Bool("mine")
							limit := cCtx.Int("limit")
							return issue.List(projectId, jsonOutput, all, mine, limit)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "project",
								Usage:   "Project key",
								Aliases: []string{"p"},
							},
							&cli.BoolFlag{
								Name:    "all",
								Usage:   "Return both resolved and unresolved issues",
								Aliases: []string{"a"},
							},
							&cli.BoolFlag{
								Name:    "json",
								Usage:   "Output as JSON",
								Aliases: []string{"j"},
							},
							&cli.BoolFlag{
								Name:    "mine",
								Usage:   "Return only issues assigned to the current user",
								Aliases: []string{"m"},
							},
							&cli.IntFlag{
								Name:    "limit",
								Usage:   "Maximum number of issues to return",
								Value:   50,
								Aliases: []string{"l"},
							},
						},
					},
				},
			},
		},
	}
}
