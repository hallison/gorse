package main

import (
  "reflect"
  "strings"
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
	return strings.ToLower(strings.Join(a, "_"))
}

func Fields(base interface{}) []string {
  var structure = reflect.ValueOf(base).Elem()
  var attribute = structure.Type()
  var columns []string

  for i := 0; i < structure.NumField(); i++ {
    columns = append(columns, attribute.Field(i).Name)
  }

  return columns
}
