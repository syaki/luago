package vm

const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

const (
	OP_MOVE = iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETTABLE
	OP_SETUPVAL
	OP_NEWTABLE
	OP_SUB
	OP_DIV
	OP_BXOR
	OP_BNOT
	OP_JMP
	OP_SELF
	OP_ADD
	OP_MUL
	OP_MOD
	OP_POW
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_RETURN
	OP_TFORLOOP
	OP_CALL
	OP_FORLOOP
	OP_FORPREP
	OP_SETLIST
	OP_TAILCALL
	OP_TFORCALL
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota
	OpArgU
	OpArgR
	OpArgK
)

type OpCode struct {
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode   byte
	name     string
}

var opcodes = []OpCode{
	OpCode{0, 1, OpArgR, OpArgN, IABC, "MOVE    "},
	OpCode{0, 1, OpArgK, OpArgN, IABx, "LOADK   "},
	OpCode{0, 1, OpArgN, OpArgN, IABx, "LOADKX  "},
	OpCode{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL"},
	OpCode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL "},
	OpCode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	OpCode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	OpCode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	OpCode{0, 0, OpArgK, OpArgK, IABC, "SETTABUP"},
	OpCode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	OpCode{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	OpCode{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	OpCode{0, 1, OpArgR, OpArgK, IABC, "SELF    "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "ADD     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "SUB     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "MUL     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "MOD     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "POW     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "DIV     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "IDIV    "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "BAND    "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "BOR     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "BXOR    "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "SHL     "},
	OpCode{0, 1, OpArgK, OpArgK, IABC, "SHR     "},
	OpCode{0, 1, OpArgR, OpArgN, IABC, "UNM     "},
	OpCode{0, 1, OpArgR, OpArgN, IABC, "BNOT    "},
	OpCode{0, 1, OpArgR, OpArgN, IABC, "NOT     "},
	OpCode{0, 1, OpArgR, OpArgN, IABC, "LEN     "},
	OpCode{0, 1, OpArgR, OpArgR, IABC, "CONCAT  "},
	OpCode{0, 0, OpArgR, OpArgN, IAsBx, "JMP     "},
	OpCode{1, 0, OpArgK, OpArgK, IAsBx, "EQ      "},
	OpCode{1, 0, OpArgK, OpArgK, IAsBx, "LT      "},
	OpCode{1, 0, OpArgK, OpArgK, IAsBx, "LE      "},
	OpCode{1, 0, OpArgR, OpArgN, IAsBx, "TEST    "},
	OpCode{1, 1, OpArgR, OpArgU, IAsBx, "TESTSET "},
	OpCode{0, 1, OpArgU, OpArgU, IABC, "CALL    "},
	OpCode{0, 1, OpArgU, OpArgU, IABC, "TAILCALL"},
	OpCode{0, 0, OpArgU, OpArgN, IABC, "RETURN  "},
	OpCode{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP "},
	OpCode{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP "},
	OpCode{0, 0, OpArgN, OpArgU, IAsBx, "TFORCALL"},
	OpCode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	OpCode{0, 0, OpArgU, OpArgU, IABC, "SETLIST "},
	OpCode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE "},
	OpCode{0, 1, OpArgU, OpArgN, IABC, "VARARG  "},
	OpCode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},
}
