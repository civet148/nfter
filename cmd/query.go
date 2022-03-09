package main

import (
	"encoding/hex"
	"github.com/civet148/log"
	"github.com/civet148/nfter/bcos/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"math/big"
)

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
