const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  
  const guardarParada = document.getElementById("guardar-parada");
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  debugger;
  if (guardarParada) {
    guardarParada.addEventListener("click", function (event) {
      if (idEnvio != "" && idEnvio != null && operacion == "FINALIZAR") {
        finalizarEnvio(idEnvio);
      }
      else if (idEnvio != "" && idEnvio != null && operacion == "PARADA") {
        generarParada(idEnvio);
      }
    });
  }

});

const urlConFiltro = `http://localhost:8080/envios`;
function generarParada(idEnvio) {
  if(confirm("¿Está seguro que desea generar esta parada?")){
    // Obtener los valores de los campos de entrada
  const ciudad = document.getElementById("ciudad").value;
  const kmRecorridos = document.getElementById("kmrecorridos").value;

  // Crear un objeto con los datos
  const data = {
    Ciudad: ciudad,
    KmRecorridos: kmRecorridos
  };

  makeRequest(
    `${urlConFiltro}/parada/${idEnvio}`,
    Method.PUT,
    data, // Enviar los datos como parte de la solicitud
    ContentType.JSON,
    CallType.PRIVATE,
    exitoEnvio,
    errorEnvio
  );
  }
  else
  {
    window.location = '/web/envios/index_envio.html';
  }
}

function finalizarEnvio(idEnvio) {
  if(confirm("¿Está seguro que desea finalizar el envío?")){
    // Obtener los valores de los campos de entrada
  const ciudad = document.getElementById("ciudad").value;
  const kmRecorridos = document.getElementById("kmrecorridos").value;

  // Crear un objeto con los datos
  const data = {
    Ciudad: ciudad,
    KmRecorridos: kmRecorridos
  };

  makeRequest(
    `${urlConFiltro}/finalizar/${idEnvio}`,
    Method.PUT,
    data, // Enviar los datos como parte de la solicitud
    ContentType.JSON,
    CallType.PRIVATE,
    exitoEnvio,
    errorEnvio
  );
  }
  else
  {
    window.location = '/web/envios/index_envio.html';
  }
}

function exitoEnvio(response) {
  alert("Envío iniciado con éxito");
  window.location = "/web/envios/index_envio.html";
}

function errorEnvio(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

