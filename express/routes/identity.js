const express = require("express");
const contracts = require("../contracts.js")

const router = express.Router();

//get current users in the contracts
router.get("/num", function(request, response) {
	console.log("/identity/num");
	contracts.IdentityContract.methods
		.users()
		.call()
		.then((value) => {
			let res = {};
			res["num"] = value;
			response.send(res);
		});
});
 

router.get("/information", function(request, response) {
	console.log("/identity/information");
	let address = request.query.address;
	console.log("Target Address is", address);
	contracts.IdentityContract.methods.getIdentity(address).call().then((value) => {
		// console.log(value);
		let res = {};
		res["name"] = value.name;
		res["ip"] = value.ip;
		res["port"] = value.port;
		res["isUsed"] = value.isUsed;
		res["points"] = value.points;
		response.send(res);
	});
});


router.get("/list", function(request, response) {
	console.log("/identity/list");
	list().then(value => {
		let res = {};
		res["list"] = value;
		response.send(res);
	});
});

async function list() {
	let users = await contracts.IdentityContract.methods.users().call();
	// console.log(users);
	let res = [];
	for (let i = 0; i < users; i++) {
		let tmp = await contracts.IdentityContract.methods.publicList(i).call();
		let temp = {};
		temp["address"] = tmp.addr;
		temp["name"] = tmp.name;
		res.push(temp);
	}
	return res;
}

router.get("/super",function(request,response){
	console.log("/identity/super");
	contracts.IdentityContract.methods.superAddress().call().then(value=>{
		let res = {};
		res["address"] = value;
		response.send(res);
	})
});

module.exports = router;