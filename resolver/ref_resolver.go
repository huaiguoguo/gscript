package resolver

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/symbol"
)

// RefResolver 引用消解和类型推断
// 函数调用，消解出是哪个 function，以及返回值类型
type RefResolver struct {
	parser.BaseGScriptListener
	at                    *AnnotatedTree
	localVariableResolver *TypeResolver
	typeResolverWalker    *antlr.ParseTreeWalker
}

func NewRefResolver(at *AnnotatedTree) *RefResolver {
	return &RefResolver{
		at:                    at,
		localVariableResolver: NewTypeResolverWithLocalVariable(at),
		typeResolverWalker:    antlr.NewParseTreeWalker(),
	}
}

// EnterVariableDeclarators 本地变量必须是添加和解析同时进行
func (s *RefResolver) EnterVariableDeclarators(ctx *parser.VariableDeclaratorsContext) {
	scope := s.at.FindEncloseScopeOfNode(ctx)
	_, isBlock := scope.(*symbol.BlockScope)
	_, isFunc := scope.(*symbol.Func)
	if isBlock || isFunc {
		s.typeResolverWalker.Walk(s.localVariableResolver, ctx)
	}
}

func (s *RefResolver) ExitIdentifierPrimary(ctx *parser.IdentifierPrimaryContext) {
	idName := ctx.IDENTIFIER().GetText()
	scope := s.at.FindEncloseScopeOfNode(ctx)
	variable := s.at.FindVariable(scope, idName)
	if variable == nil {
		// todo crossoverJie 区分返回的是函数

		// todo crossoverJie 完善编译报错信息
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		panic(fmt.Sprintf("unknown variable %s, at %d and column:%d", idName, line, column))
	}

	s.at.PutSymbolOfNode(ctx, variable)

	// 写入类型
	s.at.PutTypeOfNode(ctx, variable.GetType())
}

func (s *RefResolver) ExitLiterPrimary(ctx *parser.LiterPrimaryContext) {
	symbolType := s.at.GetTypeOfNode()[ctx]
	s.at.PutTypeOfNode(ctx, symbolType)
}

func (s *RefResolver) ExitExprPrimary(ctx *parser.ExprPrimaryContext) {
	symbolType := s.at.GetTypeOfNode()[ctx]
	s.at.PutTypeOfNode(ctx, symbolType)
}

// ExitFunctionCall 设置当前 scope 中的函数以及函数返回值
func (s *RefResolver) ExitFunctionCall(ctx *parser.FunctionCallContext) {

	// todo crossoverJie 处理内置函数
	name := ctx.IDENTIFIER().GetText()

	paramTypes := s.getParamTypes(ctx)

	scope := s.at.FindEncloseScopeOfNode(ctx)
	found := false

	// todo crossoverJie . 符号级联调用

	if !found {
		function := s.at.FindFunction(scope, name, paramTypes)
		if function != nil {
			found = true
			s.at.PutSymbolOfNode(ctx, function)
			s.at.PutTypeOfNode(ctx, function.GetReturnType())
		}
	}

	// todo crossoverJie 查找是否是类的构造函数

	// todo crossoverJie 查找是否为函数变量

}

func (s *RefResolver) ExitPrimaryExpr(ctx *parser.PrimaryExprContext) {
	symbolType := s.at.GetTypeOfNode()[ctx]
	s.at.PutTypeOfNode(ctx, symbolType)
}

func (s *RefResolver) ExitFuncCallExpr(ctx *parser.FuncCallExprContext) {
	symbolType := s.at.GetTypeOfNode()[ctx]
	s.at.PutTypeOfNode(ctx, symbolType)
}

// 查询函数的参数列表
func (s *RefResolver) getParamTypes(ctx *parser.FunctionCallContext) []symbol.Type {
	var paraTypes []symbol.Type
	for _, context := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpr() {
		symbolType := s.at.GetTypeOfNode()[context]
		paraTypes = append(paraTypes, symbolType)
	}
	return paraTypes
}
