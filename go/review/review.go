package main

/*
 * ============================================================================
 * REVIEW COMPLETO - CURSO DE GO
 * ============================================================================
 * Este arquivo contém um resumo completo e comentado de todas as aulas
 * do curso de Go, desde a introdução até tópicos avançados.
 * Use este arquivo como referência rápida para revisar os conceitos essenciais.
 * ============================================================================
 */

import "fmt"

func main() {
	fmt.Println("=== REVIEW COMPLETO DO CURSO DE GO ===")
	fmt.Println("Este arquivo contém resumos comentados de todas as aulas.")
	fmt.Println("Navegue pelos comentários para revisar os conceitos.")
}

/*
 * ============================================================================
 * MÓDULO 1: INTRODUÇÃO AO GO (Aulas 01-05)
 * ============================================================================
 */

/*
 * AULA 01: O QUE É GO E POR QUE USÁ-LO?
 * --------------------------------------
 * Go (Golang) é uma linguagem compilada, estaticamente tipada, criada no Google.
 *
 * PRINCIPAIS CARACTERÍSTICAS:
 * - Simplicidade: Sintaxe limpa, fácil de aprender
 * - Concorrência: Goroutines e channels nativos
 * - Performance: Código compilado, rápido como C++
 * - Garbage Collector: Gerencia memória automaticamente
 *
 * ONDE GO SE DESTACA:
 * - Serviços de backend e microsserviços
 * - Ferramentas CLI (linha de comando)
 * - Sistemas distribuídos e cloud computing
 * - Docker, Kubernetes, Terraform são escritos em Go
 *
 * VANTAGENS:
 * - Binário único executável (sem dependências)
 * - Compilação rápida
 * - Biblioteca padrão robusta
 * - Ferramentas poderosas (go tool)
 */

/*
 * AULA 02: HISTÓRIA DO GO
 * ------------------------
 * CRIADORES: Robert Griesemer, Rob Pike, Ken Thompson (Google, 2007)
 *
 * MOTIVAÇÃO:
 * - Frustração com complexidade de C++
 * - Tempos de compilação muito longos
 * - Falta de suporte para concorrência multi-core
 *
 * MARCOS IMPORTANTES:
 * - 2007: Início do design
 * - 2009: Anúncio público e código aberto
 * - 2012: Go 1.0 (compatibilidade garantida)
 * - 2015: Compilador reescrito em Go (self-hosted)
 * - 2018: Go 1.11 - Módulos Go (go mod)
 * - 2022: Go 1.18 - Generics adicionados
 *
 * MASCOTE: Gopher (desenhado por Renee French)
 */

/*
 * AULA 03: PREPARANDO O AMBIENTE
 * ------------------------------
 * INSTALAÇÃO:
 * 1. Baixar de https://go.dev/dl/
 * 2. Verificar: go version
 * 3. Configurar editor (VS Code com extensão Go)
 *
 * MÓDULOS GO:
 * - Sistema moderno de gerenciamento de dependências
 * - Substituiu o antigo GOPATH
 * - Permite projetos em qualquer pasta
 * - Comando: go mod init exemplo.com/projeto
 * - Cria arquivo go.mod
 */

/*
 * AULA 04: OLÁ, MUNDO!
 * --------------------
 * ESTRUTURA BÁSICA DE UM PROGRAMA GO:
 *
 * package main          // Pacote executável
 * import "fmt"         // Importar biblioteca padrão
 * func main() {        // Ponto de entrada
 *     fmt.Println("Olá, Mundo!")
 * }
 *
 * CONCEITOS:
 * - package: Organiza código em módulos
 * - package main: Especial, cria executável
 * - import: Inclui pacotes externos
 * - func main(): Função principal, executada primeiro
 * - fmt.Println: Imprime no console
 *
 * EXECUTAR:
 * - go run arquivo.go  (compila e executa)
 * - go build arquivo.go  (apenas compila)
 */

/*
 * AULA 05: COMANDOS GO
 * --------------------
 * COMANDOS ESSENCIAIS:
 *
 * go run    - Compila e executa
 * go build  - Compila para binário
 * go install - Compila e instala em $GOPATH/bin
 * go fmt    - Formata código (gofmt)
 * go mod    - Gerencia módulos e dependências
 *   - go mod init    - Inicia novo módulo
 *   - go mod tidy    - Limpa dependências
 *   - go get         - Adiciona dependência
 * go test   - Executa testes
 * go vet    - Análise estática de código
 * go doc    - Exibe documentação
 * go help   - Ajuda sobre comandos
 */

/*
 * ============================================================================
 * MÓDULO 2: VARIÁVEIS E CONSTANTES (Aulas 01-04)
 * ============================================================================
 */

