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
### Variable Scope
- The places in code where a variable can be accessed
