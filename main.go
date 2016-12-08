package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"

    // Third party code for routing
    "gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/example/{id}", ExampleId)
    router.HandleFunc("/worldspan", Worldspan)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Root directory of the API simulator")
}

func ExampleId(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    identifier := providedVars["id"]
    fmt.Fprintln(response, "Provided ID: ", identifier)
}

func Worldspan(response http.ResponseWriter, request *http.Request) {
    body, error := ioutil.ReadAll(io.LimitReader(request.Body, 1048576));
    if (error != nil) {
        log.Fatal(error);
    }
    if (strings.Contains(string(body), "<PSC5>")) {
        HandlePowerShopperRequest(response);
    } else {
        fmt.Fprintln(response, "Type of request not found");
    }
}

func HandlePowerShopperRequest(response http.ResponseWriter) {
    // @TODO Ideally, we wouldn't be using the full path here
    // @TODO Also, couldn't figure out how to break up the line below
    data, error := ioutil.ReadFile("/go/src/worldspan-simulator/data/testPowerShopperResponse");
    if (error != nil) {
        log.Fatal(error);
    }
    fmt.Fprintln(response, string(data));
}
