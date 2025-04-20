package com.javatechie.service;

import com.javatechie.Entity.Stock;
import com.javatechie.Repository.StockRepository;
import com.javatechie.grpc.StockRequest;
import com.javatechie.grpc.StockResponse;
import com.javatechie.grpc.StockTradingServiceGrpc;
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
}
