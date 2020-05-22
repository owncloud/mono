package svc

import (
	"context"
	"github.com/owncloud/ocis-settings/pkg/settings"
	store "github.com/owncloud/ocis-settings/pkg/store/filesystem"

	"github.com/owncloud/ocis-settings/pkg/config"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
)

type Service struct {
	config       *config.Config
	manager      settings.Manager
}

// NewService returns a service implementation for Service.
func NewService(cfg *config.Config) Service {
	return Service{
		config:  cfg,
		manager: store.New(cfg),
	}
}

// SaveSettingsBundle implements the BundleServiceHandler interface
func (g Service) SaveSettingsBundle(c context.Context, req *proto.SaveSettingsBundleRequest, res *proto.SaveSettingsBundleResponse) error {
	req.SettingsBundle.Identifier = getFailsafeIdentifier(req.SettingsBundle.Identifier)
	r, err := g.manager.WriteBundle(req.SettingsBundle)
	if err != nil {
		return err
	}
	res.SettingsBundle = r
	return nil
}

// GetSettingsBundle implements the BundleServiceHandler interface
func (g Service) GetSettingsBundle(c context.Context, req *proto.GetSettingsBundleRequest, res *proto.GetSettingsBundleResponse) error {
	r, err := g.manager.ReadBundle(getFailsafeIdentifier(req.Identifier))
	if err != nil {
		return err
	}
	res.SettingsBundle = r
	return nil
}

// ListSettingsBundles implements the BundleServiceHandler interface
func (g Service) ListSettingsBundles(c context.Context, req *proto.ListSettingsBundlesRequest, res *proto.ListSettingsBundlesResponse) error {
	r, err := g.manager.ListBundles(getFailsafeIdentifier(req.Identifier))
	if err != nil {
		return err
	}
	res.SettingsBundles = r
	return nil
}

// SaveSettingsValue implements the ValueServiceHandler interface
func (g Service) SaveSettingsValue(c context.Context, req *proto.SaveSettingsValueRequest, res *proto.SaveSettingsValueResponse) error {
	req.SettingsValue.Identifier = getFailsafeIdentifier(req.SettingsValue.Identifier)
	r, err := g.manager.WriteValue(req.SettingsValue)
	if err != nil {
		return err
	}
	res.SettingsValue = r
	return nil
}

// GetSettingsValue implements the ValueServiceHandler interface
func (g Service) GetSettingsValue(c context.Context, req *proto.GetSettingsValueRequest, res *proto.GetSettingsValueResponse) error {
	r, err := g.manager.ReadValue(getFailsafeIdentifier(req.Identifier))
	if err != nil {
		return err
	}
	res.SettingsValue = r
	return nil
}

// ListSettingsValues implements the ValueServiceHandler interface
func (g Service) ListSettingsValues(c context.Context, req *proto.ListSettingsValuesRequest, res *proto.ListSettingsValuesResponse) error {
	r, err := g.manager.ListValues(getFailsafeIdentifier(req.Identifier))
	if err != nil {
		return err
	}
	res.SettingsValues = r
	return nil
}

func getFailsafeIdentifier(identifier *proto.Identifier) *proto.Identifier {
	if identifier == nil {
		identifier = &proto.Identifier{}
	}
	if identifier.AccountUuid == "me" {
		identifier.AccountUuid = "5681371F-4A6E-43BC-8BB5-9C9237FA9C58"
	}
	return identifier
}
