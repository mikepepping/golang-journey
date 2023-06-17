# Package Handling

Every Go program is made up of packages.
Programs start running in the package `main`.

You can declare which package your code is in with the `package` statement.
Example, this code is inside in the main package.
```go
package main

func main() {
    return
}
```

## Importing Packages

You can import a package using the `import "<package"` statement for a single pacakage. Alternatively you can use the factored import statement `import ("package1", "package2")` for multiple imports.

Example, this code imports `fmt` and `rand` packages.

```go
package main

import (
    "fmt",
    "math/rand"
)

func main() {
    fmt.PrintLn("This is a random number ->", rand.int(10))
}

```

## Exporting

If you want to export a name, it must be in capitals.
For example the `math` package exports `Pi`.
If you tried call `math.pi` it would not be found, as lowercase names are not exported.

There is no other mechanism to export a name from a package, such as an `export` statement like in JavaScript.

Example

```go
package cool_new_package

func notExported() {
}

func ThisIsAlwaysExported() {
}

```

