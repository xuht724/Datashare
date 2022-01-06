//管理中间件
const express = require("express");
const cors = require("cors");


module.exports = app => {
	console.log('中间件');

	//cross origin 
	let corsOpt = {
		origin: "*",
		options: "*",
		optionsSuccessStatus: 200
	}
	app.use(cors(corsOpt));
	app.options('*', cors());

	//middleware
	app.use(express.json());
	app.use(express.urlencoded({
		extended: true
	}));

	//routes
	require('./routes')(app);
}