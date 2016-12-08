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

/**
    @TODO
    We should be breaking a bunch of these functions into separate classes,
    especially Worldspan handler
*/

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
        // @TODO Ideally, we wouldn't be using the full path here
        // @TODO Also, couldn't figure out how to break up the line below
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testPowerShopperResponse", response);
    } else if (strings.Contains(string(body), "<BPC9>")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testPricingResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>CK/")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testCardAuthorization", response);
    } else if (strings.Contains(string(body), "<UPC7>")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testUpdatePnrResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>*")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testNativeDisplayPnrResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>EZEI#$*")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testTicketingResponse", response);
    } else if (strings.Contains(string(body), "<HOS_RSP_SCR>F</HOS_RSP_SCR>")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testFinished", response);
    } else if (strings.Contains(string(body), "<DPC8>")) {
        HandleWorldspanRequest("/go/src/worldspan-simulator/data/testDisplayPnrResponse", response);
    } else {
        fmt.Fprintln(response, "Type of request not found");
    }
}

func HandleWorldspanRequest(responseFile string, response http.ResponseWriter) {
    data, error := ioutil.ReadFile(responseFile);
    if (error != nil) {
        log.Fatal(error);
    }
    fmt.Fprintln(response, string(data));
}
