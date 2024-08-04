package main

import (
	"context"
	"grpc-service-poc/internal/generated/model"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func BenchmarkCreateEntityWithAttributes(b *testing.B) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := model.NewEntityServiceClient(conn)

	startTime := time.Now()

	for n := 0; n < b.N; n++ {
		attributeIDs := make([]uint64, 1000)

		// Create 1000 attributes
		for i := 0; i < 1000; i++ {
			createAttributeResponse, err := client.CreateAttribute(context.Background(), &model.CreateAttributeRequest{
				Name: "Benchmark Attribute",
			})
			if err != nil {
				b.Fatalf("CreateAttribute failed: %v", err)
			}
			attributeIDs[i] = createAttributeResponse.AttributeId
		}

		// Create an entity
		createEntityResponse, err := client.CreateEntity(context.Background(), &model.CreateEntityRequest{
			Name: "Benchmark Entity",
		})
		if err != nil {
			b.Fatalf("CreateEntity failed: %v", err)
		}

		// Add all attributes to the entity
		_, err = client.AddAttributeToEntity(context.Background(), &model.AddAttributeToEntityRequest{
			EntityId:     createEntityResponse.EntityId,
			AttributeIds: attributeIDs,
		})
		if err != nil {
			b.Fatalf("AddAttributeToEntity failed: %v", err)
		}
	}

	elapsed := time.Since(startTime)
	throughput := float64(b.N) / elapsed.Seconds()

	log.Printf("Benchmark completed: %d iterations in %s", b.N, elapsed)
	log.Printf("Throughput: %.2f operations/second", throughput)
}
