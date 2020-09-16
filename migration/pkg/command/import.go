package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	revauser "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	"github.com/cs3org/reva/pkg/token"
	"github.com/cs3org/reva/pkg/token/manager/jwt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/v2/client/grpc"
	accounts "github.com/owncloud/ocis-accounts/pkg/proto/v0"
	"github.com/owncloud/ocis-migration/pkg/config"
	"github.com/owncloud/ocis-migration/pkg/flagset"
	"github.com/owncloud/ocis-migration/pkg/migrate"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type user struct {
	UserID      string `json:"userId"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

type exportData struct {
	Date         string `json:"date"`
	OriginServer string `json:"originServer"`
	User         user   `json:"user"`
}

// Import is the entrypoint for the import command.
func Import(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "import",
		Usage: "Import a user",
		Flags: flagset.ImportWithConfig(cfg),
		Action: func(c *cli.Context) error {
			logger := NewLogger(cfg)

			importPath := c.String("import-path")
			if importPath == "" {
				logger.Fatal().Msg("No import-path specified")
			}

			info, err := os.Stat(importPath)
			if err != nil {
				logger.Fatal().Err(err).Msg("Could not open export")
			}

			if !info.IsDir() {
				logger.Fatal().Msg("Import path must be a directory")
			}

			userMetaDataPath := path.Join(importPath, "user.json")
			data, err := ioutil.ReadFile(userMetaDataPath)

			if err != nil {
				logger.Fatal().Err(err).Msgf("Could not read file")
			}

			u := &exportData{}
			if err := json.Unmarshal(data, u); err != nil {
				logger.Fatal().Err(err).Msgf("Could not decode json")
			}

			gatewayClient, err := connectToReva()
			if err != nil {
				logger.Fatal().Err(err).Msg("Could not connect to reva")
			}

			tokenManager, err := jwt.New(map[string]interface{}{
				"secret":  c.String("jwt-secret"),
				"expires": int64(10),
			})

			if err != nil {
				logger.Fatal().Err(err).Msgf("Could not load token-manager")
			}

			user := &revauser.User{
				Id: &revauser.UserId{
					OpaqueId: u.User.UserID,
				},
				Username: u.User.UserID,
			}

			t, err := tokenManager.MintToken(c.Context, user)
			if err != nil {
				logger.Fatal().Err(err).Msgf("Error minting token")
			}

			ctx := token.ContextSetToken(context.Background(), t)
			ctx = metadata.AppendToOutgoingContext(ctx, token.TokenHeader, t)

			logger.Debug().Msg("Importing files-metadata")
			migrate.ForEachFile(path.Join(importPath, "files.jsonl"), func(metaData *migrate.FilesMetaData) {
				t, err := tokenManager.MintToken(c.Context, user)
				if err != nil {
					logger.Fatal().Err(err).Msgf("Error minting token")
				}

				logger.Debug().Interface("file", metaData).Msg("File imported")
				if err := migrate.ImportMetadata(token.ContextSetToken(ctx, t), gatewayClient, "/home", *metaData); err != nil {
					logger.Fatal().Err(err).Msg("Importing files metadata failed")
				}
			})

			logger.Debug().Msg("Importing shares-metadata")
			migrate.ForEachShare(path.Join(importPath, "shares.jsonl"), func(metaData *migrate.ShareMetaData) {
				t, err := tokenManager.MintToken(c.Context, user)
				if err != nil {
					logger.Fatal().Err(err).Msgf("Error minting token")
				}

				logger.Debug().Interface("share", metaData).Msg("Share imported")
				if err := migrate.ImportShare(token.ContextSetToken(ctx, t), gatewayClient, "/home", metaData); err != nil {
					logger.Fatal().Err(err).Msg("Importing shares metadata failed")
				}
			})

			logger.Debug().Msg("Creating entry in com.owncloud.accounts")
			ss := accounts.NewAccountsService("com.owncloud.accounts", grpc.NewClient())
			_, err = ss.CreateAccount(c.Context, &accounts.CreateAccountRequest{
				Account: &accounts.Account{
					// TODO really use the old username as the uuid? it would be unique, but only in the scope of this instance. shouldn't we be able to roll a new uuid?
					Id: u.User.UserID,
				},
			})

			if err != nil {
				logger.Fatal().Err(err).Msgf("Could not create entry in ocis-accounts")
			}

			return nil
		}}
}

func connectToReva() (gatewayv1beta1.GatewayAPIClient, error) {
	svc, err := registry.GetService("com.owncloud.reva")
	if err != nil {
		return nil, err
	}

	if len(svc) < 1 {
		return nil, fmt.Errorf("service 'com.owncloud.reva not found")
	}

	if len(svc[0].Nodes) < 1 {
		return nil, fmt.Errorf("no nodes for service 'com.owncloud.reva found")
	}

	addr := svc[0].Nodes[0].Address
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return gatewayv1beta1.NewGatewayAPIClient(conn), nil
}
