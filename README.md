# An API Simulator

## Commands for Running in a Docker Container
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
  6. Compile the code: `go install main.go`
  7. Run the code: `bin/main.go`
  8. To access the code, navigate to: `http://127.0.0.1:8080/worldspan`

## Compiling
* To just compile (check for errors):
  * `go build *.go`
