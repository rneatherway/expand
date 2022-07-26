package expand

import (
	"go/ast"
	"go/token"
)

type Selection struct {
	StartLine int
	StartCol  int
	EndLine   int
	EndCol    int
}

func expandSelection(
	fileSet *token.FileSet,
	file *ast.File,
	s Selection,
) Selection {
	var nextNode ast.Node
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		start := fileSet.PositionFor(n.Pos(), false)
		end := fileSet.PositionFor(n.End(), false)

		if start.Line > s.StartLine ||
			(start.Line == s.StartLine && start.Column > s.StartCol) ||
			end.Line < s.EndLine ||
			(end.Line == s.EndLine && end.Column < s.EndCol) ||
			(start.Line == s.StartLine && start.Column == s.StartCol && end.Line == s.EndLine && end.Column == s.EndCol) {

			return false
		}

		// fmt.Printf("%T %d:%d-%d:%d\n", n, start.Line, start.Column, end.Line, end.Column)

		nextNode = n
		return true
	})

	return Selection{
		StartLine: fileSet.PositionFor(nextNode.Pos(), false).Line,
		StartCol:  fileSet.PositionFor(nextNode.Pos(), false).Column,
		EndLine:   fileSet.PositionFor(nextNode.End(), false).Line,
		EndCol:    fileSet.PositionFor(nextNode.End(), false).Column,
	}
}
