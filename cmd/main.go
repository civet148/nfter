package main

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version     = "1.0.0"
	ProgramName = "nfter"
)

var (
	BuildTime = "2022-01-18"
	GitCommit = ""
)

const (
	CmdNameAccount  = "account"
	CmdNameDeploy   = "deploy"
	CmdNameMint     = "mint"
	CmdNameQuery    = "query"
	CmdNameTransfer = "transfer"
)

const (
	CmdFlagNodeUrl = "node-url"
	CmdFlagChainID = "chain-id"
	CmdFlagGroupID = "group-id"
)

const (
	DefaultNodeUrl = "192.168.2.201:8545"
	DefaultChainID = 1
	DefaultGroupID = 1
)

var (
	mintAddress     = "0x40573435A5eECAb73e6B428ca9e028178c01d77a"
	privateKey      = "01e7a043e06abf15a192585bcd5004e59ccbdc94903160ae696a3a9d01c1b1fe"
	publicKey       = "0237d17a2629880a170b26f30cdda5f4d10824049ccc65afbbe32785147bed7517"
	contractAddress = "0x5C898ee81d1BCD77178A8199832cfD7eEDD0bC9c"
	ownerAddress    = "0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03"
	ownerPrivateKey = "3e5cd186c0de12c83fa4db6b6c5a93e64572721c4e262ce1498eaa2401658cf1"
	receiveAddress  = "0x40573435A5eECAb73e6B428ca9e028178c01d77a" //receiver also is the minter
	tokenURI        = "https://cdn.pixabay.com/photo/2022/02/26/18/16/peace-7036144_960_720.png"
)

func init() {
	log.SetLevel("info")
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func main() {

	grace()

	local := []*cli.Command{
		accountCmd,
		deployCmd,
		mintCmd,
		queryCmd,
		transferCmd,
	}
	app := &cli.App{
		Name:     ProgramName,
		Usage:    "FISCO-BCOS NFT contract operator",
		Version:  fmt.Sprintf("v%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

func newHttpClient(cctx *cli.Context, pk []byte) (*client.Client, error) {
	var config = &conf.Config{
		IsHTTP:     true,
		NodeURL:    cctx.String(CmdFlagNodeUrl),
		PrivateKey: pk,
		ChainID:    cctx.Int64(CmdFlagChainID),
		GroupID:    cctx.Int(CmdFlagGroupID),
	}
	return client.Dial(config)
}
