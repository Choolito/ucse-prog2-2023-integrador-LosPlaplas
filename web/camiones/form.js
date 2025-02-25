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
    exitoCrearCamion,
    errorCrearCamion
  );
}

function actualizarCamion(idCamion) {
  const data = {
    PesoMaximo: parseInt(document.getElementById("PesoMaximo").value),
    CostoPorKilometro: parseInt(document.getElementById("CostoPorKm").value),
  };
  debugger;
  makeRequest(
    `${urlConFiltro}/${idCamion}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoActualizarCamion,
    errorActualizarCamion
  );

  }

  function exitoCrearCamion (response) {
    alert("Camion creado con exito");
    window.location = window.location.origin + "/web/camiones/index_camion.html";
  }

  function errorCrearCamion(error) {
    alert("Error en la solicitud al servidor.");
    console.log(error.json());
    throw new Error("Error en la solicitud al servidor.");
  }

  function exitoActualizarCamion (response) {
    alert("Camion actualizado con exito");
    window.location = window.location.origin + "/web/camiones/index_camion.html";
  }

  function errorActualizarCamion(error) {
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
  console.log("Datos recibidos del backend:", response);

  const patenteInput = document.getElementById("Patente");
  const pesoMaximoInput = document.getElementById("PesoMaximo");
  const costoPorKilometroInput = document.getElementById("CostoPorKm");

  if (!patenteInput || !pesoMaximoInput || !costoPorKilometroInput) {
      console.error("❌ Error: No se encontraron los elementos en el formulario.");
      return;
  }

  // Asignar valores
  patenteInput.value = response.patente || "";
  pesoMaximoInput.value = response.pesoMaximo || 0;
  costoPorKilometroInput.value = response.costoPorKilometro || 0;

  // Deshabilitar el input de la patente si estamos editando
  const urlParams = new URLSearchParams(window.location.search);
  const operacion = urlParams.get("tipo");
  if (operacion === "EDITAR") {
      patenteInput.setAttribute("disabled", "true");
  }
}



function errorObtenerCamionEdicion(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

