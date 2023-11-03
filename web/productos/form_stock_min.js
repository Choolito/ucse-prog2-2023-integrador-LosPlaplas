const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  filtrarProductos();
});

function obtenerProductosStockMenor() {
    debugger;
    const urlConFiltro = `http://localhost:8080/productos/stockminimo`;

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
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  // Llenar la tabla con los datos obtenidos
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila

      row.innerHTML = ` 
                    <td>${elemento.CodigoProducto}</td>
                    <td>${elemento.Nombre}</td>
                    <td>${elemento.TipoProducto}</td>
                    <td>${elemento.PrecioUnitario}</td>
                    <td>${elemento.PesoUnitario}</td>
                    <td>${elemento.StockMinimo}</td>
                    <td>${elemento.CantidadEnStock}</td>
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

function filtrarProductos() {
    const filtroSelect = document.getElementById("filtroSelect");
    const filtro = filtroSelect.value;

    const urlConFiltro = filtro ? `http://localhost:8080/productos/stockminimo/${filtro}` : `http://localhost:8080/productos/stockminimo`;

    makeRequest(
        urlConFiltro,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoObtenerProductos,
        errorObtenerProductos
    );
}

function quitarFiltro() {
    debugger;
    // Establecer el valor del select en vacío para mostrar todos los productos
    const filtroSelect = document.getElementById("filtroSelect");
    filtroSelect.value = "";

    // Llamar a la función de filtrado para mostrar todos los productos
    filtrarProductos();
}