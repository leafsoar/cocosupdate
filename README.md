# Cocos 游戏资源更新

## 说明

使用 Golang + HTTP/1.1 + Json

使用方法：

	* 不想架设服务器，可以使用自带的 http 服务，
	* 可以用它生成所需文件，用自己的静态文件服务器，如 H2O 活着 Nginx Apache 之类
	* 使用 docker 环境

## 特性

最好将文件打成一个 zip 包，每个 group 一个 zip


## 使用 docker 封装

参数
build 构建 更新资源 (执行完后结束)
run host:port 启动 http 服务器 ，作为最终的发布地址

run build 构建更新文件

docker 运行 使用参数决定最后的发布地址

例子：


## TODO

<!-- 生成的说明文件 version 应该为最后一个版本，不是第一个版本 -->

引擎版本可配置，或者自动读取 project.manifest 文件

在 cocosupdate build 生成热更新资源后
cocosupdate start 命令应该可以修改已经发布资源的路径
也即是修改 manifest 里面的 url， 这里应该随时可以改变的

添加效验游戏中 project.manifest 文件版本是否与目录相一致


## 遇到问题
zip 里面必须包含目录名称，并且必须以 "/" 结尾 (很坑，而不是根据文件属性)
否则不能创建目录导致子目录文件不能写入
