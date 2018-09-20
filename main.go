package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/draw", handler).Methods("GET")
	if err := http.ListenAndServe(":2003", r); err != nil {
		log.Fatal("server exited with error: ", err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	// TODO get gorilla mux queries
	w.Header().Set("Content-Type", "image/svg+xml")
	drawDO(w, 200, "blue")
}

func drawDO(w io.Writer, size int, color string) {
	canvas := svg.New(w)
	width, height := 500, 500
	canvas.Start(width, height)

	radius := size

	// Outer circle
	canvas.Circle(width/2, height/2, radius, fmt.Sprintf("fill:%s;stroke:%s", color, color))

	// Inner circle
	innerRad := int(0.6 * float64(radius))
	canvas.Circle(width/2, height/2, innerRad, fmt.Sprintf("fill:white;stroke:%s", color))

	// Square to cutout the bottom corner
	canvas.Square((width/2)-radius-1, (height/2)+1, radius, "fill:white;stroke:white")

	// Large bottom square
	Lsize := int(0.4 * float64(radius))
	canvas.Square((width/2)-Lsize, (height/2)+(innerRad-Lsize), Lsize, fmt.Sprintf("fill:%s;stroke:%s", color, color))

	// Medium square
	Msize := int(0.3 * float64(radius))
	canvas.Square((width/2)-(Msize+Lsize), (height/2)+innerRad, Msize, fmt.Sprintf("fill:%s;stroke:%s", color, color))

	// Small square
	Ssize := int(0.2 * float64(radius))
	canvas.Square((width/2)-(Msize+Lsize+Ssize), (height/2)+innerRad-Ssize, Ssize, fmt.Sprintf("fill:%s;stroke:%s", color, color))

	canvas.End()
}
