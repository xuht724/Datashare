// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./identity.sol";

/**
 * @title Data
 */

contract Data{
    uint public num; //记录Data合约中一共存储了多少共享了的数据

    IdentityControl identityControl;

    uint256[] public Datalist;
    uint256[] public Modellist;
    uint256[] public Resultlist;

    mapping(string => uint256) IdToserialNum; //数据的文件md5 - 合约序列号的映射
    mapping(uint256 => uint8) typeCheck; //文件序列号 - 文件种类

    mapping(uint256=>uint256[]) Model2Result; //记录模型已经完成的训练任务
    mapping(uint256=>uint256[]) Data2Model; //隐私数据-->模型

    mapping(uint256 => data) Dataset; //合约中文件序列号 - 数据本身的映射
    mapping(uint256 => model) Modelset; //合约中文件序列号-模型本身的映射
    mapping(uint256 => result) Resultset; //合约中文件的序列号-计算结果的映射

    mapping(address => uint256[]) ownerData; // 用户地址 - 用户拥有的data的list
    mapping(address => uint256[]) buyList; // 用户地址 - 用户购买了的datalist
    

    //type = 1
    struct data{
        uint256 serialNum;// Serial number of the data in the contract
        address ownerAddress; //The address of the owner
        string ownerName;
        string id; //id of 
        string name; //The name of the dataset
        string timestamp;
        uint256 point;
        uint256 category; //The category of the type for the dataset
        string description; // The description of the dataset
        bool isOpen;
        bool isUsed;
    }

    //type = 2
    struct model{
        uint256 serialNum;
        address ownerAddress;
        string ownerName;
        string id;
        string name;
        string timestamp;
        uint256 category;
        string description;
        uint256[] dataSets;
        bool isFinished;
        bool isUsed;
    }

    //type = 3
    struct result{
        uint256 serialNum;
        address ownerAddress;
        string ownerNamename;
        string id;
        string name;
        string timestamp;
        uint256 parentData; // the serialnum of the parentdata in this contract
        uint256 parentModel; //the serialnum of the parentModel in this contract
        bool isUsed;
    }

    constructor(address _identityControl) {
        num = 0;
        identityControl = IdentityControl(_identityControl);
    }

    event functionState(string functionName, bool state, string description);

    //上传type = 1的文件类型
    function uploadData(
        string memory _id,
        string memory _name,
        string memory _timestamp,
        uint256 _category,
        uint256 _point,
        string memory _description,
        bool _isOpen
    )public{
        if(strcmp(_id, "") || IdToserialNum[_id]!=0){
            emit functionState("upLoadData", false, "invalid id");
        }else if (!identityControl.isRegister(msg.sender)) {
            emit functionState(
                "upLoadData",
                false,
                "The address has not been registered"
            );
        }else{
            num++;
            string memory _ownerName = identityControl.getName(msg.sender);
            data memory new_data = data(
                num,
                msg.sender,
                _ownerName,
                _id,
                _name,
                _timestamp,
                _point,
                _category,
                _description,
                _isOpen,
                true
            );
            Dataset[num] = new_data;
            ownerData[msg.sender].push(num);
            Datalist.push(num);
            IdToserialNum[_id] = num;
            typeCheck[num] = 1;
            emit functionState("upLoadData", true, "upload data succeed");
        }
    }


    //上传type = 2 的模型类型
    function uploadModel(
        string memory _id,
        string memory _name, 
        string memory _timestamp,
        uint256 _category,
        string memory _description,
        uint256[] memory _dataSets
    ) public {
        if(strcmp(_id, "") || IdToserialNum[_id]!=0 ){
            emit functionState("uploadModel", false, "invalid id");
        }else if(!(identityControl.isRegister(msg.sender))){
            emit functionState("uploadModel", false, "The address has not been registered");
        }else{
            num++;
            string memory _ownerName = identityControl.getName(msg.sender);
            model memory new_model = model(
                num,
                msg.sender,
                _ownerName,
                _id,
                _name, 
                _timestamp,
                _category,
                _description,
                _dataSets,
                false,
                true
            );
            Modelset[num] = new_model;
            Modellist.push(num);
            IdToserialNum[_id] = num;
            typeCheck[num] = 2;
            for(uint256 index = 0;index < _dataSets.length; index++){
                Data2Model[_dataSets[index]].push(num);
            }
            emit functionState("uploadModel", true, "upload model succeed");
        }
    }

    // 上传 type = 3 的模型类型
    function uploadResult(
        string memory _id,
        string memory _name,
        string memory _timestamp,
        uint256 _parentData,
        uint256 _parentModel
    ) public {
        if(strcmp(_id, "") || IdToserialNum[_id]!=0 ){
            emit functionState("uploadResult", false, "invalid id");
        }else if(!(identityControl.isRegister(msg.sender))){
            emit functionState("uploadResult", false, "The address has not been registered");
        }else{
            num++;
            string memory _ownerName = identityControl.getName(msg.sender);
            result memory new_result = result(
                num,
                msg.sender,
                _ownerName,
                _id,
                _name,
                _timestamp,
                _parentData,
                _parentModel,
                true
            );
            Resultset[num] = new_result;
            Resultlist.push(num);
            IdToserialNum[_id] = num;
            typeCheck[num] = 3;
            Model2Result[_parentModel].push(num); //上传结果到模型
            emit functionState("uploadResult", true, "upload result succeed");
        }    
    }

    //删除某一个文件
    function deleteFile(uint256 _serialNum) public {
        if(Dataset[_serialNum].ownerAddress == msg.sender){
            Dataset[_serialNum].isUsed = false;
            emit functionState("deleteFile",true,"delete file succeed");
        }else{
            emit functionState("deletFile",false,"you do not have the right");
        }
    }

    //下载数据 （和积分相关）
    function buyData(uint256 _serialNum) public {
        uint8 _type = typeCheck[_serialNum];
        if(_type == 1){
            uint256 price = Dataset[_serialNum].point;
            if(Dataset[_serialNum].isOpen && identityControl.usePoints(msg.sender, price)){
                address owner = Dataset[_serialNum].ownerAddress;
                identityControl.earnPoints(owner, price);
                buyList[msg.sender].push(_serialNum);
                emit functionState("buyData",true,"Buy data succeed");
            }else{
                emit functionState("buyData",false,"Your points is not enough or the data is private");
            }
        }else{
            emit functionState("buyData",false,"The target is not a dataset!");
        }
    }

    function finishModel(uint256 _serialNum) public {
        uint _type = typeCheck[_serialNum];
        if(_type == 2){
            address target = Modelset[_serialNum].ownerAddress;
            if(target == msg.sender){
                Modelset[_serialNum].isFinished = true;
                emit functionState("finishModel",true,"Finish the model");
            }else{
                emit functionState("finishModel",false,"Only the owner has the right");
            }   
        }else{
            emit functionState("finishModel",false,"The target is not a model");
        }
    }

    function getDatalist() public view returns(uint256[] memory){
        return Datalist;
    }

    function getModellist() public view returns(uint256[] memory){
        return Modellist;
    }

    function getResultlist() public view returns(uint256[] memory){
        return Resultlist;
    }

    function resultOfmodel(uint256 _serialNum) public view returns(uint256[] memory){
        //返回模型训练已完成的result文件的序列号
        return Model2Result[_serialNum];
    }

    function modelOfdata(uint256 _serialNum) public view returns(uint256[] memory){
        //返回隐私数据对应的模型序列号
        return Data2Model[_serialNum];
    } 

    // 验证是否可以下载verifyData()
    function getBuylist(address _addr) public view returns(uint256[] memory){
        return buyList[_addr]; //获取某一个用户上传了的数据的列表
    }

    function getownerList(address _addr) public view returns(uint256[] memory){
        return ownerData[_addr]; //获取某一个用户上传了的数据的序列号列表
    }

    // 验证是否有权限下载Model
    function verifyModel(uint256 _serialNum) public view returns(uint256[] memory){
        //Model只有_datasets()的owner有权力下载
        uint8 _type = typeCheck[_serialNum];
        if(_type == 2){
            return Modelset[_serialNum].dataSets;
        }else{
            uint256[] memory res;
            return res;
        }
    }

    // 验证是否有权限下载Result;只有result的parentModel有权限下载该
    function verifyResult(address _addr,uint256 _serialNum) public view returns(bool){
        uint8 _type = typeCheck[_serialNum];
        if(_type == 3){
            uint256 index = Resultset[_serialNum].parentModel;
            if (_addr == Modelset[index].ownerAddress){
                return true;
            }else{
                return false;
            }
        }else{
            return false;
        }
    }   

    function isUsed(uint256 _serialNum) public view returns(bool){
        return Dataset[_serialNum].isUsed;
    }

    function isOpen(uint256 _serialNum) public view returns(bool){
        return Dataset[_serialNum].isOpen;
    }

    function typeConfirm(uint256 _serialNum) public view returns(uint8){
        return typeCheck[_serialNum];
    }

    //获取3种type的模型类型
    function getDataInf(uint256 _serialNum) public view returns(data memory){
        data memory res = Dataset[_serialNum];
        return res;
    }

    function getModelInf(uint256 _serialNum) public view returns(model memory){
        model memory res = Modelset[_serialNum];
        return res;
    }

    function getTaskInf(uint256 _serialNum) public view returns(result memory){
        result memory res = Resultset[_serialNum];
        return res;
    }

    function getOwnerAddress(uint256 _serialNum) public view returns(address){
        uint8 _type = typeCheck[_serialNum]; 
        if (_type == 1){
            //data
            return Dataset[_serialNum].ownerAddress;
        }else if(_type == 2){
            return Modelset[_serialNum].ownerAddress;
        }else if(_type == 3){
            return Resultset[_serialNum].ownerAddress;
        }else {
            return address(0x0); //说明该serialNum对应的数据不存在，返回空地址
        }
    }

    //internal functions
    //对比两个字符串是否相等
    function strcmp(string memory a, string memory b)
        internal
        pure
        returns (bool)
    {
        if (bytes(a).length != bytes(b).length) {
            return false;
        } else {
            bytes memory _a = bytes(a);
            bytes memory _b = bytes(b);
            return keccak256(_a) == keccak256(_b);
        }
    }

}