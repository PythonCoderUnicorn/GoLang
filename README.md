# GoLang

a repo for my Go Language learning

---

### Documentation

https://go.dev/doc/tutorial/getting-started

### Install Go 

```
sudo apt install golang-go 
```


### start a project

1. `mkdir folder; cd folder`
2. `go mod init folder/file`  this gives a pseudo link for local Go `.mod` to reference. Name it the same as the folder your code will be in with main or hello.go
3. `touch hello.go`
4. `go run . `
5. for an executable file `go build hello.go` which compiles and makes `hello` file => `./hello`




### import external package

check out packages https://pkg.go.dev/  but all of the Go packages are from GitHub.

to import a package: 

1. `go get -u github.com/`  , example:  `go get -u github.com/ttacon/chalk`
2. inside the main file import the github link inside the `import( )`, example:  `import( "fmt"; "github.com/ttacon/chalk" )` 
3. run `go mod tidy` for package maintenance 
4. run program `go run .`


#### errors

Go lang does not like when you have packages imported or variables etc declared and not used, comment the things not used to avoid any errors or warnings. 



----

Starter template

```Go
package main

import "fmt"

var pl = fmt.Println   // declare a shortcut for fmt.Println

func main() {
  pl("hello Go")
}

```