/*
 * AULA 01: DECLARANDO VARIÁVEIS
 * ------------------------------
 * DUAS FORMAS PRINCIPAIS:
 *
 * 1. var (explícito):
 *    var idade int
 *    var nome string = "João"
 *    var altura = 1.75  // Inferência de tipo
 *
 * 2. := (declaração curta - DENTRO de funções):
 *    cidade := "São Paulo"
 *    populacao := 12000000
 *
 * REGRAS:
 * - Fora de funções: SÓ pode usar var
 * - Dentro de funções: Prefira :=
 * - Para modificar variável existente: use = (não :=)
 * - Use var quando precisar de valor zero inicial
 *
 * INFERÊNCIA DE TIPO:
 * - Go "adivinha" o tipo baseado no valor
 * - Mais limpo e idiomático
 */

/*
 * AULA 02: CONSTANTES E IOTA
 * --------------------------
 * CONSTANTES:
 * - Valores imutáveis (não podem ser alterados)
 * - Definidas em tempo de compilação
 * - Sintaxe: const Nome = valor
 *
 * IOTA:
 * - Gerador de sequências numéricas
 * - Começa em 0, incrementa a cada declaração
 * - Útil para enums e flags
 *
 * EXEMPLO:
 * const (
 *     Domingo = iota  // 0
 *     Segunda         // 1
 *     Terca           // 2
 * )
 *
 * IOTA COM EXPRESSÕES:
 * const (
 *     FlagLeitura  = 1 << iota  // 1 (2^0)
 *     FlagEscrita   = 1 << iota  // 2 (2^1)
 *     FlagExecucao  = 1 << iota  // 4 (2^2)
 * )
 */

/*
 * AULA 03: VALORES ZERO
 * ----------------------
 * CONCEITO FUNDAMENTAL:
 * - Toda variável declarada SEM valor inicial recebe um "valor zero"
 * - Previsível e seguro (não há "lixo" de memória)
 *
 * VALORES ZERO POR TIPO:
 * - int, float: 0
 * - bool: false
 * - string: "" (string vazia)
 * - pointers, slices, maps, channels: nil
 *
 * VANTAGENS:
 * - Segurança: Elimina bugs de variáveis não inicializadas
 * - Simplicidade: Não precisa inicializar explicitamente
 * - Previsibilidade: Comportamento consistente
 */

/*
 * AULA 04: ESCOPO E SOMBREAMENTO
 * -------------------------------
 * ESCOPO: Região onde uma variável é visível
 *
 * NÍVEIS DE ESCOPO:
 * 1. Universo: Tipos e funções pré-declaradas
 * 2. Pacote: Variáveis no nível do arquivo
 * 3. Função: Parâmetros e variáveis locais
 * 4. Bloco: Variáveis em if, for, etc.
 *
 * SOMBREAMENTO (SHADOWING):
 * - Variável interna "esconde" variável externa com mesmo nome
 * - Pode causar bugs sutis
 * - Cuidado com := em escopos aninhados
 *
 * EXEMPLO DE BUG:
 * var numero = 100
 * if true {
 *     numero := 200  // Cria NOVA variável (sombra a externa)
 *     // numero externa não foi modificada!
 * }
 */

/*
 * ============================================================================
 * MÓDULO 3: TIPOS DE DADOS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: TIPOS DE DADOS FUNDAMENTAIS
 * -------------------------------------
 * INTEGERS (INTEIROS):
 * - Signed: int8, int16, int32, int64, int (depende da plataforma)
 * - Unsigned: uint8, uint16, uint32, uint64, uint
 * - int/uint: Tamanho depende da arquitetura (32 ou 64 bits)
 *
 * FLOATING POINTS:
 * - float32: ~7 dígitos de precisão
 * - float64: ~15-17 dígitos (PADRÃO)
 * - ⚠️ NÃO use para cálculos financeiros (erros de precisão)
 *
 * COMPLEX NUMBERS:
 * - complex64, complex128
 * - Útil para processamento de sinais, matemática
 *
 * BOOLEAN:
 * - bool: true ou false
 * - Valor zero: false
 *
 * RUNES:
 * - Alias para int32
 * - Representa ponto de código Unicode
 * - Permite trabalhar com caracteres internacionais, emojis
 *
 * STRINGS:
 * - Raw strings: `texto` (backticks) - literal, sem escape
 * - Interpreted: "texto" (aspas) - processa \n, \t, etc.
 * - Imutáveis em Go
 *
 * CONVERSÃO DE TIPOS:
 * - Go exige conversão EXPLÍCITA
 * - Sintaxe: Tipo(valor)
 * - Exemplo: int(3.14) → 3 (trunca)
 * - Para string: use strconv (Itoa, Atoi, etc.)
 */

