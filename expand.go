package expand

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

type Selection struct {
	Start int
	End   int
}

func expandSelection(
	file *token.File,
	fileAst *ast.File,
	selection Selection,
) Selection {
	path, _ := astutil.PathEnclosingInterval(fileAst, file.Pos(selection.Start), file.Pos(selection.End))

	n := path[0]
	if len(path) >= 2 && n.Pos() == token.Pos(selection.Start) && n.End() == token.Pos(selection.End) {
		n = path[1]
	}

	// fmt.Printf("%T %d-%d\n", n, n.Pos(), n.End())

	return Selection{
		Start: int(n.Pos()),
		End:   int(n.End()),
	}
}
