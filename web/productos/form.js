const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const urlConFiltro = `http://localhost:8080/productos`;
let modoEdicion = false;

document.addEventListener("DOMContentLoaded", function (event) {
  const guardarButton = document.getElementById("guardar-button");
  const urlParams = new URLSearchParams(window.location.search);
  const idProducto = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idProducto && operacion === "EDITAR") {
    modoEdicion = true;
    document.querySelector('.formulario h2').textContent = 'Editar Producto';
    obtenerProductoParaEdicion(idProducto);
  }

  if (guardarButton) {
    guardarButton.addEventListener("click", function (event) {
      event.preventDefault();
      if (validarCampos()) {
        guardarOActualizarProducto(idProducto);
      }
    });
  }
});

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
  // Rellenar el formulario con los datos del producto
  document.getElementById("codigo-producto").value = data.codigoProducto;
  document.getElementById("tipo").value = data.tipoProducto;
  document.getElementById("nombre").value = data.nombre;
  document.getElementById("precio").value = data.precioUnitario;
  document.getElementById("peso").value = data.pesoUnitario;
  document.getElementById("stock-minimo").value = data.stockMinimo;
  document.getElementById("stock-actual").value = data.cantidadEnStock;

  // Deshabilitar campos que no deberían editarse
  document.getElementById("codigo-producto").readOnly = true;
  document.getElementById("tipo").disabled = true;
}

function errorObtenerProductoEdicion(error) {
  console.error("Error al obtener el producto:", error);
  alert("Error al cargar los datos del producto.");
}

function validarCampos() {
  const campos = [
    "codigo-producto",
    "tipo",
    "nombre",
    "precio",
    "peso",
    "stock-minimo",
    "stock-actual"
  ];

  for (const campo of campos) {
    const valor = document.getElementById(campo).value;
    if (!valor.trim()) {
      alert(`El campo ${campo.replace('-', ' ')} es obligatorio.`);
      return false;
    }
  }
  return true;
}

function guardarOActualizarProducto(idProducto) {
  if (modoEdicion) {
    actualizarProducto(idProducto);
  } else {
    guardarProducto();
  }
}

function guardarProducto() {
  const data = {
    codigoProducto: document.getElementById("codigo-producto").value,
    tipoProducto: document.getElementById("tipo").value,
    nombre: document.getElementById("nombre").value,
    precioUnitario: parseInt(document.getElementById("precio").value),
    pesoUnitario: parseInt(document.getElementById("peso").value),
    stockMinimo: parseInt(document.getElementById("stock-minimo").value),
    cantidadEnStock: parseInt(document.getElementById("stock-actual").value)
  };

  makeRequest(
    urlConFiltro,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoProducto,
    errorProducto
  );
}

function actualizarProducto(idProducto) {
  const data = {
    codigoProducto: document.getElementById("codigo-producto").value,
    tipoProducto: document.getElementById("tipo").value,
    nombre: document.getElementById("nombre").value,
    precioUnitario: parseInt(document.getElementById("precio").value),
    pesoUnitario: parseInt(document.getElementById("peso").value),
    stockMinimo: parseInt(document.getElementById("stock-minimo").value),
    cantidadEnStock: parseInt(document.getElementById("stock-actual").value)
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
  window.location.href = "/web/productos/index_producto.html";
}

function errorProducto(error) {
  console.error("Error:", error);
  alert("Error en la solicitud al servidor.");
}