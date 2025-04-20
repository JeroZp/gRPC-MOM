package com.javatechie.service;

import com.javatechie.Entity.Stock;
import com.javatechie.Repository.StockRepository;
import com.javatechie.grpc.*;
import io.grpc.stub.StreamObserver;
import org.springframework.grpc.server.service.GrpcService;

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

}
