package migrate

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"path"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	collaboration "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// ShareMetaData representation in the import metadata
type ShareMetaData struct {
	Path           string `json:"path"`
	ShareType      string `json:"shareType"`
	Type           string `json:"type"`
	Owner          string `json:"owner"`
	SharedBy       string `json:"sharedBy"`
	SharedWith     string `json:"sharedWith"`
	Permissions    int    `json:"permissions"`
	ExpirationDate string `json:"expirationDate"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	Token          string `json:"token"`
}

//ImportShare from a shares.jsonl file in exportPath. The files must already be present on the storage
func ImportShare(ctx context.Context, client gateway.GatewayAPIClient, ns string, shareData *ShareMetaData) error {
	//Stat file, skip ShareMetaData creation if it does not exist on the target system
	resourcePath := path.Join(ns, shareData.Path)
	statReq := &provider.StatRequest{
		Ref: &provider.Reference{
			Spec: &provider.Reference_Path{Path: resourcePath},
		},
	}
	statResp, err := client.Stat(ctx, statReq)

	if err != nil {
		return err
	}

	if statResp.Status.Code == rpc.Code_CODE_NOT_FOUND {
		log.Print("File does not exist on target system, skipping share import: " + resourcePath)
		return nil
	}

	_, err = client.CreateShare(ctx, shareReq(statResp.Info, shareData))
	if err != nil {
		return err
	}

	return nil
}

func shareReq(info *provider.ResourceInfo, share *ShareMetaData) *collaboration.CreateShareRequest {
	return &collaboration.CreateShareRequest{
		ResourceInfo: info,
		Grant: &collaboration.ShareGrant{
			Grantee: &provider.Grantee{
				Type: provider.GranteeType_GRANTEE_TYPE_USER,
				Id: &user.UserId{
					OpaqueId: share.SharedWith,
				},
			},
			Permissions: &collaboration.SharePermissions{
				Permissions: convertPermissions(share.Permissions),
			},
		},
	}
}

// Maps oc10 permissions to roles
var ocPermToRole = map[int]string{
	1:  "viewer",
	15: "co-owner",
	31: "editor",
}

// Create resource permission-set from ownCloud permissions int
func convertPermissions(ocPermissions int) *provider.ResourcePermissions {
	perms := &provider.ResourcePermissions{}
	switch ocPermToRole[ocPermissions] {
	case "viewer":
		perms.Stat = true
		perms.ListContainer = true
		perms.InitiateFileDownload = true
		perms.ListGrants = true
	case "editor":
		perms.Stat = true
		perms.ListContainer = true
		perms.InitiateFileDownload = true

		perms.CreateContainer = true
		perms.InitiateFileUpload = true
		perms.Delete = true
		perms.Move = true
		perms.ListGrants = true
	case "co-owner":
		perms.Stat = true
		perms.ListContainer = true
		perms.InitiateFileDownload = true

		perms.CreateContainer = true
		perms.InitiateFileUpload = true
		perms.Delete = true
		perms.Move = true

		perms.ListGrants = true
		perms.AddGrant = true
		perms.RemoveGrant = true
		perms.UpdateGrant = true
	}

	return perms
}

//ForEachShare is a callback iterator for shares.jsonl created by owncloud data_exporter app
func ForEachShare(path string, fn func(metaData *ShareMetaData)) {
	sharesJSONL, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer sharesJSONL.Close()
	jsonLines := bufio.NewScanner(sharesJSONL)
	for jsonLines.Scan() {
		var f ShareMetaData
		if err := json.Unmarshal(jsonLines.Bytes(), &f); err != nil {
			log.Fatal(err)
		}

		fn(&f)
	}
}
