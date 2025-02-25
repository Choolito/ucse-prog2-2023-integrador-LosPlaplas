const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerCamiones();    
  obtenerEnvios();
  obtenerPedidos();  // Agregar el llamado a la función de obtener pedidos

  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idEnvio != "" && idEnvio != null && operacion == "INICIAR") {
    iniciarEnvio(idEnvio);
  }
}); 

const urlConFiltro = `http://localhost:8080/envios`;
let listaCamiones = [];
let listaPedidos = [];  // Nueva lista para los pedidos

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

// Modificación de exitoObtenerEnvios para esperar a los pedidos
function exitoObtenerEnvios(response) {
  // Guardamos la respuesta de envíos para procesarla cuando tengamos los pedidos
  const enviosResponse = response;
  
  // Si aún no tenemos pedidos, los obtenemos primero
  if (listaPedidos.length === 0) {
    makeRequest(
      `http://localhost:8080/pedidos`,
      Method.GET,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      (pedidosResponse) => {
        // Guardamos los pedidos
        listaPedidos = pedidosResponse;
        // Y ahora renderizamos la tabla
        renderizarTablaEnvios(enviosResponse);
      },
      errorObtenerPedidos
    );
  } else {
    // Si ya tenemos pedidos, renderizamos directamente
    renderizarTablaEnvios(enviosResponse);
  }
}

// Nueva función para renderizar la tabla
function renderizarTablaEnvios(response) {
  const elementosTable = document.getElementById("elementosTable").querySelector("tbody");
  
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr");
      console.log("Elemento: ", elemento);
      
      // Se busca el camión correspondiente usando la propiedad "id" (minúscula) del objeto camión
      const camionInfo = listaCamiones.find((camion) => camion.id === elemento.IDCamion);
      
      // Asegurémonos de que Pedidos sea tratado como array
      let pedidoIds = Array.isArray(elemento.Pedidos) ? elemento.Pedidos : [elemento.Pedidos];
      // Buscamos el primer pedido que coincida
      const pedidoInfo = listaPedidos.find((pedido) => pedidoIds.includes(pedido.ID));

      console.log("Camión encontrado: ", camionInfo);
      console.log("Pedido encontrado: ", pedidoInfo);

      row.innerHTML = `
        <td>${camionInfo ? camionInfo.patente : 'No encontrado'}</td>
        <td>${pedidoInfo ? pedidoInfo.CiudadDestinoPedido : 'Ciudad no disponible'}</td>
        <td>${elemento.Ciudad || '-'}</td>
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


function errorObtenerEnvios(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function obtenerCamiones() {
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

// Función para obtener los pedidos
function obtenerPedidos() {
  makeRequest(
    `http://localhost:8080/pedidos`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function exitoObtenerPedidos(response) {
  listaPedidos = response;
}

function errorObtenerPedidos(error) {
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
