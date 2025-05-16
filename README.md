# Welcome to Crypto Trading App 👋

## Get started

1. Install dependencies

   ```bash
   npm install
   ```

2. Create a .env file in backend/ and add the following according to you:
   - SERVER_PORT=8081
   - DB_HOST=localhost
   - DB_NAME=postgres
   - DB_USER=postgres
   - DB_PASSWORD=postgres
   - DB_SSLMODE=disable

2. Start the app

   - Open Docker

   - In backend/ :

   ```bash
   make start
   ```

   - In frontend/ : 

   ```bash
   npx expo start
   ```

4. Create an admin account (admin@gmail.com) and add some NFTs

5. Create a collector account and buy some NFTs

6. Open the app on your phone and scan the QR code that is provided from expo (on the terminal)
7. Log in as admin and scan the QR code for the NFT of the collector



## This project uses: 
1. Fiber - go get github.com/gofiber/fiber/v2
2. Golang Air for how reloading - https://github.com/air-verse/air/blob/master/air_example.toml
3. For env 
   - go get github.com/caarlos0/env 
   - go get github.com/joho/godotenv
4. Gorm for database actions - gorm.io
5. Decimal use to handle precision loss - go get github.com/shopspring/decimal
6. Frontend - https://docs.expo.dev/
