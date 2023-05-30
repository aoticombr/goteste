package lib

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logger *Logger
)

func GetArqLog() string {
	//return GetPathLOG() + time.Now().Format("2006-01-02") + ".log"
	return time.Now().Format("2006-01-02") + ".log"
}

func SalvarLog(texto string) error {
	dataHoraAtual := time.Now().Format("2006-01-02 15:04:05")
	logString := fmt.Sprintf("[%s] %s\n", dataHoraAtual, texto)
	if Getview() {
		fmt.Printf(logString)
	}
	if Getdebug() {
		//fmt.Printf("Salvar Log")

		//fmt.Printf(logString)
		arquivo, err := os.OpenFile(GetArqLog(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer arquivo.Close()

		if _, err := arquivo.WriteString(logString); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func SalvarErro(texto string, erro error) error {
	//fmt.Println("SalvarErro..")
	dataHoraAtual := time.Now().Format("2006-01-02 15:04:05")
	logString := fmt.Sprintf("[%s] Erro(%s):>> %s\n", dataHoraAtual, texto, erro.Error())
	//fmt.Println("Imprimindo erro..")
	//fmt.Println(logString)
	arquivo, err := os.OpenFile(GetArqLog(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	if _, err := arquivo.WriteString(logString); err != nil {
		log.Fatal(err)
	}
	return erro
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
