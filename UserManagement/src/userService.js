const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

// Cargar el archivo proto
const PROTO_PATH = process.env.PROTO_PATH;
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const userProto = grpc.loadPackageDefinition(packageDefinition).UserManagement;

// Simulación de una base de datos en memoria para efectos de demostración
let users = [];

// Función para crear un nuevo usuario
function createUser(call, callback) {
    const { name, email } = call.request;
    const userId = users.length + 1; // Generar un ID único
    const newUser = { id: userId, name, email };
    users.push(newUser);
    callback(null, { id: userId, name, email });
}

// Función para obtener un usuario por ID
function getUser(call, callback) {
    const userId = call.request.id;
    const user = users.find(u => u.id === userId);
    if (user) {
        callback(null, user);
    } else {
        callback({
            code: grpc.status.NOT_FOUND,
            details: 'User not found'
        });
    }
}

// Función para actualizar un usuario por ID
function updateUser(call, callback) {
    const { userId, name, email } = call.request;
    const userIndex = users.find(u => u.userId === userId);
    if (user) {
        user.name = name;
        user.email = email;
        callback(null, user);
    } else {
        callback({
            code: grpc.status.NOT_FOUND,
            details: 'User not found'
        });
    }
}

// Función para eliminar un usuario por ID
function deleteUser(call, callback) {
    const { userId } = call.request;
    const userIndex = users.findIndex(u => u.userId === userId);
    if (userIndex !== -1) {
        users.splice(userIndex, 1);
        callback(null, { message: 'User ${userId} deleted successfully' });
    } else {
        callback({
            code: grpc.status.NOT_FOUND,
            details: 'User not found'
        });
    }
}

// Crear el servidor gRPC
const server = new grpc.Server();
server.addService(userProto.UserService.service, {
    createUser,
    getUser,
    updateUser,
    deleteUser
});

// Iniciar el servidor en el puerto especificado
const PORT = process.env.GRPC_PORT || 50051;
// Aquí se puede iniciar el servidor