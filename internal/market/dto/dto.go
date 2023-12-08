package dto

type Pessoa struct {
	Nome      string `csv:"nome"`
	Idade     int    `csv:"idade"`
	Pontuacao int    `csv:"pontuacao"`
}