/*
 * ============================================================================
 * MÓDULO 4: TIPOS COMPOSTOS (Aulas 01-08)
 * ============================================================================
 */

/*
 * AULA 01: ARRAYS
 * ---------------
 * CARACTERÍSTICAS:
 * - Tamanho FIXO (definido na declaração)
 * - Tipo inclui tamanho: [5]int ≠ [10]int
 * - Value type (copia quando atribuído)
 *
 * DECLARAÇÃO:
 * var numeros [5]int
 * notas := [4]float64{7.5, 8.0, 9.2, 6.8}
 * dias := [...]string{"Seg", "Ter"}  // Tamanho automático
 *
 * ACESSO:
 * - Indexação começa em 0
 * - len(array) retorna tamanho
 * - array[indice] = valor
 *
 * LIMITAÇÕES:
 * - Não pode crescer/diminuir
 * - Cópia custosa para arrays grandes
 * - Raramente usado (prefira slices)
 *
 * QUANDO USAR:
 * - Tamanho fixo conhecido (dias da semana, coordenadas)
 * - Performance crítica com tamanho pequeno
 */

/*
 * AULA 02: SLICES
 * ---------------
 * CARACTERÍSTICAS:
 * - Tamanho DINÂMICO (pode crescer)
 * - Reference type (compartilha array subjacente)
 * - Mais usado que arrays (95% dos casos)
 *
 * COMPONENTES:
 * - Ponteiro: Aponta para array subjacente
 * - Length: Número de elementos
 * - Capacity: Capacidade máxima antes de realocar
 *
 * DECLARAÇÃO:
 * numeros := make([]int, 5)        // length=5, cap=5
 * numeros := make([]int, 5, 10)    // length=5, cap=10
 * frutas := []string{"Maçã", "Banana"}
 *
 * OPERAÇÕES:
 * - append(slice, elementos...) - Adiciona elementos
 * - copy(destino, origem) - Copia elementos
 * - slice[inicio:fim] - Cria sub-slice
 * - len(slice), cap(slice) - Tamanho e capacidade
 *
 * ⚠️ IMPORTANTE:
 * - Sub-slices compartilham memória!
 * - append() retorna novo slice (sempre atribua!)
 * - Pré-aloque capacity quando souber tamanho aproximado
 *
 * PERFORMANCE:
 * - Pré-alocação evita múltiplas realocações
 * - make([]int, 0, 1000) é muito mais rápido que []int{}
 */

/*
 * AULA 08: MAPS
 * -------------
 * CARACTERÍSTICAS:
 * - Estrutura chave-valor (dicionário)
 * - Chaves únicas
 * - Busca O(1) em média
 * - Desordenado (ordem aleatória)
 * - Reference type
 *
 * DECLARAÇÃO:
 * var mapa map[string]int        // nil (não pode escrever!)
 * mapa := make(map[string]int)   // vazio (pronto para usar)
 * cores := map[string]string{"vermelho": "#FF0000"}
 *
 * OPERAÇÕES:
 * - mapa[chave] = valor          // Adicionar/atualizar
 * - valor := mapa[chave]        // Buscar (retorna zero se não existir)
 * - valor, ok := mapa[chave]    // Verificar existência
 * - delete(mapa, chave)         // Deletar
 * - len(mapa)                   // Tamanho
 *
 * CHAVES PERMITIDAS:
 * - Tipos comparáveis: string, int, float, bool, struct, array
 * - NÃO pode: slice, map, function
 *
 * PADRÕES COMUNS:
 * - Contador: mapa[chave]++
 * - Cache/Memoization
 * - Agrupar dados
 * - Simular Set: map[T]bool
 *
 * CONCORRÊNCIA:
 * - Maps NÃO são thread-safe
 * - Use sync.Mutex ou sync.Map para acesso concorrente
 */

/*
 * ============================================================================
 * MÓDULO 5: DOCUMENTAÇÃO E COMANDOS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: DOCUMENTAÇÃO E COMANDOS
 * ---------------------------------
 * COMENTÁRIOS DE DOCUMENTAÇÃO:
 * - Devem começar imediatamente antes da declaração
 * - Devem começar com o nome do item
 * - Geram documentação automática
 *
 * EXEMPLO:
 * // CalculaMedia calcula a média de uma lista de números.
 * // Retorna erro se a lista estiver vazia.
 * func CalculaMedia(numeros []float64) (float64, error) { }
 *
 * COMANDOS:
 * - go doc pacote          - Ver documentação
 * - go doc pacote.Funcao   - Documentação específica
 * - go doc -src            - Mostra código-fonte
 * - godoc -http=:6060      - Servidor web de documentação
 *
 * BOAS PRÁTICAS:
 * - Documente funções/métodos públicos
 * - Documente tipos públicos
 * - Documente pacotes (pelo menos um arquivo)
 * - Use exemplos quando útil
 */

