package main

import (
	"fmt"
	"strings"
)

// ============================================
// EXEMPLO 1: for Loop Clássico
// ============================================

func exemploForClassico() {
	fmt.Println("=== Exemplo 1: for Loop Clássico ===")

	// Forma completa: inicialização, condição, pós-instrução
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteração %d\n", i)
	}
}

// ============================================
// EXEMPLO 2: for Loop - Apenas Condição (while-style)
// ============================================

func exemploForWhile() {
	fmt.Println("\n=== Exemplo 2: for como while ===")

	contador := 0
	for contador < 5 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}
}

// ============================================
// EXEMPLO 3: for Loop Infinito
// ============================================

func exemploForInfinito() {
	fmt.Println("\n=== Exemplo 3: for Loop Infinito ===")

	contador := 0
	for {
		contador++
		if contador >= 3 {
			break // Sair do loop
		}
		fmt.Printf("Loop infinito - iteração %d\n", contador)
	}
}

// ============================================
// EXEMPLO 4: for range - Arrays e Slices
// ============================================

func exemploForRangeSlice() {
	fmt.Println("\n=== Exemplo 4: for range com Slices ===")

	numeros := []int{10, 20, 30, 40, 50}

	// Com índice e valor
	for indice, valor := range numeros {
		fmt.Printf("Índice %d: valor %d\n", indice, valor)
	}

	// Apenas valores (ignorar índice)
	for _, valor := range numeros {
		fmt.Printf("Valor: %d\n", valor)
	}

	// Apenas índices (ignorar valor)
	for indice := range numeros {
		fmt.Printf("Índice: %d\n", indice)
	}
}

// ============================================
// EXEMPLO 5: for range - Maps
// ============================================

func exemploForRangeMap() {
	fmt.Println("\n=== Exemplo 5: for range com Maps ===")

	cores := map[string]string{
		"vermelho": "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarelo":  "#FFFF00",
	}

	// Iterar sobre map (ordem aleatória!)
	for chave, valor := range cores {
		fmt.Printf("Cor: %s = %s\n", chave, valor)
	}

	// Apenas chaves
	for chave := range cores {
		fmt.Printf("Chave: %s\n", chave)
	}

	// Apenas valores (ignorar chave)
	for _, valor := range cores {
		fmt.Printf("Valor: %s\n", valor)
	}
}

// ============================================
// EXEMPLO 6: for range - Strings (Runes)
// ============================================

func exemploForRangeString() {
	fmt.Println("\n=== Exemplo 6: for range com Strings ===")

	texto := "Olá, 世界! 🚀"

	// for range retorna índice e rune (código Unicode)
	for indice, rune := range texto {
		fmt.Printf("Posição %d: %c (Unicode: %d)\n", indice, rune, rune)
	}

	// Comparação: indexação direta (bytes, não runes!)
	fmt.Println("\nIndexação direta (bytes):")
	for i := 0; i < len(texto); i++ {
		fmt.Printf("Posição %d: byte %d (%c)\n", i, texto[i], texto[i])
	}
}

// ============================================
// EXEMPLO 7: break - Sair do Loop
// ============================================

func exemploBreak() {
	fmt.Println("\n=== Exemplo 7: break ===")

	// Sair quando encontrar o número 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Encontrei 5! Saindo do loop...")
			break
		}
		fmt.Printf("Valor: %d\n", i)
	}
}

// ============================================
// EXEMPLO 8: continue - Pular Iteração
// ============================================

func exemploContinue() {
	fmt.Println("\n=== Exemplo 8: continue ===")

	// Pular números pares
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Pula para próxima iteração
		}
		fmt.Printf("Número ímpar: %d\n", i)
	}
}

// ============================================
// EXEMPLO 9: break com Labels (Loops Aninhados)
// ============================================

