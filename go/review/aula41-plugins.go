package main

import "fmt"

/*
 * ============================================================================
 * AULA 41: PLUGINS
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - plugin package
 * - Dynamic loading
 * - Use cases
 */

func main() {
	fmt.Println("=== AULA 41: PLUGINS ===\n")
	fmt.Println("plugin package: Carregar código dinamicamente")
	fmt.Println("Build: go build -buildmode=plugin")
	fmt.Println("Load: plugin.Open()")
	fmt.Println("Use cases: Extensibilidade, hot reload")
}




