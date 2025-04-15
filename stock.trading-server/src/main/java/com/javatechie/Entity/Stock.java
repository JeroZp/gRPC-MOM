package com.javatechie.Entity;

import jakarta.persistence.*;

import java.time.LocalDateTime;

@Table(name = "stocks")
@Entity
public class Stock {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    public String getStockSymbol() {
        return stockSymbol;
    }

    public void setStockSymbol(String stockSymbol) {
        this.stockSymbol = stockSymbol;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }

    public LocalDateTime getLastUpdate() {
        return lastUpdate;
    }

    public void setLastUpdate(LocalDateTime lastUpdate) {
        this.lastUpdate = lastUpdate;
    }

    @Column(name = "stock_symbol", unique = true, nullable = false)
    private String stockSymbol;

    private double price;

    @Column(name = "Last_updated")
    private LocalDateTime lastUpdate;

}
