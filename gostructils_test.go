package main_test

import (
  "testing"
. "git.prodam.am.gov.br/dinov/gostructils"
)

type Curso struct {
  ID   int
	Nome string
}

func TestFields(t *testing.T) {
  var fields = Fields(&Curso{})

  if len(fields) == 0 {
		t.Errorf("struct deve ter pelo menos 1 atributo")
  }
}

func TestUndescore(t *testing.T) {
  var fields = map[string]string {
    "IDCurso": "id_curso",
    "CursoID": "curso_id",
    "EscolaCursoID": "escola_curso_id",
    "CSEscolaCursoID": "cs_escola_curso_id",
  }

  for name, attr := range(fields) {
    if Underscore(name) != attr {
      t.Errorf("o campo %s deve ser convertido em %s", name, attr)
    }
  }
}