/*
 * ============================================================================
 * MÓDULO 6: STRUCTS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: STRUCTS
 * ----------------
 * CONCEITO:
 * - Agrupa dados relacionados sob um tipo
 * - Modela entidades do mundo real
 * - Base para programação orientada a objetos em Go
 *
 * DECLARAÇÃO:
 * type Pessoa struct {
 *     Nome   string
 *     Idade  int
 *     Email  string
 * }
 *
 * CRIAÇÃO:
 * var p Pessoa                    // Valores zero
 * p := Pessoa{Nome: "João", Idade: 30}  // Campos nomeados (RECOMENDADO)
 * p := Pessoa{"João", 30, "email"}      // Ordem (menos seguro)
 * p := &Pessoa{Nome: "João"}     // Ponteiro
 *
 * ACESSO:
 * p.Nome = "Maria"               // Modificar campo
 * fmt.Println(p.Idade)          // Ler campo
 *
 * MÉTODOS:
 * func (p Pessoa) Apresentar() string { }      // Receiver por valor
 * func (p *Pessoa) FazerAniversario() { }      // Receiver por ponteiro
 *
 * COMPARAÇÃO:
 * - Structs são comparáveis se todos os campos forem comparáveis
 * - Não pode comparar se tiver slice, map, function
 *
 * QUANDO USAR:
 * - Agrupar dados relacionados
 * - Modelar entidades
 * - Passar múltiplos valores juntos
 * - Criar tipos customizados com comportamento
 */

/*
 * ============================================================================
 * MÓDULO 7: CONDITIONALS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: CONDITIONALS
 * ---------------------
 * IF BÁSICO:
 * if condicao {
 *     // código
 * }
 *
 * IF-ELSE:
 * if condicao {
 *     // código se verdadeiro
 * } else {
 *     // código se falso
 * }
 *
 * IF-ELSE IF:
 * if condicao1 {
 * } else if condicao2 {
 * } else {
 * }
 *
 * IF COM INICIALIZAÇÃO:
 * if valor, err := funcao(); err != nil {
 *     // err disponível apenas no escopo do if
 * }
 *
 * OPERADORES LÓGICOS:
 * - && (E): Ambas condições verdadeiras
 * - || (OU): Pelo menos uma verdadeira
 * - ! (NÃO): Inverte valor booleano
 *
 * SWITCH:
 * switch variavel {
 * case valor1:
 *     // código
 * case valor2, valor3:  // Múltiplos valores
 *     // código
 * default:
 *     // código
 * }
 *
 * SWITCH SEM EXPRESSÃO (como if-else):
 * switch {
 * case idade < 18:
 * case idade >= 65:
 * }
 *
 * CARACTERÍSTICAS:
 * - Sem parênteses na condição
 * - Chaves obrigatórias
 * - Switch não precisa de break (sem fall-through)
 * - fallthrough força próximo case
 */

/*
 * ============================================================================
 * MÓDULO 8: LOOPS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: LOOPS
 * --------------
 * FOR CLÁSSICO:
 * for inicializacao; condicao; posInstrucao {
 *     // código
 * }
 * Exemplo: for i := 0; i < 10; i++ { }
 *
 * FOR WHILE-STYLE:
 * for condicao {
 *     // código
 * }
 * Exemplo: for contador < 5 { contador++ }
 *
 * FOR INFINITO:
 * for {
 *     // código (precisa de break ou return)
 * }
 *
 * FOR RANGE (PREFERIDO para coleções):
 * for indice, valor := range slice {
 *     // código
 * }
 * for chave, valor := range mapa {
 *     // código
 * }
 * for indice, rune := range string {
 *     // código (retorna RUNES, não bytes!)
 * }
 *
 * BREAK E CONTINUE:
 * - break: Sai do loop
 * - continue: Pula para próxima iteração
 * - Labels: Controlar loops aninhados
 *
 * ⚠️ IMPORTANTE:
 * - for range sobre strings retorna RUNES (caracteres Unicode)
 * - Modificar tamanho de slice durante range pode causar problemas
 * - Maps: ordem de iteração é aleatória
 * - Maps: pode deletar durante iteração, mas não adicionar/modificar
 */

