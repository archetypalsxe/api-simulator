package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
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
    body, error := ioutil.ReadAll(io.LimitReader(self.request.Body, 1048576));
    if (error != nil) {
        log.Fatal(error);
    }

    // @TODO Should be using os.PathSeparator
    dataPath := os.Getenv("GOPATH") + "/data/";

    if (strings.Contains(string(body), "<PSC5>")) {
        // @TODO Also, couldn't figure out how to break up the line below
        self.HandleWorldspanRequest(dataPath + "testPowerShopperResponse");
    } else if (strings.Contains(string(body), "<BPC9>")) {
        self.HandleWorldspanRequest(dataPath + "testPricingResponse");
    } else if (strings.Contains(string(body), "<HOS_CMD>CK/")) {
        self.HandleWorldspanRequest(dataPath + "testCardAuthorization");
    } else if (strings.Contains(string(body), "<UPC7>")) {
        self.HandleWorldspanRequest(dataPath + "testUpdatePnrResponse");
    } else if (strings.Contains(string(body), "<HOS_CMD>*")) {
        self.HandleWorldspanRequest(dataPath + "testNativeDisplayPnrResponse");
    } else if (strings.Contains(string(body), "<HOS_CMD>EZEI#$*")) {
        self.HandleWorldspanRequest(dataPath + "testTicketingResponse");
    } else if (strings.Contains(string(body), "<HOS_RSP_SCR>F</HOS_RSP_SCR>")) {
        self.HandleWorldspanRequest(dataPath + "testFinished");
    } else if (strings.Contains(string(body), "<DPC8>")) {
        self.HandleWorldspanRequest(dataPath + "testDisplayPnrResponse");
    } else {
        fmt.Fprintln(self.response, "Type of request not found");
    }
}

func (self *worldspanConnection) HandleWorldspanRequest(responseFile string) {
    data, error := ioutil.ReadFile(responseFile);
    if (error != nil) {
        log.Fatal(error);
    }
    fmt.Fprintln(self.response, string(data));
}
