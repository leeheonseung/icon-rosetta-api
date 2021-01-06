// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// A NetworkAPIController binds http requests to an api service and writes the service results to
// the http response
type NetworkAPIController struct {
	service  NetworkAPIServicer
	asserter *asserter.Asserter
}

// NewNetworkAPIController creates a default api controller
func NewNetworkAPIController(
	s NetworkAPIServicer,
	asserter *asserter.Asserter,
) Router {
	return &NetworkAPIController{
		service:  s,
		asserter: asserter,
	}
}

// Routes returns all of the api route for the NetworkAPIController
func (c *NetworkAPIController) Routes() Routes {
	return Routes{
		{
			"NetworkList",
			strings.ToUpper("Post"),
			"/network/list",
			c.NetworkList,
		},
		{
			"NetworkOptions",
			strings.ToUpper("Post"),
			"/network/options",
			c.NetworkOptions,
		},
		{
			"NetworkStatus",
			strings.ToUpper("Post"),
			"/network/status",
			c.NetworkStatus,
		},
	}
}

// NetworkList - Get List of Available Networks
func (c *NetworkAPIController) NetworkList(w http.ResponseWriter, r *http.Request) {
	metadataRequest := &types.MetadataRequest{}
	if err := json.NewDecoder(r.Body).Decode(&metadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that MetadataRequest is correct
	if err := c.asserter.MetadataRequest(metadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.NetworkList(r.Context(), metadataRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// NetworkOptions - Get Network Options
func (c *NetworkAPIController) NetworkOptions(w http.ResponseWriter, r *http.Request) {
	networkRequest := &types.NetworkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&networkRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that NetworkRequest is correct
	if err := c.asserter.NetworkRequest(networkRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.NetworkOptions(r.Context(), networkRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// NetworkStatus - Get Network Status
func (c *NetworkAPIController) NetworkStatus(w http.ResponseWriter, r *http.Request) {
	networkRequest := &types.NetworkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&networkRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that NetworkRequest is correct
	if err := c.asserter.NetworkRequest(networkRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.NetworkStatus(r.Context(), networkRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
