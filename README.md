# Welcome 👋

## Description

1. This is a simple NFT marketplace game app that allows users to buy NFTs. When a NFT is purchased, the first person to find the admin of the app and have him scan the QR code wins the NFT and only this user will be able to have it. 

2. Another feature of this app is the real-time stock tracker that shows the market for one of the most popular companies including Apple, Amazon, Microsoft, Google and much more. 


## Get started

1. Create a .env file in backend/ and add the following according to you:
   ```
   SERVER_PORT=8081
   API_KEY=<get your api key after registrating to Finnhub>
   DB_HOST=localhost
   DB_NAME=postgres
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_SSLMODE=disable
   JWT_SECRET=secret
   ```

2. Install dependencies in the frontend/

   ```bash
   npm install
   ```

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
2. Open section Stock to see the charts.


## This project uses: 
1. Fiber
2. Golang Air for hot reloading - https://github.com/air-verse/air/blob/master/air_example.toml
3. Gorm for database actions - gorm.io
4. Postgres for the database
5. Frontend - https://docs.expo.dev/
6. API for trades from Finnhub
7. Web socket from gorilla
