package main

import (
	"fmt"
	"github.com/civet148/go-wallet"
	"github.com/urfave/cli/v2"
)

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
