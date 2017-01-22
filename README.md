# ovw

[![Build Status](https://travis-ci.org/achiku/ovw.svg?branch=master)](https://travis-ci.org/achiku/ovw)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/ovw/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/ovw)](https://goreportcard.com/report/github.com/achiku/ovw)

Merge two structs, and overwrite fields when necessary

## Why created

To make creating struct for tests easier.


## Example

- https://github.com/achiku/ovw/blob/master/ovw_test.go


## Installation

```
go get -u github.com/achiku/ovw
```


## Contributing

Pull requests for new features, bug fixes, and suggestions are welcome!

### Install gom

This project is using [gom](https://github.com/mattn/gom) for dependency management.

```
$ go get -u github.com/mattn/gom
$ gom install
```

### Test

```
$ go test -v
```
