package gostructils_test

import (
  "testing"
. "git.prodam.am.gov.br/dinov/gostructils"
  "git.prodam.am.gov.br/dinov/gostructils/dml"
)

type Curso struct {
  ID   int            `db:"primary_key"`
  Nome string         `db:"required"`
  Grau string         `db:""`
  DataInsercao string `db:""`
}

var TABLE = "curso"
var COLUMNS = []string {
  "id",
  "nome",
  "grau",
  "data_insercao",
}

func TestAttributes(t *testing.T) {
  var table, _, fields = Attributes(&Curso{})

  if len(fields) != len(COLUMNS) {
    t.Errorf("%v colunas carregadas de %v atribuídas", len(fields), len(COLUMNS))
  }

  if table != "curso" {
    t.Errorf("%s deve ser igual à 'curso'", table)
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
  var statements = map[string]string {
    "INSERT INTO curso (id, nome, grau, data_insercao) VALUES (curso_id.NEXTVAL, :NOME, :GRAU, :DATA_INSERCAO)": dml.RawSqlInsert("curso", "id", "curso_id", COLUMNS),
    "SELECT id, nome, grau, data_insercao FROM curso": dml.RawSqlSelect("curso", COLUMNS),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU), (data_insercao = :DATA_INSERCAO)": dml.RawSqlUpdate("curso", COLUMNS[1:]),
    "DELETE curso": dml.RawSqlDelete("curso"),
    "WHERE (id = :ID)": dml.RawSqlWhere("id = :ID"),
    "(nome = :NOME) AND (grau = :GRAU) AND (data_insercao = :DATA_INSERCAO)": dml.RawSqlLogical("and", COLUMNS[1:]),
    "(nome = :NOME) OR (grau = :GRAU) OR (data_insercao = :DATA_INSERCAO)": dml.RawSqlLogical("or", COLUMNS[1:]),
  }

  for fix, test := range(statements) {
    if test != fix {
      t.Errorf("SQL/DML '%s' deve ser igual a '%s'", test, fix)
    }
  }
}
