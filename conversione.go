package main
import (
  "fmt"
  "math"
  )
var decobin string;
var numerocorrente int
var decbin string
var numerin int
var numerout int
var resti int
var contatore int
var i int
var numeroutput int
var altraconversione string
var STOP bool
var ultimacifra int
func vedi_se_bin (numerodacontrollare int)  {
  fmt.Println("vedisebin chiamato")
  for numerodacontrollare>0 {
    ultimacifra=numerodacontrollare%10
    if ultimacifra!=0 && ultimacifra!=1 {
      STOP=true
    }
    numerodacontrollare=numerodacontrollare/10
  }
}
func in_bin (numeroinput int) int {
  contatore=1
  resti=0
  STOP=false
  numerocorrente=numeroinput
  for numerocorrente>0  {
    resti=resti+numerocorrente%2*contatore
    numerocorrente=numerocorrente/2
    contatore=contatore*10
  }
  return resti
}
func in_dec (numeroinput int) int {
  i=0
  numeroutput=0
  STOP=false
  vedi_se_bin(numeroinput)
  numerocorrente=numeroinput
  for numerocorrente>0 {
    numeroutput=numeroutput+numerocorrente%10*int(math.Pow(2,float64(i)))
    numerocorrente=numerocorrente/10
    i++
  }
  return numeroutput
}
func conversione(numeroin int,decbin string) int {
  if decbin=="bindec" {
    return in_dec(numeroin)
  }else{
    return in_bin(numeroin)
  }
}

func main() {
	fmt.Println("Vuoi convertire da binario ad decimale(bindec) o da decimale a binario(decbin?)")
  fmt.Printf("(bindec o decbin)")
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
    fmt.Printf("(SOLO bindec O decbin)")
    main();
  }else{
    fmt.Printf("Inserisci un numero da convertire\nNumero:")
    n, err := fmt.Scanf("%d", &numerin)
    if err != nil {
        fmt.Println(n, err)
    }
    numerout=conversione(numerin,decobin)
    if(STOP)  {
      fmt.Println("DEVI INSERIRE PER FORZA UN NUMERO BINARIO SE VERAMENTE VUOI FARE LA CONVERSIONE DA BINARIO A DECIMALE")
      main();
    }
    fmt.Printf("Il tuo numero convertito è %v\n",numerout)
  }
  fmt.Printf("Vuoi effettuare un'altra conversione?\n(Sì o No)")
  fmt.Scanf("%s",&altraconversione)
  if altraconversione=="Sì" || altraconversione=="SI" || altraconversione=="Si" || altraconversione=="S" || altraconversione=="s" || altraconversione=="s" || altraconversione=="si" || altraconversione=="sì"  {
    main();
  }
  for altraconversione!="no" && altraconversione!="No" && altraconversione!="n" && altraconversione!="N" && altraconversione!="NO" {
    fmt.Printf("Vuoi effettuare un'altra conversione?\n(Solo Sì o No)")
    fmt.Scanf("%s",&altraconversione)
    if altraconversione=="Sì" || altraconversione=="SI" || altraconversione=="Si" || altraconversione=="S" || altraconversione=="s" || altraconversione=="s" || altraconversione=="si" || altraconversione=="sì"  {
      main();
    }
  }
  }
