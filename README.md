# 🏦 Bank Assessment API

A RESTful API for basic banking operations built with **Go (Echo framework)**, **PostgreSQL**, and **Docker**. This project supports user registration, account management, and basic banking operations like deposits and withdrawals.

## 📦 Features

- User registration and retrieval
- Account top-up and withdrawal
- Saldo (balance) check
- Healthcheck endpoint
- Clean architecture with repository pattern
- Dockerized for easy setup

---

## 🚀 Getting Started

### 🔧 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) Go installed locally if you want to run/test without Docker

---

### 🐳 Run with Docker Compose

```bash
# Clone the repo
git clone https://github.com/rmrachmanfauzan/bank_assessment.git
cd bank_assessment

# (Optional) Create a .env file
cp .env.example .env

# Build and run the containers
docker-compose build --build-arg HOST=127.0.0.1 --build-arg PORT=9090
# run image
docker-compose up -d

