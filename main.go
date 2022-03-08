//Eduarda Aguiar Angelo
package main

import (
	"fmt"
  "math"
)

var option int
var num, numAux float64

func main() {
  
  menu()
  operacoes()
  
}

//menu de opções
func menu(){
  fmt.Println(" ")
  fmt.Println("============== Calculadora simples Eduarda Aguiar ==============\n") 
  
  fmt.Println("------------------------------------")
  fmt.Println("Digite qual operação deseja fazer:")
  fmt.Println("------------------------------------\n")
  
  fmt.Println("1) Soma")
  fmt.Println("2) Subtração")
  fmt.Println("3) Divisão")
  fmt.Println("4) Multiplicação")
  fmt.Println("5) Elevar número ao quadrado")
  fmt.Println("6) Elevar número ao cubo")
  fmt.Println("7) Raiz Quadrada")
  fmt.Println("8) Raiz Cubica")
  fmt.Println("9) Logaritmo\n")

}

 //opções e operações
func operacoes(){
  fmt.Scanln(&option)
  fmt.Println("")
  switch option {
    
  case 1:
   fmt.Println("")
   fmt.Println("Digite um número:")
   fmt.Scanln(&num)
   fmt.Println("Digite outro número:")
   fmt.Scanln(&numAux)
  var soma float64 = num + numAux 
   fmt.Println("")
   fmt.Println("O resultado da soma é", soma)
    menu()
    operacoes()
    
  case 2:      
   fmt.Println("Digite um número:")
   fmt.Scanln(&num)
   fmt.Println("Digite outro número:")
   fmt.Scanln(&numAux)
  var subtracao float64 = num - numAux
   fmt.Println("O resultado da subtracao é",subtracao)
    menu()
    operacoes()
    
  case 3:
   fmt.Println("Digite um número:")
   fmt.Scanln(&num)
   fmt.Println("Digite outro número:")
   fmt.Scanln(&numAux)    
  var divisao float64 = num / numAux
   fmt.Println("O resultado da divisao é",divisao)
    menu()
    operacoes()
    
  case 4:
   fmt.Println("Digite um número:")
   fmt.Scanln(&num)
   fmt.Println("Digite outro número:")
   fmt.Scanln(&numAux)    
  var multiplicacao float64 = num * numAux
   fmt.Println("O resultado da multiplicacao é",multiplicacao)
    menu()
    operacoes()
    
  case 5:
    fmt.Println("Digite um número:")
    fmt.Scanln(&num)
  var Elevar_número_ao_quadrado float64 = math.Pow(num,2)
   fmt.Println("O resultado de" ,num, "elevado ao quadrado é ",Elevar_número_ao_quadrado)
    menu()
    operacoes()
    
  case 6:    
    fmt.Println("Digite um número:")
    fmt.Scanln(&num)
  var Elevar_número_ao_cubo float64 = math.Pow(num,3)
   fmt.Println("O resultado de" ,num, "elevado ao cubo é",Elevar_número_ao_cubo)  
    menu()
    operacoes()
    
  case 7:
    fmt.Println("Digite um número:")
    fmt.Scanln(&num)
  var Raiz_Quadrada float64 = math.Sqrt(num)
   fmt.Println("O resultado da raiz quadrada de" ,num, "é",Raiz_Quadrada) 
    menu()
    operacoes()
    
  case 8:
    fmt.Println("Digite um número:")
    fmt.Scanln(&num)
  var Raiz_Cubica float64 = math.Cbrt(num)
   fmt.Println("O resultado da raiz cubica de" ,num, "é",Raiz_Cubica) 
    menu()
    operacoes()
    
    
  case 9:
   fmt.Println("Digite um número:")
   fmt.Scanln(&num)
  var Logaritmo float64 = math.Log(num)
   fmt.Println("O resultado do logaritmo de" ,num, "é",Logaritmo)
    menu()
    operacoes()

  default:
  fmt.Println("")
  fmt.Println("XXXXXXXXXXXX OPÇÃO INVÁLIDA! XXXXXXXXXXXX")
    menu()
    operacoes()
    
    
  }        
   }

 

