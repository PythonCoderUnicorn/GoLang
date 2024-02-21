

```
# Go Lang

sudo apt install golang-go 


touch file.go 
go build file.go 
```

# Go Lang

documentation
https://gobyexample.com/

online code editor !
https://go.dev/tour/welcome/1

https://go.dev/learn/


go install golang.org/x/website/tour@latest

Security with Go 
https://www.packtpub.com/product/security-with-go/9781788627917

https://github.com/packtpublishing/security-with-go


Black Hat Go 
https://nostarch.com/blackhatgo




full stack Go 
https://github.com/PacktPublishing/Full-Stack-Web-Development-with-Go

mastering Go 
https://github.com/PacktPublishing/Mastering-Go-Second-Edition




Derek Banas
https://www.youtube.com/watch?v=YzLrWHZa-Kc

https://github.com/derekbanas/Go-Tutorial/blob/main/hellogo.go



https://www.youtube.com/watch?v=S65k7Tubbck


https://www.udemy.com/course/golang-hacking/

https://github.com/parsiya/Hacking-with-Go


https://pkg.go.dev/hash     
Go pkg Hash

https://www.youtube.com/@boldlygo/videos



---

# Go Lang

- statically typed language
- strongly typed 
- compiled language 
- fast compile time
- built in concurrency  (goroutines)
- simplicity  (garbage collection)

every Go program is made of packages
programs start running in package main 

```Go
package main

import (        // must have it like this!  otherwise --> ("fmt";"math")
	"fmt"
	"math"
)

func main(){
	n := 17.0
	sqrt := math.Sqrt(n)
	fmt.Printf("sqrt of %v = %g \n", n , sqrt)
}
// sqrt of 17 = 4.123105625617661

```

exported names start with Capital letters `Pizza` , `pizza` is not exported
```Go
func main() {
	fmt.Println( math.pi ) // error!
	fmt.Println( math.Pi ) // works
}
```

### Functions

can take 0 or more arguments, `func something( [varName] [type] ) [type] { }`

```Go
package main
import "fmt"

func add(x int, y int) int {
	return x + y
}

// can also do add(x,y int) int {}

func main() {
	a := add(42, 13)
	fmt.Println("add(42, 13) = ", a)
}
// add(43, 13) = 55
```

multiple results
- a function can return any number of results
- `swap` function returns 2 strings

```Go
package main
import "fmt"

func swap(x,y string) (string, string) {
	return y,x
}

func main() {
	a,b := swap("hello", "world")
	fmt.Println("a,b swap = ", a,b)
}
// a,b swap = world hello
```

In Go when you use just `return` in a function it is called _naked return_ and should be used for only small functions
```Go
package main
import "fmt"

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return  // naked return
}

func main() {
	s := split(17)
	fmt.Println( split(s) ) // can't have "split = ", split(s)
}
// 7 10
```


### Variables  

- `var <name> <type>`
- `var c,python,ruby bool`

Variables with initializers
- `var i,j int = 3,6`
- `var c, ruby, java = true, false, "no!" `  no type needed

short variable declarations
- outside a function you need `var` 
- inside a function  `x := 8` , `bot,chat,machine := true,false,"rage" `

variables declared without a value default value is `0` , `false` or `""` for strings

`const Pi = 3.14`  const can be character, string, boolean or numeric

###  data types

```Go
//---- min - max ------------- Unsigned integers == positive only numbers
uint8   0   255  
uint16  0   65535
uint32  0   4294967295
uint64  0   18446744073709551615

//------ min ---------- max ---- signed integers
int8    -128            127
int16   -32768          32767
int32   -2147483648     2147483647
int64   -9.2e17         9223372036854775807
```

Type conversions `T(v)` T _type_ , v = value
```Go
func main() {
	var f float64 = 6786.8787
	var i int = int(f)
	fmt.Printf("%v", i)
}
// 6786
```

Type inference 
- ` var i int`  `j := i // j is also int`

```Go
func main() {
	v := 3.229
	fmt.Printf("v is type %T \n", v)
}
// v is type float64
```


### for loop

`for [init] [condition] [statement]`
- will stop when boolean condition = false
- no ( ) used just { }
- there is no `while loop`

```Go
package main
import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
// 45

//-- init and statements are optional
sum := 1
for ; sum < 1000; {
	sum += sum
}
fmt.Println(sum)  // 1024

//---- "while" loop
sum := 1
for sum < 1000 {
	sum += sum
}
fmt.Println(sum)  // 1024

//-- forever loop
func main() {
	for {
	}
}
```

