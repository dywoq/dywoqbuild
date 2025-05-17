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

func checkForCorrectnessOfLanguage(buildConfig BuildConfig) error {
	sliceOfAllowedLanguages := AllowedLanguages()
	for _, allowedLanguage := range sliceOfAllowedLanguages {
		if buildConfig.Language != allowedLanguage {
			return ErrNotAllowedLanguage
		}
	}

	return nil
}

func checkForCorrectnessOfTargetType(buildConfig BuildConfig) error {
	sliceOfAllowedTargetTypes := AllowedTargetTypes()
	for _, allowedTargetType := range sliceOfAllowedTargetTypes {
		if buildConfig.TargetType != allowedTargetType {
			return ErrWrongTargetType
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

	err = checkForCorrectnessOfLanguage(buildConfig)
	if err != nil {
		return
	}

	err = checkForCorrectnessOfTargetType(buildConfig)
	if err != nil {
		return
	}

	return
}
