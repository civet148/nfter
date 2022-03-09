pragma solidity ^0.4.24;

import "./zeppelin/ERC721Full.sol";
import "./zeppelin/ERC721MetadataMintable.sol";

contract TTPNft is ERC721Full, ERC721MetadataMintable {
    constructor() ERC721Full("TTPNFT", "NFT")
    public
    {
      //do not call self function when constructing...
    }
}