use model::entity_service_client::EntityServiceClient;
use model::{CreateEntityRequest, CreateAttributeRequest, AddAttributeToEntityRequest};

pub mod model {
    tonic::include_proto!("model");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = EntityServiceClient::connect("http://[::1]:50051").await?;

    // Test CreateEntity
    let request = tonic::Request::new(CreateEntityRequest {
        name: "Test Entity".to_string(),

    });

    let response = client.create_entity(request).await?;

    println!("RESPONSE={:?}", response);

    // Test CreateAttribute
    let request = tonic::Request::new(CreateAttributeRequest {
        name: "Test Attribute".to_string(),
        // Fill in request fields
    });

    let response = client.create_attribute(request).await?;

    println!("RESPONSE={:?}", response);

    // Test AddAttributeToEntity
    let request = tonic::Request::new(AddAttributeToEntityRequest {
        entity_id:  1,
        attribute_ids: vec![1],
        // Fill in request fields
    });

    let response = client.add_attribute_to_entity(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}
