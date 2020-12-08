// Can create methods on all types not just structs

// can call methods from pointers
// https://tour.golang.org/methods/7
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
