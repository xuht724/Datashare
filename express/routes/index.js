const dataRouter = require("./data");
const modelRouter = require("./model");
const resultRouter = require("./result");
const identityRouter = require("./identity");
const generalRouter = require("./general")
const categoryRouter = require("./category");
const logRouter = require("./log");
const assessRouter = require("./assess");


const routes = (app) => {
	app.use('',generalRouter);
	app.use('/data', dataRouter);
	app.use('/model', modelRouter);
	app.use('/result', resultRouter);
	app.use('/identity', identityRouter);
	app.use('/category', categoryRouter);
	app.use('/log', logRouter);
	app.use('/assess',assessRouter)
}

module.exports = routes;