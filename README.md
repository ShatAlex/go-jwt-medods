# Go-JWT-MEDODS

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ShatAlex/TradingApp)
![Static Badge](https://img.shields.io/badge/gin-v1.9.1-brightgreen)
![Static Badge](https://img.shields.io/badge/mongodriver-v1.12.1-yellow)
![Static Badge](https://img.shields.io/badge/jwtgo-v3.2.0-red)
![Static Badge](https://img.shields.io/badge/swagger-v1.16.1-purple)



## :sparkles: Описание проекта
Часть сервиса аутентификаци, разработанная в соответствии с правилами чистой архитектуры.  

Большое спасибо компании "MEDODS" за интересное тестовое задание.
___

## :clipboard: Использование
Для запуска приложения через make-файл в докере:
```
make run
```
___

## :pushpin: REST Endpoints

### AUTH
Эта группа конечных точек предназначена для регистрации и аутентификации пользователей. 

##### POST - /auth/sign-up
Example Input:
```
{
	"username": "user",
	"password": "qwerty"
} 
```
Example Response:
```
{
	"guid": "a2f274a6-9d53-4b55-8ee1-bcc4faf0c5a5"
} 
```
###### POST - /auth/sign-in
Example Input:
```
{
	"guid": "a2f274a6-9d53-4b55-8ee1-bcc4faf0c5a5"
} 
```
Example Response:
```
{
	"access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ3MDg3OTgsImd1aWQiOiJhMmYyNzRhNi05ZDUzLTRiNTUtOGVlMS1iY2M0ZmFmMGM1YTUifQ.gXY-Y0ci1j8IUVh55b7pAQoBCox3AZZP0yc4IToXVDEU_cTF4GVxyunXuNMr3gZFDXln-lOraKEwA2QAnX78jA",
	
  	"refresh_token": "12bd6973cf5786a70e06b4b9e172a7895edf400f22e55630993b5bd5d691500a"
} 
```
### TOKENS
Эндпоинт для выполнения refresh операции над Access и Refresh токенами.

Эндопинт доступен только аутентифицированным пользователям.
##### POST - /tokens/refresh/
Example Input:
```
{
  "refresh_token": "12bd6973cf5786a70e06b4b9e172a7895edf400f22e55630993b5bd5d691500a"
}
```
Example Response:
```
{
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ3MDg4ODcsImd1aWQiOiJhMmYyNzRhNi05ZDUzLTRiNTUtOGVlMS1iY2M0ZmFmMGM1YTUifQ.6KhAxwo592o76QefT8cVlbNJQ6artNSxuPB0jFz7EkB4gLnfPZRcNet4KNvOhoMbsH4CyOI-gS1EmjtRErAaMw",

  "refresh_token": "03abb3fd7d58624f13ae30e1275a4bbad3c723e9ab6fa8eaa06303c4f8727e13"
}
```