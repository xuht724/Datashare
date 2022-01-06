const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

router.post("/upload",function(request,response){
	console.log("/result/upload");
	let body = request.body;
	let _id = body.id;
	let _name = body.name;
  	let _timestamp = new Date().getTime().toString();
	let _parentData = body.parentData; //The serailNum of parentData;
	let _parentModel = body.parentModel;
	contracts.DataContract.methods.uploadResult(_id,_name,_timestamp,_parentData,_parentModel).send({
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
        	response["serialNum"] = -1;
      		response.send(res);
      	}
	}).catch(()=>{});
});

router.get("/information",function(request,response){
	console.log("/result/information");
	let _serialNum = request.query.serialNum;
	contracts.DataContract.methods.getTaskInf(_serialNum).call().then(value=>{
		console.log(value);
		let res = {};
		res["serialNum"] = value.serialNum;
		res["ownerAddress"] = value.ownerAddress;
		res["ownerName"] = value.ownerNamename;
		res["id"] = value.id;
		res["name"] = value.name;
		res["timestamp"] = value.timestamp;
		res["parentData"] = value.parentData;
		res["parentModel"] = value.parentModel;
		response.send(res);
	});
});

module.exports = router;