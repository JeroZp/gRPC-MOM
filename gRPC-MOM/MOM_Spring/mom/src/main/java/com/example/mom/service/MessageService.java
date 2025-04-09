package com.example.mom.service;

import org.springframework.stereotype.Service;

import com.example.mom.model.Message;

import java.util.LinkedList;
import java.util.Queue;

@Service
public class MessageService {
    private final Queue<Message> queue = new LinkedList<>();

    public void enqueue(Message message) {
        queue.add(message);
    }

    public Message dequeue() {
        return queue.poll(); // null si está vacía
    }

    public int getQueueSize() {
        return queue.size();
    }
}
