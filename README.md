# WS - Product
Product service example for baseline in golang

# Start 🚀
    1. Clone this project -> https://github.com/cristian0193/ws-baseline-golang.git
    2. Make sure port 8080.
    3. go get all 
    4. go run main.go

# Pre-requirements 📋
It is necessary to install -> https://golang.org/ 

# Dependencies 🤝
**IMPORTAT:** All dependencies will register and will be usable all dependency, however all dependency that is not used and is imported they should delete from de mod file

- [github.com/alecthomas/template] This is a fork of Go 1.4's text/template package with one addition: a backslash immediately after a closing delimiter will delete all subsequent newlines until a non-newline.
- [github.com/gin-gonic/gin] Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [github.com/gin-contrib/pprof] Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.
- [github.com/jinzhu/gorm] This is an orm framework and is used to be connected to sql server database
- [golang.org/x/crypto] This is a library for the session key encryption
- [github.com/dgrijalva/jwt-go] This is a framework that is used for make session's token for all user registrered on servicedesk system

# Project Structure 🧱

```
WS_BaseLine_Golang
    ├── application
    ├── business
    │   └── service
    ├── docs
    ├── domain
    │   ├── component
    │   │   └── errorException
    │   ├── dto
    │   │   └── apis
    │   └── entity
    ├── infrastructure
    │   ├── api
    │   ├── persistence
    │   │   └── db
    │   └── repository
    ├── middleware
    │   └── server
    └── utils
        ├── jwt
        └── language
```

# Built with 🛠️
    - Visual Studio Code

# Endpoints
    - GET    /product
    - POST   /product
    - PUT    /product
    - DELETE /product/:idProduct
    - GET    /product/:idProduct
    - GET    /products/parameters

# Authors
Christian Rodriguez
Software Developer