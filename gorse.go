package gorse

import (
. "fmt"
  "reflect"
. "strings"
  "regexp"
)

func ToUnderscore(s string) string {
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

// Attributes: returns table name, sequence name, columns
func Attributes(base interface{}) (string, string, []string) {
  var pkPattern = regexp.MustCompile("(primary[_ ]key|pk|pkey)")
  var baseObjct, baseTable = findObjects(base)
  var basePrKey, table, sequence string
  var columns []string

  for i := 0; i < baseObjct.NumField(); i++ {
    var field = baseTable.Field(i)
    var line, found = field.Tag.Lookup("db")
    var tags []string
    if found {
      line = ToLower(line)
      tags = Split(line, ",")
      if pkPattern.MatchString(line) {
        basePrKey = ToUnderscore(field.Name)
      } else if len(tags[0]) > 0 {
        columns = append(columns, ToLower(tags[0]))
      } else {
        columns = append(columns, ToUnderscore(field.Name))
      }
    }
  }

  table = ToUnderscore(baseTable.Name())

  if len(basePrKey) > 0 {
    sequence = Sprintf("%s_%s", table, basePrKey)
  } else {
    sequence = Sprintf("%s_sq", table)
  }

  return table, sequence, append([]string{basePrKey}, columns ...)
}

// Attributes: returns values
func NonemptyAttributes(base interface{}) ([]string) {
  var baseObjct, baseTable = findObjects(base)
  var columns []string

  for i := 0; i < baseObjct.NumField(); i++ {
    var field = baseTable.Field(i)
    var name = ToUnderscore(field.Name)
    var value = baseObjct.Field(i)
    if !isEmpty(value) {
      columns = append(columns, name)
    }
  }

  return columns
}

func findObjects(base interface{}) (reflect.Value, reflect.Type) {
  return reflect.ValueOf(base).Elem(), reflect.ValueOf(base).Elem().Type()
}

func isEmpty(value reflect.Value) bool {
  // https://stackoverflow.com/questions/23555241/golang-reflection-how-to-get-zero-value-of-a-field-type
  var zero bool = true
  switch value.Kind() {
    case reflect.Func, reflect.Map, reflect.Slice:
      return value.IsNil()
    case reflect.Array:
      for i := 0; i < value.Len(); i++ {
        zero = zero && isEmpty(value.Index(i))
      }
      return zero
    case reflect.Struct:
      for i := 0; i < value.NumField(); i++ {
        zero = zero && isEmpty(value.Field(i))
      }
      return zero
  }
  // Compare other types directly:
  var empty = reflect.Zero(value.Type())
  return value.Interface() == empty.Interface()
}
