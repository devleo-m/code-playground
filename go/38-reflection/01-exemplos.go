package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Exemplo 1: Inspecionar tipo básico
func exemplo1() {
	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	
	fmt.Printf("Tipo: %s\n", t)
	fmt.Printf("Kind: %s\n", t.Kind())
	fmt.Printf("Valor: %v\n", v.Int())
}

// Exemplo 2: Inspecionar struct
type Person struct {
	Name string
	Age  int
	City string
}

func exemplo2() {
	p := Person{Name: "John", Age: 30, City: "NYC"}
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	
	fmt.Printf("Tipo: %s\n", t)
	fmt.Printf("NumFields: %d\n", t.NumField())
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("  %s (%s) = %v\n", field.Name, field.Type, value.Interface())
	}
}

// Exemplo 3: Ler tags
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"min=18"`
}

func exemplo3() {
	u := User{}
	t := reflect.TypeOf(u)
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		fmt.Printf("%s: json=%s, validate=%s\n", field.Name, jsonTag, validateTag)
	}
}

// Exemplo 4: Modificar valor
func exemplo4() {
	x := 42
	fmt.Printf("Antes: %d\n", x)
	
	v := reflect.ValueOf(&x).Elem()
	v.SetInt(100)
	
	fmt.Printf("Depois: %d\n", x)
}

// Exemplo 5: Modificar struct
func exemplo5() {
	p := &Person{Name: "John", Age: 30, City: "NYC"}
	fmt.Printf("Antes: %+v\n", p)
	
	v := reflect.ValueOf(p).Elem()
	
	nameField := v.FieldByName("Name")
	nameField.SetString("Jane")
	
	ageField := v.FieldByName("Age")
	ageField.SetInt(25)
	
	fmt.Printf("Depois: %+v\n", p)
}

// Exemplo 6: Chamar método dinamicamente
type Calculator struct {
	result int
}

func (c *Calculator) Add(x int) {
	c.result += x
}

func (c *Calculator) GetResult() int {
	return c.result
}

func exemplo6() {
	calc := &Calculator{result: 10}
	v := reflect.ValueOf(calc)
	
	// Chamar Add
	addMethod := v.MethodByName("Add")
	addMethod.Call([]reflect.Value{reflect.ValueOf(5)})
	
	// Chamar GetResult
	getResultMethod := v.MethodByName("GetResult")
	results := getResultMethod.Call(nil)
	
	fmt.Printf("Resultado: %d\n", results[0].Int())
}

// Exemplo 7: Criar novo valor
func exemplo7() {
	// Criar novo int
	intType := reflect.TypeOf(0)
	intValue := reflect.New(intType).Elem()
	intValue.SetInt(42)
	fmt.Printf("Novo int: %d\n", intValue.Int())
	
	// Criar novo string
	stringType := reflect.TypeOf("")
	stringValue := reflect.New(stringType).Elem()
	stringValue.SetString("hello")
	fmt.Printf("Novo string: %s\n", stringValue.String())
}

// Exemplo 8: Validador simples
func exemplo8() {
	user := User{
		Name:  "", // Erro: required
		Email: "invalid", // Erro: email
		Age:   15, // Erro: min=18
	}
	
	errors := validate(user)
	for _, err := range errors {
		fmt.Println(err)
	}
}

func validate(x interface{}) []string {
	var errors []string
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return errors
	}
	
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		validateTag := field.Tag.Get("validate")
		
		if validateTag == "" {
			continue
		}
		
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)
			
			if rule == "required" {
				if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
					errors = append(errors, fmt.Sprintf("%s is required", field.Name))
				}
			} else if strings.HasPrefix(rule, "min=") {
				minStr := strings.TrimPrefix(rule, "min=")
				min, _ := strconv.Atoi(minStr)
				if fieldValue.Kind() == reflect.Int && int(fieldValue.Int()) < min {
					errors = append(errors, fmt.Sprintf("%s must be at least %d", field.Name, min))
				}
			} else if rule == "email" {
				email := fieldValue.String()
				if !strings.Contains(email, "@") {
					errors = append(errors, fmt.Sprintf("%s must be a valid email", field.Name))
				}
			}
		}
	}
	
	return errors
}

// Exemplo 9: JSON simples
func exemplo9() {
	p := Person{Name: "John", Age: 30, City: "NYC"}
	json := toJSON(p)
	fmt.Println(json)
}

func toJSON(x interface{}) string {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return "{}"
	}
	
	var parts []string
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fmt.Sprintf(`"%s"`, fieldValue.String())
		case reflect.Int:
			valueStr = strconv.Itoa(int(fieldValue.Int()))
		default:
			valueStr = fmt.Sprintf(`"%v"`, fieldValue.Interface())
		}
		
		parts = append(parts, fmt.Sprintf(`"%s":%s`, jsonTag, valueStr))
	}
	
	return "{" + strings.Join(parts, ",") + "}"
}

// Exemplo 10: Comparação genérica
func exemplo10() {
	// Comparar slices
	fmt.Println("Slices iguais:", reflect.DeepEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	
	// Comparar maps
	fmt.Println("Maps iguais:", reflect.DeepEqual(
		map[string]int{"a": 1},
		map[string]int{"a": 1}))
	
	// Comparar structs
	fmt.Println("Structs iguais:", reflect.DeepEqual(
		Person{Name: "John", Age: 30},
		Person{Name: "John", Age: 30}))
}

// Menu interativo
func showMenu() {
	fmt.Println("\n=== Exemplos de Reflection ===")
	fmt.Println("1. Inspecionar tipo básico")
	fmt.Println("2. Inspecionar struct")
	fmt.Println("3. Ler tags")
	fmt.Println("4. Modificar valor")
	fmt.Println("5. Modificar struct")
	fmt.Println("6. Chamar método dinamicamente")
	fmt.Println("7. Criar novo valor")
	fmt.Println("8. Validador simples")
	fmt.Println("9. JSON simples")
	fmt.Println("10. Comparação genérica")
	fmt.Println("0. Sair")
	fmt.Print("\nEscolha uma opção: ")
}

func main() {
	if len(os.Args) > 1 {
		// Modo não-interativo
		switch os.Args[1] {
		case "1":
			exemplo1()
		case "2":
			exemplo2()
		case "3":
			exemplo3()
		case "4":
			exemplo4()
		case "5":
			exemplo5()
		case "6":
			exemplo6()
		case "7":
			exemplo7()
		case "8":
			exemplo8()
		case "9":
			exemplo9()
		case "10":
			exemplo10()
		default:
			fmt.Println("Opção inválida")
		}
		return
	}

	// Modo interativo
	for {
		showMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			exemplo1()
		case 2:
			exemplo2()
		case 3:
			exemplo3()
		case 4:
			exemplo4()
		case 5:
			exemplo5()
		case 6:
			exemplo6()
		case 7:
			exemplo7()
		case 8:
			exemplo8()
		case 9:
			exemplo9()
		case 10:
			exemplo10()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}


