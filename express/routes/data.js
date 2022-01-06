const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

//Some api related to the data information 


router.post("/upload",function(request,response){
	console.log("/data/upload");
	let metaData = request.body;
  	let _id = metaData.id;
  	let _name = metaData.name;
  	let _timestamp = new Date().getTime().toString();
  	let _category = metaData.category;
  	let _point = metaData.point;
  	let _description = metaData.description;
  	let _isOpen = metaData.isOpen;


	contracts.DataContract.methods.uploadData(_id, _name, _timestamp, _category,_point,_description,_isOpen).send({
		from:contracts.hostAddress
	}).then(value=>{
	    // console.log(value);
    	// let number = await getDataNum();
    	let returnValues = value.events.functionState.returnValues;
    	// console.log(returnValues);
    	console.log(returnValues.description);
    	let res = {};
    	if (returnValues.state) {
      		res['state'] = true;
      		res['message'] = returnValues.description;
     		contracts.DataContract.methods.num().call().then(number => {
        		res['serialNum'] = number;
        		response.send(res);
      		});
    	} else {
      		res['state'] = false;
      		res['message'] = returnValues.description;
      		res['serialNum'] = -1;
      		response.send(res);
    	}
	}).catch(error=>{
		console.log(error);
	});
});

router.get("/information",function(request,response){
	console.log("/data/information");
    let serialNum = request.query.serialNum;
  	console.log("Target Data is", serialNum);
  	//调用getDataInf接口获取数据值
  	dataInf(serialNum).then(value=>{
  		response.send(value);
  	});
});

async function dataInf(_serialNum){
	let res = await contracts.DataContract.methods.getDataInf(_serialNum).call();
	console.log(res);
	let jsonRes = {};
	jsonRes["serialNum"] = res.serialNum;
	jsonRes["ownerAddress"] = res.ownerAddress;
	jsonRes["ownerName"] = res.ownerName;
	jsonRes["id"] = res.id;
	jsonRes["name"] = res.name;
	jsonRes["timestamp"] = res.timestamp;
	jsonRes["point"] = res.point;
	jsonRes["category"] = res.category;
	jsonRes["description"] = res.description;
	jsonRes["isOpen"] = res.isOpen;
	jsonRes["isUsed"] = res.isUsed;
	return jsonRes;
}

router.get("/publiclist",function(request,response){
	//返回公开可用的数据列表
	console.log("/data/publiclist");
	publiclist().then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

async function publiclist(){
	let res = await contracts.DataContract.methods.getDatalist().call();
	console.log(res);
	let list = [];
	for(let index = 0;index<res.length;index++){
		let flag1 = await contracts.DataContract.methods.isUsed(res[index]).call();
		let flag2 = await contracts.DataContract.methods.isOpen(res[index]).call();
		if(flag1){
			if(flag2){
				list.push(res[index]);
			}
		}
	}
	return list;
}

router.get("/privatelist",function(request,response){
	//返回隐私可用的数据列表
	console.log("/data/privatelist");
	privatelist().then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

async function privatelist(){
	let res = await contracts.DataContract.methods.getDatalist().call();
	let list = [];
	for(let index = 0;index < res.length;index++){
		let flag1 = await contracts.DataContract.methods.isUsed(res[index]).call();
		let flag2 = await contracts.DataContract.methods.isOpen(res[index]).call();
		if(flag1){
			if(!flag2){
				list.push(res[index]);
			}
		}
	}
	return list;
}

router.post("/buy",function(request,response){
	console.log("/data/buy");
  	let serialNum = request.body.serialNum;
  	console.log("Target Data is", serialNum);
  	contracts.DataContract.methods.buyData(serialNum).send({
    	from: contracts.hostAddress
  	}).then((value) => {
    	let returnValues = value.events.functionState.returnValues;
    	console.log(returnValues.description);
    	let res = {};
    	if (returnValues.state) {
      		res['state'] = true;
      		res['message'] = returnValues.description;
    	} else {
      		res['state'] = false;
      		res['message'] = returnValues.description;
    	}
    	response.send(res);
  });
});

router.get("/ownerList",function(request,response){
	console.log("/data/ownerList");
	let addr = request.query.address;
	usefulList(addr).then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{})
});

async function usefulList(addr){
	let list = await contracts.DataContract.methods.getownerList(addr).call();
	console.log(list);
	let res = []
	for (let index = 0; index < list.length; index++){
		let flag = await contracts.DataContract.methods.isUsed(list[index]).call();
		if(flag){
			res.push(list[index]);
		}
	}
	return res;
}

router.get("/tomodel",function(request,response){
	console.log("/data/tomodel");
	let _serialNum = request.query.serialNum;
	contracts.DataContract.methods.modelOfdata(_serialNum).call().then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	});
});

router.delete("/file",function(request,response){
	let target = request.query.serialNum;
	console.log("/data/file ",target);
	contracts.DataContract.methods.deleteFile(target).send({
		from:contracts.hostAddress
	}).then(value => {
    	let returnValues = value.events.functionState.returnValues
    	console.log(returnValues.description);
    	let res = {};
    	if (returnValues.state) {
      		res['state'] = true;
      		res['message'] = returnValues.description;
    	} else {
      	res['state'] = false;
      	res['message'] = returnValues.description;
    	}
    	response.send(res);
	});
});


router.get("/ownerPublic",function(request,response){
	console.log("/assess/ownerPublic");
	let addr = request.query.address;
	usefulPublic(addr).then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

async function usefulPublic(addr){
	let list = await contracts.DataContract.methods.getownerList(addr).call();
	console.log(list);
	let res = []
	for (let index = 0; index < list.length; index++){
		let flag1 = await contracts.DataContract.methods.isUsed(list[index]).call();
		let flag2 = await contracts.DataContract.methods.isOpen(list[index]).call()
		if(flag1){
			if(flag2){
				res.push(list[index]);
			}
		}
	}
	return res;
}

router.get("/ownerPrivate",function(request,response){
	console.log("/assess/ownerPrivate");
	let addr = request.query.address;
	usefulPrivate(addr).then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

async function usefulPrivate(addr){
	let list = await contracts.DataContract.methods.getownerList(addr).call();
	console.log(list);
	let res = []
	for (let index = 0; index < list.length; index++){
		let flag1 = await contracts.DataContract.methods.isUsed(list[index]).call();
		let flag2 = await contracts.DataContract.methods.isOpen(list[index]).call()
		if(flag1){
			if(!flag2){
				res.push(list[index]);
			}
		}
	}
	return res;
}

module.exports = router;