const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

router.get("/num", function(request, response) {
	console.log("/num");
	contracts.DataContract.methods
		.num()
		.call()
		.then((value) => {
			let res = {};
			res["num"] = value;
			response.send(res);
		});
});

router.post("/verify",function (request,response){
	console.log("/verify");
	let body = request.body;
	let _type = body.type;
	let _address = body.address;
	let _target = body.target;
	let _ip = body.ip;
	//为了测试方便目前都设置为true
	if(_type == 1){
		let res = {};
		res["state"] = true;
		response.send(res);
	}else if(_type == 2){
		let res = {};
		res["state"] = true;
		response.send(res);
	}else if(_type == 3){
		let res = {};
		res["state"] = true;
		response.send(res);
	}else{
		let res = {};
		res["state"] = false;
		response.send(res);
	}
});



module.exports = router;