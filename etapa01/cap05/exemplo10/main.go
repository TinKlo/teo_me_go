package main

import "fmt"

func main() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "hidrogem",
			"state": "Gas",
		},

		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], ":", el["state"])
	}
}
