// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Assessment
 */

import "./datacontrol.sol";

contract Assessment{
    Data data;

    mapping(uint256 => dataAssess[]) dataAssessList;
    mapping(uint256 => modelAssess[]) modelAssessList;

    struct modelAssess{
        string name;
        string metric;
        string value;
    }

    struct dataAssess{
        string name;
        string metric;
        string value;
    }

    event functionState(string functionName, bool state, string description);

    constructor(address _data){
        data = Data(_data);
    }

    function uploadDataAssess(uint256 _serialNum, string memory _name,string memory _metric, string memory _value) public {
        uint8 typeRes = data.typeConfirm(_serialNum);
        if(typeRes != 1){
            emit functionState("uploadDataAssess", false, "The target is not data");        
        }else{
            address addr = data.getDataInf(_serialNum).ownerAddress;
            if(addr!=msg.sender){
                emit functionState("uploadDataAssess", false, "Only the owner has the right");
            }else{
                dataAssess memory new_assess = dataAssess(_name, _metric,_value);
                dataAssessList[_serialNum].push(new_assess);
                emit functionState("uploadDataAssess",true,"upload assess of data successfully");
            }
        }
    }

    function getDataAssess(uint256 _serialNum) public view returns(uint256, dataAssess[] memory){
        uint256 len = dataAssessList[_serialNum].length;
        return (len, dataAssessList[_serialNum]);
    }

    function uploadModelAssess(uint256 _serialNum, string memory _name, string memory _metric, string memory _value) public {
        uint8 typeRes = data.typeConfirm(_serialNum);
        if(typeRes != 2){
            emit functionState("uploadModelAssess", false, "The target is not model");        
        }else{
            address addr = data.getModelInf(_serialNum).ownerAddress;
            if(addr!=msg.sender){
                emit functionState("uploadModelAssess", false, "Only the owner has the right");
            }else{
                modelAssess memory new_assess = modelAssess(_name, _metric, _value);
                modelAssessList[_serialNum].push(new_assess);
                emit functionState("uploadModelAssess",true,"upload assess of model successfully");
            }
        }
    }

    function getModelAssess(uint256 _serialNum) public view returns(uint256, modelAssess[] memory){
        uint256 len = modelAssessList[_serialNum].length;
        return (len, modelAssessList[_serialNum]);
    }

}
