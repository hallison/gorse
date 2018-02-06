package gorse

import (
. "fmt"
. "strings"
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

func RawSqlOrderBy(columns []string) string {
  return Sprintf("ORDER BY %s", Join(columns, ", "))
}

func RawSqlDescOrderBy(columns []string) string {
  return Sprintf("%s DESC", RawSqlOrderBy(columns))
}

type DML struct {
  HasClausule bool // WHERE, GROUP, HAVING, ORDER
  HasLogical  bool // AND, OR, NOT
// Relational bool // BETWEEN, LIKE, IN
  SQL []string // SELECT, INSERT, UPDATE, DELETE
}

func pairsOf(columns []string) []string {
  var fields = make([]string, len(columns))
  copy(fields, columns)
  for i, field := range(fields) {
    fields[i] = Sprintf("(%s = :%s)", field, ToUpper(field))
  }
  return fields
}
