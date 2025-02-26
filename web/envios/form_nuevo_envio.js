document.addEventListener("DOMContentLoaded", () => {
    cargarCamiones()
    cargarPedidos()
  
    // Configurar eventos
    document.getElementById("confirmar-envio").addEventListener("click", confirmarEnvio)
  })
  
  // Variable global para almacenar los pedidos cargados
  let pedidosCargados = []
  
  function cargarCamiones() {
    makeRequest(
      "http://localhost:8080/camiones",
      Method.GET,
      null,
      ContentType.JSON,
      true,
      exitoCargarCamiones,
      errorCargarCamiones,
    )
  }
  
  function exitoCargarCamiones(response) {
    const selectCamion = document.getElementById("camion-select")
    selectCamion.innerHTML = ""
  
    // Opción por defecto
    const defaultOption = document.createElement("option")
    defaultOption.value = ""
    defaultOption.textContent = "Seleccione un camión"
    selectCamion.appendChild(defaultOption)
  
    response.forEach((camion) => {
      const option = document.createElement("option")
      option.value = camion.id
      option.textContent = `${camion.patente}`
      selectCamion.appendChild(option)
    })
  }
  
  function errorCargarCamiones(error, responseBody) {
    console.error("Error al cargar camiones", error)
    alert("Error al obtener la lista de camiones.")
  }
  
  function cargarPedidos() {
    makeRequest(
      "http://localhost:8080/pedidos",
      Method.GET,
      null,
      ContentType.JSON,
      true,
      exitoCargarPedidos,
      errorCargarPedidos,
    )
  }
  
  function exitoCargarPedidos(response) {
    // Guardar los pedidos en la variable global
    pedidosCargados = response
    mostrarPedidosEnTabla(response)
  }
  
  function errorCargarPedidos(error, responseBody) {
    console.error("Error al cargar pedidos", error)
    alert("Error al obtener la lista de pedidos.")
  }
  
  function mostrarPedidosEnTabla(pedidos) {
    const tablaPedidos = document.getElementById("pedidosTable").getElementsByTagName("tbody")[0]
    tablaPedidos.innerHTML = ""
  
    // Filtrar pedidos por estado "Para enviar" o "Aceptado"
    pedidos = pedidos.filter((pedido) => pedido.EstadoPedido === "Aceptado")
  
    pedidos.forEach((pedido) => {
      const fila = tablaPedidos.insertRow()
  
      // Celda para checkbox
      const celdaCheckbox = fila.insertCell()
      const checkbox = document.createElement("input")
      checkbox.type = "checkbox"
      checkbox.value = pedido.ID
      checkbox.className = "pedido-checkbox"
      celdaCheckbox.appendChild(checkbox)
  
      // Celda para código de pedido
      const celdaCodigo = fila.insertCell()
      celdaCodigo.textContent = pedido.ID
  
      // Celda para productos
      const celdaProductos = fila.insertCell()
      if (pedido.ListaProductos && pedido.ListaProductos.length > 0) {
        const listaProductos = document.createElement("ul")
        listaProductos.className = "productos-lista"
  
        pedido.ListaProductos.forEach((producto) => {
          const itemProducto = document.createElement("li")
          itemProducto.textContent = `${producto.Nombre} (${producto.Cantidad || 1})`
          listaProductos.appendChild(itemProducto)
        })
  
        celdaProductos.appendChild(listaProductos)
      } else {
        celdaProductos.textContent = "No hay productos disponibles"
      }
  
      // Celda para ciudad destino
      const celdaCiudad = fila.insertCell()
      celdaCiudad.textContent = pedido.CiudadDestinoPedido
  
      // Celda para estado
      const celdaEstado = fila.insertCell()
      celdaEstado.textContent = pedido.EstadoPedido
  
      // Celda para fecha
      const celdaFecha = fila.insertCell()
      const fecha = new Date(pedido.FechaCreacion)
      celdaFecha.textContent = fecha.toLocaleDateString()
    })
  }
  
  function confirmarEnvio() {
    const camionId = document.getElementById("camion-select").value
    if (!camionId) {
      alert("Por favor, seleccione un camión.")
      return
    }
  
    const pedidosSeleccionados = obtenerPedidosSeleccionados()
    if (pedidosSeleccionados.length === 0) {
      alert("Por favor, seleccione al menos un pedido para el envío.")
      return
    }
  
    // Buscar el primer pedido seleccionado en la lista de pedidos cargados
    const primerPedidoId = pedidosSeleccionados[0]
    const pedidoEncontrado = pedidosCargados.find((pedido) => pedido.ID === primerPedidoId)
  
    if (pedidoEncontrado) {
      const ciudadDestino = pedidoEncontrado.CiudadDestinoPedido
      crearNuevoEnvio(camionId, ciudadDestino, pedidosSeleccionados)
    } else {
      alert("Error: No se encontró información del pedido seleccionado.")
    }
  }
  
  function obtenerPedidosSeleccionados() {
    const checkboxes = document.querySelectorAll(".pedido-checkbox:checked")
    return Array.from(checkboxes).map((checkbox) => checkbox.value)
  }
  
  function crearNuevoEnvio(camionId, ciudadDestino, pedidosIds) {
    const data = {
      idCamion: camionId,
      pedidos: pedidosIds,
    }
  
    console.log("Enviando datos:", data)
  
    makeRequest(
      "http://localhost:8080/envios",
      Method.POST,
      data,
      ContentType.JSON,
      true,
      exitoEnviarEnvio,
      errorEnviarEnvio,
    )
  }
  
  function exitoEnviarEnvio(response) {
    // Actualizar el estado de los pedidos seleccionados a "Enviado"
    const pedidosSeleccionados = obtenerPedidosSeleccionados()
  
    // Función para actualizar cada pedido
    function actualizarPedido(index) {
      if (index >= pedidosSeleccionados.length) {
        // Todos los pedidos han sido actualizados
        alert("Envío registrado con éxito")
        window.location.href = "index_envio.html"
        return
      }
  
      const pedidoId = pedidosSeleccionados[index]
  
      makeRequest(
        `http://localhost:8080/pedidos/${pedidoId}/estado`,
        Method.PUT,
        { estado: "Enviado" },
        ContentType.JSON,
        true,
        () => {
          // Continuar con el siguiente pedido
          actualizarPedido(index + 1)
        },
        (error, responseBody) => {
          console.error(`Error al actualizar el estado del pedido ${pedidoId}`, error)
          // Continuar con el siguiente pedido a pesar del error
          actualizarPedido(index + 1)
        },
      )
    }
  
    // Iniciar la actualización de pedidos
    actualizarPedido(0)
  }
  
  function errorEnviarEnvio(error, responseBody) {
    console.error("Error al registrar el envío", error, responseBody)
    alert("Error al registrar el envío. Intente nuevamente.")
  }
  
  