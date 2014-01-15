package kronecker

import (
    "math/rand"
    "io"
)

type generator struct {
    scale, vertexes, edges, produced  int
    a, b, c float64
}

func CreateGenerator(scale, edgeFactor int) *generator {
    g := new(generator)
    g.scale = scale
    g.a, g.b, g.c = 0.57, 0.19, 0.19

    g.vertexes = 1 << uint(scale)
    g.edges = g.vertexes*edgeFactor
    g.produced = 0

    return g
}

func YieldEdge(scale int, A, B, C float64) [2]int {
    from := 0
    to := 0

    ab := A + B
    c_norm := C/(1 - ab)
    a_norm := A/ab

    for ib := 0; ib < scale; ib++ {
        coeff := 1 << uint(ib)

        from2 := 0
        to2 := 0

        if rand.Float64() > ab { from2 = 1 }

        if rand.Float64() > (c_norm*float64(from2) + a_norm*float64((from2+1)%2)) {
            to2 = 1
        }

        from = from + coeff*from2
        to = to + coeff*to2
    }

    return [2]int{from+1, to+1}
}

func (gen *generator) Yield() ([2]int, error) {
    e := YieldEdge(gen.scale, gen.a, gen.b, gen.c)
    err := error(nil)
    if gen.produced == gen.edges { err = io.EOF }

    gen.produced++
    return e, err
}
