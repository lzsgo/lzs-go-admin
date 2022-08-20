
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

### 1.2 Contributing Guide

Hi! Thank you for choosing lzs-go-admin.

lzs-go-admin is a full-stack (frontend and backend separation) framework for developers, designers and product managers.

We are excited that you are interested in contributing to lzs-go-admin. Before submitting your contribution though, please make sure to take a moment and read through the following guidelines.

#### 1.2.1 Issue Guidelines

- Issues are exclusively for bug reports, feature requests and design-related topics. Other questions may be closed directly. If any questions come up when you are using Element, please hit [Gitter](https://gitter.im/element-en/Lobby) for help.

- Before submitting an issue, please check if similar problems have already been issued.

#### 1.2.2 Pull Request Guidelines

- Fork this repository to your own account. Do not create branches here.

- Commit info should be formatted as `[File Name]: Info about commit.` (e.g. `README.md: Fix xxx bug`)

- <font color=red>Make sure PRs are created to `develop` branch instead of `master` branch.</font>

- If your PR fixes a bug, please provide a description about the related bug.

- Merging a PR takes two maintainers: one approves the changes after reviewing, and then the other reviews and merges.

## 2. Getting started

```
- node version > v8.6.0
- golang version >= v1.14
- IDE recommendation: Goland
- initialization project: different versions of the database are not initialized. See synonyms at initialization https://www.lzs-go-admin.com/docs/first
- Replace the Qiniuyun public key, private key, warehouse name and default url address in the project to avoid data confusion in the test file.
```

### 2.1 server project

use `Goland` And other editing tools，open server catalogue，You can't open it. `lzs-go-admin` root directory

```bash
# clone the project
git clone https://github.com/flipped-aurora/lzs-go-admin.git

# open server catalogue
cd server

# use go mod And install the go dependency package
go generate

# Compile 
go build -o server main.go (windows the compile command is go build -o server.exe main.go )

# Run binary
./server (windows The run command is server.exe)
```

### 2.1 web project

```bash
# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 2.2 Server

```bash
# using go.mod

# install go modules
go list (go mod tidy)

# build the server
go build
```

### 2.3 API docs auto-generation using swagger

#### 2.3.1 install swagger 

##### (1) Using VPN or outside mainland China
````
go get -u github.com/swaggo/swag/cmd/swag
````

##### (2) In mainland China

In mainland China, access to go.org/x is prohibited，we recommend [goproxy.io](https://goproxy.io/zh/) or [goproxy.cn](https://goproxy.cn)

````bash
# If you are using a version of Go 1.13 - 1.15 Need to set up manually GO111MODULE=on, The opening mode is as follows, If your Go version is 1.16 ~ Latest edition You can ignore the following step one
# Step one、Enable Go Modules Function
go env -w GO111MODULE=on 
# Step two、Configuration GOPROXY Environment variable
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# If you dislike trouble,You can use the go generate Automatically execute code before compilation, But this can't be used command line terminal of `Goland` or `Vscode` 
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
````

#### 2.3.2 API docs generation

````
cd server
swag init
````

> After executing the above command，server directory will appear in the docs folder `docs.go`, `swagger.json`, `swagger.yaml` Three file updates，After starting the go service, type in the browser [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) You can view swagger document


## 3. Technical selection

- Frontend: using [Element](https://github.com/ElemeFE/element) based on [Vue](https://vuejs.org)，to code the page.
- Backend: using [Gin](https://gin-gonic.com/) to quickly build basic RESTful API. [Gin](https://gin-gonic.com/)is a web framework written in Go (Golang).
- DB: `MySql`(5.6.44)，using [gorm](http://gorm.io)` to implement data manipulation, added support for SQLite databases.
- Cache: using `Redis` to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.
- API: using Swagger to auto generate APIs docs。
- Config: using [fsnotify](https://github.com/fsnotify/fsnotify) and [viper](https://github.com/spf13/viper) to implement `yaml` config file。
- Log: using [zap](https://github.com/uber-go/zap) record logs。

## 4. Project Architecture

### 4.1 Architecture Diagram

![Architecture diagram](http://qmplusimg.henrongyi.top/gva/lzs-go-admin.png)

### 4.2 Front-end Detailed Design Diagram (Contributor: <a href="https://github.com/baobeisuper">baobeisuper</a>)

![Front-end Detailed Design Diagram](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 Project Layout

```
    ├── server
        ├── api             (api entrance)
        │   └── v1          (v1 version interface)
        ├── config          (configuration package)
        ├── core            (core document)
        ├── docs            (swagger document directory)
        ├── global          (global object)                    
        ├── initialize      (initialization)                        
        │   └── internal    (initialize internal function)                            
        ├── middleware      (middleware layer)                        
        ├── model           (model layer)                    
        │   ├── request     (input parameter structure)                        
        │   └── response    (out-of-parameter structure)                            
        ├── packfile        (static file packaging)                        
        ├── resource        (static resource folder)                        
        │   ├── excel       (excel import and export default path)                        
        │   ├── page        (form generator)                        
        │   └── template    (template)                            
        ├── router          (routing layer)                    
        ├── service         (service layer)                    
        ├── source          (source layer)                    
        └── utils           (tool kit)                    
            ├── timer       (timer interface encapsulation)                        
            └── upload      (oss interface encapsulation)  
            
    └─web            （frontend）
        ├─public        （deploy templates）
        └─src           （source code）
            ├─api       （frontend APIs）
            ├─assets	（static files）
            ├─components（components）
            ├─router	（frontend routers）
            ├─store     （vuex state management）
            ├─style     （common styles）
            ├─utils     （frontend common utilitie）
            └─view      （pages）

```

