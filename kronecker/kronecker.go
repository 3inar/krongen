package kronecker

import (
    "math/rand"
)

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

        if rand.Float64() > ab {
            from2 = 1
        }
        if rand.Float64() > (c_norm*float64(from2) + a_norm*float64((from2+1)%2)) {
            to2 = 1
        }

        from = from + coeff*from2
        to = to + coeff*to2
    }

    return [2]int{from+1, to+1}
}
