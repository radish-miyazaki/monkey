package evaluator

import (
	"testing"

	"github.com/radish-miyazaki/monkey/ast"
	"github.com/radish-miyazaki/monkey/lexer"
	"github.com/radish-miyazaki/monkey/object"
	"github.com/radish-miyazaki/monkey/parser"
)

func TestDefineMacro(t *testing.T) {
	input := `
	let number = 1;
	let function = fn(x, y) { x + y };
	let mymacro = macro(x, y) { x + y; };
	`

	env := object.NewEnvironment()
	program := testParseProgram(input)

	DefineMacro(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("wrong number of statements. got=%d", len(program.Statements))
	}

	_, ok := env.Get("number")
	if ok {
		t.Fatalf("number should not be defined.")
	}

	_, ok = env.Get("function")
	if ok {
		t.Fatalf("function should not be defined.")
	}

	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("macro not in environment.")
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		t.Fatalf("object is not macro. got=%T (%+v)", obj, obj)
	}

	if len(macro.Parameters) != 2 {
		t.Fatalf("wrong number of macro parameters. got=%d", len(macro.Parameters))
	}

	if macro.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not `x`. got=%q", macro.Parameters[0])
	}

	if macro.Parameters[1].String() != "y" {
		t.Fatalf("parameter is not `x`. got=%q", macro.Parameters[1])
	}

	if macro.Body.String() != "(x + y)" {
		t.Fatalf("body is not %q. got=%q", "(x + y)", macro.Body.String())
	}
}

func TestExpandMacro(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
let infixExpression = macro() { quote(1 + 2); };
infixExpression();
`,
			`(1 + 2)`,
		},
		{
			`
let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); };
reverse(2 + 2, 10 - 5);
`,
			`(10 - 5) - (2 + 2)`,
		},
		{
			`
let unless = macro(cond, conseq, alter) {
	quote(if (!(unquote(cond))) {
		unquote(conseq);
	} else {
		unquote(alter);
	});
};

unless(10 > 5, puts("not greater"), puts("greater"));
`,
			`if (!(10 > 5)) {  puts("not greater") } else { puts("greater") }`,
		},
	}

	for _, tt := range tests {
		expected := testParseProgram(tt.expected)
		program := testParseProgram(tt.input)

		env := object.NewEnvironment()
		DefineMacro(program, env)
		expanded := ExpandMacro(program, env)

		if expanded.String() != expected.String() {
			t.Errorf("not equal. got=%q, want=%q", expanded.String(), expected.String())
		}
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
