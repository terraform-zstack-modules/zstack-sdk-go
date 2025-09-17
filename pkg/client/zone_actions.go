// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryZone queries the list of zones
func (cli *ZSClient) QueryZone(params param.QueryParam) ([]view.ZoneView, error) {
	resp := make([]view.ZoneView, 0)
	return resp, cli.List("v1/zones", &params, &resp)
}

// GetZone by uuid
func (cli *ZSClient) GetZone(uuid string) (*view.ZoneView, error) {
	var resp view.ZoneView
	return &resp, cli.Get("v1/zones/", uuid, nil, &resp)
}
