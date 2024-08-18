use tonic::{Request, Response, Status};
use sonyflake::{Sonyflake, Builder};
use crate::model::entity_service_server::EntityService;
use crate::model::{
    CreateEntityRequest, CreateEntityResponse,
    CreateAttributeRequest, CreateAttributeResponse,
    AddAttributeToEntityRequest, AddAttributeToEntityResponse,
};

/// MyEntityService is responsible for handling entity-related operations.
pub struct MyEntityService {
    sonyflake_gen: Sonyflake,
}

impl MyEntityService {
    /// Creates a new instance of MyEntityService.
    pub fn new() -> Self { 
        Self {
            sonyflake_gen: Builder::new()
                .machine_id(&|| Ok(1))
                .finalize()
                .expect("Failed to create Sonyflake generator"),
        }
    }
}

#[tonic::async_trait]
impl EntityService for MyEntityService {
    /// Handles the creation of a new entity.
    async fn create_entity(
        &self,
        request: Request<CreateEntityRequest>,
    ) -> Result<Response<CreateEntityResponse>, Status> {
        let req_id = self.sonyflake_gen.next_id()
            .map_err(|e| Status::internal(format!("Sonyflake error: {}", e)))?;
        
        println!("CreateEntityRequest: {:?} with id: {}", request, req_id);

        // Respond with the generated entity ID.
        Ok(Response::new(CreateEntityResponse { entity_id: req_id }))
    }

    /// Handles the creation of a new attribute.
    async fn create_attribute(
        &self,
        _request: Request<CreateAttributeRequest>,
    ) -> Result<Response<CreateAttributeResponse>, Status> {
        // Generate attribute ID (placeholder logic).
        let attribute_id = 1;

        // Respond with the generated attribute ID.
        Ok(Response::new(CreateAttributeResponse { attribute_id }))
    }

    /// Adds an attribute to an entity.
    async fn add_attribute_to_entity(
        &self,
        _request: Request<AddAttributeToEntityRequest>,
    ) -> Result<Response<AddAttributeToEntityResponse>, Status> {
        // Logic to add the attribute to the entity (placeholder logic).

        // Respond with an empty response indicating success.
        Ok(Response::new(AddAttributeToEntityResponse {}))
    }
}
