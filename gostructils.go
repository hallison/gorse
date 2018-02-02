package gostructils

import (
. "fmt"
  "reflect"
. "strings"
  "regexp"
)

func Underscore(s string) string {
  // https://gist.github.com/vermotr/dd9cfe74169234ef7380e8f32a8fbce9
  var camel = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return ToLower(Join(a, "_"))
}

// Attributes: returns table, primaryKey, sequence, columns
func Attributes(base interface{}) (string, string, []string) {
  var baseObjct = reflect.ValueOf(base).Elem()
  var baseTable = baseObjct.Type()
  var basePrKey = ""
  var table string
  var sequence string
  var columns []string

  for i := 0; i < baseObjct.NumField(); i++ {
    var field = baseTable.Field(i)
    var tag, found = field.Tag.Lookup("db")
    if found {
      if tag == "primary_key" {
        basePrKey = ToLower(field.Name)
      } else {
        columns = append(columns, ToLower(field.Name))
      }
    }
  }

  table = ToLower(baseTable.Name())

  if len(basePrKey) > 0 {
    sequence = Sprintf("%s_%s", table, basePrKey)
  } else {
    sequence = Sprintf("%s_sq", table)
  }

  return table, sequence, append([]string{basePrKey}, columns ...)
}
