CREATE SCHEMA fabrica;

USE fabrica;

CREATE TABLE materia_prima(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    nome VARCHAR(45) NOT NULL,
    estoque INT NOT NULL
);

CREATE TABLE produto(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    nome VARCHAR(45) NOT NULL,
    valor DECIMAL(15,2) NOT NULL
);

CREATE TABLE insumo(
    id_produto INT UNSIGNED,
    id_materia_prima INT UNSIGNED,
    quantidade INT NOT NULL,
    PRIMARY KEY produto_materia_prima (`id_produto`,`id_materia_prima`),
    FOREIGN KEY(id_produto) REFERENCES produto(id),
    FOREIGN KEY(id_materia_prima) REFERENCES materia_prima(id)
);