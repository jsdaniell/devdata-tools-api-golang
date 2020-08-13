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

package scheduler

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCloudSchedulerClientHook clientHook

// CloudSchedulerCallOptions contains the retry settings for each method of CloudSchedulerClient.
type CloudSchedulerCallOptions struct {
	ListJobs  []gax.CallOption
	GetJob    []gax.CallOption
	CreateJob []gax.CallOption
	UpdateJob []gax.CallOption
	DeleteJob []gax.CallOption
	PauseJob  []gax.CallOption
	ResumeJob []gax.CallOption
	RunJob    []gax.CallOption
}

func defaultCloudSchedulerClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("cloudscheduler.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudSchedulerCallOptions() *CloudSchedulerCallOptions {
	return &CloudSchedulerCallOptions{
		ListJobs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateJob: []gax.CallOption{},
		UpdateJob: []gax.CallOption{},
		DeleteJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PauseJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ResumeJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RunJob: []gax.CallOption{},
	}
}

// CloudSchedulerClient is a client for interacting with Cloud Scheduler API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type CloudSchedulerClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	cloudSchedulerClient schedulerpb.CloudSchedulerClient

	// The call options for this service.
	CallOptions *CloudSchedulerCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCloudSchedulerClient creates a new cloud scheduler client.
//
// The Cloud Scheduler API allows external entities to reliably
// schedule asynchronous jobs.
func NewCloudSchedulerClient(ctx context.Context, opts ...option.ClientOption) (*CloudSchedulerClient, error) {
	clientOpts := defaultCloudSchedulerClientOptions()

	if newCloudSchedulerClientHook != nil {
		hookOpts, err := newCloudSchedulerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &CloudSchedulerClient{
		connPool:    connPool,
		CallOptions: defaultCloudSchedulerCallOptions(),

		cloudSchedulerClient: schedulerpb.NewCloudSchedulerClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CloudSchedulerClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudSchedulerClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudSchedulerClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListJobs lists jobs.
func (c *CloudSchedulerClient) ListJobs(ctx context.Context, req *schedulerpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListJobs[0:len(c.CallOptions.ListJobs):len(c.CallOptions.ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*schedulerpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*schedulerpb.Job, string, error) {
		var resp *schedulerpb.ListJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudSchedulerClient.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Jobs, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// GetJob gets a job.
func (c *CloudSchedulerClient) GetJob(ctx context.Context, req *schedulerpb.GetJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetJob[0:len(c.CallOptions.GetJob):len(c.CallOptions.GetJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateJob creates a job.
func (c *CloudSchedulerClient) CreateJob(ctx context.Context, req *schedulerpb.CreateJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateJob[0:len(c.CallOptions.CreateJob):len(c.CallOptions.CreateJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.CreateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateJob updates a job.
//
// If successful, the updated Job is returned. If the job does
// not exist, NOT_FOUND is returned.
//
// If UpdateJob does not successfully return, it is possible for the
// job to be in an Job.State.UPDATE_FAILED state. A job in this state may
// not be executed. If this happens, retry the UpdateJob request
// until a successful response is received.
func (c *CloudSchedulerClient) UpdateJob(ctx context.Context, req *schedulerpb.UpdateJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job.name", url.QueryEscape(req.GetJob().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateJob[0:len(c.CallOptions.UpdateJob):len(c.CallOptions.UpdateJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteJob deletes a job.
func (c *CloudSchedulerClient) DeleteJob(ctx context.Context, req *schedulerpb.DeleteJobRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteJob[0:len(c.CallOptions.DeleteJob):len(c.CallOptions.DeleteJob)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.cloudSchedulerClient.DeleteJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// PauseJob pauses a job.
//
// If a job is paused then the system will stop executing the job
// until it is re-enabled via ResumeJob. The
// state of the job is stored in state; if paused it
// will be set to Job.State.PAUSED. A job must be in Job.State.ENABLED
// to be paused.
func (c *CloudSchedulerClient) PauseJob(ctx context.Context, req *schedulerpb.PauseJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.PauseJob[0:len(c.CallOptions.PauseJob):len(c.CallOptions.PauseJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.PauseJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ResumeJob resume a job.
//
// This method reenables a job after it has been Job.State.PAUSED. The
// state of a job is stored in Job.state; after calling this method it
// will be set to Job.State.ENABLED. A job must be in
// Job.State.PAUSED to be resumed.
func (c *CloudSchedulerClient) ResumeJob(ctx context.Context, req *schedulerpb.ResumeJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ResumeJob[0:len(c.CallOptions.ResumeJob):len(c.CallOptions.ResumeJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.ResumeJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RunJob forces a job to run now.
//
// When this method is called, Cloud Scheduler will dispatch the job, even
// if the job is already running.
func (c *CloudSchedulerClient) RunJob(ctx context.Context, req *schedulerpb.RunJobRequest, opts ...gax.CallOption) (*schedulerpb.Job, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.RunJob[0:len(c.CallOptions.RunJob):len(c.CallOptions.RunJob)], opts...)
	var resp *schedulerpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudSchedulerClient.RunJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// JobIterator manages a stream of *schedulerpb.Job.
type JobIterator struct {
	items    []*schedulerpb.Job
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*schedulerpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*schedulerpb.Job, error) {
	var item *schedulerpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
