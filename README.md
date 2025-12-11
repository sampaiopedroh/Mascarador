# CSV Mascarador

Este script em Go foi desenvolvido para mascarar os dados das duas primeiras colunas de arquivos CSV, substituindo caracteres alternados por asteriscos (`*`). Ele processa automaticamente todos os arquivos `.csv` encontrados na mesma pasta, detecta o delimitador (vírgula, ponto e vírgula ou ponto) e gera um novo arquivo com o sufixo `_mascarado.csv` para cada entrada.

## Funcionalidades

*   **Processamento em Lote:** Encontra e processa todos os arquivos `.csv` na pasta de execução.
*   **Detecção Automática de Delimitador:** Identifica se o CSV usa vírgula (`,`), ponto e vírgula (`;`) ou ponto (`.`) como separador.
*   **Mascaramento Seletivo:** Aplica a máscara apenas nas **duas primeiras colunas** dos dados, ignorando o cabeçalho.
*   **Mascaramento Flexível:** Permite configurar se a máscara começa com um asterisco ou um caractere original.
*   **Suporte a UTF-8:** Lida corretamente com caracteres especiais e acentuação.
*   **Geração de Novo Arquivo:** Cria um novo arquivo com os dados mascarados, sem alterar os arquivos originais.

## Como Usar

1.  **Pré-requisitos:** Certifique-se de ter o [Go](https://go.dev/doc/install) instalado em sua máquina.
2.  **Preparação:**
    *   Crie uma pasta vazia.
    *   Salve o código Go fornecido em um arquivo chamado `main.go` dentro dessa pasta.
    *   Coloque os arquivos CSV que deseja mascarar na **mesma pasta** que o `main.go`.
3.  **Execução:**
    *   Abra o terminal ou prompt de comando na pasta onde `main.go` e seus arquivos CSV estão localizados.
    *   Execute o script com o comando:
        ```bash
        go run main.go
        ```
4.  **Resultado:**
    *   Para cada arquivo `nome_original.csv` encontrado, um novo arquivo `nome_original_mascarado.csv` será gerado na mesma pasta, contendo os dados mascarados.
    *   Arquivos que já contêm `_mascarado` no nome serão ignorados para evitar processamento duplicado.

## Configuração

A única configuração disponível no código pode ser alterada diretamente no arquivo `main.go`:

```go
// --- CONFIGURAÇÕES ---
// Define se o mascaramento começa com asterisco ou não.
// true  = *o*e (ex: "nome" vira "*o*e")
// false = n*m* (ex: "nome" vira "n*m*")
const comecaComAsterisco = false 
// ---------------------
