// dywoqbuildlib is a library that provides functionality to build own executables
// of build system dywoqbuild.
package dywoqbuildlib

type BuildConfig struct {
	IncludeDirectories []string `json:"include_directories"`
	LibraryDirectories []string `json:"library_directories"`
	Libraries          []string `json:"libraries"`
	CompilerFlags      []string `json:"compiler_flags"`
	LinkerFlags        []string `json:"linker_flags"`
	Files              []string `json:"files"`
	GlobPatterns       []string `json:"glob_patterns"`
	Standard           string   `json:"standard"`
	OutputType         string   `json:"output_type"`
	OutputName         string   `json:"output_name"`
	Language           string   `json:"language"`
	Compiler           string   `json:"compiler"`
}

// NewBuildConfig creates new instance of BuildConfig.
// It's recommended to use this method, since it runs
// all config check ups. Returns error if something is wrong.
func NewBuildConfig(buildConfig BuildConfig) (newInstance *BuildConfig, err error) {
	newInstance = &buildConfig
	err = runAllCheckUps(buildConfig)
	return
}

// CheckUp checks properties for their correctness.
// NewBuildConfig already does it, but if you still
// created instance that way:
//
// 	config := BuildConfig{...}
//
// You can use CheckUp:
//
// 	err := config.CheckUp()
// 	if err != nil {
// 		panic(err)
// 	}
func (bd BuildConfig) CheckUp() error {
	return runAllCheckUps(bd)
}
