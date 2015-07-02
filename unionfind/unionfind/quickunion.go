package unionfind

import "fmt"


type QuickUnion struct {
   id []int
}


func (this *QuickUnion) Init(size int) {
   this.id = make([]int, size)
   for index, _ := range this.id {
      this.id[index] = int(index)
   }
}


func (this *QuickUnion) root(i int) int {
   for (i != this.id[i]) {
      i = this.id[i]
   }
   return i
}


func (this *QuickUnion) Union(i, j int) {
   root_i, root_j := this.root(i), this.root(j)
   this.id[root_i] = root_j
}


func (this *QuickUnion) Find(i, j int) bool {
   return this.root(i) == this.root(j)
}


func (this *QuickUnion) Repr() (result string) {
   for _, value := range this.id {
      result += fmt.Sprint(value, " ")
   }
   return
}
