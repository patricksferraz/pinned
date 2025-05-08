# 🍽️ Pinned - Restaurant Management System

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/pinned)](https://goreportcard.com/report/github.com/yourusername/pinned)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

A modern, microservices-based restaurant management system built with Go, designed to streamline restaurant operations and enhance customer experience.

## 🌟 Features

- **Menu Management**: Comprehensive menu management system with categories, items, and pricing
- **Guest Management**: Track and manage customer information and preferences
- **Order Processing**: Real-time order tracking and management
- **Attendant System**: Staff management and task assignment
- **Place Management**: Table and seating arrangement management
- **Guest Check System**: Integrated billing and payment processing

## 🏗️ Architecture

The system is built using a microservices architecture with the following components:

- **Menu Service**: Handles menu items, categories, and pricing
- **Guest Service**: Manages customer information and preferences
- **Attendant Service**: Handles staff management and task assignment
- **Place Service**: Manages table and seating arrangements
- **Guest Check Service**: Processes orders and payments

### Technology Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Message Broker**: Apache Kafka
- **Containerization**: Docker
- **Orchestration**: Kubernetes (optional)
- **API Documentation**: Swagger/OpenAPI

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.18 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pinned.git
   cd pinned
   ```

2. Set up environment variables:
   ```bash
   cp menu/.env.example menu/.env
   # Repeat for other services as needed
   ```

3. Start the services:
   ```bash
   docker-compose up -d
   ```

### Development Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run tests:
   ```bash
   make test
   ```

3. Start development server:
   ```bash
   make run
   ```

## 📚 API Documentation

Each service exposes its own REST API. API documentation is available at:

- Menu Service: `http://localhost:8080/swagger`
- Guest Service: `http://localhost:8081/swagger`
- Attendant Service: `http://localhost:8082/swagger`
- Place Service: `http://localhost:8083/swagger`
- Guest Check Service: `http://localhost:8084/swagger`

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Patrick Ferraz - Initial work - [GitHub](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- Thanks to all contributors who have helped shape this project
- Inspired by modern restaurant management needs
- Built with best practices in mind for scalability and maintainability

## 📞 Support

For support, please open an issue in the GitHub repository or contact the maintainers.

---

⭐ Star this repository if you find it useful!
