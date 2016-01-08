# Cocos 游戏资源热更新

使用 Golang 开发的 cocos 热更新服务处理程序

## 特性

* 能够根据不同版本的资源生成相关热更新所需要的资源
* 自带轻量级 http 服务器

## 使用例子

当前目录包含资源目录 assets (默认，可改)，里面包含各个版本的资源，资源可以通过如 android 打包后从项目 assets 目录获取，每次打包版本不同，放入不同版本目录。

```
# tree -L 3 命令

─── assets
    ├── 1.0.0
    │   ├── main.jsc
    │   ├── project.json
    │   ├── res
    │   ├── script
    │   └── src
    ├── 1.0.1
    │   ├── main.jsc
    │   ├── project.json
    │   ├── res
    │   ├── script
    │   └── src
    └── 1.0.2
        ├── main.jsc
        ├── project.json
        ├── res
        ├── script
        └── src
```

上面是当前待处理资源情况，使用如下命令创建如 google (默认名称 res )渠道的热更新资源：

```
cocosupdate build --addr 192.168.1.51:8001 -n google

# 打印如下
开始生成发布资源 ...
资源目录:  assets
发布目录:  publish
发布地址:  192.168.1.51:8001
发布名称:  google
引擎版本:  3.7.1
发布完成:  http://192.168.1.51:8001/google

```

发布完成后，将会在 publish 目录生成相应的热更新资源，并且对每个版本变化添加自动打包功能：

```
─── publish
    └── google
        ├── 1.0.0-1.0.1.zip
        ├── 1.0.1-1.0.2.zip
        ├── project.manifest
        └── version.manifest
```

project.manifest 文件内容如下：

```
{
  "packageUrl": "http://192.168.1.51:8001/google",
  "remoteManifestUrl": "http://192.168.1.51:8001/google/project.manifest",
  "remoteVersionUrl": "http://192.168.1.51:8001/google/version.manifest",
  "version": "1.0.0",
  "groupVersions": {
    "1": "1.0.1",
    "2": "1.0.2"
  },
  "engineVersion": "3.7.1",
  "assets": {
    "update1": {
      "path": "1.0.0-1.0.1.zip",
      "md5": "2ad5664341debc7ae6a1327540eb7c3c",
      "compressed": true,
      "group": "1"
    },
    "update2": {
      "path": "1.0.1-1.0.2.zip",
      "md5": "5504341ade966fbe8db4255a67a7d6a9",
      "compressed": true,
      "group": "2"
    }
  },
  "searchPaths": []
}
```

## 部署

只需要将 publish 目录部署到服务器即可，注意发布资源的服务器地址要和上面生成的地址相同，你可以使用自己喜欢的服务器，如 Nginx 等，这里也提供了轻量的部署命令以供测试之用：

```
# cocosupdate start -p 8001
启动 Http 服务 path: publish port: 8001...

```

使用如上命令，可以在本机启动一个 http 服务。


## Docker


