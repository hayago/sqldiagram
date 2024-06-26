package mysql

import (
	. "github.com/hayago/sqldiagram/internal/model"
	"github.com/pingcap/tidb/parser/ast"
)

func Parse(stmts []ast.StmtNode, current []Table) []Table {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.CreateTableStmt:
			current = parseCreateTable(s, current)
		case *ast.AlterTableStmt:
			current = parseAlterTable(s, current)
		case *ast.DropTableStmt:
			current = parseDropTable(s, current)
		}
	}

	return current
}
