package main

import (
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/civet148/go-wallet"
	"github.com/civet148/log"
	"github.com/civet148/nfter/bcos/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"math/big"
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
	accountAddress  = "0x40573435A5eECAb73e6B428ca9e028178c01d77a"
	privateKey      = "01e7a043e06abf15a192585bcd5004e59ccbdc94903160ae696a3a9d01c1b1fe"
	publicKey       = "0237d17a2629880a170b26f30cdda5f4d10824049ccc65afbbe32785147bed7517"
	contractAddress = "0x7f0e7fE7d4D199b3A96cD1B8BaD7bf84c144A00E"
	ownerAddress    = "0xfcf3fee2901602b76371bded8d15c973a9fa700d"
	receiveAddress  = "0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03"
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

var accountCmd = &cli.Command{
	Name:      CmdNameAccount,
	Usage:     "create new ethereum account",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {

		var wc = wallet.NewWalletEthereum(wallet.OpType_Create)
		fmt.Printf("address: %s\n", wc.GetAddress())
		fmt.Printf("private key: %s\n", wc.GetPrivateKey())
		fmt.Printf("public key: %s\n", wc.GetPublicKey())
		fmt.Printf("phrase: %s\n", wc.GetPhrase())
		return nil
	},
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

var deployCmd = &cli.Command{
	Name:      CmdNameDeploy,
	Usage:     "deploy NFT contract",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  CmdFlagNodeUrl,
			Usage: "FISCO-BCOS node url",
			Value: DefaultNodeUrl,
		},
		&cli.Int64Flag{
			Name:  CmdFlagChainID,
			Usage: "chain id",
			Value: DefaultChainID,
		},
		&cli.IntFlag{
			Name:  CmdFlagGroupID,
			Usage: "group id",
			Value: DefaultGroupID,
		},
	},
	Action: func(cctx *cli.Context) error {

		//var wc = wallet.NewWalletEthereum(wallet.OpType_Create)
		//log.Infof("[CREATE] address [%s] private key [%s] public key [%s] phrase [%s]", wc.GetAddress(), wc.GetPrivateKey(), wc.GetPublicKey(), wc.GetPhrase())
		pk, err := hex.DecodeString(privateKey)
		if err != nil {
			log.Errorf("decode address [%s] private key [%s] error [%s]", accountAddress, publicKey, err.Error())
			return err
		}

		client, err := newHttpClient(cctx, pk)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()
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

var mintCmd = &cli.Command{
	Name:      CmdNameMint,
	Usage:     "mint NFT",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  CmdFlagNodeUrl,
			Usage: "FISCO-BCOS node url",
			Value: DefaultNodeUrl,
		},
		&cli.Int64Flag{
			Name:  CmdFlagChainID,
			Usage: "chain id",
			Value: DefaultChainID,
		},
		&cli.IntFlag{
			Name:  CmdFlagGroupID,
			Usage: "group id",
			Value: DefaultGroupID,
		},
	},
	Action: func(cctx *cli.Context) error {

		//var wc = wallet.NewWalletEthereum(wallet.OpType_Create)
		//log.Infof("[CREATE] address [%s] private key [%s] public key [%s] phrase [%s]", wc.GetAddress(), wc.GetPrivateKey(), wc.GetPublicKey(), wc.GetPhrase())
		pk, err := hex.DecodeString(privateKey)
		if err != nil {
			log.Errorf("decode address [%s] private key [%s] error [%s]", accountAddress, publicKey, err.Error())
			return err
		}

		client, err := newHttpClient(cctx, pk)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer client.Close()

		ca, err := common.NewMixedcaseAddressFromString(contractAddress)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Infof("NFT contract address [%s]", ca.Address())

		oa, err := common.NewMixedcaseAddressFromString(ownerAddress)
		if err != nil {
			log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
			return err
		}
		log.Infof("NFT owner address [%s]", oa.Address())

		ttp, err := nft.NewTTPNft(ca.Address(), client)
		if err != nil {
			log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
			return err
		}

		var tokenId = big.NewInt(1)
		tx, receipt, err := ttp.MintWithTokenURI(client.GetTransactOpts(), oa.Address(), tokenId, tokenURI)
		if err != nil {
			log.Fatal("MintWithTokenURI error [%s]", err)
			return err
		}
		log.Infof("tx [%+v]", tx)
		log.Infof("receipt [%+v]", receipt)
		return nil
	},
}

