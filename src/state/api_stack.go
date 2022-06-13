package state

func (s *LuaState) GetTop() int {
	return s.stack.top
}

func (s *LuaState) AbsIndex(idx int) int {
	return s.stack.absIndex(idx)
}
