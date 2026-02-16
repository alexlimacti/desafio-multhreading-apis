# Desafio Multithreading e APIs

Este projeto é a solução para o desafio de Multithreading da Pós-Graduação Go Expert da FullCycle. O objetivo é criar um sistema que consulte duas APIs de CEP simultaneamente e utilize a resposta daquela que for mais rápida, descartando a mais lenta.

## Requisitos do Desafio

1.  **Multithreading**: Realizar requisições simultâneas para:
    -   `https://brasilapi.com.br/api/cep/v1/{cep}`
    -   `http://viacep.com.br/ws/{cep}/json/`
2.  **Performance**: Acatar a API que entregar a resposta mais rápida.
3.  **Output**: Exibir no command line os dados do endereço e qual API enviou a resposta.
4.  **Timeout**: Limitar o tempo de resposta em 1 segundo. Caso contrário, exibir erro de timeout.

## Estrutura do Projeto

O projeto foi organizado seguindo boas práticas, com separação de responsabilidades:

-   `main.go`: Ponto de entrada da aplicação. Gerencia as goroutines, canais e o timeout.
-   `internal/api/`: Pacote contendo as implementações específicas das APIs e as estruturas de dados.
    -   `types.go`: Define a interface `Fetcher` e as structs comuns `Address` e `APIResponse`.
    -   `brasilapi.go`: Implementação da busca na BrasilAPI.
    -   `viacep.go`: Implementação da busca na ViaCEP.

## Como Executar

Para rodar o projeto, certifique-se de ter o Go instalado e execute o seguinte comando na raiz do projeto:

```bash
go run main.go
```

## Exemplo de Saída

### Sucesso (Uma das APIs responde em menos de 1 segundo)
```text
Start: 48609070
Received response from: BrasilAPI
Address: {Cep:48609070 State:BA City:Serrinha Neighborhood:Cidade Nova Street:Rua Macário Ferreira Service:open-cep}
```

### Timeout (Ambas demoram mais de 1 segundo)
```text
Timeout: Request took longer than 1 second
```

## Tecnologias Utilizadas

-   **Go**: Linguagem de programação utilizada.
-   **Goroutines**: Para execução concorrente das requisições.
-   **Channels**: Para comunicação entre as goroutines e a thread principal.
-   **Select**: Para gerenciar a corrida entre as respostas das APIs e o timeout.
