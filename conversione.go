package main
import (
  "fmt"
  )
var decobin string;
var numerocorrente int
var decbin string
var numerin int
var numerout int
func conversione(numeroin int,decbin string) int {
  numerocorrente=numeroin
  if decbin=="bindec" {
    return numeroin
  }else{
    return numeroin+50
  }
}

func main() {
	fmt.Println("Vuoi convertire da binario ad decimale(bindec) o da decimale a binario(decbin?)")
  n, err := fmt.Scanf("%s", &decobin)
  if err != nil {
      fmt.Println(n, err)
  }
  if decobin==""{
  decobin="niente"
  }
  fmt.Printf("Hai inserito %v\n", decobin)
  if decobin!="bindec" && decobin!="decbin" {
    fmt.Println("DEVI INSERIRE PER FORZA bindec O decbin a seconda di che vuoi fare")
    main();
  }else{
    fmt.Println("Inserisci un numero da convertire")
    n, err := fmt.Scanf("%d", &numerin)
    if err != nil {
        fmt.Println(n, err)
    }
    numerout=conversione(numerin,decobin)
    fmt.Printf("Il tuo numero convertito è %v\n",numerout)
  }
  }
