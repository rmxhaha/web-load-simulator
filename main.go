package main

import (
	"fmt"
	"net/http"
	"math"
	"math/rand"
)

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	f := 0.0
	for k := 0; k <= n; k++ {
		f += 4 * math.Pow(-1, float64(k)) / (2*float64(k) + 1)
	}
	return f
}


func cpu(w http.ResponseWriter, r *http.Request){
	n := pi(5000)

	fmt.Fprintf(w, "The answer is %d", n)
}

func memo(w http.ResponseWriter, r *http.Request){
	// 1MB of memory
	a := make([]int32, 250000)

	a[0] = int32(rand.Int())
	a[1] = int32(rand.Int())
	a[2] = int32(rand.Int())

	for i := 3; i < len(a); i++ {
		a[i] = a[i-3] * a[i-2] - a[i-1]
	}

	fmt.Fprintf(w, "The answer is %d", a[len(a)-1])
}

func net(w http.ResponseWriter, r *http.Request){

	for i := 0; i < 1000; i ++ {
		fmt.Fprintf(w, "1234567890")
	}
}

func main() {
	http.HandleFunc("/cpu", cpu)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/net", net)



	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3212", nil)
}