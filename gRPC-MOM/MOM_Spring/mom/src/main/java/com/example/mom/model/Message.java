package com.example.mom.model;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.persistence.*;
import lombok.*;

import java.io.IOException;
import java.util.Map;

@Entity
@Getter @Setter
@NoArgsConstructor
@AllArgsConstructor
public class Message {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String messageId;

    private String operation;

    @Lob
    private String payloadJson; // JSON en formato String

    public Map<String, String> getPayloadAsMap() throws IOException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(this.payloadJson, new TypeReference<>() {});
    }
}
