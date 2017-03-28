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

* Install the latest version of [Go](https://golang.org/dl/)
* Install [Postman](https://www.getpostman.com/)
* Install gcc ([windows](https://sourceforge.net/projects/mingw-w64/?source=typ_redirect))
* Run ` npm install ` to get all the packages required to run the following scripts
* Run ` npm run build ` or ` npm run dev ` To setup, install and build go packages 
* The executable will be built in /api/bin/api.exe
* Run ` go build ` for subsequent builds 
(you may need to set your GOPATH and GOBIN system environment variables to /api and api/bin respectively)
See [here](https://github.com/golang/go/wiki/InstallTroubleshooting) for help with getting Go to work
* To get go linting and testing to work you may need to install the golang official packages golint and gotests.
vscode asks to do this for you when you have the Go extension and start editing a Go file.
It installs these packages (gocode gopkgs go-outline go-symbols guru gorename godef goreturns golint gotests)
* Open up Postman and use it to use the api and stuff

### Front End

* Run ` npm install ` to get all the packages required to run the following scripts
* Run ` npm run build ` for the production build
* Run ` npm run dev ` To debug and run on the dev server
* Run ` npm run lint ` To run eslint

## Roadmap

- [x] Initial Setup Front End
- [x] Initial Setup Back End
- [x] DB Schema
- [x] Setup linting
- [ ] Mockups in Balsamiq
- [ ] Setup Unit Testing
 
![alt text](/img/commit.png?raw=true "Try and beat this commit")
