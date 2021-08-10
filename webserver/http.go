package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/leolegends/golang-cobra-example/app"
)

// Server is struct
type Server struct {
	Port string
}

//makeData return result.
// func makeData(r http.Request, value string) {

// 	result, err := strconv.ParseFloat(r.URL.Query().Get(value), 64)
// 	if err != nil {
// 		r = 0
// 	}
// 	return result
// }

//SumHandler is func
func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		a = 0
	}

	// a := makeData(r, "a")

	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}

	calc := app.NewCalc()
	calc.A = a
	calc.B = b

	w.Write([]byte(fmt.Sprintf("Result is %f", calc.Sum())))

}

// Serve is bootstrap func
func (s Server) Serve() {
	http.HandleFunc("/", s.SumHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
