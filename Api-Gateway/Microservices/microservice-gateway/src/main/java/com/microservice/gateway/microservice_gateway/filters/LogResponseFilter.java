package com.microservice.gateway.microservice_gateway.filters;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;


public class LogResponseFilter implements GatewayFilter {
    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        return chain.filter(exchange).then(Mono.fromRunnable(() -> {
            int statusCode = exchange.getResponse().getStatusCode().value();
            System.out.println("Response status: " + statusCode);
        }));
    }
}
