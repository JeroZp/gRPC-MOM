package com.example.mom.controller;

import com.example.mom.model.Message;
import com.example.mom.service.MessageService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/messages")
public class MessageController {

    private final MessageService messageService;

    public MessageController(MessageService messageService) {
        this.messageService = messageService;
    }

    @PostMapping("/enqueue")
    public ResponseEntity<String> enqueue(@RequestBody Message message) {
        try {
            messageService.enqueue(message);
            return ResponseEntity.ok("Mensaje procesado o encolado");
        } catch (Exception e) {
            return ResponseEntity.status(500).body("Error al procesar el mensaje");
        }
    }

    @GetMapping("/dequeue")
    public ResponseEntity<Message> dequeue() {
        Message mensaje = messageService.dequeue();
        if (mensaje != null) {
            return ResponseEntity.ok(mensaje);
        } else {
            return ResponseEntity.noContent().build();
        }
    }

    @GetMapping("/status")
    public String status() {
        return "Mensajes en cola: " + messageService.getQueueSize();
    }

    @GetMapping("/pendientes")
    public List<Message> pendientes() {
        return messageService.getAllPendingMessages();
    }

    @PostMapping("/retry")
    public ResponseEntity<String> reintentarMensajes() {
        messageService.reintentarEntregas();
        return ResponseEntity.ok("Reintento manual ejecutado");
    }
}

