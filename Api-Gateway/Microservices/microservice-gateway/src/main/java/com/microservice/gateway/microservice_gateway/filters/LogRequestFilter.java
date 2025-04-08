package com.microservice.gateway.microservice_gateway.filters;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;


public class LogRequestFilter implements GatewayFilter {
    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        String method = String.valueOf(exchange.getRequest().getMethod());
        String uri = exchange.getRequest().getURI().toString();
        String headers = exchange.getRequest().getHeaders().toString();
        System.out.println("Incoming request: " + method + " " + uri);
        System.out.println("Request headers: " + headers);

        // Continúa con el siguiente filtro en la cadena
        return chain.filter(exchange);
    }
}
