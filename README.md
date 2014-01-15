krongen
=======

A Kronecker graph generator in Go

# HOWTO

```
$ go install krongen.go
$ krongen -h
Usage of krongen:
  -edgef=8: n.o. edges = [n.o. vertexes]*edgef
  -scale=10: n.o. vertexes = 2^scale
$ krongen -scale=2 -edgef=2
1	4
4	3
2	4
4	1
1	2
4	3
4	1
4	4
```

Also comes with a generator package that can be used in other projects (TODO:
usage example)

