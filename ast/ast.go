package ast

import "github.com/radish-miyazaki/monkey/token"

// Node は全てのASTノードが実装するインターフェース
type Node interface {
	TokenLiteral() string
}

// Statement は全ての文を表すASTノードのインターフェース
type Statement interface {
	Node
	statementNode()
}

// Expression は全ての式を表すASTノードのインターフェース
type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier // 束縛の左辺にあたる識別子
	Value Expression  // 束縛の右辺にあたる識別子
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier は識別子を表すASTノード
// 本来 let 文の識別子は値を生成しないが、ノードの種類を少なくするために Expression を実装している
type Identifier struct {
	Token token.Token
	Value string // 識別子名
}

func (i *Identifier) statementNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Program は全てのASTノードのルートノードとなるもの
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
