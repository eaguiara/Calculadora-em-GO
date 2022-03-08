package main

import (
	"fmt"
)

func inicio(){
  
  fmt.Println("Digite qual operação deseja fazer:")
  fmt.Println("1) Soma")
  fmt.Println("2) Subtração")
  fmt.Println("3) Divisão")
  fmt.Println("4) Multiplicação")
 
  
  
}

func main() {

var n1, n2, resultado int
  fmt.Println("----------Calculadora simples Eduarda Aguiar----------\n")

	fmt.Println("Digite o primeiro número:")
  fmt.Scanln(&n1)
  fmt.Println("\n")

  fmt.Println("Digite o segundo número:")
  fmt.Scanln(&n2)
  fmt.Println("\n")

  inicio()
  fmt.Scanln(&resultado)
  switch resultado {
  case 1:
   var soma int = n1 + n2 
   fmt.Println(soma)
  case 2:      
  var subtracao int = n1 - n2
   fmt.Println(subtracao)
  case 3:
  var divisao int = n1 / n2
   fmt.Println(divisao)
  case 4:
  var multiplicacao int = n1 * n2
   fmt.Println(multiplicacao)  
  default:
  fmt.Println("Número inválido!")
         
  }
}
