package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "net/http"
    "strings"

    // Third party code for routing
    "github.com/gorilla/mux"
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
    fmt.Fprintf(response, "Root directory of the API simulator\n")
}

func ExampleId(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    identifier := providedVars["id"]
    fmt.Fprintln(response, "Provided ID: ", identifier)
}

func Worldspan(response http.ResponseWriter, request *http.Request) {
    // @TODO Should be using os.PathSeparator
    dataPath := os.Getenv("GOPATH") + "/data/";

    body, error := ioutil.ReadAll(io.LimitReader(request.Body, 1048576));
    if (error != nil) {
        log.Fatal(error);
    }
    if (strings.Contains(string(body), "<PSC5>")) {
        // @TODO Ideally, we wouldn't be using the full path here
        // @TODO Also, couldn't figure out how to break up the line below
        HandleWorldspanRequest(dataPath + "testPowerShopperResponse", response);
    } else if (strings.Contains(string(body), "<BPC9>")) {
        HandleWorldspanRequest(dataPath + "testPricingResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>CK/")) {
        HandleWorldspanRequest(dataPath + "testCardAuthorization", response);
    } else if (strings.Contains(string(body), "<UPC7>")) {
        HandleWorldspanRequest(dataPath + "testUpdatePnrResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>*")) {
        HandleWorldspanRequest(dataPath + "testNativeDisplayPnrResponse", response);
    } else if (strings.Contains(string(body), "<HOS_CMD>EZEI#$*")) {
        HandleWorldspanRequest(dataPath + "testTicketingResponse", response);
    } else if (strings.Contains(string(body), "<HOS_RSP_SCR>F</HOS_RSP_SCR>")) {
        HandleWorldspanRequest(dataPath + "testFinished", response);
    } else if (strings.Contains(string(body), "<DPC8>")) {
        HandleWorldspanRequest(dataPath + "testDisplayPnrResponse", response);
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
