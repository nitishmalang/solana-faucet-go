# Solana Token Faucet with Go Backend
A web-based Solana Devnet faucet for requesting 1 SOL, built by Nitish Malang.

## Live Demo
https://solana-faucet-go.vercel.app 
## Features
- Connect Phantom wallet on Solana Devnet.
- Request 1 SOL via Go backend.

## Setup
1. Clone: `git clone https://github.com/nitishmalang/solana-faucet-go.git`
2. Backend:
   ```bash
   cd solana-faucet-go
   go mod tidy
   go run main.go

## Front-end:


`cd frontend`
`npm install -g http-server`
`http-server . -p 8081`

Open http://localhost:8081, connect Phantom (Devnet).

## Deployed Backend
Hosted on Render: https://solana-faucet-go.onrender.com



   
