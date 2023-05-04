# Learn GJSON and SJSON

## References

GJSON is a Go package that provides a fast and simple way to get values from a json document.
It has features such as one line retrieval, dot notation paths, iteration, and parsing json lines.

SJSON is a Go package that provides a very fast and simple way to set a value in a json document.

[gjson api reference](https://pkg.go.dev/github.com/tidwall/gjson)

[gjson repo](https://github.com/tidwall/gjson)

[gjson path syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

[gjson getting started](https://pkg.go.dev/github.com/tidwall/gjson)

[Dynamic JSON Parsing using empty Interface and without Struct in Go Language](https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968) - basics, no libraries

[A Complete Guide to JSON in Golang](https://www.sohamkamani.com/golang/json/) - Soham Kamani - plain go

[sjson repo](https://github.com/tidwall/sjson)

[sjson api reference](https://pkg.go.dev/github.com/tidwall/sjson)

[jj repo](https://github.com/tidwall/jj) - the cli tool using gjson and sjson

## Experimental go code for applied json parsing

### gjson

`Xform` a minimal demo of JSON tansform (picking out values from input and repacking into output)

### gojson

Demo of unmarshalling json to a specific struct (static), and to arrays and maps (dynamic)
