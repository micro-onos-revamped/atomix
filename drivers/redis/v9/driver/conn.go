// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"github.com/go-redis/redis/v9"
	setv1 "github.com/micro-onos-revamped/atomix/api/runtime/set/v1"
	redissetv1 "github.com/micro-onos-revamped/atomix/drivers/redis/v9/driver/set/v1"
	"github.com/micro-onos-revamped/atomix/runtime/pkg/driver"
)

func newConn(client *redis.Client) driver.Conn {
	return &redisConn{
		client: client,
	}
}

type redisConn struct {
	client *redis.Client
}

func (c *redisConn) NewSetV1() setv1.SetServer {
	return redissetv1.NewSet(c.client)
}

func (c *redisConn) Close(ctx context.Context) error {
	return c.client.Close()
}
