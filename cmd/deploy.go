package main

import (
	"encoding/hex"
	"github.com/civet148/log"
	"github.com/civet148/nfter/bcos/nft"
	"github.com/urfave/cli/v2"
)

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
			log.Errorf("decode address [%s] private key [%s] error [%s]", mintAddress, publicKey, err.Error())
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
