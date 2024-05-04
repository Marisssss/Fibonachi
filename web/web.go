package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"main.go/logic"
)

func render(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("GG")
		return
	}
	fmt.Fprintf(w, "%s", file)
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "./web/index.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	render(w, r, "./web/result.html")

	fs := logic.FiService{}

	number, err := strconv.Atoi(r.FormValue("numberValue"))

	if err != nil || number < 0 {

		fmt.Fprintf(w, "<h2>Error</h2>")

	} else {

		fmt.Fprintf(w, "<a>The entered number: %d\n</a>", number)

		if fs.IsFi(number) {
			prev, next := fs.ValuesAroundFi(number)
			fmt.Fprintf(w, "<a>The previous Fibonacci number: %d</a>", prev)
			fmt.Fprintf(w, "<a>The next Fibonacci number: %d</a>", next)
		} else {
			closest := fs.PreviousFi(number)
			fmt.Fprintf(w, "<a>The nearest Fibonacci number: %d</a>", closest)
		}

	}

	fmt.Fprintf(w, "<a href=\"/\">To return</a>")
}

func StartServer() {
	http.HandleFunc("/", serveHandler)
	http.HandleFunc("/result", formHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
