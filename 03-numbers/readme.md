# Dealing with Numbers
The recipes in this section deal with numbers and floating points.
* printing
* comparing
* changing base representations

### Recipe 1 convert string to numbers
This can easily achieved using strconv package it has a numebr of useful util finctions
* **strconv.Atoi()** convert string to integer, this helps convert base 10 values.
* **strconv.ParseInt(str,base,byteSize)** - this can help convert binary, hex value to integers
* **strconv.ParseFloat(str, byteSize)** - this will convert string to floats.

### Recipe 2 comparing floating point numbers [important]
The floating point number can be compared using the classical == operator but it will not give you right result.
The floating point comparison is always better when using tolerance value or using the big package there to you need to use the precision values for equality

```
const TOLERANCE = 1e-8   // equivalent to a precision of 8

 func equals(a, b float64) bool{
 	delta := math.Abs(a-b)
 	if delta < TOLERANCE{
 		return true
 	}
 	return false
 }
```

We can also use the math/big package to help with float comparison.
```
prec := 16

daB := big.NewFloat(da).SetPrec(prec)
dbB := big.NewFloat(db).SetPrec(prec)
fmt.Printf("Comparing big floats with precision: %d: %v\n",prec , daB.Cmp(dbB) == 0)

```

### Recipe 3 Rounding floating point numbers
The Round function is not present in the math package in Go 1.9.2 but has been included in Go 1.10.
So for Rounding float values in 1.9.2 we can write a util function like:

```
    func Round(val float64) float64{
    	t := math.Trunc(val)
    	if math.Abs(val-t) >= 0.5{
    		return t + math.Copysign(1,val)
    	}
    	return t
    }
```

* **math.Trunc()** will only get the decimal value for the float value.
* **math.CopySign(1,val)** - will copy value 1 with the same sign as the float value is added to the float number.


### Recipe 4 Floating point arithmetics
In the case of normal calculations float64 can handle upto 64 precision values which is fine in most cases. But if you need more precision than 64
we need to use math/big package. Generally if we need to use a very high precision on the value of Pi.


### Recipe 5 Formatting floating and integer numbers

```
         fmt.Printf("%d \n", integer)
        // if we need to show the sign of the integer
        fmt.Printf("%+d \n", integer)

        // to print other bases x - base 16, o-8, b-2, d-10
        fmt.Printf("%X \n", integer)
        fmt.Printf("%#X \n", integer)

        // padding with leading zeros
        fmt.Printf("%010d \n", integer)
        // padding with leading spaces - left
        fmt.Printf("% 10d \n", integer)
        // padding with spaces - right
        fmt.Printf("% -10d \n", integer)

        // floating point precision of 5
        fmt.Printf("%.5f \n", floatNum)
        // floating point represented as scientific notation
        fmt.Printf("%e \n", floatNum)
```

### Recipe 6 Converting the number between various bases
strconv package is for this very purpose. The strconv has a number of functions that can be used to changes bases for the numbers.

### Recipe 8 Generating random numbers
The most popular package for this is the math/rand but if we use a particular seed the number could repeat. In order to get
cryptographically random number we need to use the crypto/rand

