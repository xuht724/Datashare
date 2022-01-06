//项目的中间件
const express = require("express");

const app = express();
const port = 12345;

//定义中间件
require("./middlewareConfig.js")(app);


app.listen(port, () => {
	console.log(`http://localhost:${port}`);
})