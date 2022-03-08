#SHELL=/usr/bin/env bash

BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
CUR_DIR=${PWD}
GOSDK_DIR=${GOPATH}/src/go-sdk
SOLC_VER=0.4.25
ERC20_DIR=./bcos/erc20
ERC721_DIR=./bcos/erc721
NFT_DIR=./bcos/nft
GOODS_DIR=./bcos/goods

build:
	go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o nfter cmd/main.go

zeppelin:
	rm -rf ${GOPATH}/src/v2.0.0.tar.gz
	wget -P ${GOPATH}/src https://github.com/OpenZeppelin/openzeppelin-contracts/archive/refs/tags/v2.0.0.tar.gz \
	&& cd ${GOPATH}/src && tar xvfz v2.0.0.tar.gz

solc:
	rm -rf ${GOSDK_DIR}
	git clone https://github.com/FISCO-BCOS/go-sdk ${GOSDK_DIR} \
	&& cd ${GOSDK_DIR} \
	&& go mod tidy \
	&& go build -o console cmd/console.go \
	&& go build -o abigen cmd/abigen/main.go \
	&& sudo cp abigen /usr/bin \
	&& cd ${GOSDK_DIR}/tools && ./download_solc.sh ${SOLC_VER} && sudo mv solc-${SOLC_VER} /usr/bin/solc \
	&& cd ${CUR_DIR}

erc20:
	mkdir -p ${ERC20_DIR} \
	&& abigen --sol contracts/ERC20.sol --pkg erc20 --out ./bcos/erc20/ERC20.go

erc721:
	mkdir -p ${ERC721_DIR} \
    && abigen --sol contracts/ERC721full.sol --pkg erc721 --out ./bcos/erc721/ERC721full.go

nft:
	mkdir -p ${NFT_DIR} \
    && abigen --sol ${CUR_DIR}/contracts/TTPNFT.sol --pkg nft --out ./bcos/nft/TTPNFT.go

goods:
	mkdir -p ${GOODS_DIR} \
    && abigen --sol ${CUR_DIR}/contracts/TTPGOODS.sol --pkg goods --out ./bcos/goods/TTPGOODS.go

all: erc20 erc721 nft goods