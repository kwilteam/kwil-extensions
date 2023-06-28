// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

/// @title Fractal registry v0
/// @author Antoni Dikov and Shelby Doolittle
contract FractalRegistry {
    address root;
    mapping(address => bytes32) fractalIdForAddress;
    mapping(string => mapping(bytes32 => bool)) userLists;

    constructor(address _root) {
        root = _root;

        address testAccount = 0x640568976c2CDc8789E44B39369D5Bc44B1e6Ad7;
        bytes32 testFractalId = 0xe55149bfd05867a51672a24235e3511767bd64cb1b250c33da303d5be58d2bdd;
        fractalIdForAddress[address(testAccount)] = testFractalId;
        userLists["plus"][0xe55149bfd05867a51672a24235e3511767bd64cb1b250c33da303d5be58d2bdd] = true;
    }

    /// @param addr is Eth address
    /// @return FractalId as bytes32
    function getFractalId(address addr) external view returns (bytes32) {
        return fractalIdForAddress[addr];
    }

    /// @notice Checks if a user by FractalId exists in a specific list.
    /// @param userId is FractalId in bytes32.
    /// @param listId is the list id.
    /// @return bool if the user is the specified list.
    function isUserInList(bytes32 userId, string memory listId)
        external
        view
        returns (bool)
    {
        return userLists[listId][userId];
    }

    struct Grant {
        address owner;
        address grantee;
        string dataId;
    }

    function grantsFor() external pure returns (Grant[] memory) {
        return new Grant[](3);
    }
}