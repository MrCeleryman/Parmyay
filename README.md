# Parmyay
A website dedicated to finding and rating the best chicken parmigianas in your local area

## Stack

Front end
* Vuejs 2.0
* TypeScript
* Webpack 2

API
* Go
* REST

DB
* SQLite

BANNED PACKAGES
* yarn
* lodash

## Setup

### Backend

* Install Postman https://www.getpostman.com/
* Install gcc (windows https://sourceforge.net/projects/mingw-w64/?source=typ_redirect)
* Go to api/ and create 3 folders bin, pkg and src
* Run ``` go get github.com/gin-gonic/gin ```
* Run ``` go get github.com/jinzhu/gorm ```
* Run ``` go get github.com/mattn/go-sqlite3 ```
* Run ``` go install github.com/mattn/go-sqlite3 // need gcc ```
* Run ``` go build ```
* Run ``` .\api.exe ```
* Open up Postman and use it to use the api and stuff

### Front End

## Roadmap

- [x] Initial Setup Front End
- [x] Initial Setup Back End
- [ ] Mockups in Balsamiq
- [ ] DB Schema

![alt text](/img/commit.png?raw=true "Try and beat this commit")
