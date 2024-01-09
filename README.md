- [Program structure](#program-structure)
  - [Variable declaration](#variable-declaration)
  - [Pointers](#pointers)
  - [Variable lifetime](#variable-lifetime)
  - [Tuple assignment](#tuple-assignment)
  - [Type declaration](#type-declaration)
  - [Type conversion](#type-conversion)
  - [Scope](#scope)
- [Basic data types](#basic-data-types)
- [Composite types](#composite-types)
  - [Arrays](#arrays)
  - [Slice](#slice)
  - [Maps](#maps)
  - [Structs](#structs)
  - [struct embedding](#struct-embedding)
- [Function](#function)
  - [Syntax](#syntax)
  - [Recursion](#recursion)
  - [Multiple return values](#multiple-return-values)
  - [Function values](#function-values)
  - [Anonymous function](#anonymous-function)
  - [Caveat, capturing iteration variables](#caveat-capturing-iteration-variables)
  - [Variadic functions](#variadic-functions)
  - [Deferred function calls](#deferred-function-calls)
  - [Panic](#panic)
  - [Recover](#recover)
- [Package](#package)
  - [Import path](#import-path)
  - [Example](#example)
  - [Testing](#testing)
  - [Blank import](#blank-import)
- [Reference](#reference)


## Program structure

### Variable declaration
- Variable declaration syntax is from left to right
```go
var name type = expression
```

```go
func main() {
	var age int = 100
	fmt.Println("age = ", age)
}
```

```go
package main

import "fmt"

func varaiableDeclarationSynatx1() {
	var age int = 100
	fmt.Println("synatx1, age = ", age)
}

func varaiableDeclarationSynatx2() {
	var age = 100
	fmt.Println("synatx2, age = ", age)
}

func varaiableDeclarationSynatx3() {
	age := 100
	fmt.Println("synatx3, age = ", age)
}

func main() {
	varaiableDeclarationSynatx1()
	varaiableDeclarationSynatx2()
	varaiableDeclarationSynatx3()
}

```


### Pointers
```go
func pointer() {
	x := 1
	p := &x
	fmt.Println("*p = ", *p)
	*p = 2
	fmt.Println("*p = ", *p)
	fmt.Println("x = ", x)
}
```

### Variable lifetime
- lifetime of a package level variable is the entire execution of the program
- local variables have dynamic lifetimes

```go
const globalX = 100

func lifetime() {
	localX := 4
	fmt.Println("globalX = ", globalX)
	fmt.Println("localX = ", localX)
}
```

### Tuple assignment
```go
func swap() {
	a := 100
	b := 200
	fmt.Printf("a = %v, b = %v\n", a, b)
	temp := a
	a = b
	b = temp
	fmt.Printf("a = %v, b = %v\n", a, b)
}

func swapWithTuple() {
	a := 100
	b := 200
	fmt.Printf("swapWithTuple a = %v, b = %v\n", a, b)
	b, a = a, b
	fmt.Printf("swapWithTuple a = %v, b = %v\n", a, b)
}
```

### Type declaration
```go
type name underlying-type
```

```go
type Celsius float64
type Farenheit float64

func setFurnaceTemperature(f Farenheit) {
	fmt.Println("furnace temperature set to ", f)
}
```

### Type conversion
```go
type Celsius float64
type Farenheit float64

func main() {
	cel := Celsius(2.0)
	far := Farenheit(cel)
	fmt.Println("furnace temperature set to ", far)
}
```


### Scope
- Lexical scope
- Shadowing

```go
func f() {}

var g = "g"

func main() {
  f := "f"
  fmt.Println(f) //local variable f shadows package level func f
  fmt.Println(g) //package level var
  fmt.Println(h) //compiler error: undefined: h
}
```

```go
func main() {
  x := "hello"
  for i := 0; i < len(x); i++ {
    x := x[i]
    if x != '!' {
      x := x + 'A' - 'a'
      fmt.Printf("%c", x)
    }
  }
}
```

```go
if x := f(); x > 80 {
  fmt.Println("distinction", x)
} else if x > 60 {
    fmt.Println("first class", x)
} else {
    fmt.Println("failed", x)
}

fmt.Println(x) //compiler error
```

## Basic data types
- Integers
- Floating point numbers
- Complex numbers
- Booleans
- Strings
- Constants

## Composite types

### Arrays
```go
package main

import "fmt"

func main() {
	var values [3]int
	fmt.Println("element at index 1, ", values[1])

	for index, value := range values {
		fmt.Printf("element at index %v is %v\n", index, value)
	}
}
```

```go
//array initialization
var q [3]int = [3]int{1,2,3}


r := [...]int{1,2,3}
```

### Slice
- Unlike arrays slices are not comparable
```go
months := [...]string {"", "jan", "feb", "march", "april", "may", "june", "july", "aug", "sept", "octo", "nov", "dec"}

q2 := months[4:7]
summer := months[6:9]

allMonths := months[1:] 
//months[:], 
//months[:7]
```

```go
//create an array
a := [...]int{0,1,2,3,4,5}

//create a slice
s := []int{0,1,2,3,4,5}  //this implicitly creates an array of right size and yields a pointer to it
```

```go
//slices and nil
var s []int     // len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{}     // len(s) == 0, s != nil
```

```go
//create slice using make() method
make([]T, len)
make([]T, len, cap)
```

```go
//use append method to append elements to slice
var runes []rune
for _, r := range "Hello" {
  runes = append(runes, r)
}
```

### Maps
```go
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

```go
ages := map[string]int {
  "alice": 31,
  "charlie": 34
}

fmt.Println(ages["alice"])

for name, age := range ages {
  fmt.Printf("%s: %d\n", name, age)
}
```

### Structs
```go
package main

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func main() {
	e := Employee{
		ID:     1,
		Name:   "Alice",
		Salary: 100,
	}

	e.Salary = 200
}

```

### struct embedding
```go
//problem statement
type Point struct {
  X,Y int
}

type Circle struct {
  Center Point
  Radius int
}

type Wheel struct {
  Circle Circle
  Spokes int
}

var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

```go
//solution, anonymous fields
type Circle struct {
  Point
  Radius int
}

type Wheel struct {
  Circle
  Spokes int
}

var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```


## Function

### Syntax
```go
func hypot(x, y float64) float64 {
  return math.sqrt(x*x + y*y)
} 
```

```go
func hypot(x float64, y float64) float64 {
  return math.sqrt(x*x + y*y)
} 
```

### Recursion


### Multiple return values
```go
resp, err := http.Get(url)
if err != nil {
  return err
}
```

### Function values
- Functions are first class values
- Functions have types
- They may be assigned to variables
- Passed to functions
- Returned from functions

```go
func square(n int) int { return n * n}
func negative(n int) int { return -n}
func product(n, m int) int { return n * m}

f := square
fmt.Println(f(3))  //9

f = negative
fmt.Println(f(3))  //-3

f := product //compiler error
```

### Anonymous function
```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f()) //1
    fmt.Println(f()) //4
    fmt.Println(f()) //9
    fmt.Println(f()) //16
}
```

### Caveat, capturing iteration variables


### Variadic functions

### Deferred function calls
- No exception in go
- No destructor in go
```go
func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ReadAll(f)
}
```

```go
func lookup(key string) int {
    mu.Lock()
    defer mu.Unlock()
    return m[key]
}
```

### Panic

### Recover


## Package
- There is no explict keywords like ```public``` or ```private```
- Exported identifiers start with an upper case letter

### Import path

### Example
```go
//temptype.go

//package declaration
package tempconv

//import declaration
import "fmt"

type Celsius float64
type Farenheit float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC Celsius = 0
  BoilingC Celsius = 100
)

func (c Celsius) String() string {return fmt.Sprintf("%v", c)}
func (f Farenheit) String() string {return fmt.Sprintf("%v", f)}
```

```go
//conv.go
package tempconv
func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32)}
func CToF(f Farenheit) Celsius { return Celsius((f-32) * 5/9)}
```

### Testing
- Test source file must be of the format ```*_test.go```
- Test method must be of the format ```Test*```
```go
//conv_test.go

package tempconv

import "testing"

func TestCToF(t *testing.T) {
    c := Celsius(100)
    if Farenheit(212) != CToF(c) {
        t.Error(`Celsius(100) is not Farenheit(212)`)
    }
}
```

### Blank import

## Reference
- [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Google I/O 2012 - Meet the Go Team](https://www.youtube.com/watch?v=sln-gJaURzk)