package entity

type PessoaQueueIdade struct {
	Pessoas []Pessoa
}

// verifica ordenação
func (oq *PessoaQueueIdade) Less(i, j int) bool {
	return oq.Pessoas[i].Idade < oq.Pessoas[j].Idade
}

// troca
func (oq *PessoaQueueIdade) Swap(i, j int) {
	oq.Pessoas[i], oq.Pessoas[j] = oq.Pessoas[j], oq.Pessoas[i]
}

// tamanho
func (oq *PessoaQueueIdade) Len() int {
	return len(oq.Pessoas)
}

// adiciona
func (oq *PessoaQueueIdade) Push(x interface{}) {
	oq.Pessoas = append(oq.Pessoas, x.(Pessoa))
}

// remove
func (oq *PessoaQueueIdade) Pop() interface{} {
	if len(oq.Pessoas) == 0 {
		return nil
	}

	maxIndex := 0

	for i := 1; i < len(oq.Pessoas); i++ {
		if oq.Less(i, maxIndex) {
			maxIndex = i
		}
	}

	item := oq.Pessoas[maxIndex]

	oq.Pessoas = append(oq.Pessoas[:maxIndex], oq.Pessoas[maxIndex+1:]...)
	return item
}

func NewPessoaQueueIdade() *PessoaQueueIdade {
	return &PessoaQueueIdade{}
}
