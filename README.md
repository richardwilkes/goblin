# goblin

goblin is an embeddable scripting engine written in Go.

## Notes

To regenerate `parser.go`, you need to have `goyacc` installed:

`go get -u golang.org/x/tools/cmd/goyacc/...`

## Syntax

### Comments
    # Single line, commented to end of line
    // Single line, commented to end of line
    /*
     * Multi-line comment.
     */

### Variable declaration and assignment

Note that while the descriptions below mention specific types, variables can generally
be converted transparently from one type to another, possibly with some lossiness.

    a = 1                    // Numeric 64-bit integer value
    b = 0xaf                 // Numeric 64-bit integer value
    c = 1.2                  // Numeric 64-bit float value
    d = 1.2e7                // Numeric 64-bit float value
    e = true                 // Boolean value
    f = "hello"              // String value
    g = [ 1, 2, 3 ]          // Slice containing numeric 64-bit integer values
    h = [ 1.2, 3.4 ]         // Slice containing numeric 64-bit float values
    i = [ 1, 1.2, "hello" ]  // Slice containing a mixture of values
    j = { "foo": "bar" }     // Map with key "foo"
    k = { "a": 1, "b": 2.1 } // Map with keys "a" and "b"
    l = nil                  // The nil value

### Variable types

    typeOf(1)             // int64
    typeOf(1.2)           // float64
    typeOf(true)          // bool
    typeOf([1,2])         // []interface{}
    typeOf({"foo": "bar") // map[string]interface{}
    typeOf("foo")         // string
    typeOf(nil)           // <nil>

### Operators

    =      // Assignment
    ==     // Equal
    !      // Not
    !=     // Not equal
    >      // Greater than
    >=     // Greater than or equal to
    <      // Less than
    <=     // Less than or equal to
    +      // Add
    +=     // Add and assign
    -      // Subtract
    -=     // Subtract and assign
    *      // Multiply
    *=     // Multiply and assign
    /      // Divide
    /=     // Divide and assign
    **     // Power
    %      // Remainder
    %=     // Remainder and assign
    |      // Bitwise OR
    ||     // Logical OR
    &      // Bitwise AND
    &&     // Logical AND
    >>     // Shift right
    <<     // Shift left
    ^      // Bitwise XOR
    ?      // Ternary operator, as in a == b ? c : d

### Loops

    for i = 0; i < 10; i++ {
        // C-style for loop from 0-9
    }

    for i in [1,3,5] {
        // Iterate through a slice
    }

    for i in range(0,9) {
        // Loop from 0-9
    }

    for i < 5 {
        // Loop until i >= 5
    }

    for {
        // Loop forever or until break/return
    }

### Switch

    switch x {
        case 1:
            // Do something on x == 1
        case 2:
            // Do something on x == 2
        case bar():
            // Do soemthing if x == bar()
        default:
            // Do something if no other case was hit
    }

### If

    if x == 1 {
        // Do something on x == 1
    } else if x == 2 {
        // Do something on x == 2
    } else if x == 3 {
        // Do something on x == 3
    } else {
        // Do something if none of the previous comparisons were true
    }

### Access a slice

    a[2:4]  // Return a new slice with the elements at index 2 and 3 of the original slice
    a[:4]   // Return a new slice with the elements at indexes 0-3 of the original slice
    a[2:]   // Return a new slice with the elements at index 2 and later of the original slice
    a[:]    // Return a new slice with the full contents of the original
    a[2]    // Return the element at index 2

### Define a function

    func a() {
        // Do something
    }

    func b(c) {
        // Do something... has an arg (c) that can be used
    }

    func d() {
        return 12   // Returns the value 12
    }

### Built-ins

    len(x)           // Returns the length of the slice, map or string
    keys(x)          // Returns the keys of a map
    range(limit)     // Returns a slice of int64 starting at 0 and going through limit - 1
    range(min, max)  // Returns a slice of int64 starting at min and going through max
    sort(a, func(i, j) { return a[i] < a[j] }) // Sorts an array. The second parameter is a function that should test for the element at index i being less than the element at index j.
    sleep("1s")      // Sleep for a duration, such as "100ms" for 100 milliseconds, "1s" for one second, "1m" for one minute, etc. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", and "h".
    toString(x)      // Converts the value to a string
    toInt(x)         // Converts the value to an int64
    toFloat(x)       // Converts the value to a float64
    toBool(x)        // Converts the value to a bool
    toChar(x)        // Converts the value to a single-character string
    toRune(x)        // Converts the value to a rune
    toByteSlice(x)   // Converts the value to a []byte
    toRuneSlice(x)   // Converts the value to a []rune
    toBoolSlice(x)   // Converts the value to a []bool
    toIntSlice(x)    // Converts the value to a []int64
    toFloatSlice(x)  // Converts the value to a []float64
    toStringSlice(x) // Converts the value to a []string
    typeOf(x)        // Returns the name of the type x
    defined(x)       // Returns true if the variable has been defined
    print(x...)      // Same as Go's fmt.Print
    println(x...)    // Same as Go's fmt.Println
    printf(x...)     // Same as Go's fmt.Printf
