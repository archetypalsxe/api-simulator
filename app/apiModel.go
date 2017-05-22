package main

import (
    //"strconv"
);

type apiModel struct {
    id int
    name string
    beginningEscape string
    endingEscape string
}

func (self *apiModel) getOutput() string {
    return "testing!"
    /*
    return strconv.Itoa(self.id) + " " + self.name + " " +
        self.beginningEscape + " " + self.endingEscape)
        */
}
