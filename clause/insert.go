package clause

import (
	"gorm.io/gorm/clause"
)

type Insert struct {
	Insert clause.Insert
	Values clause.Values
	From   clause.From
}

func (i Insert) Name() string {
	return "INSERT"
}

func (i Insert) Build(builder clause.Builder) {
	i.Insert.Build(builder)
	_ = builder.WriteByte(' ')
	if len(i.Values.Columns) > 0 {
		_ = builder.WriteByte('(')
		for idx, column := range i.Values.Columns {
			if idx > 0 {
				_ = builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
		_ = builder.WriteByte(')')

		_, _ = builder.WriteString(" SELECT ")
		clause.Select{Columns: i.Values.Columns}.Build(builder)
		_, _ = builder.WriteString(" FROM (")
		for idx, values := range i.Values.Values {
			if idx > 0 {
				_, _ = builder.WriteString(" UNION ")
			}
			_, _ = builder.WriteString("SELECT ")
			for j, v := range values {
				if j > 0 {
					_ = builder.WriteByte(',')
				}
				builder.AddVar(builder, v)
				col := clause.Column{Alias: i.Values.Columns[j].Name}
				builder.WriteQuoted(col)
			}
			_, _ = builder.WriteString(" FROM ")
			i.From.Build(builder)
		}
		_ = builder.WriteByte(')')
	}
}

func (i Insert) MergeClause(clause *clause.Clause) {
	clause.Name = i.Name()
	clause.Expression = i
}
