package com.javatechie.service;

import com.javatechie.Entity.Stock;
import com.javatechie.Repository.StockRepository;
import com.javatechie.grpc.*;
import io.grpc.stub.StreamObserver;
import org.springframework.grpc.server.service.GrpcService;

import java.time.LocalDateTime;

@GrpcService
public class StockTradingServiceImpl extends StockTradingServiceGrpc.StockTradingServiceImplBase{
    private final StockRepository stockRepository;

    public StockTradingServiceImpl(StockRepository stockRepository) {
        this.stockRepository = stockRepository;
    }

    @Override
    public void getStockPrice(StockRequest request, StreamObserver<StockResponse> responseObserver) {
        //StockName -> DB -> Map Response -> Return
        String stockSymbol = request.getStockSymbol();
        Stock stock = stockRepository.findByStockSymbol(stockSymbol);
        StockResponse stockResponse = StockResponse.newBuilder()
                .setStockSymbol(stock.getStockSymbol())
                .setPrice(stock.getPrice())
                .setTimestamp(stock.getLastUpdate().toString())
                .build();

        responseObserver.onNext(stockResponse);
        responseObserver.onCompleted();

    }

    @Override
    public void updateStockPrice(UpdateStockRequest request, StreamObserver<UpdateStockResponse> responseObserver) {
        String symbol = request.getStockSymbol();
        double newPrice = request.getNewPrice();
        Stock stock = stockRepository.findByStockSymbol(symbol);
        // verifica si existe
        if (stock != null) {
            stock.setPrice(newPrice);  // setter
            stockRepository.save(stock); // actualizar

            UpdateStockResponse response = UpdateStockResponse.newBuilder()
                    .setSuccess(true)
                    .setMessage("Precio actualizado correctamente")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } else {
            UpdateStockResponse response = UpdateStockResponse.newBuilder()
                    .setSuccess(false)
                    .setMessage("Símbolo no encontrado")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }

    @Override
    public void createStock(CreateStockRequest request, StreamObserver<CreateStockResponse> responseObserver) {
        try {
            // Verificar si el stock ya existe
            Stock existing = stockRepository.findByStockSymbol(request.getStockSymbol());
            if (existing != null) {
                CreateStockResponse response = CreateStockResponse.newBuilder()
                        .setSuccess(false)
                        .setMessage("Stock con ese símbolo ya existe.")
                        .build();
                responseObserver.onNext(response);
                responseObserver.onCompleted();
                return;
            }

            // Crear nuevo stock
            Stock stock = new Stock();
            stock.setStockSymbol(request.getStockSymbol());
            stock.setPrice(request.getInitialPrice());
            stock.setLastUpdate(LocalDateTime.parse(LocalDateTime.now().toString()));
            stockRepository.save(stock);

            CreateStockResponse response = CreateStockResponse.newBuilder()
                    .setSuccess(true)
                    .setMessage("Stock creado exitosamente.")
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            e.printStackTrace();
            CreateStockResponse response = CreateStockResponse.newBuilder()
                    .setSuccess(false)
                    .setMessage("Error al crear el stock: " + e.getMessage())
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }



}
