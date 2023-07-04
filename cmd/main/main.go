//to create the server, tell Go where the routers are,

package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vaibhavvvvv/goMysql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) //handles routes
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
