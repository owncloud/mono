package migrate

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	storageprovider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// FilesMetaData representation in the import data
type FilesMetaData struct {
	Type        string `json:"type"`
	Path        string `json:"path"`
	Etag        string `json:"eTag"`
	Permissions int    `json:"permissions"`
	MTime       int    `json:"mtime"`
}

//ImportMetadata from a files.jsonl file in exportPath. The files must already be present on the storage
//Will set etag and mtime
func ImportMetadata(ctx context.Context, client gateway.GatewayAPIClient, ns string, fileData FilesMetaData) error {
	m := make(map[string]string)
	if fileData.Etag != "" {
		// TODO sanitize etag? eg double quotes at beginning and end?
		m["etag"] = fileData.Etag
	}
	if fileData.MTime != 0 {
		m["mtime"] = strconv.Itoa(fileData.MTime)
	}
	//TODO permissions? is done via share? but this is owner permissions

	if len(m) > 0 {
		resourcePath := path.Join(ns, strings.TrimPrefix(fileData.Path, "/files/"))
		samReq := &storageprovider.SetArbitraryMetadataRequest{
			Ref: &storageprovider.Reference{
				Spec: &storageprovider.Reference_Path{Path: resourcePath},
			},
			ArbitraryMetadata: &storageprovider.ArbitraryMetadata{
				Metadata: m,
			},
		}
		samResp, err := client.SetArbitraryMetadata(ctx, samReq)
		if err != nil {
			return err
		}

		if samResp.Status.Code == rpc.Code_CODE_NOT_FOUND {
			log.Print("File does not exist on target system, skipping metadata import: " + resourcePath)
		}
		if samResp.Status.Code != rpc.Code_CODE_OK {
			log.Print("Error importing metadata, skipping metadata import: " + resourcePath + ", " + samResp.Status.Message)
		}
	} else {
		log.Print("no etag or mtime for : " + fileData.Path)
	}

	return nil
}

// ForEachFile is a callback iterator for files.jsonl created by owncloud data_exporter app
func ForEachFile(path string, fn func(metaData *FilesMetaData)) {
	filesJSONL, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer filesJSONL.Close()
	jsonLines := bufio.NewScanner(filesJSONL)
	for jsonLines.Scan() {
		var f FilesMetaData
		if err := json.Unmarshal(jsonLines.Bytes(), &f); err != nil {
			log.Fatal(err)
		}

		fn(&f)
	}
}
