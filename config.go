package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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

//EncodeFile Encode Config File
func EncodeFile(fpath string, v interface{}) error {
	indent, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("Marshal Config %v Error %v", fpath, err)
	}
	err = ioutil.WriteFile(fpath, formatArray(indent), 0644)
	if err != nil {
		return fmt.Errorf("Write Config %v Error %v", fpath, err)
	}
	return nil
}
func lineAndCharacter(input string, offset int) (line int, character int, err error) {
	lf := rune(0x0A)
	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("Couldn't find offset %d within the input", offset)
	}
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
func formatArray(indent []byte) []byte {
	var result string
	var pre string
	var start bool
	for _, val := range strings.Split(string(indent), "\n") {
		if strings.Contains(val, "[") {
			start = true
		}
		if strings.Contains(val, "]") {
			start = false
		}
		if strings.Contains(val, "[") && !strings.Contains(pre, "[") {
			result += strings.Trim(val, "\n")
		} else if strings.Contains(val, "[") && strings.Contains(pre, "[") {
			result += "\n" + val
		} else if strings.Contains(val, "]") && !strings.Contains(pre, "]") {
			result += strings.Trim(strings.Trim(strings.Trim(val, "\n"), " "), "	") + "\n"
		} else if start {
			result += strings.Trim(strings.Trim(strings.Trim(val, "\n"), " "), "	")
		} else if !start {
			result += val + "\n"
		}
		pre = val
	}
	return []byte(result)
}
