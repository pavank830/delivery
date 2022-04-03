package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// basic api resp codes
const (
	ResponseOK     = 0
	ResponseFailed = 1
)

// ParseRequest -- func to read and parse the http request/input
func ParseRequest(w http.ResponseWriter, r *http.Request, input interface{}) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body %+v\n", err)
		return err
	}
	if err := r.Body.Close(); err != nil {
		log.Printf("error closing body %s\n", err.Error())
		return err
	}
	if err := json.Unmarshal(body, input); err != nil {
		err = fmt.Errorf("Unmarshalling Error. %+v ", err)
		log.Printf("%v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte(err.Error()))
		w.WriteHeader(422)
		return err
	}
	return nil
}
