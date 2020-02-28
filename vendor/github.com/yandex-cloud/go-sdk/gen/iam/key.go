// Code generated by sdkgen. DO NOT EDIT.

//nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// KeyServiceClient is a iam.KeyServiceClient with
// lazy GRPC connection initialization.
type KeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements iam.KeyServiceClient
func (c *KeyServiceClient) Create(ctx context.Context, in *iam.CreateKeyRequest, opts ...grpc.CallOption) (*iam.CreateKeyResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements iam.KeyServiceClient
func (c *KeyServiceClient) Delete(ctx context.Context, in *iam.DeleteKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements iam.KeyServiceClient
func (c *KeyServiceClient) Get(ctx context.Context, in *iam.GetKeyRequest, opts ...grpc.CallOption) (*iam.Key, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.KeyServiceClient
func (c *KeyServiceClient) List(ctx context.Context, in *iam.ListKeysRequest, opts ...grpc.CallOption) (*iam.ListKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).List(ctx, in, opts...)
}

type KeyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *KeyServiceClient
	request *iam.ListKeysRequest

	items []*iam.Key
}

func (c *KeyServiceClient) KeyIterator(ctx context.Context, serviceAccountId string, opts ...grpc.CallOption) *KeyIterator {
	return &KeyIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &iam.ListKeysRequest{
			ServiceAccountId: serviceAccountId,
			PageSize:         1000,
		},
	}
}

func (it *KeyIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Keys
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *KeyIterator) Value() *iam.Key {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *KeyIterator) Error() error {
	return it.err
}

// ListOperations implements iam.KeyServiceClient
func (c *KeyServiceClient) ListOperations(ctx context.Context, in *iam.ListKeyOperationsRequest, opts ...grpc.CallOption) (*iam.ListKeyOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).ListOperations(ctx, in, opts...)
}

type KeyOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *KeyServiceClient
	request *iam.ListKeyOperationsRequest

	items []*operation.Operation
}

func (c *KeyServiceClient) KeyOperationsIterator(ctx context.Context, keyId string, opts ...grpc.CallOption) *KeyOperationsIterator {
	return &KeyOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &iam.ListKeyOperationsRequest{
			KeyId:    keyId,
			PageSize: 1000,
		},
	}
}

func (it *KeyOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *KeyOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *KeyOperationsIterator) Error() error {
	return it.err
}

// Update implements iam.KeyServiceClient
func (c *KeyServiceClient) Update(ctx context.Context, in *iam.UpdateKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Update(ctx, in, opts...)
}
