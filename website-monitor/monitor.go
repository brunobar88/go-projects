package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
) // Modularização, tem pacote pra varias coisas

const qtdMonitorings = 5
const interval = 5 // Minutes

func main() { // função principal parecido com C

	showHelloMessage()

	for { //for sem condição = loop infinito
		showMenu()

		chose := catchUserCommand()

		fmt.Println(chose)

		switch chose {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("======================= Logs =======================")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}
}

func showHelloMessage() {
	name := "Bruno B. Zanella"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("CLI Version:", version)
}

func showMenu() {
	fmt.Println("0- Sair do CLI")
	fmt.Println("1- Iniciar Scan")
	fmt.Println("2- Exibir Logs")
}

func catchUserCommand() int {
	var chose int
	fmt.Scan(&chose)

	return chose
}

func initMonitoring() {
	fmt.Println("Monitorando...")

	websites := readWebsitesFile()

	for index := 0; index < qtdMonitorings; index++ { // For e ifs nÃo ultilizam parenteses para a condição
		for _, website := range websites {
			testWebsite(website)
		}

		time.Sleep(interval * time.Minute) // pausa execução do programa
	}

	fmt.Println("")
}

func testWebsite(website string) {
	response, err := http.Get(website)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	writeLog(website, response.StatusCode)

	if response.StatusCode == 200 { // ifs so podem receber true ou false
		fmt.Println("O site:", website, "está OK (", response.StatusCode, ")")
		return
	}
	fmt.Println("O site:", website, "está fora do ar ou está com problemas (", response.StatusCode, ")")
}

func readWebsitesFile() []string {

	var websites []string

	file, err := os.Open("websites.txt") // abre um arquivo

	// tratamento de erro
	if err != nil {
		fmt.Println("occorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line) // Remove espaços no fim ou começo de uma string, inclusive os \n no fim da linha

		websites = append(websites, line)

		if err == io.EOF { // constantes de erro
			break
		}
	}

	file.Close() // fecha o arquivo aberto

	return websites
}

func writeLog(website string, status int) {
	file, err := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // abre um arquivo porém dá mais opções

	if err != nil {
		fmt.Println(err)
	}

	var statusString string

	if status == 200 {
		statusString = "online"
	} else {
		statusString = "offline"
	}

	//geito go de fazer formatação de horas.
	file.WriteString("[" + time.Now().Format("02/01/2006 15:04:05") + "] - " + website + " " + statusString + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("Log.txt") // abre e le arquios inteiros

	if err != nil {
		fmt.Println(err)
	}

	if len(file) > 0 {
		fmt.Println(string(file))
		return
	}

	fmt.Println(":Empty Log:")
}
