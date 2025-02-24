const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
   obtenerProductos();

  // Agrega un evento de escucha al formulario
  const form = document.querySelector("form");
  form.addEventListener("submit", function (event) {
    // Validar que al menos un checkbox esté seleccionado
    const checkboxes = document.querySelectorAll('input[type="checkbox"]');
    const checkedCheckboxes = Array.from(checkboxes).filter((checkbox) => checkbox.checked);

    if (checkedCheckboxes.length === 0) {
      alert("Debes seleccionar al menos un producto.");
      event.preventDefault(); // Evitar que se envíe el formulario
      return;
    }
    
    // Validar que la cantidad se proporcione para los productos seleccionados
    const cantidadInputs = document.querySelectorAll('input[type="number"]');
    for (const cantidadInput of cantidadInputs) {
      if (cantidadInput.closest("tr").querySelector('input[type="checkbox"]').checked) {
        const cantidad = parseInt(cantidadInput.value);
        if (isNaN(cantidad) || cantidad <= 0) {
          alert("La cantidad para los productos seleccionados debe ser un número mayor que 0.");
          event.preventDefault(); // Evitar que se envíe el formulario
          return;
        }
      }
    }

    // Aquí puedes llamar a la función guardarPedido si todas las validaciones son exitosas
    guardarPedido(event);
  });
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
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  // Llenar la tabla con los datos obtenidos
  if (response != null) {
    response.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila

      row.innerHTML = ` 
                    <td>${elemento.codigoProducto}</td>
                    <td>${elemento.nombre}</td>
                    <td>${elemento.tipoProducto}</td>
                    <td>${elemento.precioUnitario}</td>
                    <td class="acciones">
                    <input type="number" name="cantidad" min="1" required>
                    </td>
                    <td>
                    <input type="checkbox" name="seleccionar" value="${elemento.id}" required>
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

function guardarPedido(event) {
  event.preventDefault(); // Evita que el formulario se recargue automáticamente

  if (confirm("¿Estás seguro de que deseas guardar este pedido?")) {
    urlConFiltro = `http://localhost:8080/pedidos`;
    data = obtenerDatosParaPedido();

    console.log("Enviando solicitud a:", urlConFiltro);
    console.log("Datos enviados:", JSON.stringify(data, null, 2));

    makeRequest(
      urlConFiltro,
      Method.POST,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      (response) => {
        console.log("Respuesta del servidor:", response);
        alert("Pedido guardado correctamente.");
        window.location.href = "/web/pedidos/index_pedido.html"; // Redirigir al index
      },
      (error) => {
        alert("Error en la solicitud al servidor.");
        console.log(error);
      }
    );
  }
}



function exitoGuardadoPedido(response) {
  console.log("Respuesta del servidor:", response);
  alert("Pedido guardado correctamente.");
  window.location.href = "/web/pedidos/index_pedido.html";
}


function errorGuardadoPedido(error) {
  alert("Error en la solicitud al servidor.");
  console.log(error.json());
  throw new Error("Error en la solicitud al servidor.");
}



function obtenerDatosParaPedido() {
  // Obtener la ciudad destino del formulario
  const ciudadDestino = document.getElementById("ciudad").value;

  // Obtener los checkboxes seleccionados con sus cantidades
  const productosSeleccionados = [];
  const checkboxes = document.querySelectorAll('input[type="checkbox"]:checked');
  
  checkboxes.forEach(checkbox => {
    const idProducto = checkbox.value;
    const cantidadInput = checkbox.closest("tr").querySelector('input[name="cantidad"]');
    const cantidad = parseInt(cantidadInput.value);

    if (!isNaN(cantidad) && cantidad > 0) {
      productosSeleccionados.push({ idProducto, cantidad });
    }
  });

  // Crear el objeto JSON
  const data = {
    listaProductos: productosSeleccionados,
    ciudadDestinoPedido: ciudadDestino
  };

  return data;
}

