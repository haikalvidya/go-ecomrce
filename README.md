# Overview
It's an API of simple ecommerce to CRUD product and add image of product based on Echo framework.

## How to run it
1. Run the application using the command in the terminal:

    `make run-app`
2. Browse to localhost:7788/swagger/index.html.
3. To get access to all feature product, you need to regist and login the account.
4. After the login, copy a token from the response, then click "Authorize", enter the value for "apiKey" in a form:
"Bearer {token}". For example:

    ```
      Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTYzNTU4NDMwNiwiZXhwIjoxNjM1NTg3OTA2fQ.s5hFizdpccvA6cM6p--RhiD81qUpn_typ3LwAJgPbVU
    ```

   ![authorize button](./assets/authorize.png)
Then, click "Authorize" and close the popup.
5. To see uploads image, go to path ```localhost:7788/uploads/{filename}```

## What's feature of this project:

- Registration
- Authentication with JWT
- CRUD API for product
- Upload Image of product
- Migrations
- Swagger docs
- Docker development environment

![swagger](./assets/swagger.png)

## Tech Stack

- **Language:** [Go](https://golang.org/)
- **Framework:** [Echo](https://echo.labstack.com/)
- **Authentication:** [jwt-Go](github.com/dgrijalva/jwt-go)
- **Database:** [MySQL](https://www.mysql.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Migration:** [Go-mirations](github.com/ShkrutDenis/go-migrations)
- **Mock:** [Mocking db](https://github.com/selvatico/go-mocket)
- **API Docs:** [Swagger](https://github.com/swaggo/echo-swagger)
- **Container:** [Docker](https://www.docker.com/)

# Architecture Overview

The architecture is heavily influenced by the Clean Architecture and Hexagonal Architecture. Clean Architecture is an architecture where the business rules can be tested without the UI, database, web server, or any external element.

<p align="center">
  <img src="https://cdn-images-1.medium.com/max/719/1*ZNT5apOxDzGrTKUJQAIcvg.png" width="350"/>
  <img src="https://cdn-images-1.medium.com/max/900/0*R7uuhFwZbhcqZSvn" width="350" /> 
</p>

<p align="center">
  <img src="https://cdn-images-1.medium.com/max/1200/0*rFs1UtU4sRns5vCJ.png" width="350" />
  <img src="https://cdn-images-1.medium.com/max/1200/0*C-snK7L4sMn7b6CW.png" width="350" /> 
</p>
