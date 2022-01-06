// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title IdentityControl
 */
 
contract IdentityControl{
    uint256 public users;

    brief[] public publicList;

    mapping(address => identity) Identity;// address - identity of the network peers    

    PointsTransaction[] public transactionsInfo;

    address public superAddress;
    address public dataAddress;

    struct identity{
        string name;
        string ip; //the ip of the file server
        uint256 port;// the port used to provide the service
        bool isUsed; //judge the address if has been used
        uint256 points; //积分点
        bool isSuper;   //  是否是超级账户
    }

    struct brief{
        address addr;
        string name;
    }

    // model points transaction
    enum TransactionType {
        Earned,
        Redeemed,
        charge
    }

    struct PointsTransaction {
        uint timestamp;
        uint points;
        TransactionType transactionType;
        address identityAddress;
    }
    
    constructor(address _addr){
        users = 0;
        superAddress = _addr;
    }
    
    function isRegister(address _add) public view returns(bool){
        return Identity[_add].isUsed;   
    }    

    modifier onlySuper(address _add){
        require(_add == superAddress, "Only SuperIdentity has the right.");
        _;
    }

    // modifier hasPoints(address _add, uint _points) {
    //     // verify enough points for member
    //     require(Identity[_add].points >= _points, "Insufficient points.");
    //     _;
    // }
    
    modifier onlyData(address _addr){
        require(_addr == dataAddress, "Only Data contract has the right");
        _;
    }


    // register at firt when join the network
    function register(string memory _name, string memory _ip, uint256 _port) public{
        users ++;
        if(msg.sender != superAddress){
            identity memory new_identity = identity(_name,_ip,_port,true,100,false);
            brief memory new_brief = brief(msg.sender,_name);
            publicList.push(new_brief);
            Identity[msg.sender] = new_identity;
        }else{
            identity memory new_identity = identity(_name,_ip,_port,true,100000,true);
            brief memory new_brief = brief(msg.sender,_name);
            publicList.push(new_brief);
            Identity[msg.sender] = new_identity;
        }
    }

    function charge(address _add,uint256 _points)public
        onlySuper(msg.sender)
        returns(bool){
            Identity[_add].points = Identity[_add].points + _points;

            // add transction
            transactionsInfo.push(PointsTransaction({
                points: Identity[_add].points,
                timestamp: block.timestamp,
                transactionType: TransactionType.charge,
                identityAddress: _add
            }));

            return true;            
        }

    function earnPoints(address _add, uint256 _point)public 
        onlyData(msg.sender)
        returns(bool){
        Identity[_add].points = Identity[_add].points + _point;

        // add transction
        transactionsInfo.push(PointsTransaction({
            points: Identity[_add].points,
            timestamp: block.timestamp,
            transactionType: TransactionType.Earned,
            identityAddress: _add
        }));

        return true;
    }

    function usePoints(address _add,uint256 _point) public 
        onlyData(msg.sender)
        returns(bool){
            if(Identity[_add].points < _point){
                return false;
            }else{
                Identity[_add].points = Identity[_add].points - _point;
                // add transction
                transactionsInfo.push(PointsTransaction({
                    points: Identity[_add].points,
                    timestamp: block.timestamp,
                    transactionType: TransactionType.Redeemed,
                    identityAddress: _add
                }));
                return true;
            }
    }

    // function changeSuper(address _add) public 
    //     onlySuper(msg.sender)
    //     returns(bool){
    //         superAddress = _add;
    //         return true;
    //     }

    function changeData(address _add) public
        onlySuper(msg.sender)
        returns(bool){
            dataAddress = _add;
            return true;
        }
    
    function changeIdentity(string memory _name,string memory _ip,uint256 _port) public {
        Identity[msg.sender].name = _name;
        Identity[msg.sender].ip = _ip;
        Identity[msg.sender].port = _port;
    }

    function getIdentity(address _add) public view returns(identity memory){
        return Identity[_add];
    }

    function getName(address _add) public view returns(string memory){
        return Identity[_add].name;
    }
    
    function getUrl(address _add) public view returns(string memory,uint256){
        return (Identity[_add].ip,Identity[_add].port);
    }
    
    function getIp(address _add) public view returns(string memory){
        return (Identity[_add].ip);
    }
    
    function getPort(address _add)public view returns(uint256){
        return (Identity[_add].port);
    }
}