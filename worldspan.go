package main

import (
    "fmt"
    "net/http"
);

type worldspanConnection struct {
    // @TODO This should be in a more generic class
    request *http.Request;
    response http.ResponseWriter;
}

func (self *worldspanConnection) testing() {
    fmt.Fprintln(self.response, "Type of request not found!!!!");
}

func (self *worldspanConnection) respond() {
}
