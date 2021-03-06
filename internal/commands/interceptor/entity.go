// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package interceptor

import (
	"github.com/urfave/cli"

	"github.com/apache/skywalking-cli/api"
)

func ParseEntity(ctx *cli.Context) *api.Entity {
	service := ctx.String("service")
	normal := ctx.BoolT("isNormal")
	instance := ctx.String("instance")
	endpoint := ctx.String("endpoint")

	destService := ctx.String("destService")
	destNormal := ctx.BoolT("isDestNormal")
	destInstance := ctx.String("destServiceInstance")
	destEndpoint := ctx.String("destEndpoint")

	entity := &api.Entity{
		ServiceName:             &service,
		Normal:                  &normal,
		ServiceInstanceName:     &instance,
		EndpointName:            &endpoint,
		DestServiceName:         &destService,
		DestNormal:              &destNormal,
		DestServiceInstanceName: &destInstance,
		DestEndpointName:        &destEndpoint,
	}
	entity.Scope = parseScope(entity)

	return entity
}

// parseScope defines the scope based on the input parameters.
func parseScope(entity *api.Entity) api.Scope {
	scope := api.ScopeAll

	if *entity.DestEndpointName != "" {
		scope = api.ScopeEndpointRelation
	} else if *entity.DestServiceInstanceName != "" {
		scope = api.ScopeServiceInstanceRelation
	} else if *entity.DestServiceName != "" {
		scope = api.ScopeServiceRelation
	} else if *entity.EndpointName != "" {
		scope = api.ScopeEndpoint
	} else if *entity.ServiceInstanceName != "" {
		scope = api.ScopeServiceInstance
	} else if *entity.ServiceName != "" {
		scope = api.ScopeService
	}

	return scope
}
