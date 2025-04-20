package com.example.mom.service;

import com.example.mom.repository.MessageRepository;
import com.example.mom.model.Message;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import java.util.*;
import java.util.concurrent.TimeUnit;
import com.example.mom.grpc.UserProto.CreateUserRequest;
import com.example.mom.grpc.UserProto.DeleteUserRequest;
import com.example.mom.grpc.UserProto.GetUserRequest;
import com.example.mom.grpc.UserProto.UpdateUserRequest;
import com.example.mom.grpc.UserProto.User;
import com.example.mom.grpc.UserServiceGrpc;

@Service
public class MessageService {

    private final MessageRepository messageRepository;

    private static final long BACKOFF_INITIAL_DELAY = 1000; // 1 segundo
    private static final double BACKOFF_MULTIPLIER = 2.0;
    private static final long BACKOFF_MAX_DELAY = 32000;

    public MessageService(MessageRepository messageRepository) {
        this.messageRepository = messageRepository;
    }

    @GrpcClient("user-service")
    private UserServiceGrpc.UserServiceBlockingStub userStub;

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
        if (userStub == null) {
            System.out.println("Stub de gRPC no inicializado. Mensaje quedará en cola.");
            throw new IllegalStateException("Stub no disponible");
        }

        Map<String, String> payload = new ObjectMapper().readValue(
                message.getPayloadJson(), new TypeReference<>() {});

        // Intentos de reintentos con backoff exponencial
        int attempts = 0;
        long delay = BACKOFF_INITIAL_DELAY;

        while (attempts < 5) { // Limitar a un máximo de 5 intentos
            try {
                switch (message.getOperation()) {
                    case "CreateUser":
                        CreateUserRequest createUserReq = CreateUserRequest.newBuilder()
                                .setUser(User.newBuilder()
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
                                .setUser(User.newBuilder()
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

                // Si el intento fue exitoso, salimos del bucle
                System.out.println("Operación gRPC ejecutada con éxito: " + message.getOperation());
                return;

            } catch (StatusRuntimeException e) {
                if (e.getStatus().getCode() == Status.Code.UNAVAILABLE) {
                    // Si el servicio no está disponible, reintentamos con backoff
                    attempts++;
                    System.out.println("Error de conexión. Reintentando... Intento " + attempts);
                    Thread.sleep(delay);

                    // Aplicamos el backoff exponencial
                    delay = Math.min((long) (delay * BACKOFF_MULTIPLIER), BACKOFF_MAX_DELAY);
                } else {
                    // Si el error no es por disponibilidad, lo lanzamos de nuevo
                    throw e;
                }
            }
        }

        // Si se agotaron los intentos y sigue fallando, lanzamos un error
        System.out.println("El servicio sigue sin estar disponible después de " + attempts + " intentos.");
        throw new IllegalStateException("Servicio gRPC no disponible después de varios intentos");
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
