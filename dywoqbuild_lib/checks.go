package dywoqbuildlib

import (
	"slices"
	"strings"
)

func checkForCorrectnessOfStandard(buildConfig BuildConfig) error {
	sliceOfAllowedStandard := AllowedStandards()
	found := slices.Contains(sliceOfAllowedStandard, buildConfig.Standard)
	if !found {
		return ErrWrongStandard
	}

	return nil
}

func checkForCorrectnessOfSourceFiles(buildConfig BuildConfig) error {
	sliceOfAllowedFileExtensions := AllowedFileExtensions()
	for _, file := range buildConfig.Files {
		hasAllowedSuffix := false
		for _, allowedFileExtension := range sliceOfAllowedFileExtensions {
			if strings.HasSuffix(file, allowedFileExtension) {
				hasAllowedSuffix = true
				break
			}
		}

		if !hasAllowedSuffix {
			return ErrWrongSuffix
		}
	}
	return nil
}

func runAllCheckUps(buildConfig BuildConfig) (err error) {
	err = checkForCorrectnessOfStandard(buildConfig)
	if err != nil {
		return
	}

	err = checkForCorrectnessOfSourceFiles(buildConfig)
	if err != nil {
		return
	}

	return
}
