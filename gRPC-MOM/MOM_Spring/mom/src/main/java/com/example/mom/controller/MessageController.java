package com.example.mom.controller;

import com.example.mom.model.Message;
import com.example.mom.service.MessageService;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/messages")
public class MessageController {

    private final MessageService messageService;

    public MessageController(MessageService messageService) {
        this.messageService = messageService;
    }

    @PostMapping("/enqueue")
    public String enqueue(@RequestBody Message message) {
        messageService.enqueue(message);
        return "Mensaje encolado";
    }

    @GetMapping("/dequeue")
    public Message dequeue() {
        return messageService.dequeue();
    }

    @GetMapping("/status")
    public String status() {
        return "Mensajes en cola: " + messageService.getQueueSize();
    }
}
