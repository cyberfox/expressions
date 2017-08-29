package expressions

import (
	"testing"
)

func TestExpressions(t *testing.T) {
	all := GetExpressions("expressions.txt")
	expected := "1+(10-1)+-5"
	if all["example"].GetText() != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, all["example"].GetText())
	}
}

func TestEvaluation(t *testing.T) {
	expr := NewEvaluator()
	all := GetExpressions("expressions.txt")
	result := all["example"].Visit(expr)
	expected := int64(1 + (10 - 1) + -5)

	if result != expected {
		t.Errorf("Expected %d, got %d.", expected, result)
	}
}

func TestSecondExpression(t *testing.T) {
	expr := NewEvaluator()
	all := GetExpressions("expressions.txt")
	result := all["second"].Visit(expr)
	expected := int64(-1)

	if result != expected {
		t.Errorf("Expected %d, got %d.", expected, result)
	}
}
