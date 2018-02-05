package dml

import (
. "fmt"
. "strings"
. "git.prodam.am.gov.br/dinov/gostructils"
)

func RawSqlInsert(table, id, sequence string, columns []string) string {
  return Sprintf("INSERT INTO %s (%s) VALUES (%s.NEXTVAL, :%s)", table, Join(columns, ", "), sequence, ToUpper(Join(columns[1:], ", :")))
}

func RawSqlSelect(table string, columns []string) string {
  return Sprintf("SELECT %s FROM %s", Join(columns, ", "), table)
}

func RawSqlUpdate(table string, columns []string) string {
  return Sprintf("UPDATE %s SET %s", table, Join(pairsOf(columns), ", "))
}

func RawSqlDelete(table string) string {
  return Sprintf("DELETE %s", table)
}

func RawSqlWhere(condition string) string {
  return Sprintf("WHERE (%s)", condition)
}

func RawSqlLogical(operator string, condition string) string {
  return Sprintf("%s (%s)", ToUpper(operator), condition)
}

func RawSqlLogicalAllColumns(operator string, columns []string) string {
  return Sprintf("%s", Join(pairsOf(columns), Sprintf(" %s ", ToUpper(operator))))
}

// var statement, condition, aggregator string
type DML struct {
  Statement  string
  Clausule   string // WHERE, GROUP, HAVING, ORDER
  HasClausule bool
  HasLogical  bool // AND, OR, NOT
// Relational bool // BETWEEN, LIKE, IN
  SQL string
}

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
  if table.HasClausule && table.HasLogical {
    table.SQL = Sprintf("%s %s", table.Statement, table.Clausule)
  } else {
    table.SQL = table.Statement
  }

  table.HasClausule = false
  table.HasLogical = false

  return table.SQL
}

func (table *Table) Insert(base interface{}) *Table {
  var _, _, fields = Attributes(base)
  var columns = append([]string{fields[0]}, NonemptyAttributes(base)...)
  if !table.HasClausule {
    table.HasClausule = false
    table.Statement = RawSqlInsert(table.Name, fields[0], table.Sequence, columns)
  }

  return table
}

func (table *Table) Select(args ...string) *Table {
  if len(args) > 0 {
    table.Statement = RawSqlSelect(table.Name, args)
  } else {
    table.Statement = RawSqlSelect(table.Name, table.Columns)
  }

  table.HasClausule = true

  return table
}

func (table *Table) Update(base interface{}) *Table {
  table.HasClausule = true
  table.Statement = RawSqlUpdate(table.Name, NonemptyAttributes(base))

  return table
}

func (table *Table) Delete() *Table {
  table.HasClausule = true
  table.Statement = RawSqlDelete(table.Name)

  return table
}

func (table *Table) Where(condition string) *Table {
  if table.HasClausule {
    table.HasLogical = true
    table.Clausule = RawSqlWhere(condition)
  }

  return table
}

func (table *Table) And(condition string) *Table {
  if len(table.Statement) > 0 && table.HasClausule && table.HasLogical {
    table.Clausule = Sprintf("%s %s", table.Clausule, RawSqlLogical("and", condition))
  }

  return table
}

func (table *Table) Or(condition string) *Table {
  if len(table.Statement) > 0 && table.HasClausule && table.HasLogical {
    table.Clausule = Sprintf("%s %s", table.Clausule, RawSqlLogical("or", condition))
  }

  return table
}

func pairsOf(columns []string) []string {
  var fields = make([]string, len(columns))
  copy(fields, columns)
  for i, field := range(fields) {
    fields[i] = Sprintf("(%s = :%s)", field, ToUpper(field))
  }
  return fields
}