/*
 * ============================================================================
 * MÓDULO 9: FUNCTIONS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: FUNCTIONS
 * ------------------
 * SINTAXE BÁSICA:
 * func nomeFuncao(parametros) tipoRetorno {
 *     return valor
 * }
 *
 * MÚLTIPLOS RETORNOS:
 * func dividir(a, b float64) (float64, error) {
 *     if b == 0 {
 *         return 0, fmt.Errorf("divisão por zero")
 *     }
 *     return a / b, nil
 * }
 *
 * RETORNOS NOMEADOS:
 * func calcular() (soma int, produto int) {
 *     soma = 10
 *     produto = 20
 *     return  // Retorna soma e produto implicitamente
 * }
 *
 * VARIADIC FUNCTIONS:
 * func somar(numeros ...int) int {
 *     // numeros é um []int
 * }
 * somar(1, 2, 3)           // Chamada
 * somar(slice...)           // Desempacotar slice
 *
 * FIRST-CLASS CITIZENS:
 * - Funções podem ser atribuídas a variáveis
 * - Passadas como argumentos
 * - Retornadas de outras funções
 *
 * ANONYMOUS FUNCTIONS:
 * soma := func(a, b int) int {
 *     return a + b
 * }
 *
 * CLOSURES:
 * - Funções que capturam variáveis do escopo externo
 * - Mantêm estado entre chamadas
 * func criarContador() func() int {
 *     contador := 0
 *     return func() int {
 *         contador++
 *         return contador
 *     }
 * }
 *
 * CALL BY VALUE:
 * - Tudo é passado por valor (cópia)
 * - Slices/maps: copia a referência (compartilha dados)
 * - Use ponteiros para modificar valores originais
 */

/*
 * ============================================================================
 * MÓDULO 10: POINTERS E MEMORY (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: POINTERS E MEMORY MANAGEMENT
 * --------------------------------------
 * POINTERS:
 * - Armazenam endereços de memória
 * - Permitem modificar valores originais
 * - Evitam cópias grandes
 *
 * OPERADORES:
 * - &variavel: Obtém endereço (address-of)
 * - *ponteiro: Acessa valor apontado (dereference)
 * - *Tipo: Declara pointer para Tipo
 *
 * EXEMPLO:
 * var x int = 42
 * var p *int = &x
 * fmt.Println(*p)  // 42
 * *p = 99          // Modifica x
 *
 * ZERO VALUE: nil
 *
 * STRUCTS COM POINTERS:
 * - Go permite acesso direto: p.Nome (sem precisar (*p).Nome)
 * - Use pointer receiver quando método modifica struct
 *
 * SLICES E MAPS:
 * - Já são reference types (não precisa de * na maioria dos casos)
 * - Modificar elementos afeta original
 * - Reatribuir não afeta original (precisa de *slice)
 *
 * MEMORY MANAGEMENT:
 * - Stack: Variáveis locais, rápido, limpeza automática
 * - Heap: Dados que vivem além da função, GC gerencia
 * - Escape Analysis: Compilador decide onde alocar
 * - Garbage Collector: Automático, tri-color mark-and-sweep
 * - Você NÃO precisa gerenciar memória manualmente!
 */

/*
 * ============================================================================
 * MÓDULO 11: METHODS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: METHODS
 * ----------------
 * CONCEITO:
 * - Funções que pertencem a um tipo específico
 * - Definidas com receiver (receptor)
 * - Adicionam comportamento aos tipos
 *
 * VALUE RECEIVER:
 * func (p Pessoa) Apresentar() string {
 *     return p.Nome
 * }
 * - Recebe cópia do valor
 * - Não modifica original
 * - Use para tipos pequenos ou métodos que apenas leem
 *
 * POINTER RECEIVER:
 * func (p *Pessoa) FazerAniversario() {
 *     p.Idade++
 * }
 * - Recebe ponteiro
 * - Pode modificar original
 * - Use quando método modifica ou tipo é grande
 *
 * CONVENÇÃO:
 * - Se um método usa pointer receiver, todos devem usar
 * - Mantenha consistência
 *
 * GO É INTELIGENTE:
 * - Pode chamar value receiver em pointer: (&p).Metodo()
 * - Pode chamar pointer receiver em value: (*p).Metodo()
 * - Go faz conversão automaticamente
 */

