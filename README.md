# Welcome 👋

## Get started

1. Install dependencies

   ```bash
   npm install
   ```

2. Create a .env file in backend/ and add the following according to you:
   - SERVER_PORT=8081
   - API_KEY=<get your api key after registrating to Finnhub>
   - DB_HOST=localhost
   - DB_NAME=postgres
   - DB_USER=postgres
   - DB_PASSWORD=postgres
   - DB_SSLMODE=disable

3. Start the app

   - Open Docker

   - In backend/ run:

   ```bash
   make start
   ```

   - In frontend/ run: 

   ```bash
   npx expo start
   ```

4. Create an admin account (admin@gmail.com) and add some NFTs

5. Create a collector account and buy some NFTs

6. Open the app on your phone and scan the QR code that is provided from expo (on the terminal)

7. Log in as admin and scan the QR code for the NFT of the collector

##### Point of the NFT game is to "buy" the NFT and the first one to get it scanned by out admin wins the NFT.

## Stock trading section
1. Prime trading hours are from 9:30am to 4:00pm EST, which means 4:30pm to 11pm bulgarian time.


## This project uses: 
1. Fiber
2. Golang Air for how reloading - https://github.com/air-verse/air/blob/master/air_example.toml
3. For env 
   - go get github.com/caarlos0/env 
   - go get github.com/joho/godotenv
4. Gorm for database actions - gorm.io
5. Decimal use to handle precision loss - go get github.com/shopspring/decimal
6. Frontend - https://docs.expo.dev/
7. API for trades from Finnhub
8. Web socket from gorilla
