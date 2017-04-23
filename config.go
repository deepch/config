package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//DecodeFile Decode Config File
func DecodeFile(fpath string, v interface{}) error {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("Read Config %v Error %v", fpath, err)
	}
	err = json.Unmarshal(file, &v)
	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, _ := lineAndCharacter(string(file), int(jsonError.Offset))
		return fmt.Errorf("Decode Config Syntax Error File %v Line %v Char %v Error %v", fpath, line, character, err)
	}
	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, _ := lineAndCharacter(string(file), int(jsonError.Offset))
		return fmt.Errorf("Decode Config Type Error File %v Line %v Char %v Error %v", fpath, line, character, err)
	}
	if err != nil {
		return fmt.Errorf("Decode Config Other Error File %v Error %v", fpath, err)
	}
	return nil
}
func lineAndCharacter(input string, offset int) (line int, character int, err error) {
	lf := rune(0x0A)
	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("Couldn't find offset %d within the input.", offset)
	}
	// Humans tend to count from 1.
	line = 1
	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}
	return line, character, nil
}
