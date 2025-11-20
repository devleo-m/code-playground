1 - 

select titulo from livros

2 - 

select titulo, ano_publicacao 
from livros
where ano_publicacao > 2000

3 - 

SELECT *
FROM livros l
WHERE l.quantidade_disponivel > 0
ORDER BY l.titulo asc

4 - 

SELECT *
FROM livros l
WHERE l.quantidade_disponivel > 0
ORDER BY l.ano_publicacao ASC 
LIMIT 3

5 - 

SELECT a.nome, a.nacionalidade
FROM autores a
WHERE a.nacionalidade = 'Brasileiro'


Exercício 2:

1 - 

SELECT *
FROM livros l
WHERE l.ano_publicacao > 1990 AND l.ano_publicacao < 2000

2 - 

SELECT *
FROM livros l
WHERE l.quantidade_disponivel >= 1 AND l.quantidade_disponivel <= 5

3 - 

SELECT *
FROM autores a 
WHERE a.nacionalidade = 'Brasileiro' OR a.nacionalidade = 'Americano'


4 - 

SELECT *
FROM livros l
WHERE NOT l.categoria_id = 1 AND l.quantidade_disponivel > 1

5 - 

SELECT *
FROM livros l
WHERE l.titulo like 'D%'

// esse eu tive que colar na aula 2 infelizmente na parte do like

Exercício 3: Operadores LIKE e Padrões

1 - 

SELECT *
FROM livros l
WHERE l.titulo like 'Dom%'

2 - 

SELECT *
FROM autores a
WHERE a.nome like '%Silva%'

3 - 

SELECT *
FROM livros l 
WHERE l.titulo like '%o'

4 - 

SELECT *
FROM livros l 
WHERE LENGTH(l.titulo) = 10

exercicio 4
1-

SELECT *
FROM livros l 
WHERE l.ano_publicacao = null

2-


SELECT *
FROM livros l 
WHERE l.ano_publicacao not null

3 -

SELECT *
FROM autores a 
where a.nacionalidade is ''

entendi a charada aqui, quando eh string ou texto nao existe null
entao dessa forma se eu escrever null ele vira string de qualquer forma
ou seja, eu tenho que filtrar caso esse texto esteja vazio

4-

SELECT *
FROM livros l 
where l.editora is not ''

5 - 

INSERT INTO autores(nome, nacionalidade, data_nascimento)
VALUES ('Carlos Drummond de Andrade', 'Brasileiro', '1902-10-31')

