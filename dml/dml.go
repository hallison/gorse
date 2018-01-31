package dml

import (
. "fmt"
. "strings"
)

func Insert(table, id, sequence string, columns []string) string {
  return Sprintf("INSERT INTO %s (%s) VALUES (%s.NEXTVAL, :%s)", table, Join(columns, ", "), sequence, Join(columns[1:], ", :"))
}

func Select(table string, columns []string) string {
  return Sprintf("SELECT %s FROM %s", Join(columns, ", "), table)
}

func Update(table string, columns []string) string {
  var cols = columns[1:]
  for i, field := range(cols) {
    cols[i] = Sprintf("%s = :%s", field, field)
  }

  return Sprintf("UPDATE %s SET %s", table, Join(cols, ", "))
}

func Delete(table, column string, id interface{}) string {
  return Sprintf("DELETE %s WHERE (%s = %v)", table, column, id)
}

