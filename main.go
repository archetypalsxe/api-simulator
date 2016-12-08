package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"

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
    // @TODO Ideally, we wouldn't be using the full path here
    data, error := ioutil.ReadFile(
        "/go/src/worldspan-simulator/data/testPowerShopperResponse"
    );
    if (error != nil) {
        log.Fatal(error);
    }
    fmt.Fprintln(response, string(data));
}
