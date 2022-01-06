const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

router.post('/public',function(request,response){
	console.log('/log/public')
	let body = request.body;
	let _addr = body.address;
	let _target = body.target;
	let _ip = body.ip;
	let _time = body.time;
	contracts.LogContract.methods.uploaddataLog(_target,_addr,_ip,_time).send({
		from:contracts.hostAddress
	}).then(value=>{
    	let returnValues = value.events.functionState.returnValues;
    	let res = {};
    	if (returnValues.state) {
      		res['state'] = true;
      		res['message'] = returnValues.description;
      		response.send(res);
    	} else {
      		res['state'] = false;
      		res['message'] = returnValues.description;
      		response.send(res);
    	}
	});
})

router.get("/data",function(request,response){
	console.log("/log/data");
	let _serialNum = request.query.serialNum;
	contracts.LogContract.methods.getdataLogs(_serialNum).call().then(value=>{
		// console.log(value);
    	let num = value["0"];
    	let logs = value["1"];
    	let list = []
    	for (var i = 0; i < num; i++) {
      	let newjson = {};
      		newjson["name"] = logs[i].name;
      		newjson["time"] = logs[i].time;
      		newjson["ip"] = logs[i].ip;
      		list.push(newjson);
    	}
    	let res = {};
    		res["num"] = num;
    		res["Logs"] = list
    		response.send(res);
		})
});


router.post('/model',function(request,response){
	console.log("post /log/model");
	let body = request.body;
  	let _serialNum = body.serialNum;
  	let _timestamp = new Date().getTime().toString();
	let _participant = body.participant;
	contracts.LogContract.methods.uploadmodelLog(_serialNum, _timestamp, _participant).send({
		from:contracts.hostAddress
	}).then(value=>{
      	let returnValues = value.events.functionState.returnValues
      	// console.log(returnValues.description);
      	let res = {};
      	if (returnValues.state) {
        	res['state'] = true;
        	res['message'] = returnValues.description;
      	} else {
       		res['state'] = false;
        	res['message'] = returnValues.description;
     	}
      	response.send(res);
		}).catch(()=>{});
});

router.get("/model",function(request,response){
	console.log("get /log/model");
	let _serialNum = request.query.serialNum;
	contracts.LogContract.methods.getModelLog(_serialNum).call().then(value=>{
		let list = value.participant;
    	let res = {};
    	res["participant"] = list;
    	response.send(res);
	});
});

module.exports = router;