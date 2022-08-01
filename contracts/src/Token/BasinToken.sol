// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import './ERC20.sol';


/// @title Basin Token
/// @author Nate Sesti and Ty Dunn
/// @notice An ERC-20 token, the native currency of Basin Protocol
/// @dev Explain to a developer any extra details
contract BasinToken is ERC20 {

    constructor() {
        _name = "Basin Token";
        _symbol = "BSN";
    }

}