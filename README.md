# An API Simulator

## Requirements
* docker (if running in a docker instance)
  * docker daemon should be running
  * Either need to use sudo to the commands below, or add user to the docker group
* golang
  * Only needed if running directly on a server or
  * To run automated tests or
  * To compile without running in docker

## Commands for Running in a Docker Container
Tested in Arch Linux and Fedora
* customBin/build can be used to compile the Go code
* customBin/buildAndRun compiles and also runs the compiled code
* customBin/run will just run the already compiled code
* customBin/ssh will SSH into the docker instance
* customBin/stop will stop the docker instance

Accessed by navigating to http://localhost:8080/worldspan locally

## Running on a Server
This was done on a CentOS 6.6 server and is pretty much documentation for myself so I don't forget. Mileage may vary
* Install golang and hg
  * `sudo yum install golang hg`
* As root: `sudo su`
  * Make your directory: `mkdir /opt/go` `mkdir /opt/go/src`
  * Setup your go path: `export GOPATH=/opt/go`
  * (Optional) Add to the path: `PATH=$PATH:$GOPATH/bin`
  * Setup the go bin directory: `export GOBIN=$GOPATH/bin`
  * Clone the repo in the $GOPATH
  * From within the app folder
      * Get necessary dependencies `go get`
      * Compile the code: `go build -o $GOBIN/api-simulator`
  * Run the code: `$GOBIN/api-simulator`
  * To access the code, navigate to: `http://127.0.0.1:8080/worldspan`
    * Or the server IP if it's not local in place of 127.0.0.1
  * In the main-test.go, update the IP if necessary
    * Run automated tests (below) to test to make sure that everything is working

## Daemonizing (init.d)
* Copy the initd file to the appropriate folder
  * `cp $GOPATH/customBin/initd /etc/init.d/api-simulator`
* Add to the init.d list
  * `chkconfig --add api-simulator`

## Compiling
* To just compile (check for errors):
  * `go build *.go`

## Running Automated Tests
* `go test app/*.go`
