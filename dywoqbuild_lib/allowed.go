package dywoqbuildlib

// AllowedStandards returns a slice of allowed standard.
func AllowedStandards() []string {
	return []string{
		"c++98",
		"c++03",
		"c++11",
		"c++14",
		"c++17",
		"c++20",
		"c++23",
		"c++2a",
	}
}

// AllowedExtensionFiles returns a slice of allowed extension files.
func AllowedFileExtensions() []string {
	return []string{
		// Source file extensions
		".C",
		".cc",
		".cpp",
		".cxx",
		".c++",

		// Header file extensions
		".h",
		".H",
		".hh",
		".hpp",
		".hxx",
		".h++"}
}
