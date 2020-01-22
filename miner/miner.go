package main

import "fmt"
import "time"
import "crypto/rand"
import "sync"
import "flag"
import "runtime"
import "github.com/shirou/gopsutil/cpu"

import "github.com/deroproject/astrobwt"

func main() {
	fmt.Printf("DERO AstroBWT Miner v0.01 alpha\n")

	info, _ := cpu.Info()
	fmt.Printf("CPU: %s    PhysicalThreads:%d\n", info[0].ModelName, len(info))

	threads_ptr := flag.Int("threads", runtime.NumCPU(), "No. Of threads")
	iterations_ptr := flag.Int("iterations", 100, "No. Of DERO Stereo POW calculated/thread")
	bench_ptr := flag.Bool("bench", true, "run bench with params")
	flag.Parse()

	var wg sync.WaitGroup

	threads := *threads_ptr
	iterations := *iterations_ptr

	if threads < 1 || iterations < 1 || threads > 2048 {
		fmt.Printf("Invalid parameters\n")
		return
	}

	if *bench_ptr {

		fmt.Printf("%20s %20s %20s %20s %20s \n", "Threads", "Total Time", "Total Iterations", "Time/PoW ","Hash Rate/Sec")
		for bench := 1; bench <= threads; bench++ {
			now := time.Now()
			for i := 0; i < bench; i++ {
				wg.Add(1)
				go random_execution(&wg, iterations)
			}
			wg.Wait()
			duration := time.Now().Sub(now)

			fmt.Printf("%20s %20s %20s %20s %20s \n", fmt.Sprintf("%d", bench), fmt.Sprintf("%s", duration), fmt.Sprintf("%d", bench*iterations),
			fmt.Sprintf("%s", duration/time.Duration(bench*iterations)),fmt.Sprintf("%.1f", float32(time.Second)/ (float32(duration/time.Duration(bench*iterations)))) )

		}

	} else {
		fmt.Printf("Starting %d threads\n", threads)
		now := time.Now()
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go random_execution(&wg, iterations)
		}
		wg.Wait()
		duration := time.Now().Sub(now)
		fmt.Printf("Total iterations %d ( per thread %d)\n", threads*iterations, iterations)
		fmt.Printf("Total time %s\n", duration)
		fmt.Printf("time per PoW (avg) %s\n", duration/time.Duration(threads*iterations))
	}

}

func random_execution(wg *sync.WaitGroup, iterations int) {

	var workbuf [255]byte

	for i := 0; i < iterations; i++ {
		rand.Read(workbuf[:])
		//astrobwt.POW(workbuf[:])
		astrobwt.POW_0alloc(workbuf[:])
	}
	wg.Done()
}
