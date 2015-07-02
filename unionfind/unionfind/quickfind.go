package unionfind

import "fmt"


type QuickFind struct {
   id []int
}


func (this *QuickFind) Init(size int) {
   this.id = make([]int, size)
   for index, _ := range this.id {
      this.id[index] = int(index)
   }
}


func (this *QuickFind) Union(i, j int) {
   iid := this.id[i]
   for index, value := range this.id {
      if iid == value {
         this.id[index] = this.id[j]
      }
   }
}


func (this *QuickFind) Find(i, j int) bool {
   return this.id[i] == this.id[j]
}


func (this *QuickFind) Repr() (result string) {
   for _, value := range this.id {
      result += fmt.Sprint(value, " ")
   }
   return
}
