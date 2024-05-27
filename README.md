# Neobank Technical Test

This project is a technical test for the recruitment process at Neobank. It is designed to evaluate the technical skills and understanding of development concepts of prospective developers.

## Description

The project involves creating SPA React and backend RESTful API for user role based access and transaction management using Golang, PostgreSQL, and Docker.

## Tech Stack
Primary Techstack
- ReactJS
- Go
- PostgreSQL
- Docker

Frontend Tools
- [Vite](https://vitejs.dev/) : local development server

Golang Library
- [Gin](https://github.com/gin-gonic/gin) : Rest api framework
- [Gorm](https://github.com/go-gorm/gorm) : ORM database management
- [Wire](https://github.com/google/wire) : Automatic dependency injection
- [Zap](https://github.com/uber-go/zap) : Logging
- [GoCSV](https://github.com/gocarina/gocsv) : CSV Parser
- [MailjetClient](https://github.com/mailjet/mailjet-apiv3-go/v4) : Mailjet Golang Client
- [JWT](https://github.com/dgrijalva/jwt-go): JWT Library

# Backend
## Environment Variables

To run this project, you will need to add the following environment variables (.env file):

```env
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=postgres
DB_TIMEZONE=Asia/Jakarta
ENV=LOCAL
JWT_TTL=86400
JWT_SECRET=secret
BCRYPT_COST=10
OTP_TTL=86400
MAILJET_API_PUBLIC_KEY=<mailjet public key>
MAILJET_API_PRIVATE_KEY=<mailjet private key>
MAILJET_FROM_EMAIL=aliirsyaadn@gmail.com
MAILJET_FROM_NAME=Ali Irsyaad Nursyaban
MAILJET_SEND_OTP_TEMPLATE_ID=<mailjet template id for sending otp>
```

if you want run this project on dockerize environment, you need to change DB_HOST=host.docker.internal. Optional: ENV change to PROD to set gin to release mode.

## Installation
Clone the project to your local directory
```bash
git clone https://github.com/aliirsyaadn/neobank-technical-test
```

Change directory to the project folder
```bash
cd neobank-technical-test
cd backend
```

Run server 
```bash
make run
```

Run server on dockerize environment
```bash
make docker-up
```

The server will start running at localhost and port that specify on environment variables.

## Database Schema
```sql
CREATE TABLE "users" (
    "id" text PRIMARY KEY,
    "corporate_account_no" text NOT NULL UNIQUE,
    "corporate_name" text NOT NULL,
    "name" text NOT NULL,
    "role" text NOT NULL,
    "phone" text NOT NULL,
    "email" text NOT NULL UNIQUE,
    "password" text NOT NULL,
    "verified" boolean DEFAULT false,
    "created_at" timestamptz,
    "updated_at" timestamptz
);

CREATE TABLE "user_otps" (
    "id" serial PRIMARY KEY,
    "email" text NOT NULL,
    "otp" text NOT NULL,
    "expired_at" timestamptz NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz
);

CREATE TABLE "transactions" (
    "reference_no" text PRIMARY KEY,
    "total_transfer_amount" decimal NOT NULL,
    "total_transfer_record" bigint NOT NULL,
    "from_account_no" text NOT NULL,
    "maker_user_id" text NOT NULL,
    "transfer_date" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "status" text DEFAULT 'AWAITING_APPROVAL',
    "instruction_type" text NOT NULL,
    "transfer_type" text DEFAULT 'ONLINE',
    "estimated_service_fee" decimal DEFAULT 0,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "fk_transactions_maker_user_detail" FOREIGN KEY ("maker_user_id") REFERENCES "users"("id")
);

CREATE TABLE "transfer_records" (
    "id" serial PRIMARY KEY,
    "transaction_reference_no" text NOT NULL,
    "no" int NOT NULL,
    "to_account_no" text NOT NULL,
    "to_account_name" text NOT NULL,
    "to_bank_name" text NOT NULL,
    "amount" decimal DEFAULT 0,
    "description" text,
    "status" text DEFAULT 'AWAITING_APPROVAL',
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "fk_transfer_records_transaction" FOREIGN KEY ("transaction_reference_no") REFERENCES "transactions"("reference_no")
);
```

## Development
If you add changes on dependency, you can add your dependency on `app/provider.go` and simply run this command to automatic generate dependency injection
```bash
make wire
```


To run all of the test you can run this command
```bash
make unit-test
```

## API Documentation
The API documentation for this project is provided in the docs folder. It includes a `Postman collection`.

You can import the provided Postman collection (`postman_collection.json`) and environment file (`postman_environment.json`). These files include pre-configured requests for all the API endpoints, making it easy to test the API's functionality.

## Organization Code Explaination
The `Router-Handler-Service-Repository pattern` is a common architectural pattern used in web development, especially in Golang projects. This pattern helps to structure the code in a way that promotes separation of concerns, code reusability, and maintainability.

1. `Router`: The router is responsible for handling all the routing logic of the application. It maps HTTP requests to specific handler functions based on the request method (GET, POST, PUT, DELETE) and the URL. This makes the application's endpoints clear and easy to manage. It also allows developers to easily add new routes or modify existing ones.

2. `Handler`: Handlers are functions that are executed when a specific route is hit. They contain the business logic to handle the request and send a response. Handlers interact with services to perform operations and send the results back to the client. By separating handlers from the routing logic, we can keep our code clean and easy to read.

3. `Service`: Services contain the core business logic of the application. They interact with repositories to perform CRUD operations on the database. By separating the business logic into services, we can ensure that each function has a single responsibility, making the code easier to test and maintain.

4. `Repository`: Repositories are responsible for interacting with the database. They contain the SQL queries or ORM operations to perform CRUD operations. By isolating database operations in repositories, we can change the database or the ORM without affecting the business logic.

This pattern is beneficial as it promotes code modularity and separation of concerns. Each part of the application has a specific role and interacts with the other parts through well-defined interfaces. This makes the code easier to understand, test, and maintain. It also allows for better scalability as new features or changes can be implemented with minimal impact on the existing code.

In addition to the Router-Handler-Service-Repository pattern, this project also utilizes Dependency Injection for easier management of dependencies. This approach simplifies the process of adding or modifying dependencies, making the code more flexible and maintainable.

The `config` folder is used for managing configurations such as database connections, logging settings, and automatic migrations with Gorm. Centralizing configuration settings in this manner promotes consistency and ease of access across the application.

The `model` folder is used to store primary data structures that interact with the database and handle data parsing. This separation aids in maintaining a clean and organized codebase.

Lastly, the `entity` folder contains additional struct files that help reduce code redundancy. By reusing these entities across different parts of the application, we can ensure consistency and reduce the potential for errors.

# Frontend
## How to run
Clone the project to your local directory
```bash
git clone https://github.com/aliirsyaadn/neobank-technical-test
```

Change directory to the project folder
```bash
cd neobank-technical-test
cd frontend
```

Install dependency 
```bash
npm i
```

Run server 
```bash
npm run dev
```