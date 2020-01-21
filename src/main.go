package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


type point struct {
    x, y int
}

func main() {
	html := `<!DOCTYPE html>
    <html>
        <head>
            <title>Example</title>
        </head>
        <body>
            <p>This is an example of a simple HTML page with one paragraph.</p>
            <p>Our addtion value is %f.</p>
            <p>Our addtion type is %T.</p>
            <p>Our point value is %v.</p>
            <p>Our point type is %T.</p>
        </body>
    </html>`
	addition := 2 + 2.2

    p := point{1,2}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, fmt.Sprintf(html, addition, addition, p, p))
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
