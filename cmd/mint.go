package main

import (
	"encoding/hex"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/civet148/log"
	"github.com/civet148/nfter/bcos/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"math/big"
)

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
			log.Errorf("decode address [%s] private key [%s] error [%s]", mintAddress, publicKey, err.Error())
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
		if receipt.GetStatus() != types.Success {
			log.Errorf("tx status return [%s]", receipt.GetErrorMessage())
			return nil
		}
		log.Infof("tx [%+v]", tx)
		log.Infof("receipt [%+v]", receipt)
		return nil
	},
}
