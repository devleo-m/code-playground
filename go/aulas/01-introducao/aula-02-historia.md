# Aula 2: Uma Breve História do Go

Olá novamente! Bem-vindo(a) à nossa segunda aula.

Depois de entender o "o quê" e o "porquê" do Go, é fascinante olharmos para o "quando" e o "quem". Conhecer a história de uma linguagem nos ajuda a compreender as decisões de design e sua filosofia. Vamos embarcar nessa viagem no tempo!

## Os Criadores e a Motivação

A história do Go começa dentro dos muros do Google, por volta de 2007. Três engenheiros lendários estavam frustrados com a complexidade e a lentidão do desenvolvimento de software para os sistemas massivos da empresa. Esses engenheiros eram:

-   **Robert Griesemer:** Trabalhou em compiladores e sistemas distribuídos.
-   **Rob Pike:** Membro da equipe original do Unix na Bell Labs e co-criador do UTF-8.
-   **Ken Thompson:** Também da Bell Labs, é simplesmente um dos criadores do Unix e da linguagem C.

Eles enfrentavam problemas como:
-   **Tempos de compilação de horas:** Projetos em C++ no Google eram tão grandes que levavam uma eternidade para compilar.
-   **Gerenciamento de dependências complicado.**
-   **Código complexo e difícil de manter.**
-   **Falta de um bom suporte para concorrência** em hardware multi-core.

A solução? Criar uma nova linguagem que combinasse a **eficiência e segurança** de uma linguagem compilada como C++ com a **simplicidade e produtividade** de uma linguagem dinâmica como Python.

## Linha do Tempo e Marcos Importantes

-   **2007:** O design da linguagem começa como um projeto "20% time" no Google (um programa que permitia aos engenheiros dedicar 20% do seu tempo a projetos paralelos).
-   **2009:** Go é anunciado publicamente e seu código-fonte se torna aberto. A comunidade externa começa a se envolver e a contribuir.
-   **2012:** Lançamento do **Go 1.0**. Esta foi a primeira versão estável e, com ela, veio uma promessa crucial: a garantia de compatibilidade com versões futuras. Isso significa que um código escrito em Go 1.0 (com raras exceções) compilaria e funcionaria com Go 1.1, 1.2, e assim por diante.
-   **Go 1.5 (2015):** O compilador, que antes era escrito em C, foi totalmente reescrito em Go. A linguagem se tornou "autossuficiente" (self-hosted).
-   **Go 1.11 (2018):** Introdução dos **Módulos Go (`go mod`)**, o sistema de gerenciamento de dependências oficial, resolvendo um dos maiores desafios do ecossistema até então.
-   **Go 1.18 (2022):** Após anos de discussão e protótipos, os **Genéricos (Generics)** são finalmente adicionados à linguagem, permitindo escrever código mais flexível e reutilizável sem sacrificar a segurança de tipos.

## O Mascote: O Gopher

![Go Gopher](https://go.dev/blog/gopher/gopher.png)

Você não pode falar de Go sem mencionar seu adorável mascote, o Gopher, desenhado por Renee French. Ele se tornou um símbolo icônico e amado da comunidade.

---

E essa é a história de como o Go nasceu da necessidade de resolver problemas de engenharia de software em grande escala. É uma linguagem pragmática, criada por engenheiros para engenheiros.

Na próxima aula, vamos colocar a mão na massa de verdade: **preparar nosso ambiente de desenvolvimento** para que você possa começar a escrever seus próprios programas em Go!

Até lá!
