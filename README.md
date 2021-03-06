GMIT Software Development 3rd year.

Module: Data Representation.

GoLang Problem sheet.

### Go Examples
##### This repository contains some example code written in the programming language Go.
##### The author is Yongjin Kim, a Student at GMIT.
---
This problem set is to learn the fundamentals of creating a web application in Go. 

## How to clone
1. Open **Git Bash**.
2. Change the current working directory to the location where you want the cloned directory to be made.
3. Type 'git clone', and then paste the URL `https://github.com/kentkim84/DataRepProblemSet_03.git` or `git@github.com:kentkim84/DataRepProblemSet_03.git`
```
git clone https://github.com/kentkim84/DataRepProblemSet_03.git
```
## Coding Standards
// Version 0.2 using C standards

### 1. Naming Conventions and Style
1.1. Use Pascal casing for class and structs
    
    type Rectangle struct {
        Name string
        Width, Height float64
    }

    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

1.2. Use camel casing for local variable names and function parameters
    
    func SomeMethod(someParameter const int)
    {
        someNumber int
    }

1.3. Use verb-object pairs for method names
a.	Use pascal casing for public methods
        
    public:
    func DoSomething()

b.	Use camel casing for other methods
        
    private:
    func doSomething()

1.4. Put constant variables after import

    import (
        "fmt"
    )

    const CONSTANT_NUMBER int = 100

### References
This examples are from
* [Data Representation and Querying - Ian McLoughlin](https://data-representation.github.io/problems/go-regexp.html)

This coding standard is followed by
* [popeKim's c/c++ coding standards](https://docs.google.com/document/d/1cT8EPgMXe0eopeHvwuFmbHG4TJr5kUmcovkr5irQZmo/edit#heading=h.r2n9mhxbh2gg)