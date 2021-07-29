# pass-legitimator 
[![Go-Shield]][go-shield]
[![Apache-Shield]][apache-shield]
[![LinkedIn][linkedin-shield]][linkedin-url]
<br>
[![Testing](https://github.com/wy7-source/pass-legitimator/actions/workflows/testing_build.yml/badge.svg)](https://github.com/wy7-source/pass-legitimator/actions/workflows/testing_build.yml)
[![Linting](https://github.com/wy7-source/pass-legitimator/actions/workflows/golangci_lint.yml/badge.svg)](https://github.com/wy7-source/pass-legitimator/actions/workflows/golangci_lint.yml)

That's a smaller and efficient password validator.

  
  

## About the Project

This project, have portfolio purpose only.

He uses theses validations:
- 9 or more characters
- At least one number
- At least one uppercase letter
- At least one lowercase letter
- At least one special character, which can be: !@#$%^&*()-+
- No duplicate characters

Example:

- IsValid("") // false  
- IsValid("aa") // false  
- IsValid("ab") // false  
- IsValid("AAAbbbCc") // false  
- IsValid("AbTp9!foo") // false  
- IsValid("AbTp9!foA") // false
- IsValid("AbTp9 fok") // false
- IsValid("AbTp9!fok") // true

## Getting Started 🚀 
First of all, install Golang 1.16 or higher version, because he is required.
Than, you can clone repo and run:
``` bash
go run main.go api
```
Or, import this package and use on your application with:
``` bash
go get github.com/wy7-source/pass-legitimator
```
<br>

## Docker-way to Getting Started :whale:
You also can use the official image to run the API on a isolated container:
```bash
docker run -p 9000:9000 wy7images/pass-legitimator:latest	
```
So send the password over POST to validate on *localhost:9000/validate* with this body:
```bash
  {
      "value": "AbTp9!fok"
  } 
  ```

And, the response will seems like that:
```bash
{
    "message": "true"
}
```
<br>

## Detalhes/ Decisões

- Linguagem/ Tecnologia: 
    Optei pela Go Lang. Que além de performática, é ágil na construção de qualquer tamanho de aplicação. Aqui, fiz uma implementação o mais pura possível, visto que a propria linguagem tem muitos recursos nativos, e a comunidade da linguagem é bem purista também. Então, nada de framework mvc ou restfull. Só o jeitinho do GoLang (Go Way).

- Design/Arquitetura:
    É uma implementação de Arquitetura Hexagonal. Não optei pela Clean Architecture, pois acredito que os use-cases abrangem contextos de domínios, e o mais proximo de domínio que temos aqui, nada que um serviço não resolva.

    Levando em consideração as regras de validação de senha ( que são nossas Regras de Negócio) será possível ver 2 tipos de validações. Uma dentro do domínio, e uma como serviço de domínio.

    O Motivo é simples: Single Responsability Segregation + Domínios Ricos.

    Caso eu optasse por validações somente no domínio, ele seria em sí Rico, mas a extensibilidade, e os testes dessa opção seriam caóticos. Por isso eu implemento sim uma validação no Domínio para salientar a importancia de ter um domínio bem ou minimamente validado, e segrego as validações, para que sejam extensíveis, de fácil manutenção, e "plugáveis".

    A causa dessa decisão está no tópico abaixo.

- As Validações:
    A primeira tentativa foi fazer uma única validação com um Regex com Negative Lookahead e Lookbehind. Mas por razões de performance, a implementação de Regex no Go lang não tem suporte a essas features (Veja mais sobre [aqui](https://stackoverflow.com/questions/26771592/negative-look-ahead-in-go-regular-expressions)), sem contar que a manutenção e testes de um "Super Regex" que abarca todas as regras de negócio, seria complicada.

    Criando validações separadas e customizadas, tirando assim a possibilidade das validações serem feitas unica e exclusivamente no domínio, cria a necessidade de ter testes unitários para cada validação.

    Portanto, cada regra de negócio tem uma validação, e cada validação tem seu teste unitário.

- Multithreading:
    Com validações separadas, o serviço pode escolher como gerenciar o processamento. E com isso, optei por fazer as validações de forma multithreading com as rotinas do Go Lang, e também com o Semaphore Pattern.
    Todo o serviço foi comentado inclusive com links úteis para entendimento do fluxo do código. Bem tranquilo :)

- Execução da Aplicação:
    Foi criada uma CLI para interagir com a aplicação. 
    Com essa CLI, é possível ter uma descrição da aplicação, comandos que você pode utilizar  e também customizar a porta que a API vai subir.

- Testes: 
    Foram feitos somente de unidade. Não foram necessários testes de integração, nem mocks de serviços externos, até porque a aplicação hoje, não conta com persistência de dados por exemplo. Mas coisas assim, devida a decisão de design e arquitetura, é algo fácilmente implementável.


- Acoplamento e Injeção de Dependência:
    No Golang, a injeção de dependência é feita automaticamente quando alguém implementa uma interface, e é exatamente o que acontece aqui. O Domínio define a interface de serviço, e a implementação dos métodos da interface, elegem uma struct como injetável para aquela interface.

    O Máximo de código depende de interfaces, não de implementações. Exceto em casos onde eram necessários ponteiros do dominio, como por exemplo no momento em que fazemos um Bind/Hydrate da DTO para a struct do Domínio.

- Fluxo dentro da aplicação:
    - Primeiro, a requisição chega ao Handler responsável pelo endpoint "*/validate*", e qualquer problema de serealização no endpoint, ele retornará o erro.
    - É feito o bind/hydrate da dto para a struct de domínio. Se o domínio estiver válido, continua, se não, retorna false.
    - E enfim todas as validações entram em ação através do serviço. Se a senha for válida. retorna true, se não for, false (assim espero).

That's all folks :)
<br>

## License
Distributed under the Apache 2.0. See `LICENSE` for more information.




<!-- MARKDOWN LINKS & IMAGES -->
[go-shield]: https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go
[apache-shield]: https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/wythor
