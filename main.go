package main

import (
    "context"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/blocto/solana-go-sdk/client"
    "github.com/blocto/solana-go-sdk/rpc"
    "github.com/mr-tron/base58"
)

type AirdropRequest struct {
    WalletAddress string `json:"wallet_address"`
}

type AirdropResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    TxID    string `json:"tx_id,omitempty"`
}

func main() {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST",
        AllowHeaders: "Content-Type",
    }))

    solanaClient := client.NewClient(rpc.DevnetRPCEndpoint)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Welcome to the Solana Token Faucet Backend!"})
    })

    
    app.Post("/airdrop", func(c *fiber.Ctx) error {
       
        req := new(AirdropRequest)
        if err := c.BodyParser(req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(AirdropResponse{
                Status:  "error",
                Message: "Invalid request body: " + err.Error(),
            })
        }

       
        if req.WalletAddress == "" {
            return c.Status(fiber.StatusBadRequest).JSON(AirdropResponse{
                Status:  "error",
                Message: "Wallet address is empty",
            })
        }

        _, err := base58.Decode(req.WalletAddress)
        if err != nil || len(req.WalletAddress) < 32 || len(req.WalletAddress) > 44 {
            return c.Status(fiber.StatusBadRequest).JSON(AirdropResponse{
                Status:  "error",
                Message: "Invalid wallet address format",
            })
        }


        txID, err := solanaClient.RequestAirdrop(
            context.Background(),
            req.WalletAddress,
            1_000_000_000,
        )
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(AirdropResponse{
                Status:  "error",
                Message: fmt.Sprintf("Failed to request airdrop: %v", err),
            })
        }

        return c.JSON(AirdropResponse{
            Status:  "success",
            Message: "1 SOL successfully sent!",
            TxID:    txID,
        })
    })

    if err := app.Listen(":8082"); err != nil {
        fmt.Printf("Server startup failed: %v\n", err.Error())
    }
}