# pass-legitimator 
[![Go-Shield]][go-shield]
[![Apache-Shield]][apache-shield]
[![LinkedIn][linkedin-shield]][linkedin-url]

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

## License
Distributed under the Apache 2.0. See `LICENSE` for more information.




<!-- MARKDOWN LINKS & IMAGES -->
[go-shield]: https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go
[apache-shield]: https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/wythor
