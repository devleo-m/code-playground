package main

/*
#include <stdio.h>
void hello() {
    printf("Hello from C!\n");
}
*/
import "C"
import "fmt"

/*
 * ============================================================================
 * AULA 40: CGO
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - Integração com C
 * - cgo
 * - Quando usar
 */

func main() {
	fmt.Println("=== AULA 40: CGO ===\n")
	fmt.Println("CGO permite chamar código C do Go")
	C.hello()
	fmt.Println("Use apenas quando necessário (overhead)")
}




