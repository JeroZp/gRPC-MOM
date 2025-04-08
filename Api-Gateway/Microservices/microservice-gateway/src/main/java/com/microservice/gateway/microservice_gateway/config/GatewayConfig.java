package com.microservice.gateway.microservice_gateway.config;
import com.microservice.gateway.microservice_gateway.filters.LogRequestFilter;
import com.microservice.gateway.microservice_gateway.filters.LogResponseFilter;
import com.microservice.gateway.microservice_gateway.filters.LoggingFilter;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GatewayConfig {

    // Registrar el filtro de solicitud
    @Bean(name = "LogRequestFilter")
    public GatewayFilter logRequestFilter() {
        return new LogRequestFilter();
    }

    // Registrar el filtro de respuesta
    @Bean(name = "LogResponseFilter")
    public GatewayFilter logResponseFilter() {
        return new LogResponseFilter();
    }

    private final LoggingFilter loggingFilter;

    public GatewayConfig(LoggingFilter loggingFilter) {
        this.loggingFilter = loggingFilter;
    }

    @Bean
    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route(r -> r.path("/**")
                        .filters(f -> f.filter(loggingFilter))  // Aplicar el filtro LoggingFilter
                        .uri("http://127.0.0.1:8000"))
                .build();
    }
}
