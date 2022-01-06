const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

router.post('/upload',function(request,response){
	console.log("/category/upload");
	let body = request.body;
	let _name = body.name;
	contracts.CategoryContract.methods.uploadCategory(_name).send({
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
	})
});

router.get("/list",function(request,response){
	console.log("/category/list");
	contracts.CategoryContract.methods.getCategoryList().call().then(value=>{
		// console.log(value);
		let res = {};
		let numlist = [];
		for (let index = 1; index <= value.length;index++){
			numlist.push(index);
		}
		res['numlist'] = numlist;
		res["namelist"] = value;
		response.send(res);
	})
})

router.get("/information",function(request,response){
	console.log("/category/information");
	let _serialNum = request.query.serialNum;
	contracts.CategoryContract.methods.getCategory(_serialNum).call().then(value=>{
		let res = {};
		res["value"] = value;
		response.send(res);
	});
});
 
module.exports = router;