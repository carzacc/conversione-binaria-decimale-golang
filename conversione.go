package main
import (
  "fmt"
  )
var decobin string;
func main() {
	fmt.Println("Vuoi convertire da binario ad decimale(bindec) o da decimale a binario(decbin?)")
  n, err := fmt.Scanf("%s", &decobin)
  if err != nil {
      fmt.Println(n, err)
  }
  if decobin==""{
  decobin="wooooooooow"
  }
  fmt.Printf("Hai inserito %v\n", decobin)
}
