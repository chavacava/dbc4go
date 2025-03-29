// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ClauseExpression

import "github.com/antlr4-go/antlr/v4"

// BaseClauseExpressionListener is a complete listener for a parse tree produced by ClauseExpressionParser.
type BaseClauseExpressionListener struct{}

var _ ClauseExpressionListener = &BaseClauseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClauseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClauseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClauseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClauseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseClauseExpressionListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseClauseExpressionListener) ExitRoot(ctx *RootContext) {}

// EnterExprInParens is called when production ExprInParens is entered.
func (s *BaseClauseExpressionListener) EnterExprInParens(ctx *ExprInParensContext) {}

// ExitExprInParens is called when production ExprInParens is exited.
func (s *BaseClauseExpressionListener) ExitExprInParens(ctx *ExprInParensContext) {}

// EnterImplies is called when production Implies is entered.
func (s *BaseClauseExpressionListener) EnterImplies(ctx *ImpliesContext) {}

// ExitImplies is called when production Implies is exited.
func (s *BaseClauseExpressionListener) ExitImplies(ctx *ImpliesContext) {}

// EnterExistsIndex is called when production ExistsIndex is entered.
func (s *BaseClauseExpressionListener) EnterExistsIndex(ctx *ExistsIndexContext) {}

// ExitExistsIndex is called when production ExistsIndex is exited.
func (s *BaseClauseExpressionListener) ExitExistsIndex(ctx *ExistsIndexContext) {}

// EnterExistsElement is called when production ExistsElement is entered.
func (s *BaseClauseExpressionListener) EnterExistsElement(ctx *ExistsElementContext) {}

// ExitExistsElement is called when production ExistsElement is exited.
func (s *BaseClauseExpressionListener) ExitExistsElement(ctx *ExistsElementContext) {}

// EnterPlainGoExpression is called when production PlainGoExpression is entered.
func (s *BaseClauseExpressionListener) EnterPlainGoExpression(ctx *PlainGoExpressionContext) {}

// ExitPlainGoExpression is called when production PlainGoExpression is exited.
func (s *BaseClauseExpressionListener) ExitPlainGoExpression(ctx *PlainGoExpressionContext) {}

// EnterForallIndex is called when production ForallIndex is entered.
func (s *BaseClauseExpressionListener) EnterForallIndex(ctx *ForallIndexContext) {}

// ExitForallIndex is called when production ForallIndex is exited.
func (s *BaseClauseExpressionListener) ExitForallIndex(ctx *ForallIndexContext) {}

// EnterForallElement is called when production ForallElement is entered.
func (s *BaseClauseExpressionListener) EnterForallElement(ctx *ForallElementContext) {}

// ExitForallElement is called when production ForallElement is exited.
func (s *BaseClauseExpressionListener) ExitForallElement(ctx *ForallElementContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BaseClauseExpressionListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BaseClauseExpressionListener) ExitIterator(ctx *IteratorContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseClauseExpressionListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseClauseExpressionListener) ExitCollection(ctx *CollectionContext) {}

// EnterCompleteGoExpression is called when production completeGoExpression is entered.
func (s *BaseClauseExpressionListener) EnterCompleteGoExpression(ctx *CompleteGoExpressionContext) {}

// ExitCompleteGoExpression is called when production completeGoExpression is exited.
func (s *BaseClauseExpressionListener) ExitCompleteGoExpression(ctx *CompleteGoExpressionContext) {}

// EnterGoExpression is called when production goExpression is entered.
func (s *BaseClauseExpressionListener) EnterGoExpression(ctx *GoExpressionContext) {}

// ExitGoExpression is called when production goExpression is exited.
func (s *BaseClauseExpressionListener) ExitGoExpression(ctx *GoExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseClauseExpressionListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseClauseExpressionListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseClauseExpressionListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseClauseExpressionListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterFunctionCallArguments is called when production functionCallArguments is entered.
func (s *BaseClauseExpressionListener) EnterFunctionCallArguments(ctx *FunctionCallArgumentsContext) {
}

// ExitFunctionCallArguments is called when production functionCallArguments is exited.
func (s *BaseClauseExpressionListener) ExitFunctionCallArguments(ctx *FunctionCallArgumentsContext) {}

// EnterSliceIndex is called when production sliceIndex is entered.
func (s *BaseClauseExpressionListener) EnterSliceIndex(ctx *SliceIndexContext) {}

// ExitSliceIndex is called when production sliceIndex is exited.
func (s *BaseClauseExpressionListener) ExitSliceIndex(ctx *SliceIndexContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseClauseExpressionListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseClauseExpressionListener) ExitNumber(ctx *NumberContext) {}
