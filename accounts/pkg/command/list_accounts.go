package command

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/client/grpc"
	tw "github.com/olekukonko/tablewriter"
	"github.com/owncloud/mono/accounts/pkg/config"
	"github.com/owncloud/mono/accounts/pkg/flagset"
	accounts "github.com/owncloud/mono/accounts/pkg/proto/v0"
	"os"
	"strconv"
)

// ListAccounts command lists all accounts
func ListAccounts(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:    "list",
		Usage:   "List existing accounts",
		Aliases: []string{"ls"},
		Flags:   flagset.ListAccountsWithConfig(cfg),
		Action: func(c *cli.Context) error {
			accSvcID := cfg.GRPC.Namespace + "." + cfg.Server.Name
			accSvc := accounts.NewAccountsService(accSvcID, grpc.NewClient())
			resp, err := accSvc.ListAccounts(c.Context, &accounts.ListAccountsRequest{})

			if err != nil {
				fmt.Println(fmt.Errorf("could not list accounts %w", err))
				return err
			}

			buildAccountsListTable(resp.Accounts).Render()
			return nil
		}}
}

// buildAccountsListTable creates an ascii table for printing on the cli
func buildAccountsListTable(accs []*accounts.Account) *tw.Table {
	table := tw.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "DisplayName", "Mail", "AccountEnabled"})
	table.SetAutoFormatHeaders(false)
	for _, acc := range accs {
		table.Append([]string{
			acc.Id,
			acc.DisplayName,
			acc.Mail,
			strconv.FormatBool(acc.AccountEnabled)})
	}
	return table
}
