package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

func worker(id int, jobs <-chan idProcess, results chan <- tmp) {
	for j := range jobs {
		j.Work()
		results <- tmp{j.id}
		fmt.Print(j.id)
	}
}

type Workable interface {
	Work()
}


type tmp struct {
	id string `json:id,omitempty`
}

type idProcess struct {
	id string
}

func (idProcess *idProcess) Work() {
	idProcess.id+="100"
}

var jobs = make(chan idProcess, 10)
var results = make(chan tmp, 10)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jobs <- idProcess{params["id"]}
}

func RunHTTPServer(addr string) error {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{id}", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
	return http.ListenAndServe(addr, nil)
}

func main() {

	nWorkers := 3

	for w := 1; w <= nWorkers; w++ {
		go worker(w, jobs, results)
	}

	RunHTTPServer(":8080")

}
