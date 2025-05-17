package main

import (
	"fmt"

	dywoqbuildlib "github.com/dywoq/dywoqbuild/dywoqbuild_lib"
)

func main() {
	config := dywoqbuildlib.BuildConfig{
		Language: "C++",
		Standard: "c++11",
		Files:    []string{"main.cxx"},
	}

	fmt.Printf("config: %v\n", config)
}
