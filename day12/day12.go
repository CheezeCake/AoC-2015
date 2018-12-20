package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func sumArray(dec *json.Decoder) (sum1, sum2 uint64) {
	sum1 = 0
	sum2 = 0

	for {
		value, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		var s1, s2 uint64
		switch value.(type) {
		case json.Delim:
			d := value.(json.Delim)
			switch d.String() {
			case "{":
				s1, s2 = sumObject(dec)
			case "[":
				s1, s2 = sumArray(dec)
			case "}":
				panic("unexpected }")
			case "]":
				return
			}
		case float64:
			v := uint64(value.(float64))
			s1, s2 = v, v
		}
		sum1 += s1
		sum2 += s2
	}

	return
}

func sumObject(dec *json.Decoder) (sum1, sum2 uint64) {
	sum1 = 0
	sum2 = 0
	red := false

	for {
		property, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		switch property.(type) {
		case json.Delim:
			if red {
				sum2 = 0
			}
			return
		}

		value, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		var s1, s2 uint64
		switch value.(type) {
		case json.Delim:
			d := value.(json.Delim)
			if d.String() == "{" {
				s1, s2 = sumObject(dec)
			} else {
				s1, s2 = sumArray(dec)
			}
		case string:
			if value.(string) == "red" {
				red = true
			}
		case float64:
			v := uint64(value.(float64))
			s1, s2 = v, v
		}
		sum1 += s1
		sum2 += s2
	}

	if red {
		sum2 = 0
	}
	return
}

func main() {
	dec := json.NewDecoder(os.Stdin)

	var sum1, sum2 uint64
	t, err := dec.Token()
	if err == io.EOF {
		return
	}
	if err != nil {
		panic(err)
	}
	d := t.(json.Delim)
	if d.String() == "{" {
		sum1, sum2 = sumObject(dec)
	} else {
		sum1, sum2 = sumArray(dec)
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}
