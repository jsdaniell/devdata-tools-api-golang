// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package errorreporting

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newReportErrorsClientHook clientHook

// ReportErrorsCallOptions contains the retry settings for each method of ReportErrorsClient.
type ReportErrorsCallOptions struct {
	ReportErrorEvent []gax.CallOption
}

func defaultReportErrorsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("clouderrorreporting.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultReportErrorsCallOptions() *ReportErrorsCallOptions {
	return &ReportErrorsCallOptions{
		ReportErrorEvent: []gax.CallOption{},
	}
}

// ReportErrorsClient is a client for interacting with Stackdriver Error Reporting API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ReportErrorsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	reportErrorsClient clouderrorreportingpb.ReportErrorsServiceClient

	// The call options for this service.
	CallOptions *ReportErrorsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewReportErrorsClient creates a new report errors service client.
//
// An API for reporting error events.
func NewReportErrorsClient(ctx context.Context, opts ...option.ClientOption) (*ReportErrorsClient, error) {
	clientOpts := defaultReportErrorsClientOptions()

	if newReportErrorsClientHook != nil {
		hookOpts, err := newReportErrorsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ReportErrorsClient{
		connPool:    connPool,
		CallOptions: defaultReportErrorsCallOptions(),

		reportErrorsClient: clouderrorreportingpb.NewReportErrorsServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ReportErrorsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ReportErrorsClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ReportErrorsClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ReportErrorEvent report an individual error event.
//
// This endpoint accepts either an OAuth token,
// or an API key (at https://support.google.com/cloud/answer/6158862)
// for authentication. To use an API key, append it to the URL as the value of
// a key parameter. For example:
//
// POST https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456
func (c *ReportErrorsClient) ReportErrorEvent(ctx context.Context, req *clouderrorreportingpb.ReportErrorEventRequest, opts ...gax.CallOption) (*clouderrorreportingpb.ReportErrorEventResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ReportErrorEvent[0:len(c.CallOptions.ReportErrorEvent):len(c.CallOptions.ReportErrorEvent)], opts...)
	var resp *clouderrorreportingpb.ReportErrorEventResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.reportErrorsClient.ReportErrorEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
