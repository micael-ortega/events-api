CREATE TABLE IF NOT EXISTS curso(
  id INTEGER PRIMARY KEY,
  curso TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS instrutor(
  id INTEGER PRIMARY KEY,
  nome TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS participante(
  id INTEGER PRIMARY KEY,
  nome TEXT NOT NULL,
  cpf TEXT NOT NULl,
  funcao TEXT NOT NULL,
  diretoria TEXT NOT NULL,
  empresa  TEXT NOT NULL,
  filial TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS evento(
 id INTEGER PRIMARY KEY,
 data_ini TEXT NOT NULL,
 data_fim TEXT NOT NULL,
 modalidade TEXT NOT NULL,
 duracao REAL NOT NULL,
 instrutor_id  INTEGER NOT NULL,
 curso_id INTEGER NOT NULL,
 FOREIGN KEY (instrutor_id) REFERENCES instrutor(id),
 FOREIGN KEY (curso_id) REFERENCES curso(id)
);

CREATE TABLE IF NOT EXISTS participacao_evento(
  id INTEGER PRIMARY KEY,
  evento_id INTEGER NOT NULL,
  participante_id INTEGER NOT NULL,
  presenca INTEGER NOT NULL,
  FOREIGN KEY (evento_id) REFERENCES evento(id),
  FOREIGN KEY (participante_id) REFERENCES participante(id)
);

