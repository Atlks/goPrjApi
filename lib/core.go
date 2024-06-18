package lib

import (
	"fmt"
	"reflect"
)

// Define a global map to store function references by name
var funcMap = map[string]interface{}{
	"Add": Add,
}

// Add function example
func Add(a, b int) int {
	return a + b
}

// FunctionExists checks if a function exists in the funcMap
func function_exists(name string) bool {
	_, exists := funcMap[name]
	return exists
}

// CallFunc is a function to dynamically call a function with a list of arguments
func call_user_func_array(fn interface{}, args ...interface{}) (result []reflect.Value, err error) {
	// Get the reflection value of the function
	fnValue := reflect.ValueOf(fn)

	// Ensure the input is a function
	if fnValue.Kind() != reflect.Func {
		err = fmt.Errorf("fn is not a function")
		return
	}

	// Create a slice of reflect.Value from the provided arguments
	inputs := make([]reflect.Value, len(args))
	for i, arg := range args {
		inputs[i] = reflect.ValueOf(arg)
	}

	// Call the function with the arguments and get the result
	result = fnValue.Call(inputs)
	return
}

// cvt
func AnyToMap(v any) map[string]any {
	map1 := v.(map[string]any)
	return map1
}

func AnyToInt(v any) int32 {
	map1 := v.(int32)
	return map1
}

func AnyToNumFlt64(v any) float64 {
	map1 := v.(float64)
	return map1
}
