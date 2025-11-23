package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
 * ============================================================================
 * AULA 20: STANDARD LIBRARY
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - fmt: Formatação
 * - strings: Manipulação de strings
 * - time: Datas e horas
 * - os: Sistema operacional
 * - encoding/json: JSON
 */

func main() {
	fmt.Println("=== AULA 20: STANDARD LIBRARY ===\n")
	
	// ===== STRINGS =====
	fmt.Println("1. STRINGS:")
	texto := "  Olá, Mundo!  "
	fmt.Printf("   TrimSpace: '%s'\n", strings.TrimSpace(texto))
	fmt.Printf("   ToUpper: %s\n", strings.ToUpper(texto))
	fmt.Printf("   Contains: %t\n", strings.Contains(texto, "Mundo"))
	fmt.Printf("   Split: %v\n", strings.Split("a,b,c", ","))
	
	// ===== TIME =====
	fmt.Println("\n2. TIME:")
	agora := time.Now()
	fmt.Printf("   Agora: %s\n", agora.Format(time.RFC3339))
	fmt.Printf("   Unix: %d\n", agora.Unix())
	
	// ===== JSON =====
	fmt.Println("\n3. JSON:")
	type Pessoa struct {
		Nome string `json:"nome"`
		Idade int   `json:"idade"`
	}
	
	p := Pessoa{Nome: "João", Idade: 30}
	jsonData, _ := json.Marshal(p)
	fmt.Printf("   JSON: %s\n", jsonData)
	
	// ===== OS =====
	fmt.Println("\n4. OS:")
	hostname, _ := os.Hostname()
	fmt.Printf("   Hostname: %s\n", hostname)
}


