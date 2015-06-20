package unionfind

type DynamicConnectivityProblem interface {
   Init(size uint)
   Union(i, j uint)
   Find(i, j uint) bool
   Repr() string
}