### If

if statements don't use ( )

```Go
package main
import "fmt"

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
// 1.4142135623730951 2i
```

short `if` statement
```Go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
// pow(3, 2, 10)  = 9  20
```

if else
```Go
package main
import ("fmt" "math")

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	} else {
		fmt.Printf(" %g >= %g \n", v, lim)
	}
	return lim  // can't use v here
}

func main() {
	fmt.Println(
		pow(3,2, 10),
		pow(3,3, 20)
	)
}
// 27 >= 20
// 9 20
```


### switch 

switch statements are shorter than if -else statements, code runs on first case value is equal to condition, the `break` is not needed and switch cases can be strings (not forced to have numbers).

```Go
package main
import ("fmt" "runtime")

func main() {
	fmt.Print("Go runs on ")
	
	os := runtime.GOOS;
	
	switch os {
	case "darwin":
		fmt.Println(" OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf(" %s \n", os)
	}
}
// Go runs on Linux
```

```Go
package main
import ("fmt"; "time")

var pl = fmt.Println

func main() {
	fmt.Println("when's saturday?")
	today := time.Now().Weekday()    // Tuesday
	saturday := time.Saturday
	
	switch sat {
	case today + 0:
		pl("today!")
	case today + 1:
		pl("tomorrow")
	case today + 2:
		pl("in 2 days")
	default:
		pl(">3 days away")
	}
}
// >3 days away
```

switch with no condition
```Go
package main
import ("fmt"; "time")

var pl = fmt.Println

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		pl("morning")
	case t.Hour() < 17:
		pl("afternoon")
	default:
		pl("evening")
	}
}
```


### defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```Go
package main
import "fmt"

var pl = fmt.Println

func main() {
	defer pl("i am deferred statement")

	pl("i execute 1st!")
}
// i execute 1st!
// i am deferred statement
```

stacking defers
```Go
package main
import "fmt"

var pl = fmt.Println

func main() {
	pl("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Print( i )
	}

	pl("done")
}

// counting
// done
// 9876543210
```


## pointers

Go has pointers, which holds the memory address of a value
- type `*T` is pointer to a value, zero value is `nil`
- `var p *int` 
- `i := 42` and the & generates a pointer to its operand `p =&i`
- the `*` is the pointer's value `*p`  (dereferencing)

```Go
package main
import "fmt"

var pl = fmt.Println

func main() {
	i := 42
	p := &i // points to i
	pl("pointer *p := &i ", *p)  // 42

	*p = 288
	pl("*p & i ", *p, i) // 288 & 288

	p = &i
	*p = *p / 37    // 288/37
	pl(*p, p, i)
}
// pointer *p := &i    42
// *p & i              288 288
// *p, p, i            7 0xc000012028 7
```

## structs

a struct is a collection of fields
```
type [Name] struct {
	[var name] type
	[var name] type
}
```

struct basics
```Go
package main
import "fmt"

var pl = fmt.Println

type Vertex struct {
	X int
	Y int
}

func main() {
	pl(Vertex{1, 2})
}
// {1,2}
```

struct fields
```Go
package main
import "fmt"

var pl = fmt.Println

type Vertex struct {
	X int
	Y int
}

func main() {
	// struct literal V{X: 1, Y: 2} | V{X:1}  Y:0 implicit | V{}  ={0,0}
	v := Vertex{1,2}    
	v.X = 4
	v.Y = 35
	pl( v.X, v.Y )      // or pl( v )
	//-- pointer
	p := &v
	p.X = 1e9
	pl( p.X )
}
// 4 35
// 1000000000
```

## arrays

syntax is `[n values]Type` 
- `var b [10]int`  b is an array of 10 integers , size is fixed

```Go
package main
import "fmt"

func main() {
	var s [2]string   // array s of type string and length of 2
	s[0] = "array[0] "
	s[1] = "array[1] "
	fmt.Println( s[0], s[1] )
	fmt.Println( s )

	primes := [6]int{2,3,5,7,9,11}
	fmt.Println( primes[1:4] )     // slice[ low : high ]
	var sl = primes[1:4]        //  []int = primes[1:4]
	fmt.Println( sl )
}
// array[0]  array[1]
// [array[0] array[1]]
// [3 5 7]
// [3 5 7]
```

```Go
package main
import "fmt"

func main() {
	names := [4]string { "python","ruby","javascript","perl" }
	names[3] = "c++"
	fmt.Prinln( names )
}
// [python ruby javascript c++]
```

## slice literals / array literal
```Go
package main
import "fmt"

