const Web3 = require("web3");
const fs = require("fs");

let config = fs.readFileSync("./config.json");
let configdata = JSON.parse(config);

const port = configdata.port;
const hostAddress = configdata.hostAddress;

const identityAbi = configdata.identityAbi;
const datacontrolAbi = configdata.datacontrolAbi;
const categoryAbi = configdata.categoryAbi;
const logAbi = configdata.logAbi;
const assessmentAbi = configdata.assessmentAbi;

const identityAddress = configdata.identityAddress;
const datacontrolAddress = configdata.datacontrolAddress;
const categoryAddress = configdata.categoryAddress;
const logAddress = configdata.logAddress;
const assessmentAddress = configdata.assessmentAddress

//利用web3和链建立联系
if (typeof web3 !== "undefined") {
	console.log("No web3");
	web3 = new Web3(web3.currentProvider);
} else {
	// set the provider you want from Web3.providers
	// console.log("get!");
	web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
	// console.log(web3.eth.accounts);
}


// 根据abi编码和地址构建合约实例
let DataContract = new web3.eth.Contract(datacontrolAbi, datacontrolAddress);
let IdentityContract = new web3.eth.Contract(identityAbi, identityAddress);
let CategoryContract = new web3.eth.Contract(categoryAbi, categoryAddress);
let LogContract = new web3.eth.Contract(logAbi, logAddress);
let AssessmentContract = new web3.eth.Contract(assessmentAbi,assessmentAddress);

let contracts = {
	hostAddress,
	DataContract,
	IdentityContract,
	CategoryContract,
	LogContract,
	AssessmentContract
}

module.exports = contracts;