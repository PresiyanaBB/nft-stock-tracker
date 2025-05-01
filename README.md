# crypto-price-tracker

This project uses: 
1. Fiber - go get github.com/gofiber/fiber/v2
2. Golang Air for how reloading - https://github.com/air-verse/air/blob/master/air_example.toml
3. For env 
   - go get github.com/caarlos0/env 
   - go get github.com/joho/godotenv
4. Gorm for database actions - gorm.io
5. Decimal use to handle precision loss - go get github.com/shopspring/decimal


How to run:
1. create a .env file in backend/ and add the following according to you:
   - SERVER_PORT=8081
   - DB_HOST=localhost
   - DB_NAME=postgres
   - DB_USER=postgres
   - DB_PASSWORD=postgres
   - DB_SSLMODE=disable
2. run `make start`