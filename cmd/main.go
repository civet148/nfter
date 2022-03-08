package main

import (
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"github.com/civet148/nfter/bcos/nft"
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
	CmdNameRun = "run"
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
		runCmd,
	}
	app := &cli.App{
		Name:     ProgramName,
		Usage:    "NFT contract committer",
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

var runCmd = &cli.Command{
	Name:      CmdNameRun,
	Usage:     "run NFT committer",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		var strAddress = "0x40573435A5eECAb73e6B428ca9e028178c01d77a"
		var strPrivateKey = "01e7a043e06abf15a192585bcd5004e59ccbdc94903160ae696a3a9d01c1b1fe"
		var strPublicKey = "0237d17a2629880a170b26f30cdda5f4d10824049ccc65afbbe32785147bed7517"
		//var wc = wallet.NewWalletEthereum(wallet.OpType_Create)
		//log.Infof("[CREATE] address [%s] private key [%s] public key [%s] phrase [%s]", wc.GetAddress(), wc.GetPrivateKey(), wc.GetPublicKey(), wc.GetPhrase())
		pk, err := hex.DecodeString(strPrivateKey)
		if err != nil {
			log.Errorf("decode address [%s] private key [%s] error [%s]", strAddress, strPublicKey, err.Error())
			return err
		}
		var config = &conf.Config{
			IsHTTP:     true,
			NodeURL:    "192.168.2.201:8545",
			PrivateKey: pk,
			ChainID:    1,
			GroupID:    1,
		}
		client, err := client.Dial(config)
		if err != nil {
			log.Fatal(err)
		}
		address, tx, instance, err := nft.DeployTTPNft(client.GetTransactOpts(), client)
		if err != nil {
			log.Errorf("DeployTTPNft error [%s]", err)
			return err
		}
		log.Infof("contract address: %s", address.Hex()) // the address should be saved
		log.Infof("transaction hash: %s", tx.Hash().Hex())
		_ = instance
		session := &nft.TTPNftSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
		_ = session
		//session.MintWithTokenURI()
		return nil
	},
}
