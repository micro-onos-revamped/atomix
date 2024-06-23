// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	rsmapiv1 "github.com/micro-onos-revamped/atomix/protocols/rsm/api/v1"
	"github.com/micro-onos-revamped/atomix/runtime/pkg/driver"
	"github.com/micro-onos-revamped/atomix/runtime/pkg/network"
)

func New(network network.Driver) driver.Driver {
	return &sharedMemoryDriver{
		network: network,
	}
}

type sharedMemoryDriver struct {
	network network.Driver
}

func (d *sharedMemoryDriver) Connect(ctx context.Context, config *rsmapiv1.ProtocolConfig) (driver.Conn, error) {
	conn := newConn(d.network)
	if err := conn.Connect(ctx, config); err != nil {
		return nil, err
	}
	return conn, nil
}
