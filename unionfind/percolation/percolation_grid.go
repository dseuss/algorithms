package percolation


import (
    "errors"
    "fmt"
    "github.com/dseuss/algorithms/unionfind/unionfind"
)


type Indices2d struct {
    I, J int
}


type PercolationGrid struct {
    size int
    topnode, bottomnode int
    site_open [][]bool
    conngraph unionfind.DCP
}


var neighboars = []Indices2d { {0, 1}, {0, -1}, {1, 0}, {-1, 0}, }

func (this* PercolationGrid) Init(size int, dcp unionfind.DCP) {
    this.size = size
    this.site_open = make([][]bool, size)
    for i := range this.site_open {
        this.site_open[i] = make([]bool, size)
    }

    this.conngraph = dcp
    this.conngraph.Init(size * size + 2)
    this.topnode = size * size
    this.bottomnode = size * size + 1

    // initialize the connectivity graph with virtual top/bottom node
    for j := 0; j < this.size; j++ {
        index, err := this.indices2node(0, j)
        if err != nil {
            panic(err.Error())
        }
        this.conngraph.Union(this.topnode, index)

        index, err = this.indices2node(this.size - 1, j)
        if err != nil {
            panic(err.Error())
        }
        this.conngraph.Union(this.bottomnode, index)
    }
}


func (this* PercolationGrid) indices2node(i, j int) (int, error) {
    if (i < 0) || (i >= this.size) {
        return -1, errors.New(fmt.Sprintf("Index i=%v out of bounds", i))
    }
    if (j < 0) || (j >= this.size) {
        return -1, errors.New(fmt.Sprintf("Index j=%v out of bounds", j))
    }

    return this.size * i + j, nil
}


func (this* PercolationGrid) node2indices(node int) (i, j int, err error) {
    if (node < 0) || (node >= this.size * this.size) {
        return -1, -1, errors.New(fmt.Sprintf("Node index=%v out of bounds", node))
    }

    j = node % this.size
    i = (node - j) / this.size
    return i, j, nil
}



func (this* PercolationGrid) Open(i, j int) {
    index, err := this.indices2node(i, j)
    if err != nil {
        panic(err.Error())
    }
    this.site_open[i][j] = true

    for _, nb := range neighboars {
        index_nb, err := this.indices2node(i + nb.I, j + nb.J)
        if (err == nil) && this.IsOpen(i + nb.I, j + nb.J) {
            this.conngraph.Union(index, index_nb)
        }

    }
}


func (this *PercolationGrid) IsOpen(i, j int) bool {
    return this.site_open[i][j]
}


func (this *PercolationGrid) IsFull(i, j int) bool {
    index, err := this.indices2node(i, j)
    if err != nil {
        panic(err.Error())
    }
    return this.conngraph.Find(this.topnode, index)
}


func (this* PercolationGrid) Percolates() bool {
    return this.conngraph.Find(this.topnode, this.bottomnode)
}


// TODO This should be done cleaner
func (this* PercolationGrid) Repr() (repr string) {
    for _, row := range this.site_open {
        for _, is_open := range row {
            if is_open {
                repr += "X "
            } else {
                repr += "  "
            }
        }
        repr += "\n"
    }
    return repr
}
