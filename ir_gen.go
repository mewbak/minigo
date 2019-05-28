package main

import "fmt"

type IrRoot struct {
	vars           []*DeclVar
	funcs          []*DeclFunc
	packages       []*AstPackage
	methodTable    map[int][]string
	uniquedDTypes  []string
	importOS       bool
}

var groot *IrRoot

var vars []*DeclVar
var funcs []*DeclFunc

func collectDecls(pkg *AstPackage) {
	for _, f := range pkg.files {
		for _, decl := range f.topLevelDecls {
			if decl.vardecl != nil {
				vars = append(vars, decl.vardecl)
			} else if decl.funcdecl != nil {
				funcs = append(funcs, decl.funcdecl)
			}
		}
	}
}


func setStringLables(pkg *AstPackage, prefix string) {
	for id, sl := range pkg.stringLiterals {
		sl.slabel = fmt.Sprintf("%s.S%d", prefix, id+1)
	}
}

func makeIR(internalUniverse *AstPackage, internalRuntime *AstPackage, csl *compiledStdlib, mainPkg *AstPackage) *IrRoot {
	var packages []*AstPackage

	importedPackages := csl.getPackages()
	for _, pkg := range importedPackages {
		packages = append(packages, pkg)
		collectDecls(pkg)
		setStringLables(pkg, string(pkg.name))
	}

	packages = append(packages, internalUniverse)
	collectDecls(internalUniverse)
	setStringLables(internalUniverse, "")

	packages = append(packages, internalRuntime)
	collectDecls(internalRuntime)
	setStringLables(internalRuntime, "iruntime")

	packages = append(packages, mainPkg)
	collectDecls(mainPkg)
	setStringLables(mainPkg, string(mainPkg.name))

	var dynamicTypes []*Gtype
	for _, pkg := range packages {
		for _, dt := range pkg.dynamicTypes {
			dynamicTypes = append(dynamicTypes, dt)
		}
	}

	root := &IrRoot{}
	root.packages = packages
	root.vars = vars
	root.funcs = funcs
	root.setDynamicTypes(dynamicTypes)
	root.importOS = in_array("os", csl.uniqImportedPackageNames)
	root.methodTable = composeMethodTable(funcs)
	return root
}

func (ir *IrRoot) setDynamicTypes(dynamicTypes []*Gtype) {
	var uniquedDTypes []string = builtinTypesAsString
	for _, gtype := range dynamicTypes {
		gs := gtype.String()
		if !in_array(gs, uniquedDTypes) {
			uniquedDTypes = append(uniquedDTypes, gs)
		}
	}

	ir.uniquedDTypes = uniquedDTypes
}

func composeMethodTable(funcs []*DeclFunc)  map[int][]string  {
	var methodTable map[int][]string = map[int][]string{} // receiverTypeId : []methodTable

	for _, funcdecl := range funcs {
		if funcdecl.receiver == nil {
			continue
		}
			//debugf("funcdecl:%v", funcdecl)
			gtype := funcdecl.receiver.getGtype()
			if gtype.kind == G_POINTER {
				gtype = gtype.origType
			}
			if gtype.relation == nil {
				errorf("no relation for %#v", funcdecl.receiver.getGtype())
			}
			typeId := gtype.relation.gtype.receiverTypeId
			symbol := funcdecl.getSymbol()
			methods := methodTable[typeId]
			methods = append(methods, symbol)
			methodTable[typeId] = methods
	}
	debugf("set methodTable")
	return methodTable
}
