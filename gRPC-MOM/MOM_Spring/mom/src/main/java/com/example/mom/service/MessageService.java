package com.example.mom.service;

import com.example.mom.MessageRepository;
import com.example.mom.model.Message;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import java.util.*;

@Service
public class MessageService {

    private final RestTemplate restTemplate = new RestTemplate();

    private final MessageRepository messageRepository;

    // Mapa de destinos dinámico
    private final Map<String, String> destinoMap = Map.of(
            "servicioA", "http://localhost:8081/api/procesar",
            "servicioB", "http://localhost:8082/api/procesar"
    );

    public MessageService(MessageRepository messageRepository) {
        this.messageRepository = messageRepository;
    }

    public void enqueue(Message message) {
        String url = destinoMap.get(message.getDestination());

        if (url == null) {
            System.out.println("Destino desconocido: " + message.getDestination());
            return;
        }

        try {
            restTemplate.postForEntity(url, message, Void.class);
            System.out.println("Mensaje entregado a " + message.getDestination());
        } catch (RestClientException e) {
            messageRepository.save(message);
            System.out.println("Fallo entrega. Mensaje guardado para " + message.getDestination());
        }
    }

    @Scheduled(fixedRate = 10000)
    public void reintentarEntregas() {
        List<Message> pendientes = messageRepository.findAll();
        for (Message msg : pendientes) {
            String url = destinoMap.get(msg.getDestination());

            if (url == null) {
                System.out.println("Destino desconocido: " + msg.getDestination());
                messageRepository.deleteById(msg.getId());
                continue;
            }

            try {
                restTemplate.postForEntity(url, msg, Void.class);
                messageRepository.deleteById(msg.getId());
                System.out.println("Mensaje reenviado y eliminado: " + msg.getMessageId());
            } catch (RestClientException e) {
                System.out.println("Fallo reenvío a: " + msg.getDestination());
            }
        }
    }

    public int getQueueSize() {
        return messageRepository.findAll().size();
    }

    public Message dequeue() {
        List<Message> all = messageRepository.findAll();
        if (!all.isEmpty()) {
            Message m = all.get(0);
            messageRepository.delete(m);
            return m;
        }
        return null;
    }

    public List<Message>  getAllPendingMessages(){
        return  messageRepository.findAll();
    }


}
