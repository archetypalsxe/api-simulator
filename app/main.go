package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    // Third party code for routing
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/", Index)
    http.HandleFunc("/settings", Settings)
    http.HandleFunc("/example/{id}", ExampleId)
    http.HandleFunc("/worldspan", Worldspan)
    http.Handle("/css/", http.StripPrefix("/css/",
        http.FileServer(http.Dir("../css"))))
    log.Fatal(http.ListenAndServe(":8080", nil))


    /*
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/settings", Settings)
    router.HandleFunc("/example/{id}", ExampleId)
    router.HandleFunc("/worldspan", Worldspan)

    fileServer := http.FileServer(http.Dir("/go/src/api-simulator/css"))
    router.Handle("/css/", http.StripPrefix("/css/", fileServer))
    /*
    router.Handle("/css/", http.StripPrefix("/css/",
        http.FileServer(http.Dir("../css"))))
        */
        /*
    router := mux.NewRouter().StrictSlash(true)
    router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("/go/src/api-simulator/css"))))
    */
    //log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/go/src/api-simulator/css"))))
    //log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Root directory of the API simulator\n")
}

func Settings(response http.ResponseWriter, request *http.Request) {
    settingsPage := settingsPage{response: response}
    settingsPage.respond()
}

func ExampleId(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    identifier := strings.TrimSpace(providedVars["id"])
    fmt.Fprintln(response, "Provided ID: ", identifier)
}

func Worldspan(response http.ResponseWriter, request *http.Request) {
    worldspan := worldspanConnection{response: response, request: request}
    worldspan.respond()
}

