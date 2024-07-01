# Online-Shop  
Online shop is a task to finished submission from Synapsis.id.

## Main Feature
### 1. Auth <br>
  This feature allows users to authenticate and register their accounts, ensuring that only registered users can access and interact.
### 2. Product <br>
  Users can create and get product within the Online shop.

### 3. Transaction <br>
Users can perform transactions related to products or services within the online shop. 

## Technology Architecture

### 
- Golang
- Fiber
- Postman
- PostgreSQL

  
##  Developer
- [Husada Hutasoit](https://github.com/husadahts)          

## Future Feature
1. Chat


The service available:
> Base url of this service is: http://localhost:4000/

# Authentications

This service utilizes token-based authentication, requiring users to have an account for accessing its features. If you don't have an account, please create a new one. Once registered, you can generate an authentication token. This token serves as a means of logging in, requiring you to authenticate yourself using your email and password. If the authentication is successful, you will receive an access token, enabling you to access the service. If the authentication fails, an error message will be displayed.

The provided tokens are the accessToken. The accessToken remains valid for one hour.
By following this authentication process, you can securely access the service and enjoy its functionalities.

# Instructions
This project run in go1.20.5  

1. Open the `.env.example file` in the root directory of your project
```bash
# HTTP SERVER
DB_HOST=YOUR_HOST
DB_PORT=YOUR_PORT
DB_NAME=YOUR_NAME
DB_USER=YOUR_USER
DB_PASSWORD=YOUR_PASSWORD
ENCRYPTION_SALT=SALTNUMBER
JWT_SECRET=YOUR_SECRET

```
   
2. Copy and paste `.env.example` file into `.env` file in your project 
3. Create config.yaml to configuration of application
```bash 
mkdir config/database/test.json  
```
```bash
{
  "user": YOUR_USER_POSTGRES,
  "password": YOUR_PASSWORD_POSTGRES,
  "host": YOUR_HOST_POSTGRES,
  "port": YOUR_PORT_POSTGRES,
  "database": YOUR_DATABASE_POSTGRES
}

```

5. Run server:
```bash 
go run cmd/api/main.go  
```


I'd be happy to review any pull requests that may better the Online Shop API, in particular if you have a bug fix, enhancement, or a new idea, you can contact us.
