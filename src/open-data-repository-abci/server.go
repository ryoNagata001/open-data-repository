package main

import (
	"open-data-repository/src/open-data-repository-abci/route"
	"os"
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
)

func main() {
	initApp()
}

func initApp() error {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Create the application
	var app types.Application

	app = route.NewJSONStoreApplication()

	// Start the listener
	srv, err := server.NewServer("tcp://127.0.0.1:26658", "socket", app)
	if err != nil {
		return err
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		return err
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
	return nil
}