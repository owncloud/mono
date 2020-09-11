package flagset

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis-reva/pkg/config"
)

// DriverOwnCloudWithConfig applies cfg to the root flagset
func DriverOwnCloudWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "storage-owncloud-datadir",
			Value:       "/var/tmp/reva/data",
			Usage:       "the path to the owncloud data directory",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_DATADIR"},
			Destination: &cfg.Reva.Storages.OwnCloud.Root,
		},
		&cli.StringFlag{
			Name:        "storage-owncloud-uploadinfo-dir",
			Value:       "/var/tmp/reva/uploadinfo",
			Usage:       "the path to the tus upload info directory",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_UPLOADINFO_DIR"},
			Destination: &cfg.Reva.Storages.OwnCloud.UploadInfoDir,
		},
		&cli.StringFlag{
			Name:        "storage-owncloud-share-folder",
			Value:       "/Shares",
			Usage:       "name of the shares folder",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_SHARE_FOLDER"},
			Destination: &cfg.Reva.Storages.OwnCloud.ShareFolder,
		},
		&cli.BoolFlag{
			Name:        "storage-owncloud-scan",
			Value:       true,
			Usage:       "scan files on startup to add fileids",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_SCAN"},
			Destination: &cfg.Reva.Storages.OwnCloud.Scan,
		},
		&cli.StringFlag{
			Name:        "storage-owncloud-redis",
			Value:       ":6379",
			Usage:       "the address of the redis server",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_REDIS_ADDR"},
			Destination: &cfg.Reva.Storages.OwnCloud.Redis,
		},
		&cli.BoolFlag{
			Name:        "storage-owncloud-enable-home",
			Value:       false,
			Usage:       "enable the creation of home storages",
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_ENABLE_HOME"},
			Destination: &cfg.Reva.Storages.OwnCloud.EnableHome,
		},
		&cli.StringFlag{
			Name:        "storage-owncloud-layout",
			Value:       "{{.Id.OpaqueId}}",
			Usage:       `"layout of the users home dir path on disk, in addition to {{.Username}}, {{.Mail}}, {{.Id.OpaqueId}}, {{.Id.Idp}} also supports prefixing dirs: "{{substr 0 1 .Username}}/{{.Username}}" will turn "Einstein" into "Ei/Einstein" `,
			EnvVars:     []string{"REVA_STORAGE_OWNCLOUD_LAYOUT"},
			Destination: &cfg.Reva.Storages.OwnCloud.UserLayout,
		},
	}
}
