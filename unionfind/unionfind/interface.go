package unionfind

type DCP interface {
   Init(size int)
   Union(i, j int)
   Find(i, j int) bool
   Repr() string
}
