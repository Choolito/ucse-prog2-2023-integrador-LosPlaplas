const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const baseUrl = "http://localhost:8080/productos";

document.addEventListener("DOMContentLoaded", function (event) {
    // Cargar productos iniciales con stock mínimo
    obtenerProductosStockMenor();
});

function obtenerProductosStockMenor() {
    console.log("Obteniendo productos con stock mínimo...");
    makeRequest(
        `${baseUrl}/stockminimo`,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoObtenerProductos,
        errorObtenerProductos
    );
}

function exitoObtenerProductos(response) {
    console.log("Respuesta del servidor:", response);
    const elementosTable = document
        .getElementById("elementosTable")
        .querySelector("tbody");

    // Limpiar la tabla antes de agregar nuevos datos
    elementosTable.innerHTML = '';

    // Llenar la tabla con los datos obtenidos
    if (response != null && response.length > 0) {
        response.forEach((elemento) => {
            // Verificar si el stock actual es menor o igual al stock mínimo
            if (elemento.cantidadEnStock <= elemento.stockMinimo) {
                const row = document.createElement("tr");
                row.innerHTML = ` 
                    <td>${elemento.codigoProducto || ''}</td>
                    <td>${elemento.nombre || ''}</td>
                    <td>${elemento.tipoProducto || ''}</td>
                    <td>${elemento.precioUnitario || 0}</td>
                    <td>${elemento.pesoUnitario || 0}</td>
                    <td>${elemento.stockMinimo || 0}</td>
                    <td>${elemento.cantidadEnStock || 0}</td>
                `;
                elementosTable.appendChild(row);
            }
        });

        // Si después de filtrar no hay productos, mostrar mensaje
        if (elementosTable.children.length === 0) {
            mostrarMensajeNoProductos(elementosTable);
        }
    } else {
        mostrarMensajeNoProductos(elementosTable);
    }
}

function mostrarMensajeNoProductos(elementosTable) {
    const row = document.createElement("tr");
    row.innerHTML = '<td colspan="7" style="text-align: center;">No se encontraron productos con stock mínimo</td>';
    elementosTable.appendChild(row);
}

function errorObtenerProductos(error) {
    console.error("Error al obtener productos:", error);
    alert("Error en la solicitud al servidor.");
}

function filtrarProductos() {
    const filtroSelect = document.getElementById("filtroSelect");
    const filtro = filtroSelect.value;

    console.log("Aplicando filtro:", filtro);

    // Si hay un filtro seleccionado, usar el endpoint de filtrado
    if (filtro) {
        makeRequest(
            `${baseUrl}/stockminimo/${filtro}`,
            Method.GET,
            null,
            ContentType.JSON,
            CallType.PRIVATE,
            exitoObtenerProductos,
            errorObtenerProductos
        );
    } else {
        // Si no hay filtro, obtener todos los productos con stock mínimo
        obtenerProductosStockMenor();
    }
}

function quitarFiltro() {
    console.log("Quitando filtros...");
    // Restablecer el select
    const filtroSelect = document.getElementById("filtroSelect");
    filtroSelect.value = "";

    // Volver a cargar los productos con stock mínimo
    obtenerProductosStockMenor();
}

// Agregar los enums necesarios si no están definidos en request.js
const Method = {
    GET: "GET",
    POST: "POST",
    PUT: "PUT",
    DELETE: "DELETE"
};

const ContentType = {
    JSON: "application/json",
    FORM_URL_ENCODED: "application/x-www-form-urlencoded"
};

const CallType = {
    PRIVATE: "PRIVATE",
    PUBLIC: "PUBLIC"
};

// Función auxiliar para hacer las peticiones si no está definida en request.js
async function makeRequest(url, method, data, contentType, callType, onSuccess, onError) {
    try {
        console.log("Realizando petición a:", url);
        const options = {
            method: method,
            headers: {
                "Content-Type": contentType,
                ...customHeaders
            },
            body: data ? JSON.stringify(data) : null
        };

        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const responseData = await response.json();
        console.log("Respuesta exitosa:", responseData);
        onSuccess(responseData);
    } catch (error) {
        console.error("Error en la petición:", error);
        onError(error);
    }
}

function volverAProductos() {
  // Navegar a la página de productos
  window.location.href = '/web/productos/index_producto.html';
}