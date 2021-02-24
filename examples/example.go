package main

import (
	"github.com/ChristianStefaniw/cgr"
	"math"
	"net/http"
	"strconv"
)

func main() {
	r := cgr.NewRouter()
	squareConf := cgr.NewRouteConf()

	// Configuration will be passed to all routes
	r.SkipClean(true)

	// Configuration will be passed to the route it is assigned to
	squareConf.AppendSlash(true)

	r.Route("/").Method("GET").Handler(home)
	r.Route("/square/:num/").SetConf(squareConf).Method("GET").Handler(square).AppendSlash(true)
	r.Route("/routes").Method("GET").Handler(func(writer http.ResponseWriter, request *http.Request) {
		for _, route := range r.ViewRouteTree(){
			writer.Write([]byte(route))
		}
	})
	helloRoute := r.Route("/hello/:name/").Handler(hello).Method("GET")

	// Configure route after declaration
	helloRoute.AppendSlash(true)

	cgr.Run("8000", r)
}

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("home"))
	if err != nil {
		panic("error")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := cgr.GetVars(r)["name"]
	_, err := w.Write([]byte("Hello " + name))
	if err != nil {
		panic("error")
	}
}

func square(w http.ResponseWriter, r *http.Request) {
	num, _ := strconv.ParseFloat(cgr.GetVars(r)["num"], 32)
	pow := strconv.FormatFloat(math.Pow(num, 2), 'f', -1, 32)
	_, err := w.Write([]byte(pow))
	if err != nil {
		panic("error")
	}
}