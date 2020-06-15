package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	pbtime "google.golang.org/protobuf/types/known/timestamppb"

	"od_go"
)

func main() {
	addr := "localhost:9000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := od_go.NewOutliersClient(conn)
	req := &od_go.OutliersRequest{
		Metrics: dummyData(),
	}

	resp, err := client.Detect(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("outliers at: %v", resp.Indices)
}

func dummyData() []*od_go.Metric {
	const size = 1000
	out := make([]*od_go.Metric, size)
	t := time.Date(2020, 6, 15, 16, 20, 15, 0, time.UTC)
	for i := 0, i < size; i++ {
		m := od_go.Metric{
			Time: Timestamp(t),
			Name: "CPU",
			Value: rand.Float64() * 40
		}
		out[i] = &m
		t.Add(time.Second)
	}
	out[7].Value = 97.3
	out[113].Value = 92.1
	out[835].Value = 93.2
	return out
}

func Timestamp(t time.Time) *pbtime.Timestamp {
	return &pbtime.Timestamp{
		Seconds: t.Unix(),
		Nanos: int32(t.Nanosecond()),
	}
}