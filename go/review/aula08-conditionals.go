package main

import "fmt"

/*
 * ============================================================================
 * AULA 08: CONDITIONALS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - if: Executa código se condição for verdadeira
 * - if-else: Escolhe entre dois caminhos
 * - if-else if: Múltiplas condições
 * - switch: Compara variável contra múltiplos valores
 * - switch sem expressão: Como if-else if
 * - Operadores lógicos: &&, ||, !
 */

func main() {
	fmt.Println("=== AULA 08: CONDITIONALS ===\n")
	
	// ===== IF BÁSICO =====
	fmt.Println("1. IF BÁSICO:")
	idade := 20
	if idade >= 18 {
		fmt.Println("   Você é maior de idade")
	}
	
	// ===== IF-ELSE =====
	fmt.Println("\n2. IF-ELSE:")
	temperatura := 25
	if temperatura > 30 {
		fmt.Println("   Está muito quente!")
	} else {
		fmt.Println("   Temperatura agradável")
	}
	
	// ===== IF-ELSE IF =====
	fmt.Println("\n3. IF-ELSE IF:")
	nota := 85
	if nota >= 90 {
		fmt.Println("   Nota A - Excelente!")
	} else if nota >= 80 {
		fmt.Println("   Nota B - Muito bom!")
	} else if nota >= 70 {
		fmt.Println("   Nota C - Bom")
	} else if nota >= 60 {
		fmt.Println("   Nota D - Regular")
	} else {
		fmt.Println("   Nota F - Reprovado")
	}
	
	// ===== IF COM INICIALIZAÇÃO =====
	fmt.Println("\n4. IF COM INICIALIZAÇÃO:")
	if num := 42; num > 40 {
		fmt.Printf("   O número %d é maior que 40\n", num)
		// num está disponível apenas aqui
	}
	
	// Com função
	if resultado, err := dividir(10, 2); err != nil {
		fmt.Printf("   Erro: %v\n", err)
	} else {
		fmt.Printf("   Resultado: %.2f\n", resultado)
	}
	
	// ===== OPERADORES LÓGICOS =====
	fmt.Println("\n5. OPERADORES LÓGICOS:")
	salario := 3000
	horasTrabalhadas := 45
	
	// && (E)
	if salario > 2500 && horasTrabalhadas > 40 {
		fmt.Println("   Você recebeu horas extras!")
	}
	
	// || (OU)
	if salario < 2000 || horasTrabalhadas < 30 {
		fmt.Println("   Atenção: salário ou horas abaixo do esperado")
	}
	
	// ! (NÃO)
	chovendo := false
	if !chovendo {
		fmt.Println("   Pode sair sem guarda-chuva")
	}
	
	// ===== SWITCH BÁSICO =====
	fmt.Println("\n6. SWITCH BÁSICO:")
	dia := "segunda"
	switch dia {
	case "segunda":
		fmt.Println("   Início da semana")
	case "terça", "quarta", "quinta":
		fmt.Println("   Meio da semana")
	case "sexta":
		fmt.Println("   Final da semana")
	case "sábado", "domingo":
		fmt.Println("   Fim de semana!")
	default:
		fmt.Println("   Dia inválido")
	}
	
	// ===== SWITCH SEM EXPRESSÃO =====
	fmt.Println("\n7. SWITCH SEM EXPRESSÃO:")
	idade2 := 25
	switch {
	case idade2 < 18:
		fmt.Println("   Menor de idade")
	case idade2 >= 18 && idade2 < 65:
		fmt.Println("   Adulto")
	case idade2 >= 65:
		fmt.Println("   Idoso")
	}
	
	// ===== SWITCH COM INICIALIZAÇÃO =====
	fmt.Println("\n8. SWITCH COM INICIALIZAÇÃO:")
	switch hora := 14; {
	case hora < 12:
		fmt.Println("   Bom dia!")
	case hora < 18:
		fmt.Println("   Boa tarde!")
	default:
		fmt.Println("   Boa noite!")
	}
	
	// ===== TYPE SWITCH =====
	fmt.Println("\n9. TYPE SWITCH:")
	var valor interface{} = 42
	switch v := valor.(type) {
	case int:
		fmt.Printf("   É um inteiro: %d\n", v)
	case string:
		fmt.Printf("   É uma string: %s\n", v)
	case bool:
		fmt.Printf("   É um booleano: %v\n", v)
	default:
		fmt.Printf("   Tipo desconhecido: %T\n", v)
	}
	
	// ===== VALIDAÇÃO COM CONDITIONALS =====
	fmt.Println("\n10. VALIDAÇÃO:")
	email := "usuario@email.com"
	senha := "senha123"
	
	if email == "" {
		fmt.Println("   Email não pode ser vazio")
	} else if len(senha) < 6 {
		fmt.Println("   Senha muito curta")
	} else {
		fmt.Println("   Dados válidos!")
	}
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

