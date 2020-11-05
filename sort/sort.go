package sort

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/ihtkas/go-libs/sort/bubble"
	"github.com/ihtkas/go-libs/sort/heap"
	"github.com/ihtkas/go-libs/sort/insertion"
	"github.com/ihtkas/go-libs/sort/merge"
	"github.com/ihtkas/go-libs/sort/quick"
	"github.com/ihtkas/go-libs/sort/radix"
	"github.com/ihtkas/go-libs/sort/selection"
	"github.com/ihtkas/go-libs/sort/shell"
)

//TODO: make this as test file
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	defer func() {
		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			defer f.Close() // error handling omitted for example
			runtime.GC()    // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
		}
	}()
	// arr := []int{29, 16, 23, 42, 14, 46, 50, 6, 18, 56}
	// clone := make([]int, len(arr))
	// copy(clone, arr)
	// quickSort3Way(clone)
	// fmt.Println(clone)
	// if !checkSort(arr, clone) {
	// 	fmt.Println("quickSort3Way failed")
	// 	fmt.Println(arr)
	// 	fmt.Println(clone)
	// }
	// return
	for i := 0; i < 1000; i++ {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		src = rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(src)
		x := r.Intn(1000)
		y := r2.Intn(1000)
		arr := append(r.Perm(x), r.Perm(y)...)
		clone := make([]int, len(arr))
		copy(clone, arr)
		// fmt.Println(arr)

		// fmt.Println(clone)
		// fmt.Println(arr)
		insertion.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("insertionSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		bubble.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("bubleSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		selection.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("selectionSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		quick.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("quickSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		merge.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("mergeSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		heap.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("heapSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		radix.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("radixSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		shell.Sort(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("shellSort failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}

		copy(clone, arr)
		quick.Sort3Way(clone)
		// fmt.Println(clone)
		if !checkSort(arr, clone) {
			fmt.Println("quickSort3Way failed")
			fmt.Println(arr)
			fmt.Println(clone)
		}
	}
}

func checkSort(orig, sorted []int) bool {
	if len(orig) == 0 {
		return len(sorted) == 0
	}
	x := make(map[int]int, len(orig))
	for _, e := range orig {
		x[e]++
	}

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			return false
		}
		x[sorted[i]]--
	}

	x[sorted[len(sorted)-1]]--

	for _, e := range x {
		if e != 0 {
			return false
		}
	}
	return true
}
