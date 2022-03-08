import "./zeppelin/ERC721Full.sol";
import "./zeppelin/ERC721MetadataMintable.sol";

contract TTPGoods is ERC721Full, ERC721MetadataMintable {
    constructor() ERC721Full("TTPGOODS", "GOODS")
    public
    {
      //do not call self function when constructing...
    }
}