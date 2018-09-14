# exopulse files package
Golang package contains common file operations.

[![CircleCI](https://circleci.com/gh/exopulse/files.svg?style=svg)](https://circleci.com/gh/exopulse/files)
[![Build Status](https://travis-ci.org/exopulse/files.svg?branch=master)](https://travis-ci.org/exopulse/files)
[![GitHub license](https://img.shields.io/github/license/exopulse/files.svg)](https://github.com/exopulse/files/blob/master/LICENSE)

# Overview

This package contains various file related functions and helpers.

# Using files package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/files
 
Include files in your application.
```go
import "github.com/exopulse/files"
```

## Function list
```go
// FileExists return true if specified file exists.
func FileExists(path string) bool

// ReadLines read file lines into a slice.
func ReadLines(path string) ([]string, error)

// ReadBytes reads bytes from specified file or from STDIN if file name is empty.
func ReadBytes(path string) ([]byte, error)
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Files package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/files/blob/master/LICENSE)
