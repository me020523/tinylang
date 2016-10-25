package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
)

const REGEX_PAT = "\\s*(//.*)|" +
	"([1-9][0-9]*)|" +
	`("(?:\\"|\\\\|[^"])*")|` +
	"(" +
	"[_a-zA-Z])[0-9a-zA-Z]*|" +
	"==|>=|<=|&&|\\|\\||" +
	"+|-|*|/|=" +
	")"

type Lexer struct {
	source    io.Reader     //源代码输入
	tokens    []interface{} //源码中所有的token列表
	parserBuf string        //存放一些中间字符串的缓存
}

func NewLexer(source io.Reader) *Lexer {
	return &Lexer{
		source:    source,
		tokens:    []interface{}{}, //源代码中所有的token列表
		parserBuf: "",
	}
}

func (lexer *Lexer) Parse() error {
	bufferedReader := bufio.NewReader(lexer.source)
	var current uint32 = 0 //当前所在行
	var newLine bool = true
	for {
		if newLine {
			current = current + 1
		}
		line, prefix, err := bufferedReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		//判断当前行是否已经读完
		if !prefix {
			//当前行已经读完，即将开始新的一行
			newLine = true
		} else {
			//当前行还未读完
			newLine = false
		}

		//处理当前行
		more, err := lexer.processLine(current, string(line))
		if err != nil {
			return err
		}
		if newLine && more {
			errMsg := fmt.Sprintf("Line num: %d, Error: unknown token: %s",
				current, lexer.parserBuf)
			return errors.New(errMsg)
		}
	}
	return nil
}

/**
** 返回值:
**   bool: 指示是否当前行还有未处理的完的字符
**   error: 是否发生错误
**/
func (lexer *Lexer) processLine(lineNum uint32, line string) (bool, error) {
	matcher, err := regexp.Compile(REGEX_PAT)
	if err != nil {
		return false, err
	}
	lastPos := 0
	for lastPos != len(line) {
		indexes := matcher.FindStringSubmatchIndex(line)
		if len(indexes) == 0 {
			//剩余字符串中不包含有效信息
			lexer.parserBuf = line[lastPos:len(line)]
			return true, nil
		}
		//获取到一个token
	}

	return true, nil
}
