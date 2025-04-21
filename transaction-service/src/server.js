// src/server.js
import path from "path";
import { fileURLToPath } from "url";
import grpc from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";

//
// Calcular __dirname en ESM
//
const __filename = fileURLToPath(import.meta.url);
const __dirname  = path.dirname(__filename);

//
// Opciones para proto-loader
//
const loaderOpts = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

//
// Rutas a los .proto
//
const USER_PROTO_PATH        = path.resolve(__dirname, "../proto/user.proto");
const TRANSACTION_PROTO_PATH = path.resolve(__dirname, "../proto/transaction.proto");

//
// Cargar definiciones
//
const userPackageDef = protoLoader.loadSync(USER_PROTO_PATH, loaderOpts);
const txPackageDef   = protoLoader.loadSync(TRANSACTION_PROTO_PATH, loaderOpts);

const userProto = grpc.loadPackageDefinition(userPackageDef).user;
const txProto   = grpc.loadPackageDefinition(txPackageDef).transaction;

//
// Cliente gRPC para UserService
//
const userClient = new userProto.UserService(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

// Implementa la lógica de transferencia

function transfer(call, callback) {
    const { from_id, to_id, amount } = call.request;
  
    // Obtener remitente
    userClient.GetUser({ id: from_id }, (err, resp) => {
      if (err) {
        return callback(null, {
          success: false,
          message: `Error al obtener usuario remitente: ${err.message}`,
        });
      }

      const fromUser = resp.user;
  
      // Verificar saldo suficiente
      if (fromUser.credits < amount) {
        return callback(null, {
          success: false,
          message: "Saldo insuficiente",
        });
      }
  
      // Obtener destinatario
      userClient.GetUser({ id: to_id }, (err2, resp2) => {
        if (err2) {
          return callback(null, {
            success: false,
            message: `Error al obtener usuario destino: ${err2.message}`,
          });
        }
        const toUser = resp2.user;
  
        // Ajustar balances
        fromUser.credits -= amount;
        toUser.credits   += amount;
  
        // Actualizar remitente
        userClient.UpdateUser({ user: fromUser }, err3 => {
          if (err3) {
            return callback(null, {
              success: false,
              message: `Error al actualizar remitente: ${err3.message}`,
            });
          }
  
          // Actualizar destinatario
          userClient.UpdateUser({ user: toUser }, err4 => {
            if (err4) {
              return callback(null, {
                success: false,
                message: `Error al actualizar destinatario: ${err4.message}`,
              });
            }
  
            // Responder éxito
            callback(null, {
              success: true,
              message: `Transferidos ${amount} créditos`,
              from_user: fromUser,
              to_user:   toUser,
            });
          });
        });
      });
    });
}  

// Montar y arrancar el servidor gRPC

function main() {
  const server = new grpc.Server();
  server.addService(txProto.TransactionService.service, { Transfer: transfer });

  server.bindAsync(
    "0.0.0.0:50052",
    grpc.ServerCredentials.createInsecure(),
    (err, port) => {
      if (err) {
        console.error("Error al bindear TransactionService:", err);
        process.exit(1);
      }
      console.log(`TransactionService escuchando en :${port}`);
    }
  );
}

main();