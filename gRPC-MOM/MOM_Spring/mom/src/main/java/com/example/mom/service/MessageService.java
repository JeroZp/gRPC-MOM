package com.example.mom.service;

import com.example.mom.MessageRepository;
import com.example.mom.model.Message;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.grpc.StatusRuntimeException;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import user.*;
import java.util.*;

@Service
public class MessageService {

    private final MessageRepository messageRepository;

    @GrpcClient("user-service") // nombre del microservicio gRPC (definido en application.yml)
    private UserServiceGrpc.UserServiceBlockingStub userStub;

    public MessageService(MessageRepository messageRepository) {
        this.messageRepository = messageRepository;
    }

    public void enqueue(Message message) {
        try {
            processGrpc(message);
        } catch (Exception e) {
            messageRepository.save(message);
            System.out.println("Fallo gRPC. Mensaje guardado para reintento.");
        }
    }

    @Scheduled(fixedRate = 10000)
    public void reintentarEntregas() {
        List<Message> pendientes = messageRepository.findAll();
        for (Message msg : pendientes) {
            try {
                processGrpc(msg);
                messageRepository.deleteById(msg.getId());
                System.out.println("Mensaje reenviado y eliminado: " + msg.getMessageId());
            } catch (Exception e) {
                System.out.println("Fallo reenvío gRPC. Se mantiene en BD.");
            }
        }
    }

    private void processGrpc(Message message) throws Exception {
        Map<String, String> payload = new ObjectMapper().readValue(
                message.getPayloadJson(), new TypeReference<>() {});

        switch (message.getOperation()) {
            case "CreateUser":
                CreateUserRequest createUserReq = CreateUserRequest.newBuilder()
                        .setUser(UserOuterClass.User.newBuilder()
                                .setId(payload.get("id"))
                                .setName(payload.get("name"))
                                .setEmail(payload.get("email"))
                                .setCredits(Integer.parseInt(payload.get("credits")))
                                .build())
                        .build();
                userStub.createUser(createUserReq);
                break;

            case "GetUser":
                GetUserRequest getUserReq = GetUserRequest.newBuilder()
                        .setId(payload.get("id"))
                        .build();
                userStub.getUser(getUserReq);
                break;

            case "DeleteUser":
                DeleteUserRequest deleteUserReq = DeleteUserRequest.newBuilder()
                        .setId(payload.get("id"))
                        .build();
                userStub.deleteUser(deleteUserReq);
                break;

            case "UpdateUser":
                UpdateUserRequest updateUserReq = UpdateUserRequest.newBuilder()
                        .setUser(UserOuterClass.User.newBuilder()
                                .setId(payload.get("id"))
                                .setName(payload.get("name"))
                                .setEmail(payload.get("email"))
                                .setCredits(Integer.parseInt(payload.get("credits")))
                                .build())
                        .build();
                userStub.updateUser(updateUserReq);
                break;

            default:
                System.out.println("Operación desconocida: " + message.getOperation());
                throw new IllegalArgumentException("Operación no soportada");
        }

        System.out.println("Operación gRPC ejecutada: " + message.getOperation());
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

    public List<Message> getAllPendingMessages() {
        return messageRepository.findAll();
    }
}
