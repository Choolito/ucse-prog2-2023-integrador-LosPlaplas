const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function () {
  const guardarParada = document.getElementById("guardar-parada");
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (guardarParada) {
    guardarParada.addEventListener("click", function () {
      console.log("Botón 'guardar-parada' presionado. ID:", idEnvio, "Operación:", operacion);
      
      if (idEnvio && operacion === "FINALIZAR") {
        finalizarEnvio(idEnvio);
      } else if (idEnvio && operacion === "PARADA") {
        generarParada(idEnvio);
      }
    });
  }
});

const urlConFiltro = `http://localhost:8080/envios`;

function generarParada(idEnvio) {
  if (confirm("¿Está seguro que desea generar esta parada?")) {
    // Obtener los valores de los campos de entrada
    const ciudad = document.getElementById("ciudad").value;
    const kmRecorridos = document.getElementById("kmrecorridos").value;

    // Crear un objeto con los datos
    const data = {
      ciudad: ciudad, // Corregido: la clave debe coincidir con el backend
      kmRecorridos: parseInt(kmRecorridos, 10) // Asegurar que es un número
    };

    console.log("Enviando solicitud para generar parada:", data);

    makeRequest(
      `${urlConFiltro}/parada/${idEnvio}`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoGenerarParada,
      errorGenerarParada
    );
  } else {
    window.location = "/web/envios/index_envio.html";
  }
}

function finalizarEnvio(idEnvio) {
  if (confirm("¿Está seguro que desea finalizar el envío?")) {
    // Obtener los valores de los campos de entrada
    const ciudad = document.getElementById("ciudad").value;
    const kmRecorridos = document.getElementById("kmrecorridos").value;

    // Crear un objeto con los datos
    const data = {
      ciudad: ciudad, // Corregido
      kmRecorridos: parseInt(kmRecorridos, 10)
    };

    console.log("Enviando solicitud para finalizar envío:", data);

    makeRequest(
      `${urlConFiltro}/finalizar/${idEnvio}`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoFinalizarEnvio,
      errorFinalizarEnvio
    );
  } else {
    window.location = "/web/envios/index_envio.html";
  }
}

function exitoGenerarParada(response) {
  alert("Parada generada con éxito.");
  console.log("Respuesta del servidor (parada):", response);
  window.location.href = "http://localhost:8080/web/envios/index_envio.html"; // Redirección absoluta
}

function exitoFinalizarEnvio(response) {
  alert("Envío finalizado con éxito.");
  console.log("Respuesta del servidor (finalizar envío):", response);
  window.location.href = "http://localhost:8080/web/envios/index_envio.html"; // Redirección absoluta
}




// ❌ Función de error al generar parada
function errorGenerarParada(error) {
  alert("Error al generar la parada.");
  console.error("Error en generar parada:", error);
}


// ❌ Función de error al finalizar envío
function errorFinalizarEnvio(error) {
  alert("Error al finalizar el envío.");
  console.error("Error en finalizar envío:", error);
}
