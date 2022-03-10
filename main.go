package main

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	fmt.Println("fetch SOL total supply:", solana.SolMint)

	c := rpc.New("https://api.mainnet-beta.solana.com")

	r, err := c.GetTokenSupply(context.Background(), solana.SolMint, rpc.CommitmentFinalized)

	fmt.Println(r, err)
}
