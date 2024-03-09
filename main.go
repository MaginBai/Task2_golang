package main

import (
	"Convert/internal"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFileName := flag.String("input", "", "Имя входного файла в формате JSON")
	flag.Parse()
	if *inputFileName == "" {
		log.Fatal("Необходимо указать имя входного файла с помощью флага -input")
	}
	file, err := os.ReadFile(*inputFileName)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	yamlData, err := internal.Convert(file)
	if err != nil {
		log.Fatalf("Ошибка конвертации: %v", err)
	}
	outputFileName := "output.yaml"
	fileOutput, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fileOutput.Write(yamlData)
	defer fileOutput.Close()
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}

	fmt.Printf("Конвертация успешно завершена. Результат записан в %s\n", outputFileName)

}
