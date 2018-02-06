package gorse_test

import (
  "testing"
. "github.com/hallison/gorse"
)

type Curso struct {
  ID   int            `db:"ID,PRIMARY KEY"`
  Nome string         `db:"NOME,REQUIRED"`
  Grau string         `db:"GRAU,"`
  DataInsercao string `db:"DATA_INSERCAO"`
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
    if ToUnderscore(name) != attr {
      t.Errorf("o campo %s deve ser convertido em %s", name, attr)
    }
  }
}

func TestDMLRawSql(t *testing.T) {
  var statements = map[string]string {
    "INSERT INTO curso (id, nome, grau, data_insercao) VALUES (curso_id.NEXTVAL, :NOME, :GRAU, :DATA_INSERCAO)": RawSqlInsert("curso", "id", "curso_id", COLUMNS),
    "SELECT id, nome, grau, data_insercao FROM curso": RawSqlSelect("curso", COLUMNS),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU), (data_insercao = :DATA_INSERCAO)": RawSqlUpdate("curso", COLUMNS[1:]),
    "DELETE curso": RawSqlDelete("curso"),
    "WHERE (id = :ID)": RawSqlWhere("id = :ID"),
    "(nome = :NOME) AND (grau = :GRAU) AND (data_insercao = :DATA_INSERCAO)": RawSqlLogicalAllColumns("and", COLUMNS[1:]),
    "(nome = :NOME) OR (grau = :GRAU) OR (data_insercao = :DATA_INSERCAO)": RawSqlLogicalAllColumns("or", COLUMNS[1:]),
    "ORDER BY nome": RawSqlOrderBy(COLUMNS[1:2]),
    "ORDER BY nome, grau": RawSqlOrderBy(COLUMNS[1:3]),
    "ORDER BY nome DESC": RawSqlDescOrderBy(COLUMNS[1:2]),
    "ORDER BY nome, grau DESC": RawSqlDescOrderBy(COLUMNS[1:3]),
  }

  for fix, test := range(statements) {
    if test != fix {
      t.Errorf("SQL/DML: '%s' deve ser igual a '%s'", test, fix)
    }
  }
}

func TestDMLTable(t *testing.T) {
  var table = NewTable(&Curso{})
  var curso = &Curso{ Nome: "Teste", Grau: "1", DataInsercao: "20180205" }
  var statements = map[string]string {
    "INSERT INTO curso (id, nome, grau, data_insercao) VALUES (curso_id.NEXTVAL, :NOME, :GRAU, :DATA_INSERCAO)": table.Insert(curso).Sql(),
    "INSERT INTO curso (id, nome, grau) VALUES (curso_id.NEXTVAL, :NOME, :GRAU)": table.Insert(&Curso{ Nome: "Outro Teste", Grau: "2" }).Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso": table.Select().Sql(),
    "SELECT nome, grau FROM curso": table.Select("nome", "grau").Sql(),
    "SELECT nome FROM curso WHERE (grau = :GRAU)": table.Select("nome").Where("grau = :GRAU").Sql(),
    "SELECT nome FROM curso WHERE (grau = :GRAU) AND (data_insercao = :DATA_INSERCAO)": table.Select("nome").Where("grau = :GRAU").And("data_insercao = :DATA_INSERCAO").Sql(),
    "SELECT nome FROM curso WHERE (grau = :GRAU) AND NOT (data_insercao = :DATA_INSERCAO)": table.Select("nome").Where("grau = :GRAU").AndNot("data_insercao = :DATA_INSERCAO").Sql(),
    "SELECT nome FROM curso WHERE (grau = :GRAU) OR (data_insercao = :DATA_INSERCAO)": table.Select("nome").Where("grau = :GRAU").Or("data_insercao = :DATA_INSERCAO").Sql(),
    "SELECT nome FROM curso WHERE (grau = :GRAU) OR NOT (data_insercao = :DATA_INSERCAO)": table.Select("nome").Where("grau = :GRAU").OrNot("data_insercao = :DATA_INSERCAO").Sql(),
    "SELECT nome FROM curso WHERE (grau IN (:GRAUS))": table.Select("nome").Where("grau IN (:GRAUS)").Sql(),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU), (data_insercao = :DATA_INSERCAO)": table.Update(curso).Sql(),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU), (data_insercao = :DATA_INSERCAO) WHERE (id = :ID)": table.Update(curso).Where("id = :ID").Sql(),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU)": table.Update(&Curso{ Nome: "Outro Teste", Grau: "2" }).Sql(),
    "UPDATE curso SET (nome = :NOME), (grau = :GRAU) WHERE (id = :ID)": table.Update(&Curso{ Nome: "Outro Teste", Grau: "2" }).Where("id = :ID").Sql(),
    "DELETE curso": table.Delete().Sql(),
    "DELETE curso WHERE (id = :ID)": table.Delete().Where("id = :ID").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso ORDER BY nome": table.Select().OrderBy("nome").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso ORDER BY nome, grau": table.Select().OrderBy("nome", "grau").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso ORDER BY nome DESC": table.Select().DescOrderBy("nome").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso ORDER BY nome, grau DESC": table.Select().DescOrderBy("nome", "grau").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso WHERE (grau = :GRAU) ORDER BY nome": table.Select().Where("grau = :GRAU").OrderBy("nome").Sql(),
    "SELECT id, nome, grau, data_insercao FROM curso WHERE (grau = :GRAU) AND (data_insercao = :DATA_INSERCAO) ORDER BY nome": table.Select().Where("grau = :GRAU").And("data_insercao = :DATA_INSERCAO").OrderBy("nome").Sql(),
  }

  for fix, test := range(statements) {
    if test != fix {
      t.Errorf("SQL/DML: '%s' deve ser igual a '%s'", test, fix)
    }
  }
}
