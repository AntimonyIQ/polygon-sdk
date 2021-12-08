package util

import (
	"os"

	"github.com/0xPolygon/polygon-sdk/command/dev"
	"github.com/0xPolygon/polygon-sdk/command/genesis"
	"github.com/0xPolygon/polygon-sdk/command/helper"
	"github.com/0xPolygon/polygon-sdk/command/ibft"
	"github.com/0xPolygon/polygon-sdk/command/loadbot"
	"github.com/0xPolygon/polygon-sdk/command/monitor"
	"github.com/0xPolygon/polygon-sdk/command/peers"
	"github.com/0xPolygon/polygon-sdk/command/secrets"
	"github.com/0xPolygon/polygon-sdk/command/server"
	"github.com/0xPolygon/polygon-sdk/command/status"
	"github.com/0xPolygon/polygon-sdk/command/txpool"
	"github.com/0xPolygon/polygon-sdk/command/version"
	"github.com/mitchellh/cli"
)

// Commands returns a mapping of all available commands
func Commands() map[string]cli.CommandFactory {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	meta := helper.Meta{
		UI: ui,
	}

	// Grab a reference to the commands
	serverCmd := server.ServerCommand{UI: ui}
	devCmd := dev.DevCommand{UI: ui}
	genesisCmd := genesis.GenesisCommand{UI: ui}
	monitorCmd := monitor.MonitorCommand{Meta: meta}
	statusCmd := status.StatusCommand{Meta: meta}
	versionCmd := version.VersionCommand{UI: ui}

	ibftCmd := ibft.IbftCommand{}
	ibftSnapshotCmd := ibft.IbftSnapshot{Meta: meta}
	ibftStatusCmd := ibft.IbftStatus{Meta: meta}

	peersCmd := peers.PeersCommand{}
	peersAddCmd := peers.PeersAdd{Meta: meta}
	peersListCmd := peers.PeersList{Meta: meta}
	peersStatusCmd := peers.PeersStatus{Meta: meta}

	txPoolCmd := txpool.TxPoolCommand{}
	txPoolAddCmd := txpool.TxPoolAdd{Meta: meta}
	txPoolStatusCmd := txpool.TxPoolStatus{Meta: meta}

	loadbotCmd := loadbot.LoadbotCommand{Meta: meta}

	secretsManagerCmd := secrets.SecretsCommand{}
	secretsGenerateCmd := secrets.SecretsGenerate{Meta: meta}
	secretsInitCmd := secrets.SecretsInit{Meta: meta}

	return map[string]cli.CommandFactory{

		// GENERIC SDK COMMANDS //

		serverCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &serverCmd, nil
		},
		devCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &devCmd, nil
		},
		genesisCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &genesisCmd, nil
		},

		// PEER COMMANDS //

		peersCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &peersCmd, nil
		},
		peersAddCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &peersAddCmd, nil
		},
		peersStatusCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &peersStatusCmd, nil
		},
		peersListCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &peersListCmd, nil
		},

		// IBFT COMMANDS //

		ibftCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &ibftCmd, nil
		},
		ibftSnapshotCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &ibftSnapshotCmd, nil
		},
		ibftStatusCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &ibftStatusCmd, nil
		},

		// TXPOOL COMMANDS //

		txPoolCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &txPoolCmd, nil
		},
		txPoolAddCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &txPoolAddCmd, nil
		},
		txPoolStatusCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &txPoolStatusCmd, nil
		},

		// BLOCKCHAIN COMMANDS //

		statusCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &statusCmd, nil
		},
		monitorCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &monitorCmd, nil
		},
		versionCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &versionCmd, nil
		},

		// SECRETS MANAGER COMMANDS //
		secretsManagerCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &secretsManagerCmd, nil
		},
		secretsGenerateCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &secretsGenerateCmd, nil
		},
		secretsInitCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &secretsInitCmd, nil
		},

		// LOADBOT COMMANDS //

		loadbotCmd.GetBaseCommand(): func() (cli.Command, error) {
			return &loadbotCmd, nil
		},
	}
}
