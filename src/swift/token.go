package swift

// Token is an enum of lexical tokens of the Swift programming language.
type Token int

//TODO: associativity and precedence

// Token values.
const (
	INVALID Token = iota
	EOF

	// DO WE NEED DYNAMIC prefix postfix operators ?
	// prefix operator
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	// TODO .... other op's

	// punctuations
	ARROW       //->
	COLON       // :
	COMMA       // ,
	DOT         // .
	SEMICOLON   // ;
	UNDERSCORE  // _
	AT          // @
	HASH        // #
	BACKLASH    // `/`
	LEFTPAREN   // (
	RIGHTPAREN  // )
	LEFTBRACE   // {
	RIGHTBRACE  // }
	LEFTSQUARE  // [
	RIGHTSQUARE // ]

	// Literals
	INT
	FLOAT
	CHAR
	BOOLEAN
	STRING

	// Modifier
	CONVENIENCE
	DYNAMIC
	FINAL
	LAZY
	MUTATING
	NONMUTATING
	OPTIONAL
	OVERRIDE
	REQUIRED
	STATIC
	UNOWNED
	WEAK
	INTERNAL
	PRIVATE
	PUBLIC
	FILEPRIVATE
	// OPEN - no need for this.

	// Keywords
	ANY
	SELF
	AS
	//associativity
	BREAK
	CATCH
	CASE
	CLASS
	CONTINUE
	DEFAULT
	DEFER
	DEINIT
	DIDSET
	DO
	ENUM
	EXTENSION
	ELSE
	FALLTHROUGH
	FOR
	FUNC
	GET
	GUARD
	IF
	IMPORT
	IN
	INDIRECT
	INFIX
	INIT
	INOUT
	IS
	LET
	LEFT
	NIL
	NONE
	TRUE  // true
	FALSE // false

	OPERATOR // ?
	POSTFIX  // ?
	PREFIX   // ?

	//precedence

	PROTOCOL

	REPEAT
	RETHROWS
	RETURN

	RIGHT //?
	SAFE  //?

	SET
	STRUCT
	SUBSCRIPT
	SUPER
	SWITCH
	THROW
	THROWS
	TRY
	TYPEALIAS

	UNSAFE //?

	VAR
	WHERE
	WHILE
	WILLSET

	TYPE
)

var tokens = [...]string{
	INVALID: "INVALID",
	EOF:     "EOF",
}
