package unionfind

import "fmt"


type QuickFind struct {
   id []uint
}


func (this *QuickFind) Init(size uint) {
   this.id = make([]uint, size)
   for index, _ := range this.id {
      this.id[index] = uint(index)
   }
}


func (this *QuickFind) Union(i, j uint) {
   iid := this.id[i]
   for index, value := range this.id {
      if iid == value {
         this.id[index] = this.id[j]
      }
   }
}


func (this *QuickFind) Find(i, j uint) bool {
   return this.id[i] == this.id[j]
}


func (this *QuickFind) Repr() (result string) {
   for _, value := range this.id {
      result += fmt.Sprint(value, " ")
   }
   return
}
