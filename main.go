package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	file, err := os.Open("./AABLNCRYJG.txt")
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("MemAlloc (before): %v bytes\n", memStats.Alloc)

	var fileSize int
	if fileStat, err := file.Stat(); err == nil {
		fileSize = int(fileStat.Size())
	}
	fmt.Printf("Size of 'file' variable: %v bytes\n", fileSize)

	runtime.ReadMemStats(&memStats)
	fmt.Printf("MemAlloc (after): %v bytes\n", memStats.Alloc)

	//========================================================

	// Profilisanje memorije pre otvaranja datoteke
	profileBefore := pprof.Lookup("heap")
	profileBefore.WriteTo(os.Stdout, 0)

	// Profilisanje memorije nakon otvaranja datoteke
	profileAfter := pprof.Lookup("heap")
	profileAfter.WriteTo(os.Stdout, 0)
}
