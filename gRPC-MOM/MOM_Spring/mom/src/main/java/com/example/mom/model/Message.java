package com.example.mom.model;

import java.util.Map;

public class Message {
    private String id;
    private String content;
    private String destination; // Ej: "servicioA", "servicioB"

    private String status; // Ej: "ENVIADO", "FALLIDO"

    public Message(String id, String content, String destination) {
        this.id = id;
        this.content = content;
        this.destination = destination;
        this.status = "ENVIADO"; // Estado inicial
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getContent() {
        return content;
    }

    @GrpcClient("servicioA")
private ServicioAGrpc.ServicioABlockingStub stubA;

@GrpcClient("servicioB")
private ServicioBGrpc.ServicioBBlockingStub stubB;

private final Map<String, Object> stubs = Map.of(
    "servicioA", stubA,
    "servicioB", stubB
);

}    