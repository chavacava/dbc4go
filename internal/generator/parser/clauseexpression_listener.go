// Code generated from ClauseExpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ClauseExpression

import "github.com/antlr4-go/antlr/v4"

// ClauseExpressionListener is a complete listener for a parse tree produced by ClauseExpressionParser.
type ClauseExpressionListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterExprInParens is called when entering the ExprInParens production.
	EnterExprInParens(c *ExprInParensContext)

	// EnterImplies is called when entering the Implies production.
	EnterImplies(c *ImpliesContext)

	// EnterExistsIndex is called when entering the ExistsIndex production.
	EnterExistsIndex(c *ExistsIndexContext)

	// EnterExistsElement is called when entering the ExistsElement production.
	EnterExistsElement(c *ExistsElementContext)

	// EnterPlainGoExpression is called when entering the PlainGoExpression production.
	EnterPlainGoExpression(c *PlainGoExpressionContext)

	// EnterIff is called when entering the Iff production.
	EnterIff(c *IffContext)

	// EnterForallIndex is called when entering the ForallIndex production.
	EnterForallIndex(c *ForallIndexContext)

	// EnterForallElement is called when entering the ForallElement production.
	EnterForallElement(c *ForallElementContext)

	// EnterExistsIterator is called when entering the ExistsIterator production.
	EnterExistsIterator(c *ExistsIteratorContext)

	// EnterForallIterator is called when entering the ForallIterator production.
	EnterForallIterator(c *ForallIteratorContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterCompleteGoExpression is called when entering the completeGoExpression production.
	EnterCompleteGoExpression(c *CompleteGoExpressionContext)

	// EnterGoExpression is called when entering the goExpression production.
	EnterGoExpression(c *GoExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// EnterFunctionCallArguments is called when entering the functionCallArguments production.
	EnterFunctionCallArguments(c *FunctionCallArgumentsContext)

	// EnterSliceIndex is called when entering the sliceIndex production.
	EnterSliceIndex(c *SliceIndexContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitExprInParens is called when exiting the ExprInParens production.
	ExitExprInParens(c *ExprInParensContext)

	// ExitImplies is called when exiting the Implies production.
	ExitImplies(c *ImpliesContext)

	// ExitExistsIndex is called when exiting the ExistsIndex production.
	ExitExistsIndex(c *ExistsIndexContext)

	// ExitExistsElement is called when exiting the ExistsElement production.
	ExitExistsElement(c *ExistsElementContext)

	// ExitPlainGoExpression is called when exiting the PlainGoExpression production.
	ExitPlainGoExpression(c *PlainGoExpressionContext)

	// ExitIff is called when exiting the Iff production.
	ExitIff(c *IffContext)

	// ExitForallIndex is called when exiting the ForallIndex production.
	ExitForallIndex(c *ForallIndexContext)

	// ExitForallElement is called when exiting the ForallElement production.
	ExitForallElement(c *ForallElementContext)

	// ExitExistsIterator is called when exiting the ExistsIterator production.
	ExitExistsIterator(c *ExistsIteratorContext)

	// ExitForallIterator is called when exiting the ForallIterator production.
	ExitForallIterator(c *ForallIteratorContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitCompleteGoExpression is called when exiting the completeGoExpression production.
	ExitCompleteGoExpression(c *CompleteGoExpressionContext)

	// ExitGoExpression is called when exiting the goExpression production.
	ExitGoExpression(c *GoExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitFunctionCallArguments is called when exiting the functionCallArguments production.
	ExitFunctionCallArguments(c *FunctionCallArgumentsContext)

	// ExitSliceIndex is called when exiting the sliceIndex production.
	ExitSliceIndex(c *SliceIndexContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}
