package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			saindoDoPrograma()
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}

}

// ----- Functions------

func exibeIntroducao() {
	name := "Felipe"
	fmt.Println("Olá", name)
	time.Sleep(1 * time.Second)
}

func leComando() int {
	var comandoLido int
	fmt.Printf("Selecione o número do que deseja: ")
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("preparando para monitorar....")
	time.Sleep(2 * time.Second)
	fmt.Println("")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("testando site:", i, ":", site)

			time.Sleep(1 * time.Second)
			testaSite(site)
			fmt.Println("")
		}
		fmt.Println("Testando novamente....")
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}

}
func imprimeLogs() {
	fmt.Println("Exibindo Logs....")
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
func saindoDoPrograma() {
	fmt.Println("Saindo do programa....")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso! StatusCode: ", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas... StatusCode: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	return sites
}

func registraLog(sites string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 - ") + sites + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
