// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import"./datacontrol.sol";

contract Log{

	Data data;

	//serialNum => public dataLog list
	mapping(uint256 => dataLog[]) publicLog;
	mapping(uint256 => modelLog) taskLog;

	struct dataLog{
		uint256 serialNum;
		address sender;
        string name;
		string ip;
		string time;
	}

	struct modelLog{
		uint256 serialNum;
		string timestamp;
		string[] participant;
	}

    event functionState(string functionName, bool state, string description);

	//初始化的时候声明data合约
	constructor(address _data){
		data = Data(_data);
	}

	function uploaddataLog(uint256 _serialNum, address _sender, string memory _ip, string memory _time) public {
		uint8 datatype = data.typeConfirm(_serialNum);
		if(datatype != 1){
			emit functionState("uploaddataLog",false, "Target is not data");
		}else{
			bool isOpen = data.isOpen(_serialNum);
			if(!isOpen){
				emit functionState("uploaddataLog",false, "Target is not open");
			}else{
                string memory _name = data.getDataInf(_serialNum).ownerName;
				dataLog memory new_dataLog = dataLog(_serialNum,_sender,_name,_ip,_time);
				publicLog[_serialNum].push(new_dataLog);
				emit functionState("uploaddataLog",true,"Upload Log successfully");
			}
		}
	}

	function getdataLogs(uint256 _serialNum) public view returns(uint256, dataLog[] memory){
		uint256 len = publicLog[_serialNum].length;
		return (len, publicLog[_serialNum]);
	}

	function uploadmodelLog(uint256 _serialNum, string memory _timestamp, string[] memory _participant) public{
		uint8  datatype = data.typeConfirm(_serialNum);
		if(datatype != 2){
			emit functionState("uploadmodelLog",false,"Target is not model");
		}else{
			address _addr = data.getModelInf(_serialNum).ownerAddress;
			if(_addr!=msg.sender){
				emit functionState("uploadmodelLog",false,"Only the Owner has the right");
			}else{
				modelLog memory new_modelLog = modelLog(_serialNum,_timestamp, _participant);
				taskLog[_serialNum] = new_modelLog;
				emit functionState("uploadmleodelLog",true,"Up load model log successfully");
			}
		}
	}

	function getModelLog(uint256 _serialNum) public view returns(modelLog memory){
		return taskLog[_serialNum];
	}
}
