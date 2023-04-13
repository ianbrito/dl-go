package tag

type Tag int

const (
	PROGRAM Tag = iota
	BEGIN
	END
	INT
	REAL
	BOOL
	WRITE
	IF
	ASSIGN
	SUM
	SUB
	MUL
	OR
	LT
	LE
	GT
	SEMI
	DOT
	LPAREN
	RPAREN
	LIT_INT
	LIT_REAL
	ID
	TRUE
	FALSE
	EOF
	UNK
	TEMP
)

func (t Tag) String() string {
	names := [...]string{
		"PROGRAM",
		"BEGIN",
		"END",
		"INT",
		"REAL",
		"BOOL",
		"WRITE",
		"IF",
		"ASSIGN",
		"SUM",
		"SUB",
		"MUL",
		"OR",
		"LT",
		"LE",
		"GT",
		"SEMI",
		"DOT",
		"LPAREN",
		"RPAREN",
		"LIT_INT",
		"LIT_REAL",
		"ID",
		"TRUE",
		"FALSE",
		"EOF",
		"UNK",
		"TEMP",
	}

	if t < PROGRAM || t > TEMP {
		return "unknown"
	}

	return names[t]
}
