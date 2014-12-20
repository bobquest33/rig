package main

import (
	"os"
	"os/signal"

	"github.com/fujin/rig/command"
	"github.com/mattbaird/gosaml"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Rig commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{

		"proxy": func() (cli.Command, error) {
			return &command.ProxyCommand{
				ShutdownCh:      makeShutdownCh(),
				UI:              ui,
				AccountSettings: saml.NewAccountSettings("cert", "http://www.onelogin.net"),
				AppSettings:     saml.NewAppSettings("http://www.onelogin.net", "issuer"),
			}, nil
		},
		// "exec": func() (cli.Command, error) {
		// 	return &command.ExecCommand{
		// 		ShutdownCh: makeShutdownCh(),
		// 		Ui:         ui,
		// 	}, nil
		// },

		"version": func() (cli.Command, error) {
			ver := Version
			rel := VersionPrerelease
			if GitDescribe != "" {
				ver = GitDescribe
			}
			if GitDescribe == "" && rel == "" {
				rel = "dev"
			}

			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           ver,
				VersionPrerelease: rel,
				UI:                ui,
			}, nil
		},
	}
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
