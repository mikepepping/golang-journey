# Variables and Types

## Variable creation

Variables can only use Letter, Numbers and Underscores.
i.e a-Z, 0-9, _

Cannot start with a Number.

```go
var variable_name <type> // declares a zero initialized variable
var other_variable <type> = <value> // declares and assigns with a explicit type
var infered_variable = <value> // declares and assigns with infered type
infered_shorhand := <value> // declares and assigns with infered type
```

## Types

### Unsigned Integers
 - `uint8` -> `uint64`
 - `uint` (machine dependant i.e 32bit/64bit depending on machine/os)

### Signed Integers
 - `int8` -> `int64`
 - `int` (machine dependant i.e 32bit/64bit depending on machine/os)

### Floats
 - `float32`
 - `float64`

### Strings
 - `"double qoutes"` - interprets escape characters
 - `` `back ticks` `` - ignores escape characters

### Runes
 - `'s'` - single quotes are for single characers (runes)

### Booleans
 - `true`
 - `false`

### Exported Variables
Covered also in package handling, if a variable is capitialized it will automatically be exported from the package.
