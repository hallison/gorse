package gorse

import (
. "strings"
)

type Table struct {
  Name, Sequence string
  Columns []string
  DML
}

func NewTable(base interface{}) *Table {
  var name, sequence, columns = Attributes(base)
  return &Table {
    Name: name,
    Sequence: sequence,
    Columns: columns,
  }
}

func (table *Table) Sql() string {
  var sql = Join(table.SQL, " ")

  table.SQL = []string{}
  table.HasClausule = false
  table.HasLogical = false

  return sql
}

func (table *Table) Insert(base interface{}) *Table {
  var _, _, fields = Attributes(base)
  var columns = append([]string{fields[0]}, NonemptyAttributes(base)...)
  if !table.HasClausule {
    table.HasClausule = false
    table.SQL = append(table.SQL, RawSqlInsert(table.Name, fields[0], table.Sequence, columns))
  }

  return table
}

func (table *Table) Select(args ...string) *Table {
  if len(args) > 0 {
    table.SQL = append(table.SQL, RawSqlSelect(table.Name, args))
  } else {
    table.SQL = append(table.SQL, RawSqlSelect(table.Name, table.Columns))
  }

  table.HasClausule = true

  return table
}

func (table *Table) Update(base interface{}) *Table {
  table.HasClausule = true
  table.SQL = append(table.SQL, RawSqlUpdate(table.Name, NonemptyAttributes(base)))

  return table
}

func (table *Table) Delete() *Table {
  table.HasClausule = true
  table.SQL = append(table.SQL, RawSqlDelete(table.Name))

  return table
}

func (table *Table) Where(condition string) *Table {
  if table.HasClausule {
    table.HasLogical = true
    table.SQL = append(table.SQL, RawSqlWhere(condition))
  }

  return table
}

func (table *Table) And(condition string) *Table {
  if table.HasClausule && table.HasLogical {
    table.SQL = append(table.SQL, RawSqlLogical("and", condition))
  }

  return table
}

func (table *Table) AndNot(condition string) *Table {
  if table.HasClausule && table.HasLogical {
    table.SQL = append(table.SQL, RawSqlLogical("and not", condition))
  }

  return table
}

func (table *Table) Or(condition string) *Table {
  if table.HasClausule && table.HasLogical {
    table.SQL = append(table.SQL, RawSqlLogical("or", condition))
  }

  return table
}

func (table *Table) OrNot(condition string) *Table {
  if table.HasClausule && table.HasLogical {
    table.SQL = append(table.SQL, RawSqlLogical("or not", condition))
  }

  return table
}

func (table *Table) OrderBy(fields ...string) *Table {
  if table.HasClausule {
    table.SQL = append(table.SQL, RawSqlOrderBy(fields))
  }

  return table
}

func (table *Table) DescOrderBy(fields ...string) *Table {
  if table.HasClausule {
    table.SQL = append(table.SQL, RawSqlDescOrderBy(fields))
  }

  return table
}