/*
 * ============================================================================
 * MÓDULO 12: INTERFACES (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: INTERFACES
 * -------------------
 * CONCEITO:
 * - Contratos que especificam métodos
 * - Tipos satisfazem interfaces IMPLICITAMENTE
 * - Não precisa declarar "implements"
 *
 * SINTAXE:
 * type Forma interface {
 *     Area() float64
 *     Perimetro() float64
 * }
 *
 * SATISFAÇÃO IMPLÍCITA:
 * - Se tipo tem os métodos, satisfaz automaticamente
 * - Permite polimorfismo
 * - Desacopla código
 *
 * EMPTY INTERFACE (interface{}):
 * - Aceita qualquer tipo
 * - Útil antes de generics (Go 1.18+)
 * - Hoje: prefira generics
 *
 * TYPE ASSERTION:
 * valor, ok := variavel.(Tipo)
 * - Verifica e converte tipo
 * - ok indica sucesso
 *
 * TYPE SWITCH:
 * switch v := variavel.(type) {
 * case int:
 * case string:
 * default:
 * }
 *
 * VANTAGENS:
 * - Polimorfismo
 * - Testabilidade (fácil criar mocks)
 * - Flexibilidade
 * - Desacoplamento
 */

/*
 * ============================================================================
 * MÓDULO 13: GENERICS (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: GENERICS
 * -----------------
 * CONCEITO:
 * - Código reutilizável com type safety
 * - Introduzido no Go 1.18
 * - Resolve problema de duplicação de código
 *
 * SINTAXE:
 * func Max[T comparable](a, b T) T {
 *     if a > b {
 *         return a
 *     }
 *     return b
 * }
 *
 * TYPE PARAMETERS:
 * - Definidos entre [] antes dos parâmetros
 * - Podem ter múltiplos: [T, U any]
 *
 * CONSTRAINTS:
 * - any: Qualquer tipo
 * - comparable: Tipos que podem ser comparados (==, !=)
 * - constraints.Ordered: Tipos ordenados (<, >, etc.)
 * - Interfaces: Constraints customizadas
 *
 * TYPE INFERENCE:
 * - Go infere tipos automaticamente
 * - Max(10, 20) → T é int
 * - Max[int](10, 20) → Explícito
 *
 * TIPOS GENÉRICOS:
 * type Stack[T any] struct {
 *     items []T
 * }
 *
 * VANTAGENS:
 * - Type safety
 * - Performance (sem overhead)
 * - Reutilização
 * - Legibilidade
 */

/*
 * ============================================================================
 * MÓDULO 14: ERROR HANDLING (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: TRATAMENTO DE ERROS
 * -----------------------------
 * FILOSOFIA GO:
 * - Erros são valores
 * - Sem exceções (try/catch)
 * - Verificação explícita de erros
 *
 * PADRÃO IDIOMÁTICO:
 * func funcao() (Tipo, error) {
 *     if erro {
 *         return zeroValue, fmt.Errorf("mensagem")
 *     }
 *     return valor, nil
 * }
 *
 * VERIFICAÇÃO:
 * resultado, err := funcao()
 * if err != nil {
 *     // tratar erro
 *     return err
 * }
 * // usar resultado
 *
 * CRIAR ERROS:
 * - fmt.Errorf("formato %s", args)
 * - errors.New("mensagem")
 * - errors.Is(err, target) - Verificar tipo
 * - errors.As(err, &target) - Extrair tipo específico
 *
 * WRAPPING:
 * return fmt.Errorf("contexto: %w", err)
 * - Preserva erro original
 * - Permite unwrap
 *
 * BOAS PRÁTICAS:
 * - Sempre verifique erros
 * - Retorne erros descritivos
 * - Use wrapping para contexto
 * - Não ignore erros com _
 */

/*
 * ============================================================================
 * MÓDULO 15: CODE ORGANIZATION (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: ORGANIZAÇÃO DE CÓDIGO
 * -------------------------------
 * PACOTES:
 * - Organizam código relacionado
 * - package nome (no topo do arquivo)
 * - Nome do pacote = último diretório
 *
 * VISIBILIDADE:
 * - Maiúscula: Exportado (público)
 * - Minúscula: Não exportado (privado)
 * - Exemplo: FuncaoPublica vs funcaoPrivada
 *
 * IMPORTS:
 * - import "pacote"
 * - import ( "pacote1" "pacote2" )
 * - import alias "pacote" - Renomear
 * - import _ "pacote" - Importar apenas init()
 *
 * INIT():
 * - Executado automaticamente
 * - Pode ter múltiplos init() por arquivo
 * - Ordem: imports → vars → init() → main()
 *
 * ESTRUTURA DE PROJETO:
 * - cmd/: Aplicações executáveis
 * - pkg/: Bibliotecas reutilizáveis
 * - internal/: Código privado do módulo
 * - api/: Definições de API
 * - web/: Assets web
 */

