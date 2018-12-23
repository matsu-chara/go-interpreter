package evaluator

import (
	"github.com/matsu-chara/go-interpreter/monkey/ast"
	"github.com/matsu-chara/go-interpreter/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
