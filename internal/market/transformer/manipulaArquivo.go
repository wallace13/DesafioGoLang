package transformer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/wallace13/imersaodev/go/internal/market/entity"
)

// lê as linhas do arquivo CSV.
func LerLinhasDoArquivo(filePath string) ([]string, error) {
	// Abre o arquivo CSV
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var stringLines []string
	for _, line := range lines {
		stringLines = append(stringLines, strings.Join(line, ","))
	}

	return stringLines, nil
}

// ConverterParaPessoas as linhas do CSV
func ConverterParaPessoas(lines []string) ([]entity.Pessoa, error) {
	var pessoas []entity.Pessoa

	for i, line := range lines {

		//ignora o cabeçalho
		if i == 0 {
			continue
		}

		fields := strings.Split(line, ",")

		nome := fields[0]
		idade := 0
		pontuacao := 0

		if len(fields) > 1 {
			fmt.Sscanf(fields[1], "%d", &idade)
		}
		if len(fields) > 2 {
			fmt.Sscanf(fields[2], "%d", &pontuacao)
		}

		//padroniza nomes para todos começarem com letra Maiuscula
		//nome = strings.Title(nome)

		pessoa := entity.Pessoa{Nome: nome, Idade: idade, Pontuacao: pontuacao}
		pessoas = append(pessoas, pessoa)
	}

	return pessoas, nil
}
