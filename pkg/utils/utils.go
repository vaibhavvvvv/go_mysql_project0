//we recieve request in JSON but we need to unmarshall it to use it.

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

//In Go, an interface is a collection of method signatures.
//A type that implements all the methods in an interface is said to implement that interface.
//Interfaces are used to abstract away the implementation details of a type.
// type Shape interface {
// 	Area() float64
//   }

//   type Rectangle struct {
// 	Width, Height float64
//   }

//   func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
//   }
