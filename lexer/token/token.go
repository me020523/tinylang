package token

/**
* 在tinylang中存在三种token:
* 1) 字符串常量
* 2) 整数常量
* 3) 标识符
**/
func NewIntegerToken(value int, lineNum int) IntegerToken {
	return IntegerToken{
		Token{
			lineNum,
		},
		value,
	}
}

func NewStringToken(literal string, lineNum int) StringToken {
	return StringToken{
		Token{
			lineNum,
		},
		literal,
	}
}

func NewIdentifierToken(identifier string, lineNum int) IdentifierToken {
	return IdentifierToken{
		Token{
			lineNum,
		},
		identifier,
	}
}

type Token struct {
	LineNum int
}

func (token Token) isInteger() bool {
	return false
}

func (token Token) isString() bool {
	return false
}

func (token Token) isIdentifier() bool {
	return false
}

//////////////////////definitions of IntegerToken///////////////////////
type IntegerToken struct {
	Token
	Value int
}

func (token IntegerToken) isInteger() bool {
	return true
}

/////////////////////definitions of StringToken/////////////////////////////////////
type StringToken struct {
	Token
	Literal string
}

func (token StringToken) isString() bool {
	return true
}

////////////////definitions of IdentifierToken /////////////////////////
type IdentifierToken struct {
	Token
	Identifier string
}

func isIdentifier(token IdentifierToken) bool {
	return true
}
