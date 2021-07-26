# pass-legitimator
Esse é um pequeno validador de senhas, porém eficiente.

## How to run 🚀 
- Você pode executar a aplicação das seguintes formas:
	- Docker:
Com o comando abaixo: 
        ```bash
        docker run -p 9000:9000 wy7images/iti-challenge:latest
        ```
    - Executável: 
        Você pode baixar um executável para seu linux nesse [link](https://github.com/wy7-source/iti-challenge/releases)

    - Go way: 
        Com o proprio GoLang na sua máquina ( versão ^1.16.0 ) no diretório do repo:
        ``` bash
        go run main.go api
        ```

- Com a aplicação rodando, você deve enviar um POST para *localhost:9000/validate* com o seguinte body:
    ```bash
    {
        "value": "AbTp9!fok"
    } 
    ```

- E logo em seguida você recebera uma resposta no seguinte formato:

    ```bash
    {
        "message": "true"
    }
    ```
<br>
