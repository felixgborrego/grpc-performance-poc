package service

import (
	"context"
	"grpc-service-poc/internal/generated/model"
	"grpc-service-poc/internal/snowflake"
)

// EntityServiceServer implements the EntityServiceServer interface defined in the model package.
type EntityServiceServer struct {
	model.UnimplementedEntityServiceServer
}

// NewEntityServiceServer creates a new instance of EntityServiceServer.
func NewEntityServiceServer() *EntityServiceServer {
	return &EntityServiceServer{}
}

// CreateEntity handles the creation of a new entity.
func (s *EntityServiceServer) CreateEntity(ctx context.Context, req *model.CreateEntityRequest) (*model.CreateEntityResponse, error) {
	return &model.CreateEntityResponse{
		EntityId: snowflake.GenerateID(),
	}, nil
}

// CreateAttribute handles the creation of a new attribute.
func (s *EntityServiceServer) CreateAttribute(ctx context.Context, req *model.CreateAttributeRequest) (*model.CreateAttributeResponse, error) {
	return &model.CreateAttributeResponse{
		AttributeId: snowflake.GenerateID(),
	}, nil
}

// AddAttributeToEntity adds an attribute to an existing entity.
func (s *EntityServiceServer) AddAttributeToEntity(ctx context.Context, req *model.AddAttributeToEntityRequest) (*model.AddAttributeToEntityResponse, error) {
	return &model.AddAttributeToEntityResponse{}, nil
}
