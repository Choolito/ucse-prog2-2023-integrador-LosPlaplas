const customHeaders = new Headers()
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0")
customHeaders.append("Accept", "*/*")
customHeaders.append("Accept-Encoding", "gzip, deflate, br")
customHeaders.append("Connection", "keep-alive")


document.addEventListener("DOMContentLoaded", () => {
  obtenerCamiones()
  obtenerEnvios()
  obtenerPedidos()

  const urlParams = new URLSearchParams(window.location.search)
  const idEnvio = urlParams.get("id")
  const operacion = urlParams.get("tipo")

  if (idEnvio && operacion === "INICIAR") {
    console.log(`Iniciando envío con ID: ${idEnvio}`)
    setTimeout(() => iniciarEnvio(idEnvio), 100)
  }
})

const urlConFiltro = `http://localhost:8080/envios`
let listaCamiones = []
let listaPedidos = []

function obtenerEnvios() {
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerEnvios,
    errorObtenerEnvios,
  )
}

function exitoObtenerEnvios(response) {
  const enviosResponse = response

  if (listaPedidos.length === 0) {
    makeRequest(
      `http://localhost:8080/pedidos`,
      Method.GET,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      (pedidosResponse) => {
        listaPedidos = pedidosResponse
        renderizarTablaEnvios(enviosResponse)
      },
      errorObtenerPedidos,
    )
  } else {
    renderizarTablaEnvios(enviosResponse)
  }
}

function errorObtenerEnvios(error) {
  alert("Error en la solicitud de envíos.")
  console.error(error)
}

function obtenerCamiones() {
  makeRequest(
    `http://localhost:8080/camiones`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerCamiones,
    errorObtenerCamiones,
  )
}

function exitoObtenerCamiones(response) {
  listaCamiones = response
}

function errorObtenerCamiones(error) {
  alert("Error en la solicitud de camiones.")
  console.error(error)
}

function obtenerPedidos() {
  makeRequest(
    `http://localhost:8080/pedidos`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorObtenerPedidos,
  )
}

function exitoObtenerPedidos(response) {
  listaPedidos = response
}

function errorObtenerPedidos(error) {
  alert("Error en la solicitud de pedidos.")
  console.error(error)
}

function obtenerCiudadActual(envio, pedidoInfo) {
  // Si hay paradas, tomar la última ciudad
  if (envio.Paradas && envio.Paradas.length > 0) {
    return envio.Paradas[envio.Paradas.length - 1].Ciudad
  }
  // Si no hay paradas, tomar la ciudad destino del pedido
  return pedidoInfo ? pedidoInfo.CiudadDestinoPedido : "Ciudad no disponible"
}

function calcularKilometrosRecorridos(paradas) {
  if (!paradas || paradas.length === 0) return 0
  return paradas.reduce((total, parada) => total + parada.KmRecorridos, 0)
}

function renderizarTablaEnvios(response) {
  const elementosTable = document.getElementById("elementosTable").querySelector("tbody")
  elementosTable.innerHTML = "" // Limpia la tabla antes de agregar nuevos datos

  if (response) {
    response.forEach((elemento) => {
      const row = document.createElement("tr")

      // Obtener información del camión
      const camionInfo = listaCamiones.find((camion) => camion.id === elemento.IDCamion)

      // Obtener información del primer pedido (asumiendo que necesitamos uno para la ciudad destino)
      const pedidoIds = Array.isArray(elemento.Pedidos) ? elemento.Pedidos : [elemento.Pedidos]
      const pedidoInfo = listaPedidos.find((pedido) => pedidoIds.includes(pedido.ID))

      // Calcular ciudad actual (última parada registrada)
      const ciudadActual = elemento.Paradas && elemento.Paradas.length > 0
        ? elemento.Paradas[elemento.Paradas.length - 1].Ciudad
        : "Sin paradas"

      // Obtener la ciudad destino del pedido
      const ciudadDestino = pedidoInfo ? pedidoInfo.CiudadDestinoPedido : "Desconocida"

      // Calcular kilómetros recorridos
      const kilometrosRecorridos = calcularKilometrosRecorridos(elemento.Paradas)

      row.innerHTML = `
        <td>${camionInfo ? camionInfo.patente : "No encontrado"}</td>
        <td>${ciudadActual}</td>
        <td>${ciudadDestino}</td>
        <td>${kilometrosRecorridos}</td>
        <td>${elemento.Estado}</td>
        <td class="acciones">
          <a href="form_parada.html?id=${elemento.ID}&tipo=PARADA">Generar Parada</a>
        </td>
        <td class="acciones">
          <a href="#" onclick="iniciarEnvio('${elemento.ID}')">Iniciar</a> | 
          <a href="form_parada.html?id=${elemento.ID}&tipo=FINALIZAR">Finalizar</a>
        </td>
      `

      elementosTable.appendChild(row)
    })
  }
}


function iniciarEnvio(idEnvio) {
  if (confirm("¿Estás seguro de que deseas iniciar este envío?")) {
    console.log(`Llamando a la API para iniciar el envío: ${idEnvio}`)

    makeRequest(
      `http://localhost:8080/envios/iniciar/${idEnvio}`,
      Method.PUT,
      JSON.stringify({}),
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio,
    )
  }
}

function exitoEnvio(response) {
  alert("Envío iniciado con éxito")
  console.log("Respuesta exitosa:", response)
  window.location.href = "/web/envios/index_envio.html"
}

function errorEnvio(error) {
  alert("Error al iniciar el envío.")
  console.error("Error en la API:", error)
}

