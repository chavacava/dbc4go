// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ClauseExpression

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ClauseExpressionParser.
type ClauseExpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ClauseExpressionParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ExprInParens.
	VisitExprInParens(ctx *ExprInParensContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#Implies.
	VisitImplies(ctx *ImpliesContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ExistsIndex.
	VisitExistsIndex(ctx *ExistsIndexContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ExistsElement.
	VisitExistsElement(ctx *ExistsElementContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#PlainGoExpression.
	VisitPlainGoExpression(ctx *PlainGoExpressionContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#Iff.
	VisitIff(ctx *IffContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ForallIndex.
	VisitForallIndex(ctx *ForallIndexContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ForallElement.
	VisitForallElement(ctx *ForallElementContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ExistsIterator.
	VisitExistsIterator(ctx *ExistsIteratorContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#ForallIterator.
	VisitForallIterator(ctx *ForallIteratorContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#iterator.
	VisitIterator(ctx *IteratorContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#collection.
	VisitCollection(ctx *CollectionContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#completeGoExpression.
	VisitCompleteGoExpression(ctx *CompleteGoExpressionContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#goExpression.
	VisitGoExpression(ctx *GoExpressionContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#functionCallArguments.
	VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#sliceIndex.
	VisitSliceIndex(ctx *SliceIndexContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ClauseExpressionParser#string.
	VisitString(ctx *StringContext) interface{}
}
