package main

import "fmt"

func emit(format string, v ...interface{}) {
	fmt.Printf("\t"+format+"\n", v...)
}

func emitLabel(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func emitDataSection() {
	emit(".data")

	// put strings
	for _, ast := range strings {
		emitLabel(".%s:", ast.slabel)
		emit(".string \"%s\"", ast.sval)
	}
}

func emitFuncPrologue(funcdef *Ast) {
	emitLabel("# funcdef %s", funcdef.fname)
	emit(".text")
	emit(".globl	%s", funcdef.fname)
	emitLabel("%s:", funcdef.fname)
	emit("push %%rbp")
	emit("mov %%rsp, %%rbp")

	if len(funcdef.localvars) > 0 {
		emit("# allocate stack area")
		stack_size := 8 * len(funcdef.localvars)
		emit("sub $%d, %%rsp", stack_size)
	}

	emit("# end of prologue")
}

func emitFuncEpilogue() {
	emit("#")
	emit("leave")
	emit("ret")
}

func emitExpr(ast *Ast) {
	switch ast.typ {
	case "int":
		emit("movl	$%d, %%eax", ast.ival)
	case "binop":
		emitExpr(ast.left)
		emit("push %%rax")
		emitExpr(ast.right)
		emit("push %%rax")
		emit("pop %%rbx")
		emit("pop %%rax")
		if ast.op == "+" {
			emit("addl	%%ebx, %%eax")
		} else if ast.op == "-" {
			emit("subl	%%ebx, %%eax")
		} else if ast.op == "*" {
			emit("imul	%%ebx, %%eax")
		}
	case "lvar":
		emit("mov %d(%%rbp), %%eax", ast.offset)
	case "assign":
		emitExpr(ast.right)
		emit("push %%rax")
		emit("mov %%eax, %d(%%rbp)", ast.left.offset)
	case "string":
		emit("lea .%s(%%rip), %%rax", ast.slabel)
	case "funcall":
		emitFuncall(ast)
	case "compound":
		for _, stmt := range ast.stmts {
			emitExpr(stmt)
		}
	case "decl":
		;
	default:
		panic(fmt.Sprintf("unexpected ast type %s", ast.typ))
	}
}

var regs = []string{"rdi", "rsi"}

func emitFuncall(funcall *Ast) {
	fname := funcall.fname
	emit("# funcall %s", fname)
	args := funcall.args
	for i, _ := range args {
		emit("push %%%s", regs[i])
	}

	emit("# setting arguments")
	for _, arg := range args {
		emitExpr(arg)
		emit("push %%rax")
	}

	for i, _ := range args {
		j := len(args) - 1 - i
		emit("pop %%%s", regs[j])
	}
	emit("mov $0, %%rax")
	emit("call %s", fname)

	for i, _ := range args {
		j := len(args) - 1 - i
		emit("pop %%%s", regs[j])
	}
}

func emitTopLevel(toplevel *Ast) {
	if toplevel.typ == "funcdef" {
		emitFuncPrologue(toplevel)
		emitExpr(toplevel.body)
		emit("mov $0, %%eax") // return 0
		emitFuncEpilogue()
	}
}

func generate(a *AstFile) {
	emitDataSection()
	for _, toplevel := range a.asts {
		emitTopLevel(toplevel)
	}
}
