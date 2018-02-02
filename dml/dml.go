package dml

import (
. "fmt"
. "strings"
. "git.prodam.am.gov.br/dinov/gostructils"
)

func RawSqlInsert(table, id, sequence string, columns []string) string {
  return Sprintf("INSERT INTO %s (%s) VALUES (%s.NEXTVAL, :%s)", table, Join(columns, ", "), sequence, ToUpper(Join(columns[1:], ", :")))
}

func Insert(base interface{}) string {
  var table, sequence, columns = Attributes(base)
  var id = columns[0]

  return RawSqlInsert(table, id, sequence, columns[1:])
}

func RawSqlSelect(table string, columns []string) string {
  return Sprintf("SELECT %s FROM %s", Join(columns, ", "), table)
}

func Select(base interface{}) string {
  var table, _, columns = Attributes(base)
  return RawSqlSelect(table, columns)
}

func RawSqlUpdate(table string, columns []string) string {
  return Sprintf("UPDATE %s SET %s", table, Join(pairsOf(columns), ", "))
}

func Update(base interface{}) string {
  var table, _, columns = Attributes(base)
  return RawSqlUpdate(table, columns)
}

func RawSqlDelete(table string) string {
  return Sprintf("DELETE %s", table)
}

func Delete(base interface{}) string {
  var table, _, _ = Attributes(base)
  return RawSqlDelete(table)
}

func RawSqlWhere(condition string) string {
  return Sprintf("WHERE (%s)", condition)
}

func RawSqlLogical(operator string, columns []string) string {
  return Sprintf("%s", Join(pairsOf(columns), Sprintf(" %s ", ToUpper(operator))))
}

func pairsOf(columns []string) []string {
  var fields = make([]string, len(columns))
  copy(fields, columns)
  for i, field := range(fields) {
    fields[i] = Sprintf("(%s = :%s)", field, ToUpper(field))
  }
  return fields
}
