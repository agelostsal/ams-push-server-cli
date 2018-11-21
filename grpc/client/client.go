package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	amsPb "github.com/ARGOeu/ams-push-server-cli/grpc/proto"
)

func main() {

	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure())

	if err != nil {
		fmt.Print(err.Error())
	}

	defer conn.Close()

	cl := grpc_health_v1.NewHealthClient(conn)

	r, err := cl.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{Service: "api.v1.grpc.PushService"})
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Printf("%+v\n", r.GetStatus())

	retry1 := amsPb.RetryPolicy{Type: "linear", Period: 30}
	pCfg1 := amsPb.PushConfig{PushEndpoint: "https://127.0.0.1:5000/receive_here", RetryPolicy: &retry1}
	sub1 := amsPb.Subscription{FullName: "projects/p1/subscription/sub1", FullTopic: "projects/p1/topics/topic1", PushConfig: &pCfg1}

	cl2 := amsPb.NewPushServiceClient(conn)
	r2, e2 := cl2.ActivateSubscription(context.Background(), &amsPb.ActivateSubscriptionRequest{Subscription: &sub1})
	if e2 != nil {
		fmt.Print(err)
	}

	w1,w2 := status.FromError(e2)

	fmt.Printf("\n %v \n %v", w1.Code(),w2)

	fmt.Printf("%+v\n", r2.String())
}
