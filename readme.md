# Go Lang Notes  <img src="./logo.png" height="40" width="auto">  

###Variable Scope
- The places in code where a variable can be accessed
```go
 var x = 4
 func f(){
     fmt.Printf("%d", x)
 }
 func g(){
     fmt.Printf("%d", x)
 }

```
```go
 func f(){
     var x = 4
     fmt.Printf("%d", x)
 }
 func g(){
     fmt.Printf("%d", x) //variable x cannot be accessed here
 }

```
###Blocks
- A sequence of declarations and statements within matching brackets, {}
   - Including function definitions
- Hierarchy of implicit blocks also
- Universe block, all go source
- Package block, all source in a package
- File block, all source in a file
- "```if```", "```for```", "```switch```" all code inside the statement
- Clause in "```switch```" or "```select```", individual clauses each get a block

###Lexical Scoping
- Go is lexically scoped using blocks
- b1 >= b2 if b2 is defined inside b1
  - "defined inside" is transitive
```go
//b1
 var x = 4
//b2
 func f(){
     fmt.Printf("%d", x)
 }
//b3
 func g(){
     fmt.Printf("%d", x)
 }

```
###Deallocating Memory

- Whenever a variable is no longer needed it should be deallocated
  - Memory space is made available
- Otherwise, we will everntually run out of memory
```go
func f(){
    var x = 4
    fmt.Printf("%f", x)
}
```
- Each call ```f()``` creates an integer

###Stack Vs. Heap
- Stack is dedicated to function calls
   - Local variables are stored here
   - Deallocated after funciton completes
- Heap is persistent
  
 ```go
  var x = 4 //heap
  func f(){
      fmt.Printf("%d", x)
  }
  func g(){
      fmt.Printf("%d", x)
  }

 ```
 ```go
  func f(){
      var x = 4 //stack
      fmt.Printf("%d", x)
  }
  func g(){
      fmt.Printf("%d", x) 
  }

 ```

###Manual Deallocation
- Data on the heap must be deallocated when it is done being used
- In most comiled languages(i.e. C), this is done manually 
  ```c
  x = malloc (32);
  free(x);
  ```
- Error-prone, but fast

##Garbage Collection

### Pointers and Deallocation
- Hard to determine when a variable is no longer in use
  ```go
  func foo() *int{
      x:=1
      return &x
  }
  func main (){
      var y *int  
      y = foo()  //foo returns a pointer to x
      fmt.Printf("%d", *y)
  }
  ```
###Variable Scope
- The places in code where a variable can be accessed
- In interpreted languages, this is doen by the interpreter
   - java virtual machine
   - python interpreter
- Easy for the programmer
- Go is a compiled language which enables garbage collection
- Implementation is fast
- Compiler determines stack vs heap
- Garbage collection in the background

###Printing
- Format strings are good for formatting
  - conversion characters for each argument
  ``` fmt.Printf("Hi %s", x)```

- Generic int declaration
  - ```var x int ```
- Different lengths and signs
  - ```int8```, ```int 16```, ```int 32```, ```int 64```,
    ```uint 8```, ```uint 16```, ```uint 32```, ```uint 64```

###Type Conversions
- Most binary operations need operands for the same type
  - Including assignments
  - ```go 
       var x int32 = 1
       var y int16 = 2
       x = y //this will give an error
    ```
- Convert type with ```T()``` operation
      ``` x = int32(y)``` 

###Floating point
- ```float32``` ~6 digits of precision
- ```float64``` ~15 digits of precision
- Expressed using decimals or scientific notaion
  ```go
      var x float64 = 123.45
      var y float64 = 1.2345e2
  ```
- Complet umbers represented as two floats: real and imaginary
  ```go
      var z complex128 = complex(2,3)
  ```
###ASCII and Unicode
- American Standard Code for Information Exchange
- Character coding - each character is associated with an (7)8-bit number
    -  'A" = 0x41
- **Unicode** is a 32-bit character code
- **UTF-8** is variable length
   - 8-bit UTF codes are same as ASCII
- **Code points** - Unicode characters
- **Rune** - acode point in Go

###Strings
- Sequesnce of arbitary bytes
  - Read-only
  - often meant to be printed
- String literal - notated by double quotes
  ``` x := "hi there"```
- Each byte is a rune(UTF-8 code point)