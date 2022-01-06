// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Category
 */

contract Category{
    uint256 public num;
    
    string[] categoryList;
    mapping(uint256=>string) categorySet;

    event functionState(string functionName, bool state, string description);

    function uploadCategory(string memory _name) public{
        num++;
        categorySet[num] = _name;
        categoryList.push(_name);
		emit functionState("uploadCategory",true,"upload category succeedd");
    }

    function getCategory(uint256 _serialNum) public view returns(string memory){
        return categorySet[_serialNum];
    }

    function getCategoryList() public view returns(string[] memory){
        return categoryList;
    }
}