## Golang generics turned into FP

In this package I dip my toes in the generics added to golang in 1.18. I really enjoy the fp style of scala and js. However, there are caveats to the golang way.

Lets say you have a generic struct `MyStruct[A]` and you want to turn it into a `MyStruct[B]`. You would define a function like `func (m *MyStruct[A]) [B any]Map(func(A)B) *MyStruct[B]`, but that's not allowed. Functions on generic objects are not allowed to have their own type parameters. What you can get away with is to define `B` as a `B any` in your package, but that leads to uglier code when your API is consumed. As the consumer might have to cast their precious values to `yourpackage.B`. The best example of this can be seen in the `Test_FlatMap` vs `Test_FreestandingFlatMap`.

Instead what you can do, is to define that function, in your package, instead of on the struct (like I've done). Then that function can do anything, at the cost of passing the "container object" as a param to the function.

For all your needs in this department, look at the slice package.

## Todo

On the horizon im looking to add a box package with subpackages for Optional, Operation & perhaps one more.