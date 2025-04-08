package com.microservice.gateway.microservice_gateway.filters;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

@Component
public class LoggingFilter implements GatewayFilter {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        // Log de información de la solicitud
        System.out.println("Request Path: " + exchange.getRequest().getPath());
        System.out.println("Request Method: " + exchange.getRequest().getMethod());

        // Puedes agregar más detalles si lo necesitas

        return chain.filter(exchange).then(Mono.fromRunnable(() -> {
            // Log de la respuesta
            System.out.println("Response Status Code: " + exchange.getResponse().getStatusCode());
        }));
    }
}
