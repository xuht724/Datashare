const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

router.post("/data",function(request,response){
	console.log("post /assess/data");
	let body = request.body;
	dataAssess(body).then(returnValues=>{
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

async function dataAssess(body){
  	let _list = body.list;
    let returnValues = {};
    let _serialNum = body.serialNum;
    for(let index = 0;index < _list.length;index++){
    	let _name = _list[index].name;
    	let _metric = _list[index].metric;
    	let _value = _list[index].value;
    	returnValues = await upLoadOneData(_serialNum,_name,_metric,_value);
    }
    if(_list.length == 0){
    	returnValues["state"] = false;
    	returnValues["description"] = "No assess";
    }
    return returnValues;
}

async function upLoadOneData(_serialNum, _name, _metric, _value) {
  let value = await contracts.AssessmentContract.methods.uploadDataAssess(_serialNum, _name, _metric, _value).send({
    from: contracts.hostAddress
  });
  let returnValues = value.events.functionState.returnValues;
  // console.log(returnValues.description);
  return returnValues;
}

router.get("/data",function(request,response){
	console.log("get /assess/data");
	let _serialNum = request.query.serialNum;
	contracts.AssessmentContract.methods.getDataAssess(_serialNum).call().then(value=>{
	    let len = value['0'];
    	let metrics = value['1'];
    	// console.log(metrics);
    	let list = []
    	for (let i = 0; i < len; i++) {
      		let newjson = {};
      		newjson["name"] = metrics[i].name;
      		newjson["metric"] = metrics[i].metric;
      		newjson["value"] = metrics[i].value;
      		list.push(newjson);
    	}
    	let res = {};
    	res["list"] = list;
    	response.send(res);
	});
})

router.post("/model",function(request,response){
	console.log("post /assess/model");
	let body = request.body;
	modelAssess(body).then(returnValues=>{
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

async function modelAssess(body){
  	let _list = body.list;
    let returnValues = {};
    let _serialNum = body.serialNum;
    for(let index = 0;index < _list.length;index++){
    	let _name = _list[index].name;
    	let _metric = _list[index].metric;
    	let _value = _list[index].value;
    	returnValues = await upLoadOneModel(_serialNum,_name,_metric,_value);
    }
    if(_list.length == 0){
    	returnValues["state"] = false;
    	returnValues["description"] = "No assess";
    }
    return returnValues;
}

async function upLoadOneModel(_serialNum, _name, _metric, _value) {
  let value = await contracts.AssessmentContract.methods.uploadModelAssess(_serialNum, _name, _metric, _value).send({
    from: contracts.hostAddress
  });
  let returnValues = value.events.functionState.returnValues;
  // console.log(returnValues.description);
  return returnValues;
}


router.get("/model",function(request,response){
	console.log("get /assess/model");
	let _serialNum = request.query.serialNum;
	contracts.AssessmentContract.methods.getModelAssess(_serialNum).call().then(value=>{
	    let len = value['0'];
    	let metrics = value['1'];
    	console.log(metrics);
    	let list = []
    	for (let i = 0; i < len; i++) {
      		let newjson = {};
      		newjson["name"] = metrics[i].name;
      		newjson["metric"] = metrics[i].metric;
      		newjson["value"] = metrics[i].value;
      		list.push(newjson);
    	}
    	let res = {};
    	res["list"] = list;
    	response.send(res);
	});
});





module.exports = router;