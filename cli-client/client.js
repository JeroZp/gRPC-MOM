#!/usr/bin/env node
import inquirer from "inquirer";
import axios from "axios";

const api = axios.create({ baseURL: "http://localhost:8080" });

async function main() {
  const { action } = await inquirer.prompt({
    type: "list",
    name: "action",
    message: "¿Qué quieres hacer?",
    choices: [
      { name: "Crear usuario", value: "create" },
      { name: "Listar usuarios",  value: "read"   },
      { name: "Actualizar usuario", value: "update" },
      { name: "Borrar usuario", value: "delete" },
      { name: "Salir",       value: "exit"   },
    ],
  });

  if (action === "exit") {
    console.log("¡Hasta luego!");
    process.exit(0);
  }

  try {
    switch (action) {
      case "create": {
        const { name, email } = await inquirer.prompt([
          { name: "name",  message: "Nombre:"  },
          { name: "email", message: "Email:"   },
        ]);
        const res = await api.post("/register", { name, email });
        console.log("→ Usuario creado:", res.data);
        break;
      }
      case "read": {
        const res = await api.get("/users");
        console.log("→ Usuarios existentes:");
        console.table(res.data);
        break;
      }
      case "update": {
        const { id, name, email, credits } = await inquirer.prompt([
          { name: "id",      message: "ID de usuario:"       },
          { name: "name",    message: "Nuevo nombre:"        },
          { name: "email",   message: "Nuevo email:"         },
          { name: "credits", message: "Nuevos créditos:"     },
        ]);
        const res = await api.put("/user", { id, name, email, credits: +credits });
        console.log("→ Usuario actualizado:", res.data);
        break;
      }
      case "delete": {
        const { id } = await inquirer.prompt([{ name: "id", message: "ID a borrar:" }]);
        const res = await api.delete(`/user/${id}`);
        console.log("→", res.data);
        break;
      }
    }
  } catch (err) {
    console.error("× Error:", err.response?.data || err.message);
  }

  main();
}

main();