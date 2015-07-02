package main

import (
    "fmt"
    "runtime"
    "math/rand"
    "github.com/dseuss/algorithms/unionfind/percolation"
    "github.com/dseuss/algorithms/unionfind/unionfind"
)


func generate_indices(gridsize int) []percolation.Indices2d {
    indices := make([]percolation.Indices2d, gridsize * gridsize)
    perm := rand.Perm(gridsize * gridsize)

    for i := 0; i < gridsize; i++ {
        for j := 0; j < gridsize; j++ {
            index := perm[i * gridsize + j]
            indices[index].I = i
            indices[index].J = j
        }
    }

    return indices
}


func run_experiment(gridsize int) int {
    var grid percolation.PercolationGrid
    grid.Init(gridsize, &unionfind.QuickFind{})

    for n, index := range generate_indices(gridsize) {
        grid.Open(index.I, index.J)
        if grid.Percolates() {
            return n
        }
    }

    panic("Experiment failed.")
    return -1
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    const gridsize = 200
    const trials = 10

    critical_points := make(chan int, trials)
    rand.Seed(0)

    for i := 0; i < trials; i++ {
        go func() { critical_points <- run_experiment(gridsize) }()
    }

    for i := 0; i < trials; i++ {
        fmt.Println(<-critical_points)
    }
}