func exemploBreakLabel() {
	fmt.Println("\n=== Exemplo 9: break com Labels ===")

	// Sem label: break sai apenas do loop interno
	fmt.Println("Sem label:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break // Sai apenas do loop j
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// Com label: break sai do loop externo também
	fmt.Println("\nCom label:")
LoopExterno:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break LoopExterno // Sai do loop externo!
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

// ============================================
// EXEMPLO 10: continue com Labels
// ============================================

func exemploContinueLabel() {
	fmt.Println("\n=== Exemplo 10: continue com Labels ===")

LoopExterno:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue LoopExterno // Pula para próxima iteração do loop externo
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

// ============================================
// EXEMPLO 11: Loops Aninhados
// ============================================

func exemploLoopsAninhados() {
	fmt.Println("\n=== Exemplo 11: Loops Aninhados ===")

	// Tabela de multiplicação
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}

// ============================================
// EXEMPLO 12: Modificar Slice Durante Iteração
// ============================================

func exemploModificarSlice() {
	fmt.Println("\n=== Exemplo 12: Modificar Slice Durante Iteração ===")

	numeros := []int{1, 2, 3, 4, 5}

	// Modificar elementos durante iteração (SEGURO)
	for i := range numeros {
		numeros[i] *= 2 // Dobrar cada número
	}
	fmt.Println("Números dobrados:", numeros)

	// CUIDADO: Não adicionar/remover durante iteração range
	// Use índice tradicional se precisar modificar tamanho
}

// ============================================
// EXEMPLO 13: Iterar e Filtrar
// ============================================

func exemploFiltrar() {
	fmt.Println("\n=== Exemplo 13: Filtrar Elementos ===")

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares := []int{}

	for _, num := range numeros {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	fmt.Printf("Números originais: %v\n", numeros)
	fmt.Printf("Números pares: %v\n", pares)
}

// ============================================
// EXEMPLO 14: Converter String para Rune Slice
// ============================================

func exemploStringParaRunes() {
	fmt.Println("\n=== Exemplo 14: String para Rune Slice ===")

	texto := "Olá, 世界! 🚀"

	// Converter para slice de runes para acesso aleatório
	runes := []rune(texto)

	fmt.Println("Acesso por índice (rune):")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("Posição %d: %c\n", i, runes[i])
	}

	// Comparação: acesso direto em string (bytes!)
	fmt.Println("\nAcesso direto em string (bytes - INCORRETO para Unicode):")
	for i := 0; i < len(texto); i++ {
		fmt.Printf("Posição %d: byte %d\n", i, texto[i])
	}
}

// ============================================
// EXEMPLO 15: Buscar em Slice
// ============================================

func exemploBuscar() {
	fmt.Println("\n=== Exemplo 15: Buscar em Slice ===")

	nomes := []string{"João", "Maria", "Pedro", "Ana"}

	// Buscar "Pedro"
	encontrado := false
	indice := -1

	for i, nome := range nomes {
		if nome == "Pedro" {
			encontrado = true
			indice = i
			break // Encontrou, pode sair
		}
	}

	if encontrado {
		fmt.Printf("Encontrado 'Pedro' no índice %d\n", indice)
	} else {
		fmt.Println("'Pedro' não encontrado")
	}
}

// ============================================
// EXEMPLO 16: Contar Elementos
// ============================================

func exemploContar() {
	fmt.Println("\n=== Exemplo 16: Contar Elementos ===")

	numeros := []int{1, 2, 2, 3, 2, 4, 2, 5}
	alvo := 2
	contador := 0

	for _, num := range numeros {
		if num == alvo {
			contador++
		}
	}

	fmt.Printf("O número %d aparece %d vezes\n", alvo, contador)
}

// ============================================
// EXEMPLO 17: Soma de Elementos
// ============================================

func exemploSoma() {
	fmt.Println("\n=== Exemplo 17: Soma de Elementos ===")

	numeros := []int{10, 20, 30, 40, 50}
	soma := 0

	for _, num := range numeros {
		soma += num
	}

	fmt.Printf("Soma: %d\n", soma)
	fmt.Printf("Média: %.2f\n", float64(soma)/float64(len(numeros)))
}

// ============================================
// EXEMPLO 18: Iterar Map e Modificar
// ============================================

func exemploModificarMap() {
	fmt.Println("\n=== Exemplo 18: Modificar Map Durante Iteração ===")

	// CUIDADO: Não pode modificar map durante iteração range
	// Mas pode deletar!

	cores := map[string]int{
		"vermelho": 1,
		"verde":   2,
		"azul":    3,
	}

	fmt.Println("Map original:", cores)

	// Deletar durante iteração (SEGURO)
	for chave, valor := range cores {
		if valor == 2 {
			delete(cores, chave) // Deletar é seguro
		}
	}

	fmt.Println("Após deletar:", cores)

	// Para modificar valores, criar novo map ou usar índice
	novasCores := make(map[string]int)
	for chave, valor := range cores {
		novasCores[chave] = valor * 2
	}
	fmt.Println("Novo map com valores dobrados:", novasCores)
}

// ============================================
// EXEMPLO 19: Processar Linhas de Texto
// ============================================

func exemploProcessarLinhas() {
	fmt.Println("\n=== Exemplo 19: Processar Linhas de Texto ===")

	texto := "Linha 1\nLinha 2\nLinha 3\nLinha 4"
	linhas := strings.Split(texto, "\n")

	for indice, linha := range linhas {
		fmt.Printf("Linha %d: %s\n", indice+1, linha)
	}
}

// ============================================
// EXEMPLO 20: Loop com Múltiplas Variáveis
// ============================================

func exemploMultiplasVariaveis() {
	fmt.Println("\n=== Exemplo 20: Loop com Múltiplas Variáveis ===")

	// Inicializar múltiplas variáveis
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
}

// ============================================
// EXEMPLO 21: Loop Descendente
// ============================================

func exemploLoopDescendente() {
	fmt.Println("\n=== Exemplo 21: Loop Descendente ===")

	// Contar de 10 até 1
	for i := 10; i >= 1; i-- {
		fmt.Printf("Contagem regressiva: %d\n", i)
	}
}

// ============================================
// EXEMPLO 22: Loop com Step Personalizado
// ============================================

func exemploLoopStep() {
	fmt.Println("\n=== Exemplo 22: Loop com Step ===")

	// Incrementar de 2 em 2
	for i := 0; i < 10; i += 2 {
		fmt.Printf("Número par: %d\n", i)
	}

	// Decrementar de 3 em 3
	fmt.Println()
	for i := 15; i >= 0; i -= 3 {
		fmt.Printf("Valor: %d\n", i)
	}
}

// ============================================
// EXEMPLO 23: goto (DESENCORAJADO)
// ============================================

func exemploGoto() {
	fmt.Println("\n=== Exemplo 23: goto (DESENCORAJADO) ===")

	// goto existe em Go, mas NÃO é recomendado!
	// Use break, continue, ou funções ao invés

	contador := 0

Inicio:
	contador++
	if contador < 3 {
		fmt.Printf("goto - iteração %d\n", contador)
		goto Inicio
	}

	fmt.Println("NOTA: Use break/continue ao invés de goto!")
}

// ============================================
// EXEMPLO 24: Loop com Condição Complexa
// ============================================

func exemploCondicaoComplexa() {
	fmt.Println("\n=== Exemplo 24: Loop com Condição Complexa ===")

	contador := 0
	maxIteracoes := 5
	continuar := true

	for continuar && contador < maxIteracoes {
		contador++
		fmt.Printf("Iteração %d\n", contador)

		if contador == 3 {
			continuar = false // Mudar condição
		}
	}
}

// ============================================
// EXEMPLO 25: Iterar String com Indexação Correta
// ============================================

func exemploStringIndexacao() {
	fmt.Println("\n=== Exemplo 25: Indexação Correta de Strings ===")

	texto := "Café"

	// ERRADO: Indexação direta (bytes)
	fmt.Println("Indexação direta (ERRADO para Unicode):")
	for i := 0; i < len(texto); i++ {
		fmt.Printf("texto[%d] = %d (byte)\n", i, texto[i])
	}

	// CORRETO: for range (runes)
	fmt.Println("\nfor range (CORRETO para Unicode):")
	for i, r := range texto {
		fmt.Printf("Posição %d: %c (rune %d)\n", i, r, r)
	}

	// CORRETO: Converter para []rune primeiro
	fmt.Println("\n[]rune (CORRETO para acesso aleatório):")
	runes := []rune(texto)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("runes[%d] = %c\n", i, runes[i])
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	exemploForClassico()
	exemploForWhile()
	exemploForInfinito()
	exemploForRangeSlice()
	exemploForRangeMap()
	exemploForRangeString()
	exemploBreak()
	exemploContinue()
	exemploBreakLabel()
	exemploContinueLabel()
	exemploLoopsAninhados()
	exemploModificarSlice()
	exemploFiltrar()
	exemploStringParaRunes()
	exemploBuscar()
	exemploContar()
	exemploSoma()
	exemploModificarMap()
	exemploProcessarLinhas()
	exemploMultiplasVariaveis()
	exemploLoopDescendente()
	exemploLoopStep()
	exemploGoto()
	exemploCondicaoComplexa()
	exemploStringIndexacao()
}

