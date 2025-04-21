// src/server.js
import path from "path";
import { fileURLToPath } from "url";
import grpc from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";

//
// 1) Calcular __dirname en ESM
//
const __filename = fileURLToPath(import.meta.url);
const __dirname  = path.dirname(__filename);

//
// 2) Opciones comunes para proto-loader
//
const loaderOpts = {
  keepCase:  true,
  longs:     String,
  enums:     String,
  defaults:  true,
  oneofs:    true,
};

//
// 3) Rutas a tus .proto
//
const USER_PROTO_PATH        = path.resolve(__dirname, "../proto/user.proto");
const TRANSACTION_PROTO_PATH = path.resolve(__dirname, "../proto/transaction.proto");
const MOM_PROTO_PATH         = path.resolve(__dirname, "../proto/mom.proto");

//
// 4) Cargar defs
//
const userPkgDef = protoLoader.loadSync(USER_PROTO_PATH, loaderOpts);
const txPkgDef   = protoLoader.loadSync(TRANSACTION_PROTO_PATH, loaderOpts);
const momPkgDef  = protoLoader.loadSync(MOM_PROTO_PATH, loaderOpts);

const userProto = grpc.loadPackageDefinition(userPkgDef).user;
const txProto   = grpc.loadPackageDefinition(txPkgDef).transaction;
const momProto  = grpc.loadPackageDefinition(momPkgDef).mom;

//
// 5) Clientes gRPC
//
const userClient = new userProto.UserService(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

let txClient;   // lo inicializamos tras arrancar el server
const momClient = new momProto.MOMService(
  "localhost:50054",
  grpc.credentials.createInsecure()
);

//
// 6) Implementación de la lógica de Transferencia
//
function transfer(call, callback) {
  const { from_id, to_id, amount } = call.request;

  // 6.1) Obtener remitente
  userClient.GetUser({ id: from_id }, (err, resp) => {
    if (err) return callback(null, { success: false, message: err.message });
    const fromUser = resp.user;

    // 6.2) Verificar créditos
    if (fromUser.credits < amount) {
      return callback(null, { success: false, message: "Saldo insuficiente" });
    }

    // 6.3) Obtener destinatario
    userClient.GetUser({ id: to_id }, (err2, resp2) => {
      if (err2) return callback(null, { success: false, message: err2.message });
      const toUser = resp2.user;

      // 6.4) Ajustar balances
      fromUser.credits -= amount;
      toUser.credits   += amount;

      // 6.5) Actualizar remitente
      userClient.UpdateUser({ user: fromUser }, err3 => {
        if (err3) return callback(null, { success: false, message: err3.message });

        // 6.6) Actualizar destinatario
        userClient.UpdateUser({ user: toUser }, err4 => {
          if (err4) return callback(null, { success: false, message: err4.message });

          // 6.7) Responder éxito
          callback(null, {
            success: true,
            message: `Transferidos ${amount} créditos`,
            from_user: fromUser,
            to_user:   toUser
          });
        });
      });
    });
  });
}

//
// 7) Worker que consume de MOM ("tx_ops")
//    y repite las transferencias vía gRPC sobre este mismo servicio
//
function startWorker() {
  const call = momClient.Subscribe();
  call.write({ queue: "tx_ops" });

  call.on("data", msg => {
    const { from_id, to_id, amount } = JSON.parse(msg.data.toString());

    // 7.1) Llamada gRPC local a Transfer
    txClient.Transfer({ from_id, to_id, amount }, (err, res) => {
      if (err) {
        console.error("Error al procesar Transfer vía MOM:", err);
      } else {
        console.log("Transfer vía MOM OK:", res);
        // 7.2) ACK
        call.write({ queue: "tx_ops", id: msg.id });
      }
    });
  });

  call.on("error", err => {
    console.error("Error en suscripción MOM:", err);
  });
}

//
// 8) Arranque del servidor y del worker
//
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
      console.log(`TransactionService gRPC escuchando en :${port}`);

      // Inicializar cliente local para Transfer vía gRPC 
      txClient = new txProto.TransactionService(
        `localhost:${port}`,
        grpc.credentials.createInsecure()
      );

      // Arrancar el worker MOM una vez el server está listo
      startWorker();
    }
  );
}

main();