package transformer

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/wallace13/imersaodev/go/internal/market/entity"
)

func GerarArquivo(filePath string, pessoas []entity.Pessoa) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Define o cabeçalho
	writer.Write([]string{"Nome", "Idade", "Pontuação"})

	// Escreve os dados do arquivo
	for _, pessoa := range pessoas {
		writer.Write([]string{pessoa.Nome, fmt.Sprint(pessoa.Idade), fmt.Sprint(pessoa.Pontuacao)})
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println("Erro ao escrever arquivo:", err)
	}
}
