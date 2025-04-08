package com.miempresa.inventoryservice;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import java.io.IOException;

import com.miempresa.inventoryservice.InventoryServiceProto.AddProductRequest;
import com.miempresa.inventoryservice.InventoryServiceProto.AddProductResponse;
import com.miempresa.inventoryservice.InventoryServiceProto.GetProductRequest;
import com.miempresa.inventoryservice.InventoryServiceProto.GetProductResponse;
import com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductRequest;
import com.miempresa.inventoryservice.InventoryServiceProto.UpdateProductResponse;
import com.miempresa.inventoryservice.InventoryServiceGrpc.InventoryServiceImplBase;

public class InventoryServiceServer {

    private Server server;
    private final int port = 50054;

    private void start() throws IOException {
        server = ServerBuilder.forPort(port)
                .addService(new InventoryServiceImpl())
                .build()
                .start();
        System.out.println("Inventory Service started, listening on " + port);

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.err.println("*** Shutting down gRPC server since JVM is shutting down");
            InventoryServiceServer.this.stop();
            System.err.println("*** Server shut down");
        }));
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        InventoryServiceServer server = new InventoryServiceServer();
        server.start();
        server.blockUntilShutdown();
    }

    static class InventoryServiceImpl extends InventoryServiceImplBase {
        @Override
        public void addProduct(AddProductRequest request, StreamObserver<AddProductResponse> responseObserver) {
            System.out.println("Adding product: " + request.getProductId());
            AddProductResponse response = AddProductResponse.newBuilder()
                    .setProductId(request.getProductId())
                    .setName(request.getName())
                    .setStatus("Product added successfully")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }

        @Override
        public void getProduct(GetProductRequest request, StreamObserver<GetProductResponse> responseObserver) {
            System.out.println("Getting product: " + request.getProductId());
            GetProductResponse response = GetProductResponse.newBuilder()
                    .setProductId(request.getProductId())
                    .setName("Sample Product")
                    .setDescription("This is a sample product")
                    .setQuantity(100)
                    .setPrice(29.99)
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }

        @Override
        public void updateProduct(UpdateProductRequest request, StreamObserver<UpdateProductResponse> responseObserver) {
            System.out.println("Updating product: " + request.getProductId());
            UpdateProductResponse response = UpdateProductResponse.newBuilder()
                    .setProductId(request.getProductId())
                    .setName(request.getName())
                    .setStatus("Product updated successfully")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }
}
