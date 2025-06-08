# Solana Token Faucet with Go Backend
A web-based Solana Devnet faucet for requesting 1 SOL on Devnet.

## Live Demo
[Try it here](https://solana-faucet-go.vercel.app/)

## Features
- Connects to Phantom Wallet on Solana Devnet.
- Requests 1 SOL via a Go backend.
- Built with HTML, CSS, JavaScript, Go, Fiber, and solana-go-sdk.

## Setup
### Prerequisites
- [Go](https://go.dev/dl/) (v1.24+)
- [Node.js](https://nodejs.org/) (for `http-server`)
- [Phantom Wallet](https://phantom.app/) (set to Devnet)

### Backend
1. Clone the repo:
 
   `git clone https://github.com/nitishmalang/solana-faucet-go.git`
   `cd solana-faucet-go`

Install dependencies:
   `go mod tidy`

Run backend:

   `go run main.go`

Runs on `http://localhost:8082`
  Verify:

`curl http://localhost:8082`

Expected: {"message":"Welcome to the Solana Token Faucet Backend!"}

## Front-End

Navigate to front-end:

`cd frontend`

Install http-server:

`npm install -g http-server`

Run front-end:

`http-server . -p 8081`

Open 

`http://localhost:8081`

Testing

Open `http://localhost:8081` or `https://solana-faucet-go.vercel.app/`.

Connect Phantom Wallet (Devnet).

Request 1 SOL. Check Phantom for +1 SOL.

If airdrop fails (429 error), use a new Devnet wallet or QuickNode RPC:
Update main.go:
go

solanaClient := client.NewClient("https://your-quicknode-devnet-endpoint")

