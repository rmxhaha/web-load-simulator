package main

import (
	"fmt"
	"net/http"
	"math"
	"math/rand"
	"bufio"
	"os"
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


func cpuLoad() string {
	return fmt.Sprintf("The answer is %d", pi(10000))
}

func memoryLoad() string {
	// 1MB of memory
	a := make([]int32, 250000)

	a[0] = int32(rand.Int())
	a[1] = int32(rand.Int())
	a[2] = int32(rand.Int())

	for i := 3; i < len(a); i++ {
		a[i] = a[i-3] * a[i-2] - a[i-1]
	}

	return fmt.Sprintf("The answer is %d", a[len(a)-1])
}

func netLoad() string {
	str := ""
	for i := 0; i < 1000; i ++ {
		str = fmt.Sprintf( "1234567890")
	}
	return str
}

func fsLoad() string {
	f, err := os.Create(fmt.Sprintf("somefile%d", rand.Int31()))
	if err != nil {
		return err.Error()
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < 1000; i ++ {
		w.Write([]byte("1234567890"))
	}

	w.Flush()
	return "OK"
}

func noLoad() string {
	return "nothing to do here"
}

func cpu(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, cpuLoad())
}

func memo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, memoryLoad())
}

func net(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, netLoad())
}

func disk(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, fsLoad())
}

func none(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, noLoad())
}

func ScenarioFactory(funcs []func()string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, funcs[rand.Intn(len(funcs))]())
	}
}



func main() {
	http.HandleFunc("/cpu", cpu)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/net", net)
	http.HandleFunc("/disk", disk)
	http.HandleFunc("/none", none)



	scenarios := [][]func()string {
		{noLoad},
		{cpuLoad},
		{memoryLoad},
		{netLoad},
		{fsLoad},
		{cpuLoad,noLoad},
		{memoryLoad,noLoad},
		{netLoad,noLoad},
		{fsLoad,noLoad},
		{cpuLoad,memoryLoad,netLoad,fsLoad},
	}

	for i := 0; i < len(scenarios); i ++ {
		http.HandleFunc(fmt.Sprintf("/s%d", i), ScenarioFactory(scenarios[i]))
	}


	http.ListenAndServe(":3212", nil)
}
