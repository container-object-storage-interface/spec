package fake

import (
	"context"
	"github.com/container-object-storage-interface/spec/lib/go/cosi"
	"google.golang.org/grpc"
)

// this ensures that the mock implements the client interface
var _ cosi.CosiControllerClient = (*MockCosiControllerClient)(nil)

// MockBucketPolicyClient is a type that implements all the methods for RolePolicyAttachmentClient interface
type MockCosiControllerClient struct {
	MockCreateBucket    func(ctx context.Context, in *cosi.CreateBucketRequest, opts ...grpc.CallOption) (*cosi.CreateBucketResponse, error)
	MockDeleteBucket    func(ctx context.Context, in *cosi.DeleteBucketRequest, opts ...grpc.CallOption) (*cosi.DeleteBucketResponse, error)
	MockGrantBucketAccess func(ctx context.Context, in *cosi.GrantBucketAccessRequest, opts ...grpc.CallOption) (*cosi.GrantBucketAccessResponse, error)
	MockRevokeBucketAccess func(ctx context.Context, in *cosi.RevokeBucketAccessRequest, opts ...grpc.CallOption) (*cosi.RevokeBucketAccessResponse, error)
}

// GetBucketPolicyRequest mocks GetBucketPolicyRequest method
func (m *MockCosiControllerClient) CreateBucket(ctx context.Context, in *cosi.CreateBucketRequest, opts ...grpc.CallOption) (*cosi.CreateBucketResponse, error) {
	return m.CreateBucket(ctx, in, opts...)
}

// PutBucketPolicyRequest mocks PutBucketPolicyRequest method
func (m *MockCosiControllerClient) DeleteBucket(ctx context.Context, in *cosi.DeleteBucketRequest, opts ...grpc.CallOption) (*cosi.DeleteBucketResponse, error) {
	return m.DeleteBucket(ctx, in, opts...)
}

// DeleteBucketPolicyRequest mocks DeleteBucketPolicyRequest method
func (m *MockCosiControllerClient) GrantBucketAccess(ctx context.Context, in *cosi.GrantBucketAccessRequest, opts ...grpc.CallOption) (*cosi.GrantBucketAccessResponse, error) {
	return m.GrantBucketAccess(ctx, in, opts...)
}

// DeleteBucketPolicyRequest mocks DeleteBucketPolicyRequest method
func (m *MockCosiControllerClient) RevokeBucketAccess(ctx context.Context, in *cosi.RevokeBucketAccessRequest, opts ...grpc.CallOption) (*cosi.RevokeBucketAccessResponse, error) {
	return m.RevokeBucketAccess(ctx, in, opts...)
}