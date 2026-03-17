# auth-gateway
================

## Description
------------

Auth Gateway is a robust and scalable authentication gateway that provides a secure and convenient way to authenticate users for various applications. It supports multiple authentication protocols, including OAuth, JWT, and Basic Auth, and allows for easy integration with existing authentication systems.

## Features
----------

### Core Features

* **Multi-protocol support**: Supports OAuth, JWT, and Basic Auth protocols for authentication
* **Scalable architecture**: Designed to handle high traffic and large user bases
* **Flexible configuration**: Allows for customization of authentication settings and workflows
* **User management**: Provides a user database for storing and managing user information
* **API documentation**: Includes a comprehensive API documentation for easy integration

### Additional Features

* **User registration and login**: Supports user registration and login functionality
* **Password reset and account recovery**: Allows users to reset their passwords and recover their accounts
* **Authentication tokens**: Issues authentication tokens for secure and valid authentication
* **Rate limiting**: Implements rate limiting to prevent brute-force attacks
* **Logging and analytics**: Includes logging and analytics for monitoring and troubleshooting

## Technologies Used
-----------------

* **Language**: Java 11
* **Framework**: Spring Boot
* **Database**: PostgreSQL
* **Authentication**: OAuth 2.0, JWT, Basic Auth
* **Testing framework**: JUnit
* **Code quality**: Plovl and Jacoco

## Installation
------------

### Prerequisites

* Java 11
* Maven
* PostgreSQL

### Installation Steps

1. Clone the repository: `git clone https://github.com/auth-gateway/auth-gateway.git`
2. Create a new PostgreSQL database: `create database auth-gateway`
3. Create a new application.properties file and configure the database connection settings
4. Build the project: `mvn clean package`
5. Run the application: `java -jar target/auth-gateway.jar`
6. Access the API documentation at `http://localhost:8080/swagger-ui`

### Running the Application

To run the application, execute the following command:

```bash
java -jar target/auth-gateway.jar
```

### Using the API

To use the API, you can access the API documentation at `http://localhost:8080/swagger-ui`. You can generate authentication tokens using the `POST /auth/token` endpoint and use them to access protected resources.

### Contributing
--------------

Contributions are welcome! If you'd like to contribute to the project, please fork the repository and submit a pull request. All contributions should be made to the `dev` branch.

### Issues
--------

If you encounter any issues or have questions, please submit an issue to the repository. I'll do my best to respond in a timely manner.

### License
--------

The auth-gateway project is licensed under the MIT License. Please see the LICENSE file for details.