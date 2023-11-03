const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");
let modoEdicion = false;

document.addEventListener("DOMContentLoaded", function (event) {
  
  debugger;
  const guardarButton = document.getElementById("guardar-button");
  const urlParams = new URLSearchParams(window.location.search);
  const idProducto = urlParams.get("id");
  const operacion = urlParams.get("tipo");
  if (
    idProducto != "" &&
    idProducto != null &&
    operacion == "EDITAR"
  ) {
        modoEdicion = true;
        obtenerProductoParaEdicion(idProducto);
  }
  if (guardarButton) {
    guardarButton.addEventListener("click", function (event) {
      event.preventDefault(); // Prevenir el envío del formulario si estás utilizando uno
      if (validarCampos()) {
        guardarOActualizarProducto(idProducto);
      }
    });
  }
});
const urlConFiltro = `http://localhost:8080/productos`;

function validarCampos() {
  const codigoProducto = document.getElementById("codigo-producto").value;
  const tipoProducto = document.getElementById("tipo").value;
  const nombre = document.getElementById("nombre").value;
  const precio = document.getElementById("precio").value;
  const peso = document.getElementById("peso").value;
  const stockMinimo = document.getElementById("stock-minimo").value;
  const stockActual = document.getElementById("stock-actual").value;

  if (
    codigoProducto.trim() === "" ||
    tipoProducto.trim() === "" ||
    nombre.trim() === "" ||
    precio.trim() === "" ||
    peso.trim() === "" ||
    stockMinimo.trim() === "" ||
    stockActual.trim() === ""
  ) {
    alert("Todos los campos son obligatorios. Por favor, complete todos los campos.");
    return false;
  }

  return true;
}

function guardarOActualizarProducto(idProducto) {
  if (modoEdicion) {
    // Si estamos en modo de edición, ejecuta el método actualizarProducto
    actualizarProducto(idProducto);
  } else {
    // Si no estamos en modo de edición, ejecuta el método guardarProducto
    guardarProducto();
  }
}

function guardarProducto() {
  //armo la data a enviar
  const data = {
    CodigoProducto: document.getElementById("codigo-producto").value,
    TipoProducto: document.getElementById("tipo").value,
    Nombre: document.getElementById("nombre").value,
    PesoUnitario: parseInt(document.getElementById("peso").value),
    PrecioUnitario: parseInt(
      document.getElementById("precio").value
    ),
    StockMinimo: parseInt(document.getElementById("stock-minimo").value),
    CantidadEnStock: parseInt(document.getElementById("stock-actual").value),
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoProducto,
    errorProducto
  );
}


function actualizarProducto(idProducto) {
  debugger;
  const data = {
    // Obtiene los datos del formulario, puedes actualizar solo nombre y precio aquí
    Nombre: document.getElementById("nombre").value,
    PrecioUnitario: parseInt(document.getElementById("precio").value),
  };

  makeRequest(
    `${urlConFiltro}/${idProducto}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoProducto,
    errorProducto
  );
}


function exitoProducto(data) {
  debugger;
  window.location = window.location.origin + "/web/productos/index_producto.html";
}

function errorProducto(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function obtenerProductoParaEdicion(idProducto) {

  makeRequest(
    `${urlConFiltro}/${idProducto}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProductoEdicion,
    errorObtenerProductoEdicion
  );
}

function exitoObtenerProductoEdicion(data) {
  document.getElementById("codigo-producto").value = data.CodigoProducto;
  document.getElementById("tipo").value = data.TipoProducto;
  document.getElementById("nombre").value = data.Nombre;
  document.getElementById("peso").value = data.PesoUnitario;
  document.getElementById("precio").value = data.PrecioUnitario;
  document.getElementById("stock-minimo").value = data.StockMinimo;
  document.getElementById("stock-actual").value = data.CantidadEnStock;

  // Deshabilitar campos que no se pueden editar
  document.getElementById("codigo-producto").disabled = true;
  document.getElementById("tipo").disabled = true;
  document.getElementById("peso").disabled = true;
  document.getElementById("stock-minimo").disabled = true;
  document.getElementById("stock-actual").disabled = true;
}

function errorObtenerProductoEdicion(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}