package state

import (
	"reflect"

	"github.com/syaki/luago/src/api"
)

type LuaValue interface{}

func typeOf(val LuaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64, float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		panic("todo! " + reflect.TypeOf(val).String())
	}
}
