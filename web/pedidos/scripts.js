const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerPedidos();
  const urlParams = new URLSearchParams(window.location.search);
  const idPedido = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idPedido != "" && idPedido != null && operacion == "ACEPTAR") {
    aceptarPedido(idPedido);
  } else if (idPedido != "" && idPedido != null && operacion == "CANCELAR") {
    cancelarPedido(idPedido);
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarPedido(event);
      });
  }

});

function obtenerPedidos() {
  urlConFiltro = `http://localhost:8080/pedidos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function exitoObtenerPedidos(response) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");
  
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila

      row.innerHTML=`
                    <td>${elemento.ID}</td>
                    <td>${elemento.ListaProductos ?
                      elemento.ListaProductos
                      .map(
                      (producto) => `
                      ${producto.Nombre}`
                       )
                      .join(", ")
                      : `Nohay productos disponibles`}
                    </td>
                    <td>${elemento.CiudadDestinoPedido}</td>
                    <td>${elemento.EstadoPedido}</td>
                    <td>${elemento.FechaCreacion}</td>
                    <td class="acciones"> <a href="index_pedido.html?id=${elemento.ID}&tipo=ACEPTAR">Aceptar Pedido</a> | <a href="index_pedido.html?id=${elemento.ID}&tipo=CANCELAR">Cancelar Pedido</a></td>
                    `;
      elementosTable.appendChild(row);
    });
  }
}

function errorObtenerPedidos(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function aceptarPedido(id) {
  if (confirm("¿Estás seguro de que deseas aceptar este pedido?")) {
    debugger;
    makeRequest(
      `http://localhost:8080/pedidos/aceptar/${id}`,
      Method.PUT,
      id,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoAceptarPedido,
      errorAceptarPedido
    );
  } else {
    window.location.href = "/web/pedidos/index_pedido.html";
  }
}

function exitoAceptarPedido(response) {
  window.location.href = window.location.origin + "/web/pedidos/index_pedido.html";
}


function errorAceptarPedido(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function cancelarPedido(id) {
  if (confirm("¿Estás seguro de que deseas cancelar el pedido?")) {
    debugger;
    makeRequest(
      `http://localhost:8080/pedidos/cancelar/${id}`,
      Method.PUT,
      id,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoCancelarPedido,
      errorCancelarPedido
    );
  } else {
    window.location = "/web/pedidos/index_pedido.html";
  }
}

function exitoCancelarPedido(response) {
  window.location = window.location.origin + "/web/pedidos/index_pedido.html";
}

function errorCancelarPedido(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}