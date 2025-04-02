const axios = require('axios');

// API URL
const API_GATEWAY_URL = 'http://localhost:80801';

// Function to do a GET request
async function getRequest(endpoint) {
    const url = `${API_GATEWAY_URL}/${endpoint}`;
    try {
        const response = await axios.get(url);
        console.log('Response:', response.data);
    } catch (error) {
        console.error('Error making GET request:', error.message);
    }
}

// Function to do a POST request
async function postRequest(endpoint, data) {
    const url = `${API_GATEWAY_URL}/${endpoint}`;
    try {
        const response = await axios.post(url, data);
        console.log('Response:', response.data);
    } catch (error) {
        console.error('Error making POST request:', error.message);
    }
}

// Ejemplo de uso de las funciones
/*async function main() {
    // Solicitar datos con GET
    await getData("microservicio-endpoint");

    // Enviar datos con POST
    const dataToSend = { name: "Nuevo Usuario", email: "usuario@example.com" };
    await postData("usuario/create", dataToSend);
}

main();*/