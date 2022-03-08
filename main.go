package main

import (
	"fmt"
  "math"
)

var resultado int
var num, numaux float64

func inicio(){
  
   

	fmt.Println("Digite o primeiro número:")
  fmt.Scanln(&num)
  fmt.Println("\n")

  fmt.Println("Digite o segundo número:")
  fmt.Scanln(&numaux)
  fmt.Println("\n")

 
  
 }

func main() {

  fmt.Println("============== Calculadora simples Eduarda Aguiar ==============\n") 
  
 fmt.Println("==================================")
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
  
 

  inicio()
  fmt.Scanln(&resultado)
  switch resultado {
  case 1:
   var soma float64 = num + numaux 
   fmt.Println("O resultado é", soma)
  case 2:      
  var subtracao float64 = num - numaux
   fmt.Println("O resultado é",subtracao)
  case 3:
  var divisao float64 = num / numaux
   fmt.Println("O resultado é",divisao)
  case 4:
  var multiplicacao float64 = num * numaux
   fmt.Println("O resultado é",multiplicacao)
  case 5:
  var Elevar_número_ao_quadrado float64 = math.Pow(num,2)
   fmt.Println("O resultado é",Elevar_número_ao_quadrado)
  case 6:
  var Elevar_número_ao_cubo float64 = math.Pow(num,3)
   fmt.Println("O resultado é",Elevar_número_ao_cubo)  
  case 7:
  var Raiz_Quadrada float64 = math.Sqrt(num)
   fmt.Println("O resultado é",Raiz_Quadrada) 
  case 8:
  var Raiz_Cubo float64 = math.Cbrt(num)
   fmt.Println("O resultado é",Raiz_Cubo) 
  case 9:
  var Logaritmo float64 = num * numaux
   fmt.Println("O resultado é",Logaritmo)
    
    
    
  default:
  fmt.Println("Número inválido!")
         
  }
}
