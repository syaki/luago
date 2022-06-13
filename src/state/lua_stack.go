package state

type LuaStack struct {
	slots []LuaValue
	top   int
}

func newLuaStack(size int) *LuaStack {
	return &LuaStack{
		slots: make([]LuaValue, size),
		top:   0,
	}
}

func (s *LuaStack) check(n int) {
	free := len(s.slots) - s.top
	for i := free; i < n; i++ {
		s.slots = append(s.slots, nil)
	}
}

func (s *LuaStack) push(val LuaValue) {
	if s.top == len(s.slots) {
		panic("stack overflow! ")
	}
	s.slots[s.top] = val
	s.top++
}

func (s *LuaStack) pop() LuaValue {
	if s.top < 1 {
		panic("stack underflow! ")
	}
	s.top--
	val := s.slots[s.top]
	s.slots[s.top] = nil
	return val
}

func (s *LuaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + s.top + 1
}

func (s *LuaStack) isValid(idx int) bool {
	absIdx := s.absIndex(idx)
	return absIdx > 0 && absIdx <= s.top
}

func (s *LuaStack) get(idx int) LuaValue {
	absIdx := s.absIndex(idx)
	if absIdx > 0 && absIdx <= s.top {
		return s.slots[absIdx-1]
	}
	return nil
}

func (s *LuaStack) set(idx int, val LuaValue) {
	absIdx := s.absIndex(idx)
	if absIdx > 0 && absIdx <= s.top {
		s.slots[absIdx-1] = val
		return
	}
	panic("invalid index! ")
}
