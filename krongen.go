/*
    This is an implementation of a kronecker edge list generator v. closely
    based on the graph500 octave example (http://graph500.org)
 */
package main

import (
    "fmt"
    "krongen/kronecker"
    "math"
    "math/rand"
    "time"
    "runtime"
    "flag"
)



func parallel_generator(scale, edgefactor int) {
    // The number of vertices is 2^scale
    N := int(math.Exp2(float64(scale)))

    // N.o. edges
    M := edgefactor*N

    // initiator probabilities
    A, B, C := 0.57, 0.19, 0.19

    results := make(chan [2]int)
    dudes := 0
    for i := 0; i < 100; i++ {
        dudes++
        go func () {
            for {
                edge := kronecker.YieldEdge(scale, A, B, C)
                results <- edge
            }
        } ()
    }

    // permutation of edge labels, the kronecker generator has a greater
    // probability of creating edges in the lower ints
    perm := rand.Perm(N)
    for i:=0;i<M;i++ {
        tst := <-results

        // edges from the generator are 1-indexed,
        // those of the permutation are 0-indexed
        tst[0] = perm[tst[0] - 1] + 1
        tst[1] = perm[tst[1] - 1] + 1

        fmt.Printf("%d\t%d\n", tst[0], tst[1])
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    rand.Seed(time.Now().UTC().UnixNano())

    var scale = flag.Int("scale", 10, "n.o. vertexes = 2^scale")
    var edgefactor = flag.Int("edgef", 8, "n.o. edges = [n.o. vertexes]*edgef")
    flag.Parse()

    parallel_generator(*scale, *edgefactor)
}
