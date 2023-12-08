package main

import (
	"fmt"
	"sort"

	"github.com/wallace13/imersaodev/go/internal/market/entity"
	"github.com/wallace13/imersaodev/go/internal/market/transformer"
)

func main() {
	// Caminho do arquivo CSV
	filePath := "arquivo/arquivo-origem.csv"

	// Lê arquivo
	lines, err := transformer.LerLinhasDoArquivo(filePath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Converte em Pessoa
	pessoas, err := transformer.ConverterParaPessoas(lines)
	if err != nil {
		fmt.Println("Erro ao converter para pessoas:", err)
		return
	}

	// Exibe as pessoas lidas do CSV
	fmt.Println("Pessoas Lidas:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d, Pontuação: %d\n", pessoa.Nome, pessoa.Idade, pessoa.Pontuacao)
	}

	// Ordena por nome
	sort.Slice(pessoas, func(i, j int) bool {
		return pessoas[i].Nome < pessoas[j].Nome
	})

	fmt.Println("\nPessoas ordenadas por nome:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d, Pontuação: %d\n", pessoa.Nome, pessoa.Idade, pessoa.Pontuacao)
	}

	transformer.GerarArquivo("arquivo/arquivo-destino.csv", pessoas)

	// Ordena por Idade
	queueIdade := entity.NewPessoaQueueIdade()
	queueIdade.Pessoas = pessoas

	fmt.Println("\nPessoas ordenadas por idade:")
	for queueIdade.Len() > 0 {
		pessoa := queueIdade.Pop().(entity.Pessoa)
		fmt.Printf("Idade: %d, Nome: %s, Pontuação: %d\n", pessoa.Idade, pessoa.Nome, pessoa.Pontuacao)
	}

}
