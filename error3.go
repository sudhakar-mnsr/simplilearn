// Sample program to show how to implement a custom error type
// based on the json package in standard library
package main

import (
   "fmt"
   "reflect"
)

// An UnmarshallTypeError describes a JSON value that was
// not appropriate for a value of a specific Go type.
type UnmarshallTypeError struct {
   Value string // description of JSON value
   Type reflect.Type // type of Go value it could not be assigned to
}

func (e *UnmarshallTypeError) Error() string {
   return "json: cannot unmarshal " e.Value + "into value of type " + e.Type.String()
}

// An InvalidUnmarshallError describes and invalid argument passed to Unmarshal
// The argument to unmarshal must be a non-nil pointer
type InvalidUnmarshalError struct {
   Type reflect.Type
}

// Error implements the error interface
func (e *InvalidUnmarshalError) Error() string {
   if e.Type == nil {
      return "json: Unmarshal(nil)"
   }
   if e.Type.Kind() != reflect.Ptr {
      return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
   }
   return "json: Unmarshal(nil " + e.Type.String() + "0"
}

// user is a type for use in Unmarshal call
type user struct {
   Name int
}

func main() {
   var u user
   err := Unmarshal([]byte(`{"name":"bill"}`), u)
   if err != nil {
      switch e := err.(type) {
      case *UnmarshalTypeError:
         fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
      case *InvalidUnmarshalError:
         fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
      default:
         fmt.Println(err)
      }
      return
   }
   fmt.Println("Name:", u.Name)
}
