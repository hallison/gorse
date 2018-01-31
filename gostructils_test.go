package main_test

import (
  "testing"
. "git.prodam.am.gov.br/dinov/gostructils"
  "git.prodam.am.gov.br/dinov/gostructils/dml"
)

type Curso struct {
  ID   int
	Nome string
  Grau string
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

func TestDML(t *testing.T) {
  var columns = []string{ "id", "nome", "grau" }
  var statements = map[string]string {
    "INSERT INTO curso (id, nome, grau) VALUES (curso_sq.NEXTVAL, :nome, :grau)": dml.Insert("curso", "id", "curso_sq", columns),
    "SELECT id, nome, grau FROM curso": dml.Select("curso", columns),
    "UPDATE curso SET nome = :nome, grau = :grau": dml.Update("curso", columns),
    "DELETE curso WHERE (id = 1)": dml.Delete("curso", "id", 1),
  }

  for fix, test := range(statements) {
    if test != fix {
        t.Errorf("SQL/DML '%s' deve ser igual a '%s'", test, fix)
    }
  }
}