func main() {
	L := []int{4,5,6,7,9}
	fmt.Println( L )
}
// [4 5 6 7 8 9]
```

slice defaults
```Go
s := []int{3,4,5,2,1,6}

s[1:4]
s[:5]
s[:] // prints everything in array

// array length
len(s)  // 6
```

slices with make()  which creates dynamically sized arrays
- `m := make( []int, [length], [capacity] `

```Go
package main
import "fmt"

func main() {
	m := make( []int, 5, 7)  // len 5 cap 7 [0 0 0 0 0]
}
```

slices of slices
```Go
package main
import "fmt"

func main() {
	// tic-tac board
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	fmt.Println( board )

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// [ [_ _ _] [_ _ _] [_ _ _] ]
/*

X _ X
O _ X
_ _ O

*/
```


append to a slice
- `append( [var], [type], [type values] )`

```Go
var s []int           // []
s = append(s,1)       // [1]
s = append(s,2,3,4)   // [1 2 3 4]
```

slice and range
```Go
var pow = []int{1,2,3,4,5}

func main() {
	for i,v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
// 2**0 =1
// 2**1 =2
// 2**2 =4 ...
```

range shorthand
```Go
for i, _ := range variable
for _, v := range variable
```

## maps

a `map` maps keys to values

```Go
package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	v := Vertex{10, 12}    

	m = make( map[string]Vertex )  // need to keep as is =>   map[]

	m["Bell_Labs"] = Vertex { 40.68, -74.39}
	fmt.Println( m )
	fmt.Println( m["Bell_Labs"] )
	fmt.Println( m["Bell_Labs"].Lat )
}

// {10,12}
// map[]
// map[Bell_Labs:{40.68 -74.39}]
// {40.68 -74.39}
// 40.68
```

map literals
```Go
package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs": { 40.68433, -74.39967 },
	"Google": { 37.42202, -122.08408   }
}

func main(){
	fmt.Println( m )
}
// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
```

mutating map
```
m := make( map[string]int )
m[key] = value
delete(m, key)
ok := m[key]
fmt.Println("value exist = ", ok) // false
```


## function values

```Go
package main
import "fmt"

func modFunc(a, b int) int {
	return a % b
}

func main() {
	mod := modFunc        // assign to variable
	result := mod(8,3)
	fmt.Println( result )
}
```

function closures
```Go
package main
import "fmt"


func main() {
 	counter := createCounter()

    // call the closure multiple times to see the count increment
    fmt.Println(counter())  // Output: 1
    fmt.Println(counter())  // Output: 2
    fmt.Println(counter())  // Output: 3
}

func createCounter() func() int {
    count := 0  // variable captured by the closure

    return func() int {
        count++  // increment the captured count variable
        return count
    }
}
```

Go lang example
```Go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
/*
0 0
1 -2
3 -6
6 -12 ...
*/
```

## methods

Go <span style="color:#ffc000">does not have Class</span> data types but does have <span style="color:#26fd3f">methods</span> on types
- a method is a function with a special receiver argument

style 1
```
func (v Vertex) Abs() float64 {
	vx := v.X * v.X
	vy := v.Y * v.Y
	return math.Sqrt(vx + vy) 
}
```

style 2
```
func Abs( v Vertex ) float64 {
	vx := v.X * v.X
	vy := v.Y * v.Y
	return math.Sqrt(vx + vy) 
}
```

```Go
package main
import ("fmt"; "math")

//-- struct  Vertex{X: float64, Y: float64}
type Vertex struct { X, Y float64 }

//-- method Abs() belongs to Vertex struct
// (v Vertex) is method receiver
func Abs( v Vertex ) float64 {
	vx := v.X * v.X
	vy := v.Y * v.Y
	return math.Sqrt(vx + vy) 
}

func main() {
	v := Vertex{3, 4}   
	// 3*3 + 4*4 -> sqrt(25) = 5
	fmt.Println("v.abs() = ", v.Abs())
}
```

pointer receivers
```Go
package main
import ("fmt";"math")

type Vertex struct { X, Y float64 }

// can't change order here
func (v Vertex) Abs() float64 {
	vx := v.X * v.X
	vy := v.Y * v.Y
	return math.Sqrt(vx + vy)
}

// *Vertex is the pointer     -- if you rm *Vertex == 5
// scale() belongs vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println( v )          // {3 4}
	
	fmt.Println( v.Abs() )    // 5
	
	v.Scale(10)
	fmt.Println( v.Abs() )   // 50  | 5 if you rm *Vertex
}

```

choosing a value or pointer



























