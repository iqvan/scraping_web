package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getContenidoPaginaFiltro(urlRuta string) string {
	// Request the HTML page.
	res, err := http.Get(urlRuta)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	band := (doc.Find("a").Text())

	return band
}

func getContenidoPagina(urlRuta2 string) string {
	// Request the HTML page.
	res, err := http.Get(urlRuta2)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	return newStr

}

func operacionMatematica(urlRuta2 string, valorInicial int) int {

	words := strings.Fields(urlRuta2)

	operation := words[0]
	numberInicial := words[1]

	number, err := strconv.Atoi(numberInicial)
	if err != nil {
		log.Fatal(err)
	}

	switch operation {
	case "divide":
		valorInicial = valorInicial / number
	case "minus":
		valorInicial = valorInicial - number
	case "multiply":
		valorInicial = valorInicial * number
	case "add":
		valorInicial = valorInicial + number
	default:
		fmt.Println("Este falto")
	}

	return valorInicial
}

func main() {

	var urlRuta string = "http://10.10.88.82:3010/"

	respBody := getContenidoPaginaFiltro(urlRuta)
	fmt.Print("La respuesta final es: \n")
	fmt.Printf(respBody + "\n")
	//for respBody == "9765" {
	var urlRuta2 string = "http://10.10.88.82:" + respBody + "/"

	var valorInicial int = 0

	respBody2 := getContenidoPagina(urlRuta2)

	fmt.Printf("La operación solicitada es: " + respBody2 + "\n")

	respBody3 := operacionMatematica(respBody2, valorInicial)

	respBody4 := strconv.Itoa(respBody3)

	fmt.Printf("El resultado de la operacion es: " + respBody4 + "\n")

	//}
	/*timeout := time.After(4 * time.Second)
	for ; condition ; {
	    // Code to repeatedly execute
	}
	*/
}
