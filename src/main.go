package main

import (
	"flag"
	"runtime"
	"fmt"
)


var (
	num_cpus int
)

func initCmdLineParams()  {
	flag.IntVar(&num_cpus, "cpus", runtime.NumCPU(), "")
	flag.Parse()
}

func init() {
	initCmdLineParams()
}

func main() {
	fmt.Println(num_cpus)
}
