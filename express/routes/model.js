const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();


//和模型相关的api
router.post("/upload",function(request,response){
	console.log("/model/upload");
  	let body = request.body;
  	let _id = body.id;
  	let _name = body.name;
  	let _timestamp = new Date().getTime().toString();
  	let _category = body.category;
  	let _description = body.description;
  	let _dataSets = body.dataSets;
  	contracts.DataContract.methods.uploadModel(_id,_name,_timestamp,_category,_description,_dataSets).send({
  		from:contracts.hostAddress
  	}).then(value=>{
      	let returnValues = value.events.functionState.returnValues
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

  	});
});

router.get("/information",function(request,response){
	console.log("/model/information");
	let _serialNum = request.query.serialNum;
	contracts.DataContract.methods.getModelInf(_serialNum).call().then(value=>{
		// console.log(value);
		let res = {};
		res["serialNum"] = value.serialNum;
		res["ownerAddress"] = value.ownerAddress;
		res["ownerName"] = value.ownerName;
		res["id"] = value.id;
		res["name"] = value.name;
		res["timestamp"] = value.timestamp;
		res["category"] = value.category;
		res["description"] = value.description;
		res["dataSets"] = value.dataSets
		res["isFinished"] = value.isFinished;
		response.send(res);
	});
});

router.get("/list",function(request,response){
	console.log("/model/list");
	contracts.DataContract.methods.getModellist().call().then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

router.get("/ownerlist",function(request,response){
	console.log("/model/ownerlist");
	let addr = request.query.address;
	ownerModel(addr).then(value=>{
		let res = {};
		res["list"] = value;
		response.send(res);
	}).catch(()=>{});
});

async function ownerModel(addr){
	let list = await contracts.DataContract.methods.getModellist().call();
	let res = [];
	for(let i = 0;i<list.length;i++){
		let testaddr = await contracts.DataContract.methods.getOwnerAddress(list[i]).call();
		if(addr == testaddr){
			res.push(list[i]);
		}
	}
	return res;
}

router.get("/finishResult",function(request,response){
	console.log("/model/finishResult");
	let _serialNum = request.query.serialNum;
	Results(_serialNum).then(value=>{
		response.send(value);
	});
});

async function Results(_serialNum){
	let list1 = []; // Data serialNum list
	let list2 = await contracts.DataContract.methods.resultOfmodel(_serialNum).call(); //result serialNum list
	let list3 = [];
	for(let i = 0;i < list2.length;i++){
		let obj = await contracts.DataContract.methods.getTaskInf(list2[i]).call();
		list1.push(obj.parentData);
		list3.push(obj.id);
	}
	let res = {};
	res["datalist"] = list1;
	res["resultlist"] = list2;
	res['idlist'] = list3;
	return res;
}	

router.post("/finish",function(request,response){
	console.log("/model/finish");
	let _serialNum = request.query.serialNum;
	contracts.DataContract.methods.finishModel(_serialNum).send({
		from:contracts.hostAddress
	}).then(value=>{
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
	})
});


module.exports = router;