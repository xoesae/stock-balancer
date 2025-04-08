# Stock Balancer

![Go CLI](https://img.shields.io/badge/Go-CLI-blue) ![License](https://img.shields.io/badge/license-MIT-green)

Stock Balancer é uma ferramenta de linha de comando (CLI) desenvolvida em Go para ajudar investidores a rebalancear suas carteiras de ações. O programa permite calcular rapidamente as quantidades de ações necessárias para atingir um percentual ideal definido para cada ativo, com base em um valor de investimento.

## Recursos

- **Rebalanceamento automatizado:** Calcula as quantidades necessárias para atingir as porcentagens ideais da carteira.
- **Integração com a Bolsa de Valores Brasileira:** Atualiza os valores das ações em tempo real.
- **Configuração flexível:** Define as porcentagens ideais das ações em um arquivo JSON.
- **Dois comandos principais:**
    - `list`: Exibe as ações cadastradas na carteira.
    - `rebalance`: Calcula os valores necessários para rebalancear a carteira com base em um valor de investimento informado.

## Como Funciona

1. O programa busca informações sobre as ações definidas em um arquivo JSON.
2. Atualiza os valores das ações com dados da Bolsa de Valores Brasileira.
3. Calcula a quantidade de ações necessárias para rebalancear a carteira com base no investimento informado.

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/xoesae/stock-balancer.git
   cd stock-balancer
   ```

2. Compile o projeto:
   ```bash
   go build -o stock-balancer
   ```

3. Execute o programa:
   ```bash
   ./stock-balancer
   ```

## Uso

### Listar as ações cadastradas
```bash
./stock-balancer list
```

### Calcular rebalanceamento com base em um valor de investimento
```bash
./stock-balancer rebalance 1000
```

### Exemplo de JSON de Configuração
```json
{
  "stocks": [
    {
      "ticker": "BBAS3",
      "ideal_ratio": 0.5,
      "current_price": 27.78,
      "amount": 1,
      "updated_at": "2025-02-07T21:35:24.02599019-03:00"
    },
    {
      "ticker": "BBDC4",
      "ideal_ratio": 0.5,
      "current_price": 11.99,
      "amount": 1,
      "updated_at": "2025-02-07T21:35:24.257580504-03:00"
    }
  ]
}
```

## Contribuição

Sinta-se à vontade para abrir issues e enviar pull requests com melhorias. Todas as sugestões são bem-vindas!

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.

---

