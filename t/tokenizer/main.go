package main

var GENERATION int = 2

var debugMode = true
var debugToken = false

func f3() {
	path := S("t/min/min.go")
	bs := NewByteStreamFromFile(path)

	var c byte
	c, _ = bs.get()
	tn := &Tokenizer{
		bs: bs,
	}
	ident := tn.readIdentifier(c)
	fmtPrintf(S("%s\n"), []byte(ident))
}

func f4() {
	path := S("t/min/min.go")
	bs := NewByteStreamFromFile(path)
	tokens := Tokenize(bs)
	fmtPrintf(S("%d\n"), len(tokens)) // 26
	fmtPrintf(S("----------\n"))
	// disable befow for now
	return
	for _, tok := range tokens {
		fmtPrintf(S("%d:%s\n"), int(tok.typ), tok.getSval())
	}
}

func f5() {
	debugToken = false
	path := S("t/data/string.txt")
	bs := NewByteStreamFromFile(path)

	tokens := Tokenize(bs)
	tok := tokens[0]
	fmtPrintf(S("----------\n"))
	fmtPrintf(S("[%s]\n"), []byte(tok.sval))
}

func main() {
	f3()
	f4()
	f5()
}
