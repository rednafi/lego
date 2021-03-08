package main

import "fmt"

// Dict emulates a map type
type Dict map[string]string

// Search takes a map an return the key
func (d Dict) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", fmt.Errorf("can't find the value for key %q", key)
	}
	return val, nil
}

// Add takes a map and a key-value pair and updates the map
func (d Dict) Add(key, value string) error {
	// Checking if the key already exists
	_, ok := d.Search(key)

	if ok == nil {
		return fmt.Errorf("key %q already exists in Dict", key)
	}

	d[key] = value
	return nil

}

// Update method updates a key:value pair if the key already exists
func (d Dict) Update(key, value string) error {
	// Checking if the key exists or not
	_, ok := d.Search(key)
	if ok != nil {
		return fmt.Errorf("key %q does not exit", key)
	}

	d[key] = value
	return nil
}

// Delete method deletes a key:value pair if the key already exists
func (d Dict) Delete(key string) error {
	// Checking if the key exists or not
	_, ok := d.Search(key)
	if ok != nil {
		return fmt.Errorf("key %q does not exit", key)
	}
	delete(d, key)
	return nil

}

// func main() {
// 	d := Dict{}
// 	// d["test"] = "this is just a test"
// 	// err := d.Add("test", "ss")
// 	// fmt.Println(d, err)

// 	val, err := d.Search("test")
// 	fmt.Println(val, err)

// }
