# An API Simulator

## Requirements
* docker (if running in a docker instance)
  * docker daemon should be running
  * Either need to use sudo to the commands below, or add user to the docker group
* golang
  * Only needed if running directly on a server or
  * To run automated tests  or
  * To compile without running in docker

## Commands for Running in a Docker Container
Tested in Arch Linux and Fedora
* bin/build can be used to compile the Go code
* bin/buildAndRun compiles and also runs the compiled code
* bin/run will just run the already compiled code
* bin/ssh will SSH into the docker instance
* bin/stop will stop the docker instance

Accessed by navigating to http://localhost:6060/worldspan locally

## Running on a Server
This was done on a CentOS 6.6 server and is pretty much documentation for myself so I don't forget. Mileage may vary
1. git clone the repo somewhere on the server
  1. Quite a bit can be deleted from the repo as it won't be necessary, especially the docker stuff
2. `sudo yum install golang hg`
3. As root: `sudo su`
  1. Make your directory: `mkdir /home/httpd/htdocs/vendor/go`
  2. Setup your go path: `export GOPATH=/home/httpd/htdocs/vendor/go`
  3. (Optional) Add to the path: `PATH=$PATH:$GOPATH/bin`
  4. Setup the go bin directory: `export GOBIN=$GOPATH/bin`
  5. Get necessary dependencies `go get`
  6. Compile the code: `go install`
  7. Run the code: `$GOBIN/api-simulator`
  8. To access the code, navigate to: `http://127.0.0.1:8080/worldspan`
    * Or the server IP if it's not local in place of 127.0.0.1
  9. If not already in the same folder, move the data folder into the $GOPATH
    * `mv data $GOBIN`
  10. Run `go test` to test to make sure that everything is working

## Compiling
* To just compile (check for errors):
  * `go build *.go`

## Running Automated Tests
* `go test`