/*
 * ============================================================================
 * MÓDULO 16: CONCURRENCY (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: CONCORRÊNCIA
 * ---------------------
 * CONCEITOS:
 * - Concorrência: Estruturar para múltiplas tarefas
 * - Paralelismo: Executar simultaneamente
 * - Go suporta ambos
 *
 * GOROUTINES:
 * - Threads leves (muito baratas)
 * - Criar: go funcao()
 * - Gerenciadas pelo runtime
 * - Múltiplas goroutines em poucas threads OS
 *
 * CHANNELS:
 * - Comunicação entre goroutines
 * - make(chan Tipo) - Unbuffered (síncrono)
 * - make(chan Tipo, n) - Buffered (assíncrono)
 * - ch <- valor - Enviar
 * - valor := <-ch - Receber
 * - close(ch) - Fechar channel
 *
 * SELECT:
 * - Multiplexa channels
 * - Similar a switch
 * - default: não bloqueia
 *
 * SYNC:
 * - sync.Mutex: Exclusão mútua
 * - sync.RWMutex: Leitura concorrente
 * - sync.WaitGroup: Esperar goroutines
 * - sync.Once: Executar uma vez
 * - sync.Map: Map thread-safe
 *
 * PADRÕES:
 * - Worker pools
 * - Fan-out/Fan-in
 * - Pipeline
 * - Context para cancelamento
 *
 * FILOSOFIA:
 * "Don't communicate by sharing memory;
 *  share memory by communicating"
 */

/*
 * ============================================================================
 * MÓDULO 17: CONTEXT (Aula 01)
 * ============================================================================
 */

/*
 * AULA 01: CONTEXT
 * ----------------
 * PROPÓSITO:
 * - Cancelamento de operações
 * - Timeouts
 * - Valores de requisição
 * - Propagação de escopo
 *
 * CONTEXT PADRÃO:
 * ctx := context.Background()  // Context raiz
 * ctx := context.TODO()         // Placeholder
 *
 * DERIVAÇÃO:
 * ctx, cancel := context.WithCancel(parent)
 * ctx, cancel := context.WithTimeout(parent, duration)
 * ctx, cancel := context.WithDeadline(parent, time)
 * ctx := context.WithValue(parent, key, value)
 *
 * VERIFICAÇÃO:
 * select {
 * case <-ctx.Done():
 *     return ctx.Err()
 * default:
 *     // continuar
 * }
 *
 * BOAS PRÁTICAS:
 * - Primeiro parâmetro em funções
 * - Não armazene context em structs
 * - Passe explicitamente
 * - Use para cancelamento/timeout
 */

/*
 * ============================================================================
 * MÓDULO 18-43: TÓPICOS AVANÇADOS
 * ============================================================================
 */

/*
 * MÓDULO 18: CONCURRENCY PATTERNS
 * --------------------------------
 * - Worker pools
 * - Fan-out/Fan-in
 * - Pipeline
 * - Rate limiting
 * - Semáforos
 */

/*
 * MÓDULO 19: RACE DETECTION
 * --------------------------
 * - go test -race
 * - Detecção de race conditions
 * - Ferramenta de debugging
 */

/*
 * MÓDULO 20: STANDARD LIBRARY
 * ----------------------------
 * - fmt: Formatação
 * - strings: Manipulação de strings
 * - strconv: Conversão string/número
 * - time: Datas e horas
 * - os: Sistema operacional
 * - io: I/O operations
 * - net/http: HTTP client/server
 * - encoding/json: JSON
 * - encoding/xml: XML
 * - crypto: Criptografia
 */

/*
 * MÓDULO 21: TESTING E BENCHMARKING
 * ----------------------------------
 * - go test
 * - Testes unitários
 * - Testes de tabela
 * - Benchmarks
 * - Exemplos como testes
 * - Coverage
 */

/*
 * MÓDULO 22: BUILDING CLIs
 * -------------------------
 * - Flags (flag package)
 * - Argumentos
 * - Cobra (biblioteca popular)
 * - Output formatting
 */

/*
 * MÓDULO 23: WEB DEVELOPMENT
 * ---------------------------
 * - net/http server
 * - Handlers
 * - Routing
 * - Middleware
 * - Templates
 */

/*
 * MÓDULO 24: GIN FRAMEWORK
 * -------------------------
 * - Framework web popular
 * - Routing
 * - Middleware
 * - JSON binding
 * - Validação
 */

/*
 * MÓDULO 25: gRPC E PROTOBUF
 * ---------------------------
 * - Protocol Buffers
 * - gRPC
 * - Service definition
 * - Code generation
 */

/*
 * MÓDULO 26: ORMs E DB ACCESS
 * ----------------------------
 * - database/sql
 * - GORM
 * - Migrations
 * - Connection pooling
 */

/*
 * MÓDULO 27: LOGGING
 * -------------------
 * - log package
 * - logrus
 * - zap
 * - Structured logging
 */

