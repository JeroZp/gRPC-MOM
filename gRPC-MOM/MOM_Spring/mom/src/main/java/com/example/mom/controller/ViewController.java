package com.example.mom.controller;

import com.example.mom.model.Message;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import com.example.mom.service.MessageService;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.List;

@Controller
public class ViewController {

    private final MessageService messageService; // Suponiendo que ya tienes un servicio que maneja los mensajes

    public ViewController(MessageService messageService) {
        this.messageService = messageService;
    }

    @GetMapping("/view/pending")
    public String showPendingMessages(Model model) {
        List<Message> pendingMessages = messageService.getAllPendingMessages();
        model.addAttribute("pendingMessages", pendingMessages);
        return "pending";
    }

}
