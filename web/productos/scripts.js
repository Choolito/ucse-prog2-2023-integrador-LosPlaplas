const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerProductos();

  const urlParams = new URLSearchParams(window.location.search);
  const codProducto = urlParams.get("id");
  const operacion = urlParams.get("tipo");
  if (codProducto != "" && codProducto != null && operacion == "ELIMINAR") {
    eliminarProducto(codProducto);
  } 
});

function obtenerProductos() {
  urlConFiltro = `http://localhost:8080/productos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProductos,
    errorObtenerProductos
  );
}
  

function exitoObtenerProductos(response) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  // Llenar la tabla con los datos obtenidos
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr");

      row.innerHTML = ` 
                    <td>${elemento.codigoProducto}</td>
                    <td>${elemento.nombre}</td>
                    <td>${elemento.tipoProducto}</td>
                    <td>${elemento.precioUnitario}</td>
                    <td>${elemento.pesoUnitario}</td>
                    <td>${elemento.stockMinimo}</td>
                    <td>${elemento.cantidadEnStock}</td>
                    <td class="acciones"> 
                    <a href="form_nuevo_producto.html?id=${elemento.id}&tipo=EDITAR">Editar</a> |
                    <a href="index_producto.html?id=${elemento.id}&tipo=ELIMINAR">Eliminar</a> 
                    </td>
                    `;

      elementosTable.appendChild(row);
    });
  }
}


function errorObtenerProductos(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function eliminarProducto(codProducto) {
  if (confirm("¿Estás seguro de que deseas eliminar este producto?")) {
    debugger;
    makeRequest(
      `${urlConFiltro}/${codProducto}`,
      Method.DELETE,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoProducto,
      errorProducto
    );
  } else {
    window.location = "/web/productos/index_producto.html";
  }
}

function exitoProducto(data) {
  debugger;
  window.location.href =`/web/productos/index_producto.html`;
}

function errorProducto(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}



// function obtenerProductoFiltrado(tipo) {
//   var url = new URL(urlConFiltro);

//   switch (tipo) {
//     case "stock":
//       url.searchParams.set("filtrarPorStockMinimo", true);
//       break;
//     case "estado":
//       url.searchParams.set(
//         "tipoProducto",
//         document.getElementById("tipo").value
//       );
//       break;
//     default:
//       url = `http://localhost:8080/productos`;
//       break;
//   }

//   makeRequest(
//     `${url}`,
//     Method.GET,
//     null,
//     ContentType.JSON,
//     CallType.PRIVATE,
//     exitoObtenerEnvio,
//     errorEnvio
//   );
// }

