package com.example.mom.service;

import com.example.mom.model.Message;
import jakarta.annotation.PostConstruct;
import org.springframework.http.ResponseEntity;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import java.util.Iterator;
import java.util.LinkedList;
import java.util.Queue;

@Service
public class MessageService {

    private final Queue<Message> queue = new LinkedList<>();
    private final RestTemplate restTemplate = new RestTemplate();
    

    // URL del microservicio destino
    private final String destinoUrl = "http://localhost:8081/api/procesar"; // Cámbialo a tu microservicio real

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
            queue.add(message);
            System.out.println("Fallo entrega. Mensaje encolado para " + message.getDestination());
        }
    }
    

    public Message dequeue() {
        return queue.poll();
    }

    public int getQueueSize() {
        return queue.size();
    }

    // Reintento automático cada 10 segundos
    @Scheduled(fixedRate = 10000)
    public void reintentarEntregas() {
        System.out.println("Revisando cola para reintento...");
        Iterator<Message> iterator = queue.iterator();
        while (iterator.hasNext()) {
            Message msg = iterator.next();
            String url = destinoMap.get(msg.getDestination());
    
            if (url == null) {
                System.out.println("Destino desconocido: " + msg.getDestination());
                iterator.remove(); // o mantenlo si prefieres
                continue;
            }
    
            try {
                restTemplate.postForEntity(url, msg, Void.class);
                iterator.remove();
                System.out.println("Mensaje reenviado exitosamente a " + msg.getDestination());
            } catch (RestClientException e) {
                System.out.println("Fallo reenvío a " + msg.getDestination() + ". Se mantiene en cola.");
            }
        }
    }
}
