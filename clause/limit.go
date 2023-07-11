package clause

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// RewriteLimitOld only support limit.
func RewriteLimitOld(c clause.Clause, builder clause.Builder) {
	if limitExp, ok1 := c.Expression.(clause.Limit); ok1 {
		if stmt, ok2 := builder.(*gorm.Statement); ok2 {
			buf := strings.Builder{}
			if _, ok3 := stmt.Clauses["ORDER BY"]; ok3 {
				oldSql := stmt.SQL.String()
				idx := strings.LastIndex(strings.ToUpper(oldSql), "ORDER BY")
				stmt.SQL.Reset()
				stmt.SQL.WriteString(strings.TrimSuffix(oldSql[:idx], " "))
				buf.WriteString(oldSql[idx:])
			}

			if limit := limitExp.Limit; limit != nil && *limit > 0 {
				if _, ok := stmt.Clauses["WHERE"]; !ok {
					_, _ = stmt.SQL.WriteString(" WHERE ")
				} else {
					_, _ = stmt.SQL.WriteString(" AND ")
				}
				_, _ = stmt.SQL.WriteString("ROWNUM <= ")
				_, _ = stmt.SQL.WriteString(strconv.Itoa(*limit))
				_ = stmt.SQL.WriteByte(' ')
			}

			_, _ = stmt.SQL.WriteString(buf.String())
		}
	}
}

// RewriteLimitNew oracle >= 12c support <offset xx rows fetch next xx rows only>
func RewriteLimitNew(c clause.Clause, builder clause.Builder) {
	if limitExp, ok1 := c.Expression.(clause.Limit); ok1 {
		if stmt, ok2 := builder.(*gorm.Statement); ok2 {
			if _, ok3 := stmt.Clauses["ORDER BY"]; !ok3 {
				// using default order
				if stmt.Schema != nil && stmt.Schema.PrioritizedPrimaryField != nil {
					_, _ = builder.WriteString("ORDER BY ")
					builder.WriteQuoted(stmt.Schema.PrioritizedPrimaryField.DBName)
					_, _ = builder.WriteString(" ")
				}
			}
		}

		if offset := limitExp.Offset; offset > 0 {
			_, _ = builder.WriteString("OFFSET ")
			_, _ = builder.WriteString(strconv.Itoa(offset))
			_, _ = builder.WriteString(" ROWS ")
		}

		if limit := limitExp.Limit; limit != nil && *limit > 0 {
			_, _ = builder.WriteString("FETCH NEXT ")
			_, _ = builder.WriteString(strconv.Itoa(*limit))
			_, _ = builder.WriteString(" ROWS ONLY ")
		}
	}
}
