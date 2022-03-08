package main

import (
	"fmt"
  "math"
)

var escolha int
var num, numaux float64

func main() {

  menu() 
  escolher_funcao()
  
  fmt.Println("         ")
	fmt.Println("Digite o primeiro número:")
  fmt.Scanln(&num)
  fmt.Println("\n")

  fmt.Println("Digite o segundo número:")
  fmt.Scanln(&numaux)
  fmt.Println("\n")
  }

  


func menu() {
  fmt.Println("============== Calculadora simples Eduarda Aguiar ==============\n") 
  

  fmt.Println("Digite qual operação deseja fazer:\n")
  
  fmt.Println("1) Soma")
  fmt.Println("2) Subtração")
  fmt.Println("3) Divisão")
  fmt.Println("4) Multiplicação")
  fmt.Println("5) Elevar número ao quadrado")
  fmt.Println("6) Elevar número ao cubo")
  fmt.Println("7) Raiz Quadrada")
  fmt.Println("8) Raiz Cubica")
  fmt.Println("9) Logaritmo\n")

  fmt.Scanln(&escolha)

}
 
func escolher_funcao(int escolha, float64 num, float64 numaux){
  var resultado float64

  switch escolha {
  case 1:
   resultado = soma(num, numaux)
   fmt.Println("O resultado é", soma)
  case 2:      
   resultado = subtracao(num, numaux)
   fmt.Println("O resultado é",subtracao)
  case 3:
   resultado = divisao(num, numaux)
   fmt.Println("O resultado é",divisao)
  case 4:
   resultado = multiplicacao(num, numaux)
   fmt.Println("O resultado é",multiplicacao)
  case 5:
   resultado = Elevar_número_ao_quadrado(num)
   fmt.Println("O resultado é",Elevar_número_ao_quadrado)
  case 6:
   resultado = Elevar_número_ao_cubo(num)
   fmt.Println("O resultado é",Elevar_número_ao_cubo)  
  case 7:
   resultado = Raiz_Quadrada(num)
   fmt.Println("O resultado é",Raiz_Quadrada) 
  case 8:
   resultado = Raiz_Cubo(num)
   fmt.Println("O resultado é",Raiz_Cubo) 
  case 9:
   resultado = Logaritmo(num, numaux)
   fmt.Println("O resultado é",Logaritmo)
  default:
  fmt.Println("Número inválido!")     
  }
  
  