package kronecker

import (
    "testing"
    "math/rand"
    "fmt"
)


func TestGeneratorCreation(t *testing.T) {
    scale, edgeFactor := 10, 8
    if gen := CreateGenerator(scale, edgeFactor); gen == nil {
        t.Error("generator could not be created")
    }
}

func TestGenerator_Yield(t *testing.T) {
    scale, edgeFactor := 10, 8
    gen := CreateGenerator(scale, edgeFactor)
    A, B, C := 0.57, 0.19, 0.19

    rand.Seed(66)
    tmp := YieldEdge(scale, A, B, C)
    from, to := tmp[0], tmp[1]
    rand.Seed(66)

    // yield an edge
    e, _ := gen.Yield()
    if e[0] != from || e[1] != to {
        t.Errorf("output was %d %d, expected %d %d", e[0], e[1], from, to)
    }
}

func TestGenerator_Yield_halts(t *testing.T) {
    scale, edgeFactor := 3, 5

    halt := (1 << uint(scale))*edgeFactor
    gen := CreateGenerator(scale, edgeFactor)
    fmt.Println(halt, gen.edges)

    i := 0
    for {
        if i > halt { t.Error("generator produces edges indefinitely"); break }
        i ++

        _, err := gen.Yield()
        if err != nil { break }
    }
}
