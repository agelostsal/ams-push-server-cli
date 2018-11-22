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

	r3, e3 :=  cl2.DeactivateSubscription(context.Background(), &amsPb.DeactivateSubscriptionRequest{FullName:"projects/p1/subscription/sub1"})

	if e3 != nil {
		fmt.Print(err)
	}

	w3,w4 := status.FromError(e3)

	fmt.Printf("\n %v \n %v", w3.Code(),w4)

	fmt.Printf("%+v\n", r3.String())

	r4, e4 :=  cl2.DeactivateSubscription(context.Background(), &amsPb.DeactivateSubscriptionRequest{FullName:"projects/p1/subscription/sub1"})

	if e4 != nil {
		fmt.Print(err)
	}

	w5,_:= status.FromError(e4)

	fmt.Printf("\n %v \n %v", w5.Code(),w5.Message())
	fmt.Printf("%+v\n", r4.String())
}
