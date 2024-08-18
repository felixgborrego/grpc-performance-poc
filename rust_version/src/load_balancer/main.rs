use tonic::transport::Server;
use crate::entity_service::MyEntityService;
use crate::model::entity_service_server::EntityServiceServer;

pub mod model {
    tonic::include_proto!("model");
}

mod entity_service;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Initialize the logger
    env_logger::init();

    // Define the address for the gRPC server to listen on
    let addr = "[::1]:50051".parse()?;
    println!("EntityServiceServer listening on {}", addr);

    // Create an instance of MyEntityService
    let entity_service = MyEntityService::new();

    // Start the gRPC server using the Tokio runtime
    Server::builder()
        .add_service(EntityServiceServer::new(entity_service))
        .serve(addr)
        .await?;

    Ok(())
}