var queryCmd = &cli.Command{
	Name:      CmdNameQuery,
	Usage:     "query NFT",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  CmdFlagNodeUrl,
			Usage: "FISCO-BCOS node url",
			Value: DefaultNodeUrl,
		},
		&cli.Int64Flag{
			Name:  CmdFlagChainID,
			Usage: "chain id",
			Value: DefaultChainID,
		},
		&cli.IntFlag{
			Name:  CmdFlagGroupID,
			Usage: "group id",
			Value: DefaultGroupID,
		},
	},
	Action: func(cctx *cli.Context) error {

		pk, err := hex.DecodeString(privateKey)
		if err != nil {
			log.Errorf("decode address [%s] private key [%s] error [%s]", accountAddress, publicKey, err.Error())
			return err
		}

		client, err := newHttpClient(cctx, pk)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer client.Close()

		ca, err := common.NewMixedcaseAddressFromString(contractAddress)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Infof("NFT contract address [%s]", ca.Address())
		oa, err := common.NewMixedcaseAddressFromString(ownerAddress)
		if err != nil {
			log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
			return err
		}
		log.Infof("NFT owner address [%s]", oa.Address())
		ttp, err := nft.NewTTPNft(ca.Address(), client)
		if err != nil {
			log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
			return err
		}

		var tokenId = big.NewInt(1)
		owner, err := ttp.OwnerOf(client.GetCallOpts(), tokenId)
		if err != nil {
			log.Fatal("OwnerOf error [%s]", err)
			return err
		}
		supply, err := ttp.TotalSupply(client.GetCallOpts())
		if err != nil {
			log.Fatal("TotalSupply error [%s]", err)
			return err
		}
		uri, err := ttp.TokenURI(client.GetCallOpts(), tokenId)
		if err != nil {
			log.Fatal("TokenURI error [%s]", err)
			return err
		}
		log.Infof("NFT token id [%d] owner address [%s] supply [%d] uri [%s]", tokenId.Int64(), owner.String(), supply, uri)
		return nil
	},
}

var transferCmd = &cli.Command{
	Name:      CmdNameTransfer,
	Usage:     "transfer NFT",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  CmdFlagNodeUrl,
			Usage: "FISCO-BCOS node url",
			Value: DefaultNodeUrl,
		},
		&cli.Int64Flag{
			Name:  CmdFlagChainID,
			Usage: "chain id",
			Value: DefaultChainID,
		},
		&cli.IntFlag{
			Name:  CmdFlagGroupID,
			Usage: "group id",
			Value: DefaultGroupID,
		},
	},
	Action: func(cctx *cli.Context) error {

		//pk, err := hex.DecodeString(privateKey)
		//if err != nil {
		//	log.Errorf("decode address [%s] private key [%s] error [%s]", accountAddress, publicKey, err.Error())
		//	return err
		//}
		//
		//client, err := newHttpClient(cctx, pk)
		//if err != nil {
		//	log.Fatal(err)
		//	return err
		//}
		//defer client.Close()
		//
		//ca, err := common.NewMixedcaseAddressFromString(contractAddress)
		//if err != nil {
		//	log.Fatal(err)
		//	return err
		//}
		//log.Infof("NFT contract address [%s]", ca.Address())
		//log.Infof("NFT owner address [%s]", ca.Address())
		//oa, err := common.NewMixedcaseAddressFromString(ownerAddress)
		//if err != nil {
		//	log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
		//	return err
		//}
		//
		//ttp, err := nft.NewTTPNft(ca.Address(), client)
		//if err != nil {
		//	log.Fatal("NewMixedcaseAddressFromString error [%s]", err)
		//	return err
		//}
		//
		//var tokenId = big.NewInt(1)
		//tx, receipt, err := ttp.MintWithTokenURI(client.GetTransactOpts(), oa.Address(), tokenId, tokenURI)
		//if err != nil {
		//	log.Fatal("MintWithTokenURI error [%s]", err)
		//	return err
		//}
		//log.Infof("tx [%+v]", tx)
		//log.Infof("receipt [%+v]", receipt)
		return nil
	},
}
