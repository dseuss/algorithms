package main

import (
   "fmt"
   "dseuss/algorithms/unionfind"
)


func test_problem(problem unionfind.DynamicConnectivityProblem) {
   problem.Init(10)
   problem.Union(4, 3)
   problem.Union(3, 8)
   problem.Union(6, 5)
   problem.Union(9, 4)
   problem.Union(2, 1)

   fmt.Println(problem.Find(3, 9))
   fmt.Println(problem.Find(4, 9))
   fmt.Println(problem.Find(1, 2))
}


func main() {
   var problem unionfind.QuickFind
   test_problem(&problem)
}
