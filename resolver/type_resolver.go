package resolver

import (
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/symbol"
)

type TypeResolver struct {
	parser.BaseGScriptListener
	at *AnnotatedTree
}

func NewTypeResolver(at *AnnotatedTree) *TypeResolver {
	return &TypeResolver{at: at}
}

func (t *TypeResolver) ExitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) {
}

func (t *TypeResolver) EnterVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) {
	id := ctx.IDENTIFIER().GetText()
	variable := symbol.NewVariable(ctx, id, nil)
	t.at.PutSymbolOfNode(ctx, variable)
}
