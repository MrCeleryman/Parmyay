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
	* `GOROOT = INSTALL LOCATION`
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

* `node build\build.js` in root folder restores packages and build api into `dist/api.exe`
* Run `api.exe` and api will run on port 8900
* `node build\test.js` in root folder runs tests and coverage and outputs to html coverage in browser

## Roadmap

- [x] Initial Setup Front End
- [x] Initial Setup Back End
- [x] DB Schema
- [x] Setup linting
- [x] Setup Unit Testing
- [x] Code Coverage Setup
- [ ] Code Coverage 100% currently at 98.8% 


## Elixir Api

To start your Phoenix app:

  * Install [elixir](http://elixir-lang.org/install.html)
  * Install dependencies with `mix deps.get`
  * 

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.
