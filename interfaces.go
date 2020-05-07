package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`
func addNum(num1, num2 interface{})int{
    // num1 and num2 are of interface{} type only (not any type)
    // but any type can be converted to interface{} type
    // This is really tricky, esp. coming from Python world
    // Need type assertion to ensure types are as expected
    m := num1.(int)
    n := num2.(int)
    return m + n    // returns an integer
}
func main() {
    // our target will be of type map[string]interface{}, which is a
    // pretty generic type that will give us a hashtable whose keys
    // are strings, and whose values are of type interface{}
    var val map[string]interface{}

    if err := json.Unmarshal([]byte(input), &val); err != nil {
        panic(err)
    }

    fmt.Println(val)
    for k, v := range val {
        fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(addNum(1,2))
}
