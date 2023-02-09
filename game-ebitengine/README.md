Ebitengine - 2D 游戏引擎使用
--------------------------

[ebiten](https://github.com/hajimehoshi/ebiten)是一个简单的 Go 语言 2D 游戏引擎。

### 运行方式

```shell
// 生成必要的素材
go generate

// 运行
go run .

// 构建
go build . alien_invasion
```

### 参考流程

1. [一起用Go做一个小游戏（上）](https://mp.weixin.qq.com/s/5HfZ2TrnUl2pfBft5-CAJg)
2. [一起用Go做一个小游戏（中）](https://mp.weixin.qq.com/s/UXpekTlUcK6nxKOYGZfP2A)
3. [一起用Go做一个小游戏（下）](https://mp.weixin.qq.com/s/Hw2GFSTY9Sgv2SPgYypreQ)

### 打包资源

[file2byteslice](github.com/hajimehoshi/file2byteslice) 可以将图片、配置文件等文件打包成一个 Go 文件，并以 byte slice 的方式提供程序使用。

使用`go generate`命令和`//go:generate`伪指令可以更方便的完成资源的打包处理。

安装：

```shell
go install -v github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@latest
```

### 网页运行

使用 Go 内置对 WASM 的支持，将游戏编译成 WASM 格式，从而可以在网页上运行该游戏：

```shell
GOOS=js GOARCH=wasm go build -o ./wasm/alien_invasion.wasm

cd wasm
go run main.go
```

然后就可以在网页中访问 `http://localhost:8088/wasm_exec.html` 来运行游戏了。
