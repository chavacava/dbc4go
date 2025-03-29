// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ClauseExpression

import "github.com/antlr4-go/antlr/v4"

type BaseClauseExpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseClauseExpressionVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitExprInParens(ctx *ExprInParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitImplies(ctx *ImpliesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitExistsIndex(ctx *ExistsIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitExistsElement(ctx *ExistsElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitPlainGoExpression(ctx *PlainGoExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitForallIndex(ctx *ForallIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitForallElement(ctx *ForallElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitIterator(ctx *IteratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitCollection(ctx *CollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitCompleteGoExpression(ctx *CompleteGoExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitGoExpression(ctx *GoExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitSliceIndex(ctx *SliceIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClauseExpressionVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}
