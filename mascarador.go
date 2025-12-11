package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// --- CONFIGURAÇÕES ---
// Define se o mascaramento começa com asterisco ou não.
const comecaComAsterisco = false

// ---------------------

func main() {
	// 1. Encontrar todos os arquivos .csv na pasta atual
	files, err := filepath.Glob("*.csv")
	if err != nil {
		fmt.Println("Erro ao buscar arquivos:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Nenhum arquivo CSV encontrado na pasta.")
		return
	}

	fmt.Printf("Encontrados %d arquivos CSV. Iniciando processamento...\n\n", len(files))

	// 2. Loop para processar cada arquivo encontrado
	for _, filename := range files {
		// Evita processar arquivos que já são resultado de um processamento anterior
		if strings.Contains(filename, "_mascarado") {
			continue
		}

		processarArquivo(filename)
	}
}

func processarArquivo(inputFile string) {
	fmt.Printf("-> Processando: %s... ", inputFile)

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("\n[ERRO] Falha ao abrir '%s': %v\n", inputFile, err)
		return
	}
	defer file.Close()

	// Detectar delimitador
	scanner := bufio.NewScanner(file)
	var firstLine string
	if scanner.Scan() {
		firstLine = scanner.Text()
	} else {
		fmt.Println("[PULADO] Arquivo vazio.")
		return
	}

	delimiter := detectDelimiter(firstLine)
	
	// Voltar ao inicio do arquivo
	file.Seek(0, 0)

	reader := csv.NewReader(file)
	reader.Comma = delimiter
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Criar nome do arquivo de saída dinamicamente
	// Ex: "dados.csv" vira "dados_mascarado.csv"
	outputFile := strings.TrimSuffix(inputFile, ".csv") + "_mascarado.csv"
	
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("\n[ERRO] Falha ao criar arquivo de saída: %v\n", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	writer.Comma = delimiter
	defer writer.Flush()

	lineCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Ignora linhas com erro mas continua o arquivo
			continue
		}

		// Cabeçalho (linha 0) apenas copia
		if lineCount == 0 {
			writer.Write(record)
			lineCount++
			continue
		}

		// Processa as duas primeiras colunas
		for i := 0; i < len(record); i++ {
			if i <= 1 {
				record[i] = applyMask(record[i], comecaComAsterisco)
			}
		}

		writer.Write(record)
		lineCount++
	}

	fmt.Printf("Feito! (Gerado: %s)\n", outputFile)
}

func detectDelimiter(line string) rune {
	separators := []rune{',', ';', '.'}
	maxCount := 0
	bestSep := ',' 

	for _, sep := range separators {
		count := strings.Count(line, string(sep))
		if count > maxCount {
			maxCount = count
			bestSep = sep
		}
	}
	return bestSep
}

func applyMask(s string, starFirst bool) string {
	runes := []rune(s)
	result := make([]rune, len(runes))

	for i, r := range runes {
		isEven := (i % 2 == 0)
		if starFirst {
			if isEven { result[i] = '*' } else { result[i] = r }
		} else {
			if !isEven { result[i] = '*' } else { result[i] = r }
		}
	}
	return string(result)
}
