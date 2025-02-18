const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerCamiones();

  const urlParams = new URLSearchParams(window.location.search);
  const codCamion = urlParams.get("id");
  const operacion = urlParams.get("tipo");
  if (codCamion != "" && codCamion != null && operacion == "ELIMINAR") {
    eliminarCamion(codCamion);
  } 

});

const urlConFiltro = `http://localhost:8080/camiones`;

function obtenerCamiones() {

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerCamiones,
    errorObtenerCamiones
  )
}

function exitoObtenerCamiones(response) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  // Llenar la tabla con los datos obtenidos
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila

      row.innerHTML = ` 
                    <td>${elemento.patente}</td>
                    <td>${elemento.pesoMaximo}</td>
                    <td>${elemento.costoPorKilometro}</td>
                    <td class="acciones"> 
                    <a href="form_nuevo_camion.html?id=${elemento.ID}&tipo=EDITAR">Editar</a> |
                    <a href="index_camion.html?id=${elemento.ID}&tipo=ELIMINAR">Eliminar</a> 
                    </td>
                    `;

      elementosTable.appendChild(row);
    });
  }
}

function errorObtenerCamiones(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

function eliminarCamion(idCamion) {
  makeRequest(
    `${urlConFiltro}/${idCamion}`,
    Method.DELETE,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoCamion,
    errorCamion
  )
}

function exitoCamion(response) {
  alert("Camion eliminado con exito");
  location.reload();
}

function errorCamion(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}

