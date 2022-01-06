# 基于以太坊区块链的数据共享项目介绍

## 项目简介

本项目是一个针对有限信任的实体之间共享自己拥有数据信息的项目，将数据分为公开数据（Public Data）和隐私数据（Private Data），针对公开数据和私有数据，按照不同的流程来分别对于数据进行共享、使用。

## 项目架构

![image-20220104223328852](https://raw.githubusercontent.com/xuht724/blog-img/main/img/image-20220104223328852.png)

单独客户端各组件功能介绍：

1. Eth Blockchain
   - 自动执行数据共享业务流程
   - 记录数据共享、传输信息
2. Chain Server
   - 为前端提供操作链上数据的API
3. GoFastDfs
   - 提供文件存储服务，可部署成集群
   - 提供文件在内网的下载服务
4. File Server
   - 为前端提供操作文件服务器中文件的API
5. 前端页面
   - 可视化操作展示

## 代码简介

### Chain Server

[API文档](https://www.apifox.cn/apidoc/shared-b68b0b0c-1c7a-4c27-ba46-9a5a0203255f)

启动方法`node server.js`

```
express
|-routes 记录各功能的Api
	|- result.js
	|- model.js
	|- log.js
	|- index.js
	|- identity.js
	|- general.js
	|- data.js
	|- category.js
	|- assets.js
|-server.js 服务器启动文件
|-contract.js 合约实例初始化文件
|-middlewareConfig.js 中间件
|-config.json 用于web3合约实例的创建，包括链上合约的地址和abi接口，以及coinbase地址
|-package.json 依赖项
```

### Eth Blockchain

#### Eth client

本次使用的是Geth客户端，目前项目处于Dev环境下进行测试。运行的客户端需要安装好geth。eth文件夹中存放了当前开发使用的链的数据。启动客户端的命令：

```bash
geth --datadir devdata --allow-insecure-unlock --http --dev --http.corsdomain "*" console 2>>devgeth.log
```

可以直接 `bash devstart.sh`

#### Contract

为部署在`Eth Blockchain`上面的合约代码。架构如图所示，箭头函数表示的是合约之间的依赖关系。

![image-20220106162607472](https://raw.githubusercontent.com/xuht724/blog-img/main/img/image-20220106162607472.png)

Identity合约：最基础的合约，管理地址和账户信息之间的映射关系，Data合约上的所有操作都要经过identity合约管理。

Data合约：核心的数据共享逻辑合约。管理三类文件：数据集、模型以及结果。

Assessment合约：评估合约，用来记录评估数据集、模型的metric和value

Category合约：类别合约。功能为将数据集进行分类。

Log合约：日志合约。目前实现为针对数据文件，记录下载请求；针对模型文件，记录参与计算用户。







