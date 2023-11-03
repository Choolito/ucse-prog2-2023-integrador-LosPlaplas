const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");
let modoEdicion = false;

document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const guardarButton = document.getElementById("guardar-button");
  const urlParams = new URLSearchParams(window.location.search);
  const idCamion = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  debugger;
  if (
    idCamion != "" &&
    idCamion != null &&
    operacion == "EDITAR"
  ) {
        modoEdicion = true;
        obtenerCamionParaEdicion(idCamion);
  }
  if (guardarButton) {
    guardarButton.addEventListener("click", function (event) {
      event.preventDefault(); // Prevenir el envío del formulario si estás utilizando uno
      if (validarCampos()) {
        guardarOActualizarCamion(idCamion);
      }
    });
  }
});
const urlConFiltro = `http://localhost:8080/camiones`;


function validarCampos() {
  const patente = document.getElementById("Patente").value;
  const pesoMaximo = document.getElementById("PesoMaximo").value;
  const costoPorKm = document.getElementById("CostoPorKm").value;

  if (!patente || !pesoMaximo || !costoPorKm) {
    alert("Por favor, complete todos los campos obligatorios.");
    return false;
  }

  return true;
}

function guardarOActualizarCamion(idCamion) {
  if (modoEdicion) {
    actualizarCamion(idCamion);
  }
  else {
    guardarCamion();
  }
}

function guardarCamion() {
  const data = {
    Patente: document.getElementById("Patente").value,
    PesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    CostoPorKilometro: parseInt(document.getElementById("CostoPorKm").value),
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );
}

function actualizarCamion(idCamion) {
  const data = {
    CostoPorKilometro: parseInt(document.getElementById("CostoPorKm").value),
  };
  debugger;
  makeRequest(
    `${urlConFiltro}/${idCamion}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  );

  }

  function exitoCamion (response) {
    alert("Camion actualizado con exito");
    window.location = window.location.origin + "/web/productos/index_producto.html";
  }

  function errorCamion(error) {
    alert("Error en la solicitud al servidor.");
    console.log(error.json());
    throw new Error("Error en la solicitud al servidor.");
  }

function obtenerCamionParaEdicion(idCamion) {
  debugger;
  makeRequest(
    `${urlConFiltro}/${idCamion}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerCamionEdicion,
    errorObtenerCamionEdicion
  );
}

function exitoObtenerCamionEdicion(response) {
  document.getElementById("Patente").value = response.Patente;
  document.getElementById("PesoMaximo").value = response.PesoMaximo;
  document.getElementById("CostoPorKm").value = response.CostoPorKilometro;

  //Deshabilitar
  document.getElementById("Patente").disabled = true;
  document.getElementById("PesoMaximo").disabled = true;
}

function errorObtenerCamionEdicion(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

