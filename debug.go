package main

import "fmt"
import "os"

func debugf(format string, v... interface{}) {
	debugPrint(fmt.Sprintf(format, v...))
}

func debugPrint(s string) {
	spaces := ""
	for i:=0;i<nest;i++ {
		spaces += "  "
	}

	fmt.Fprintf(os.Stderr, "|%s %s\n", spaces, s)
}

func debugPrintV(v interface{}) {
	debugPrint(fmt.Sprintf("%v", v))
}

func debugPrintVar(name string, v interface{}) {
	debugPrint(fmt.Sprintf("%s=%v", name, v))
}

func dumpToken(tok *Token) {
	debugPrint(fmt.Sprintf("tok: type=%-8s, sval=\"%s\"", tok.typ, tok.sval))
}

var nest int

func (a *AstPkgClause) dump() {
	debugf("package %s", a.name)
}

func (a *AstFuncDef) dump() {
	debugf("funcdef %s", a.fname)
	nest++
	for _, stmt := range a.body.stmts {
		stmt.dump()
	}
	nest--
}

func (ast *AstAssignment) dump() {
	debugf("assign")
	nest++
	ast.left.dump()
	ast.right.dump()
	nest--
}

func (a *AstDeclVar) dump() {
	if a.initval == nil {
		debugf("var %s", a.variable.varname)
	} else {
		debugf("var %s =", a.variable.varname)
		a.initval.dump()
	}
}

func (a *AstStmt) dump() {
	if a.declvar != nil {
		a.declvar.dump()
	} else if a.assignment != nil {
		a.assignment.dump()
	} else if a.expr != nil {
		a.expr.dump()
	}
}
func (a *AstFile) dump() {
	debugPrint("==== Dump AstExpr Start ===")
	a.pkg.dump()
	for _, imprt := range a.imports {
		debugf("import \"%v\"", imprt.paths)
	}
	for _, f := range a.funcdefs {
		f.dump()
	}
	debugPrint("==== Dump AstExpr End ===")
}

func (ast *ExprFuncall) dump() {
	debugf(ast.fname)
	nest++
	for _, arg := range ast.args {
		arg.dump()
	}
	nest--
}

func (ast *ExprVariable) dump() {
	debugf("var %s", ast.varname)
}

func (ast *ExprNumberLiteral) dump() {
	debugf("int %d", ast.val)
}

func (ast *ExprStringLiteral) dump() {
	debugf("\"%s\"", ast.val)
}

func (ast *ExprBinop) dump() {
	debugf("binop %s", ast.op)
	nest++
	ast.left.dump()
	ast.right.dump()
	nest--
}

func errorf(format string, v ...interface{}) {
	/*
		currentTokenIndex := ts.index - 1
		fmt.Printf("%v %v %v\n",
			ts.getToken(currentTokenIndex-2), ts.getToken(currentTokenIndex-1), ts.getToken(currentTokenIndex))
	*/
	s := fmt.Sprintf(format, v...)
	panic(s)
}

func assert(cond bool, msg string) {
	if !cond {
		panic(fmt.Sprintf("assertion failed: %s", msg))
	}
}