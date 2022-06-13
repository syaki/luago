package state

type LuaState struct {
	stack *LuaStack
}

func New() *LuaState {
	return &LuaState{
		stack: newLuaStack(20),
	}
}
