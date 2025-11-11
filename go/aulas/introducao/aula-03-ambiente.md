# Aula 3: Preparando seu Ambiente de Desenvolvimento

Chegou a hora de sujar as mãos! Nesta aula, vamos preparar tudo o que você precisa para começar a escrever, compilar e executar código Go no seu próprio computador. Um ambiente bem configurado é o primeiro passo para uma jornada de desenvolvimento produtiva.

## Passo 1: Instalar o Go

A primeira coisa a fazer é instalar o compilador e as ferramentas do Go.

1.  **Acesse o site oficial:** Vá para a página de downloads oficial do Go: [https://go.dev/dl/](https://go.dev/dl/)
2.  **Baixe o instalador:** Escolha o instalador apropriado para o seu sistema operacional (Windows, macOS ou Linux).
3.  **Execute o instalador:** Siga as instruções do instalador. Ele cuidará de tudo, incluindo a configuração da variável de ambiente `PATH`, que permite executar os comandos do Go de qualquer lugar no seu terminal.

## Passo 2: Verificar a Instalação

Para garantir que a instalação foi bem-sucedida, abra um novo terminal (ou Prompt de Comando no Windows) e digite o seguinte comando:

```bash
go version
```

Se tudo estiver correto, você verá uma saída parecida com esta (a versão pode ser diferente):

```
go version go1.22.3 darwin/amd64
```

Isso confirma que o Go está instalado e pronto para ser usado!

## Passo 3: Configurar um Editor de Código

Você pode escrever Go em qualquer editor de texto, mas usar um editor com suporte específico para a linguagem torna a vida muito mais fácil.

**Nossa recomendação:** **Visual Studio Code (VS Code)**
-   **É gratuito, poderoso e muito popular.**
-   Possui uma extensão oficial do Go que oferece funcionalidades incríveis.

**Como configurar o VS Code:**
1.  Se você ainda não tem, baixe e instale o [VS Code](https://code.visualstudio.com/).
2.  Abra o VS Code, vá para a aba de **Extensões** (ícone de blocos no menu lateral).
3.  Procure por "Go" e instale a extensão oficial publicada pela `Go Team at Google`.
4.  Após a instalação, a extensão pode pedir para instalar algumas ferramentas adicionais (como `gopls`, o servidor de linguagem, e `dlv`, o debugger). Clique em "Install All" para aceitar.

Com isso, você terá autocompletar, formatação automática ao salvar, depuração e muito mais!

*Alternativas populares incluem o **GoLand** da JetBrains (uma IDE paga e extremamente completa) ou editores como Vim e Emacs com as configurações adequadas.*

## Passo 4: Entendendo Módulos Go

Antigamente, o Go usava um sistema chamado `GOPATH` que forçava todos os projetos a ficarem dentro de uma única pasta. Isso mudou!

Hoje, o padrão moderno é usar **Módulos Go**. Com módulos, você pode criar um projeto em **qualquer pasta** do seu computador. Um módulo é uma coleção de pacotes Go que são versionados juntos como uma única unidade.

Para iniciar um novo projeto (módulo), você usa o comando `go mod init`. Vamos fazer um teste rápido que servirá para a nossa próxima aula:

1.  **Abra seu terminal.**
2.  **Crie uma pasta para seus projetos.** Vamos chamá-la de `projetos-go`.
    ```bash
    mkdir projetos-go
    cd projetos-go
    ```
3.  **Crie uma pasta para nosso primeiro programa.**
    ```bash
    mkdir ola-mundo
    cd ola-mundo
    ```
4.  **Inicie o módulo.**
    ```bash
    go mod init exemplo.com/ola-mundo
    ```
    *O nome `exemplo.com/ola-mundo` é o "caminho do módulo". É uma convenção usar um nome que se pareça com uma URL para garantir que ele seja único. Se você tivesse um projeto no GitHub, usaria algo como `github.com/seu-usuario/seu-projeto`.*

Ao executar esse comando, você verá que um novo arquivo chamado `go.mod` foi criado. Ele rastreia as dependências do seu projeto.

---

Excelente! Seu ambiente está pronto e configurado. Você tem o Go instalado, um editor de código inteligente e um novo projeto (módulo) inicializado.

Na próxima aula, vamos finalmente escrever e executar nosso primeiro programa em Go: o famoso "Olá, Mundo!".

Até lá!
