# Structs

Structs are custom types in Go.
Go does not have classes, and no inheritance (Yay!)

## Defining a struct

Structs follow the same export rules as anything else.
A struct with a Capital letter will be automaticaly exported, the same for the fields inside.

```go
type User struct {
    Name    string
    Email   string
    Active  bool
    Age     unt8
}
```

## Using instantiating a Struct

Structs are instantiated similar to C syntax, where you name the struct then supply all fields inside following squirly braces.

```go
var new_user User = User{ "Name", "email@email.com", true, 35 }
```

## Methods

Given our user struct, we can attach a method to is by adding the `(name Type)` statement into our function definition, between `func` and `MethodName()`.

Example:

```go
func (u User) PrintStatus() {
    fmt.Printf("User is active?", u.Active)
}
```

We can then call the method on a user instance like below:
```go
var new_user User = User{ "Name", "email@email.com", true, 35 }

new_user.PrintStatus()
```

### Methods pass by pointer or by value

When you attach a method to a Struct, unless you express the struct instance as a pointer, the instance will be passed to the method by value.

If passed by value you will no be able to update the original, only the copy given. To alter the state of the original, you must pass by pointer by expressing the type is a "pointer type" using the syntax `*Type`, in our case `(u *User)` instead of `(u User)`.

Example

```go
func (u *User) Deactivate() {
    u.Active = false
}
```

## Learning Resources used
1.  Hitesh Choudhary
    - structs - https://www.youtube.com/watch?v=NMTN543WVQY
    - methods - https://www.youtube.com/watch?v=GhYIKwMxz_Y

