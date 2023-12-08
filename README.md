# Ordenação de Dados em Go

## FASE 01

### Informações do Desafio

**Objetivo:**
Criar um programa em Go que lê informações de um arquivo. Realizar a ordenação dos dados lidos por diferentes critérios, como nome e idade. A ordenação por nome deve seguir a tabela ASCII. Salvar o resultado ordenado em outro arquivo.

**Requisitos Técnicos:**
Utilizar a biblioteca padrão do Go para manipulação de arquivos. Considere trabalhar com dados estruturados, como linhas de um arquivo CSV.

### Estrutura do Programa

1. Crie um arquivo de entrada contendo dados (por exemplo, CSV, cada linha representando uma entidade com algumas informações).
2. Implemente funções que leiam o arquivo, realizem a ordenação dos dados por nome e por idade, e retornem os resultados.
3. Salve os resultados ordenados em dois novos arquivos (um para cada critério de ordenação).

### Exemplo

**Arquivo de Entrada (arquivo-origem.csv):**

- Nome,Idade,Pontuação
- Carlos,30,80
- Carlos,22,75
- Maria,30,95
- Joao,25,80
- carlos,15,90


**Exemplo de Arquivo de Saída (arquivo-destino.csv):**

- Nome,Idade,Pontuação
- Carlos,22,75
- Carlos,30,80
- carlos,15,90
- Joao,25,80
- Maria,30,95


### Observações:

- Utilize a lógica de manipulação de arquivos em Go.
- Implemente a ordenação por nome e por idade.
- Se desejar, crie funções separadas para leitura, processamento e escrita de arquivos.
- Lide com erros de maneira apropriada.

### Comando de Execução:

Precisamos rodar o comando go run main.go ./arquivo-origem.csv ./arquivo-destino.csv. Será usado um arquivo CSV aleatório nosso para testar sua aplicação.
