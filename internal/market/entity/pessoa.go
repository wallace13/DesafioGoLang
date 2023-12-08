package entity

type Pessoa struct {
	Nome      string
	Idade     int
	Pontuacao int
}

func NewPessoa(nome string, idade, pontuacao int) *Pessoa {
	return &Pessoa{
		Nome:      nome,
		Idade:     idade,
		Pontuacao: pontuacao,
	}
}
