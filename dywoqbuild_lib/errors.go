package dywoqbuildlib

import "errors"

// ErrWrongStandard is a error indicating that the chosen standard by a user may be wrong.
var ErrWrongStandard = errors.New("standard c++ is not correct, values: c++98, c++03, c++11, c++14, c++17, c++20, c++23, c++2a")

// ErrWrongSuffix is a error that means the one of source files may have wrong suffix.
var ErrWrongSuffix = errors.New("suffix of one source files is not correct, allowed suffixes: .C, .cc, .cpp, .cxx, c++, .h, .H, .hh, .hpp, .hxx, h++")

// ErrNotFoundSourceFile is a error that occurs when the files (source and header files) are not found.
var ErrNotFoundFiles = errors.New("can't find source, header files; make sure you did add them")

// ErrNotAllowedLanguage is a error meaning the chose language is wrong.
var ErrNotAllowedLanguage = errors.New("not allowed language, values: C++, C")
