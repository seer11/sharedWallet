package command

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seer11/sharedWallet/config"
	"github.com/seer11/sharedWallet/pkg/blockchain"
	"github.com/spf13/cobra"
)

// NewDeployCommand creates the deploy command
func NewDeployCommand(ctx context.Context) *cobra.Command {
	deployCommand := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy contract to blockchain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deploy(ctx)
		},
	}
	return deployCommand
}

func deploy(ctx context.Context) error {
	log.Println("deploying contract")
	ctx, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()
	client, err := ethclient.DialContext(ctx, config.App.Blockchain.Address)
	if err != nil {
		return err
	}
	deployer := blockchain.NewDeployer()
	err = deployer.Deploy(ctx, client)
	if err != nil {
		return err
	}
	log.Printf("contract deployed at address %s\n", deployer.ContractAddress())
	return nil
}