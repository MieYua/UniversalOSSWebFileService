package xorm

import (
	"reflect"

	"UniversalOSSWebFileService/Godeps/_workspace/src/github.com/go-xorm/core"
)

var (
	ptrPkType = reflect.TypeOf(&core.PK{})
	pkType    = reflect.TypeOf(core.PK{})
)