/*
 * MÓDULO 28: REALTIME COMMUNICATION
 * ----------------------------------
 * - WebSockets
 * - Server-Sent Events
 * - gorilla/websocket
 */

/*
 * MÓDULO 29: CORE GO COMMANDS
 * ----------------------------
 * - go build
 * - go run
 * - go test
 * - go mod
 * - go get
 * - go install
 * - go fmt
 * - go vet
 * - go doc
 */

/*
 * MÓDULO 30: CODE QUALITY ANALYSIS
 * ---------------------------------
 * - go vet
 * - golangci-lint
 * - staticcheck
 * - gofmt
 * - goimports
 */

/*
 * MÓDULO 31: LINTERS
 * -------------------
 * - golangci-lint
 * - Configuração
 * - Regras customizadas
 * - Integração CI/CD
 */

/*
 * MÓDULO 32: SECURITY
 * --------------------
 * - Input validation
 * - SQL injection prevention
 * - XSS prevention
 * - Secrets management
 * - crypto package
 */

/*
 * MÓDULO 33: CODE GENERATION E BUILD TAGS
 * ----------------------------------------
 * - go generate
 * - Build tags
 * - Conditional compilation
 */

/*
 * MÓDULO 34: PERFORMANCE E DEBUGGING
 * ------------------------------------
 * - pprof
 * - Benchmarking
 * - Memory profiling
 * - CPU profiling
 * - Trace
 */

/*
 * MÓDULO 35: DEPLOYMENT TOOLING
 * ------------------------------
 * - Docker
 * - Kubernetes
 * - CI/CD
 * - Build optimization
 */

/*
 * MÓDULO 36: MEMORY MANAGEMENT
 * -----------------------------
 * - Stack vs Heap
 * - Escape analysis
 * - GC tuning
 * - Memory profiling
 */

/*
 * MÓDULO 37: ESCAPE ANALYSIS
 * ---------------------------
 * - Como funciona
 * - go build -gcflags="-m"
 * - Otimizações
 */

/*
 * MÓDULO 38: REFLECTION
 * ----------------------
 * - reflect package
 * - TypeOf, ValueOf
 * - Type assertions
 * - Use cases
 */

/*
 * MÓDULO 39: UNSAFE PACKAGE
 * --------------------------
 * - unsafe.Pointer
 * - Conversões de baixo nível
 * - Quando usar (raramente)
 */

/*
 * MÓDULO 40: BUILD CONSTRAINTS/TAGS
 * ----------------------------------
 * - // +build tags
 * - //go:build
 * - Conditional compilation
 */

/*
 * MÓDULO 41: COMPILER E LINKER FLAGS
 * -----------------------------------
 * - -gcflags
 * - -ldflags
 * - Otimizações
 */

/*
 * MÓDULO 42: CGO BASICS
 * ----------------------
 * - Integração com C
 * - cgo
 * - Quando usar
 */

/*
 * MÓDULO 43: PLUGINS E DYNAMIC LOADING
 * -------------------------------------
 * - plugin package
 * - Dynamic loading
 * - Use cases
 */

/*
 * ============================================================================
 * RESUMO GERAL - CONCEITOS ESSENCIAIS
 * ============================================================================
 *
 * 1. SIMPLICIDADE:
 *    - Go prioriza simplicidade sobre "features"
 *    - Menos é mais
 *    - Código claro e legível
 *
 * 2. CONCORRÊNCIA:
 *    - Goroutines e channels são first-class
 *    - "Share memory by communicating"
 *    - Fácil escrever código concorrente
 *
 * 3. TIPAGEM:
 *    - Estaticamente tipada
 *    - Type safety em tempo de compilação
 *    - Generics desde Go 1.18
 *
 * 4. PERFORMANCE:
 *    - Compilado para código nativo
 *    - GC eficiente
 *    - Baixo overhead
 *
 * 5. FERRAMENTAS:
 *    - go tool é completo
 *    - gofmt garante estilo consistente
 *    - go test integrado
 *
 * 6. COMUNIDADE:
 *    - Código aberto
 *    - Documentação excelente
 *    - Ecossistema maduro
 *
 * ============================================================================
 * DICAS FINAIS
 * ============================================================================
 *
 * - Sempre verifique erros
 * - Use go fmt antes de commitar
 * - Escreva testes
 * - Documente código público
 * - Siga convenções Go (Effective Go)
 * - Prefira composição sobre herança
 * - Use interfaces para desacoplar
 * - Aproveite a concorrência do Go
 * - Mantenha funções pequenas e focadas
 * - Nomes descritivos são importantes
 *
 * ============================================================================
 * FIM DO REVIEW
 * ============================================================================
 */

