const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerCamiones();  
  obtenerEnvios();

  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idEnvio != "" && idEnvio != null && operacion == "INICIAR") {
    iniciarEnvio(idEnvio);
  }
}); 

const urlConFiltro = `http://localhost:8080/envios`;
let listaCamiones = [];


function obtenerEnvios() {
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerEnvios,
    errorObtenerEnvios
  );
}

function exitoObtenerEnvios(response) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr");

      // Busca el camión correspondiente en la lista de camiones
      const camionInfo = listaCamiones.find((camion) => camion.ID === elemento.IDCamion);

      row.innerHTML = `
        <td>${camionInfo ? camionInfo.Patente : 'No encontrado'}</td>
        <td>${elemento.Nombre}</td>
        <td>${elemento.TipoProducto}</td>
        <td>${elemento.Estado}</td>
        <td class="acciones">
          <a href="form_parada.html?id=${elemento.ID}&tipo=PARADA">Generar Parada</a>
        </td>
        <td class="acciones">
          <a href="index_envio.html?id=${elemento.ID}&tipo=INICIAR">Iniciar</a> | 
          <a href="form_parada.html?id=${elemento.ID}&tipo=FINALIZAR">Finalizar</a>
        </td>
      `;

      elementosTable.appendChild(row);
    });
  }
}


function errorObtenerEnvios(error)
{
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function obtenerCamiones() {
  debugger;
  makeRequest(
    `http://localhost:8080/camiones`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerCamiones,
    errorObtenerCamiones
  );
}

function exitoObtenerCamiones(response) {
  listaCamiones = response;
}

function errorObtenerCamiones(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function iniciarEnvio(idEnvio) {
  if (confirm("¿Estás seguro de que deseas iniciar este envío?")) {
    makeRequest(
      `${urlConFiltro}/iniciar/${idEnvio}`,
      Method.PUT,
      idEnvio,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  }
  else{
    window.location = "/web/envios/index_envio.html";
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
