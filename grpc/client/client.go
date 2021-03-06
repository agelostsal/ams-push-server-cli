package grpcclient

import (
	"context"
	"fmt"

	"github.com/ARGOeu/ams-push-server-cli/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// grpcClient is a wrapper of the two existing clients
type grpcClient struct {
	hsc grpc_health_v1.HealthClient
	psc ams.PushServiceClient
}

// grpcClientStatus holds the outcome of a grpc request
type grpcClientStatus struct {
	err     error
	message string
}

// Result prints the result of an grpc request
func (st *grpcClientStatus) Result() string {

	grpcStatus := status.Convert(st.err)

	if grpcStatus.Code() == codes.OK {
		return fmt.Sprintf("Success: %v", st.message)
	}

	return fmt.Sprintf("Error: %v", grpcStatus.Message())
}

// New instantiates a grpcClient
func New(uri string) *grpcClient {

	cl := new(grpcClient)

	conn, err := grpc.Dial(uri, grpc.WithInsecure())

	if err != nil {
		fmt.Printf("Error while connecting to %v, %v", uri, err.Error())
		return cl
	}

	cl.hsc = grpc_health_v1.NewHealthClient(conn)
	cl.psc = ams.NewPushServiceClient(conn)

	return cl
}

// HealthCheck is a wrapper over the grpc health Check call
func (cl *grpcClient) HealthCheck() *grpcClientStatus {

	r, err := cl.hsc.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{
		Service: "api.v1.grpc.PushService"},
	)

	return &grpcClientStatus{
		err:     err,
		message: r.GetStatus().String(),
	}
}

// ActivateSubscription is a wrapper over the grpc ActivateSubscription call
func (cl *grpcClient) ActivateSubscription(fullName, fullTopic, pushEndpoint, retryType string, retryPeriod uint32) *grpcClientStatus {

	actSubR := &ams.ActivateSubscriptionRequest{
		Subscription: &ams.Subscription{
			FullName:  fullName,
			FullTopic: fullTopic,
			PushConfig: &ams.PushConfig{
				PushEndpoint: pushEndpoint,
				RetryPolicy: &ams.RetryPolicy{
					Type:   retryType,
					Period: retryPeriod,
				},
			},
		}}

	r, err := cl.psc.ActivateSubscription(context.Background(), actSubR)

	return &grpcClientStatus{
		err:     err,
		message: r.GetMessage(),
	}
}

// DeactivateSubscription is a wrapper over the grpc DeactivateSubscription call
func (cl *grpcClient) DeactivateSubscription(fullName string) *grpcClientStatus {

	deActSubR := &ams.DeactivateSubscriptionRequest{
		FullName: fullName,
	}

	r, err := cl.psc.DeactivateSubscription(context.Background(), deActSubR)

	return &grpcClientStatus{
		err:     err,
		message: r.GetMessage(),
	}
}

// GetSubscription is a wrapper over the grpc GetSubscription call
func (cl *grpcClient) GetSubscription(fullName string) *grpcClientStatus {

	getSubR := &ams.GetSubscriptionRequest{
		FullName: fullName,
	}

	r, err := cl.psc.GetSubscription(context.Background(), getSubR)

	return &grpcClientStatus{
		err:     err,
		message: r.String(),
	}
}

// ListSubscriptions is a wrapper over the grpc ListSubscriptions call
func (cl *grpcClient) ListSubscriptions() *grpcClientStatus {

	r, err := cl.psc.ListSubscriptions(context.Background(), &ams.ListSubscriptionsRequest{})

	return &grpcClientStatus{
		err:     err,
		message: fmt.Sprintf("%v", r.Subscriptions),
	}
}
