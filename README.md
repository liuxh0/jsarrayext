# jsarrayext

[![Build Status](https://travis-ci.com/liuxh0/jsarrayext.svg?branch=master)](https://travis-ci.com/liuxh0/jsarrayext)
[![codecov](https://codecov.io/gh/liuxh0/jsarrayext/branch/master/graph/badge.svg)](https://codecov.io/gh/liuxh0/jsarrayext)

**jsarrayext** = javascript-style array extension for golang slice

This library lets you use Golang slices like JavaScript arrays by implementing the same methods. Generally, chaining methods, i.e. methods that don't modify the array itself and return the modified array, and some very common and useful methods are implemented.

Due to Golang language restrictions, optional parameters are not supported. Further, index is provided in callback functions while array not.

## Development Progress

Already implemented:

- every()
- fill()
- filter()
- find()
- findIndex()
- forEach()
- includes()
- indexOf()
- lastIndexOf()
- map() *Only method returning interface{} slice is implemented for now.*
- reduce()
- reduceRight()
- reverse()
- some()
- sort()


To be implemented:

- concat()

Won't implement:

- entries(), keys(), values() *since iterator is returned*
- flat(), flatMap() *due to experimental feature*
- join()
- pop(), unshift() *use append() instead*
- push(), shift(), splice() *see https://github.com/golang/go/wiki/SliceTricks*
- slice()
- toLocalString()
- toSource()
- toString()

Maybe:

- copyWithin()
- join()
