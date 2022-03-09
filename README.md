# test account for BCOS

```text

1. NFT拥有者

address: [0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03]
private key: [3e5cd186c0de12c83fa4db6b6c5a93e64572721c4e262ce1498eaa2401658cf1]
public key: [03bada18ca7161ce5e8c21ef6c1b32f72524ac6a9d7e3328848b00274cf2a9c958]
phrase: [between miracle thank electric lady fee maid unfair relax rent hunt useful]

2. NFT接收者(同时也是铸币者)
address: [0x40573435A5eECAb73e6B428ca9e028178c01d77a]
private key: [01e7a043e06abf15a192585bcd5004e59ccbdc94903160ae696a3a9d01c1b1fe]
public key: [0237d17a2629880a170b26f30cdda5f4d10824049ccc65afbbe32785147bed7517]
phrase: [soap question avoid invite stool prize cotton access choose stand artefact online]

```

# 合约预编译

生成的abi和bin文件在./bcos/abi目录, go代码文件在./bcos/nft目录

```shell
# 编译TTPNFT.sol合约
$ make nft

# 编译TTPGOODS.sol合约
$ make goods
```

# 合约运行

```shell
# 部署合约
$ make deploy

# 铸币
$ make mint

# 查询NFT信息
$ make query

# 转移NFT所有权
$ make transfer
```