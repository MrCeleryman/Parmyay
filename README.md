# :poultry_leg: Parmyay :poultry_leg:

An API dedicated to finding and rating the best chicken parmigianas in your local area
* [ParmyayWeb](https://github.com/MrCeleryman/ParmyayWeb)
* [ParmyayMobile](https://github.com/MrCeleryman/ParmyayMobile)

## Stack

API
* Go
* REST

DB
* SQLite

## Setup

### Prerequisites

* Install the latest version of [Go](https://golang.org/dl/)
* Set the following System Environment Variables
	* `GOROOT = INSTALL LOCATION/bin`
	* `GOPATH = Parmyay/api`
	* `GOBIN = Parmyay/api/bin`
* Install [Postman](https://www.getpostman.com/)
* Install gcc ([windows](https://sourceforge.net/projects/mingw-w64/?source=typ_redirect))

### Backend

* NodeJS is used for cross platform enviroment variable management
* Run `cd api && go get` to build the Go package
	* The executable will be built in src/src.exe
	* A convenience script `node build\build.js` will bootstrap the GO enviroment variables and build the go package, then put the output into `dist/api.exe`
* Run `cd api && go test ` to run the Go test suite
	* A convenience script `node build\test.js` will bootstrap the GO enviroment variables and run the test suite.
* To get go linting and testing to work you may need to install the golang official packages golint and gotests.
vscode asks to do this for you when you have the Go extension and start editing a Go file.
It installs these packages (gocode gopkgs go-outline go-symbols guru gorename godef goreturns golint gotests)
* Open up Postman and use it to use the api and stuff

## Usage

* `go build` in `api\` builds api and creates api.exe
* `go get` in `api\` gets packages requires and builds exe in bin/
* `go test` in `api\src\parmyay` runs unit tests
* `go test -cover` in `api\src\parmyay` runs unit tests and coverage
* `go tool cover -html=coverage` in `api\src\parmyay` runs unit tests, coverage and launches web browser to view coverage report

## Roadmap

- [x] Initial Setup Front End
- [x] Initial Setup Back End
- [x] DB Schema
- [x] Setup linting
- [x] Setup Unit Testing
- [ ] Setup Dev Server
