# Cocos 游戏资源更新

## 说明

使用 Golang + HTTP/1.1 + Json


## 特性

最好将文件打成一个 zip 包，每个 group 一个 zip


## 使用 docker 封装

参数
build 构建 更新资源 (执行完后结束)
run host:port 启动 http 服务器 ，作为最终的发布地址

run build 构建更新文件

docker 运行 使用参数决定最后的发布地址

例子：


