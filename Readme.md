# golang + Utilities = goutils

goutils (Golang UTILitieS) is a Golang implementation of some functions and data types to help other libraries.  
You can see an extended doc in [godocs](https://godoc.org/github.com/fmorenovr/goutils).

Install it writing in terminal:

    go get github.com/fmorenovr/goutils

Example:

If you wan to compare different interfaces with an Integer parse:

  ```go
    var a, b interface{}
    a = 2
    b = 5
    goutils.IntComparator(a, b)
    // this should return 1 if a>b
    // this should return -1 if a<b
    // this should return 0 if a=b
  ```
    
Or define your own comparator function:

  ```go
    // supose that this is your struct
    type MyDataType struct{
      ID int
      name string
    }

    // your comparator, you cna use this function as TypeComparator type function
    func MyComparator(a, b interface{}) int {
      AaAsserted := a.(MyDataType)
      bAsserted := b.(MyDataType)
      switch {
        case aAsserted.ID > bAsserted.ID:
          return 1
        case aAsserted.ID < bAsserted.ID:
          return -1
        default:
          return 0
      }
    }
  ```
