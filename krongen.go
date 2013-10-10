/*
    This is an implementation of a kronecker edge list generator v. closely
    based on the graph500 octave example (http://graph500.org)
 */
package main

import "fmt"
import "math"
import "math/rand"
import "time"
import "runtime"
import "flag"


func create_edge(scale int, A, B, C float64) []int {
    from := 0
    to := 0

    ab := A + B
    c_norm := C/(1 - ab)
    a_norm := A/ab

    for ib := 1; ib <= scale; ib++ {
        coeff := int(math.Exp2(float64(ib - 1)))
        from2 := 0
        to2 := 0

        if rand.Float64() > ab {
            from2 = 1
        }
        if rand.Float64() > (c_norm*float64(from2) + a_norm*float64((from2+1)%2)) {
            to2 = 1
        }

        from = from + coeff*from2
        to = to + coeff*to2
    }
    edge := make([]int,2)
    edge[0], edge[1] = from+1, to+1
    return edge
}

func k_generator(scale, edgefactor int) {
    //fmt.Println("ey", scale, edgefactor)

    // The number of vertices is 2^scale
    N := int(math.Exp2(float64(scale)))

    // N.o. edges
    M := edgefactor*N

    // initiator probabilities
    A, B, C := 0.57, 0.19, 0.19

    results := make(chan []int)
    dudes := 0
    for i := 0; i < 100; i++ {
        dudes++
        go func () {
            for {
                edge := create_edge(scale, A, B, C)
                results <- edge
            }
        } ()
    }

    for i:=0;i<M;i++ {
        tst := <-results

        // permute edge labels
        //perm := rand.Perm(N)
        //tst[0] = perm[tst[0] - 1]
        //tst[1] = perm[tst[1] - 1]

        fmt.Printf("%d\t%d\n", tst[0], tst[1])
    }

}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    rand.Seed(time.Now().UTC().UnixNano())
    var scale = flag.Int("scale", 10, "n.o. vertexes = 2^scale")
    var edgefactor = flag.Int("edgef", 8, "n.o. edges = [n.o. vertexes]*edgef")


    flag.Parse()

    //fmt.Println(*scale, *edgefactor)
    k_generator(*scale, *edgefactor)
}
/*
    the below is old rubbish code that I might need later
 */
    //fmt.Println(A, B, C, M)
    //ij := make([]int, 2*M)
    //for i := range(ij) { ij[i] = 1 }

    //ab := A + B
    //c_norm := C/(1 - ab)
    //a_norm := A/ab

    //fmt.Println(ab, c_norm, a_norm)
    //for ib := 1; ib <= scale; ib++ {
    //    bits := make([]int, 2*M)
    //    for i := 0; i < M; i++ {
    //        bit := 0
    //        if rand.Float64() > ab {
    //            bit = 1
    //        }

    //        notbit := (bit + 1) % 2

    //        other_bit := 0
    //        if rand.Float64() > (c_norm*float64(bit) + a_norm*float64(notbit)) {
    //            other_bit = 1
    //        }

    //        bits[i] = bit
    //        bits[M + i] = other_bit

    //    }

    //    for i := range(bits) {
    //        coeff := int(math.Exp2(float64(ib - 1)))
    //        ij[i] = ij[i] + coeff*bits[i]
    //    }
    //}
