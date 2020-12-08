// Can create methods on all types not just structs

// can call methods from pointers
// https://tour.golang.org/methods/7
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK

// null receiver
//https://tour.golang.org/methods/12
/*
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
Note that an interface value that holds a nil concrete value is itself non-nil. 
*/
