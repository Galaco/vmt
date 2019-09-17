[![GoDoc](https://godoc.org/github.com/golang-source-engine/vmt?status.svg)](https://godoc.org/github.com/golang-source-engine/vmt)
[![Go report card](https://goreportcard.com/badge/github.com/golang-source-engine/vmt)](https://goreportcard.com/badge/github.com/golang-source-engine/vmt)
[![GolangCI](https://golangci.com/badges/github.com/golang-source-engine/vmt.svg)](https://golangci.com)
[![CircleCI](https://circleci.com/gh/golang-source-engine/vmt/tree/master.svg?style=svg)](https://circleci.com/gh/golang-source-engine/vmt/tree/master)
[![codecov](https://codecov.io/gh/golang-source-engine/vmt/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-source-engine/vmt)

# vmt

> Vmt is a parser for Valve Material Format (`.vmt`) files.

Vmt's in Source can be a little painful; large number of possible parameters, many being
specific to a particular shader. This package exists to reduce the pain of parsing materials.

You can use either the `Properties` struct provided, or a custom definition to only get the properties desired. Properties must
be correctly typed in whatever definition is used. 

If there are parameters missing from the standard `Properties` definition, please submit a PR to add them.


**For now, nested properties are unsupported (e.g. `Proxies`).**

### Usage:
There are 3 ways to read material definitions using this package:

##### From a github.com/golang-source-engine/filesystem
This is the recommended solution, as it is the only method of resolving a material that
uses the `include` property.
```go
var fs *filesystem.Filesystem // create this elsewhere
result := vmt.NewProperties()
mat,err := vmt.FromFilesystem("metal/metal_01.vmt", fs, result)
```

##### From a stream
```go
stream,_ := os.Open("metal/metal_01.vmt")
result := vmt.NewProperties()
mat,err := vmt.FromStream(stream, result)
```

##### From existing github.com/galaco/keyvalues
```go
var kv *keyvalues.KeyValue // create this elsewhere
result := vmt.NewProperties()
mat,err := vmt.FromKeyValues(kv, result)
```


##### Material Definitions
If you've ever used the build-in json package, this should be familiar. See the `Properties` struct in this package for the default material. However, you can define your own definition below:
```go
package foo

type CustomShader struct {
    MyProperty string `vmt:$myproperty`
}
```