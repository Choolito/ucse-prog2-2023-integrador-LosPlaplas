const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const baseUrl = "http://localhost:8080/productos";

const TIPOS_PRODUCTO = {
    GOLOSINAS: "Golosinas",
    BEBIDAS: "Bebidas",
    CIGARRILLOS: "Cigarrillos",
    COMESTIBLES: "Comestibles",
    HIGIENE_Y_SALUD: "Higiene y Salud",
    SIN_FILTRO: "" // Para traer todos los productos
};

document.addEventListener("DOMContentLoaded", function () {
    obtenerProductosStockMinimo(TIPOS_PRODUCTO.SIN_FILTRO);
});

let productosStockMinimo = []; // Variable global

function obtenerProductosStockMinimo(tipoProducto = "") {
    let tipoProductoCorregido = Object.values(TIPOS_PRODUCTO).includes(tipoProducto) ? tipoProducto : "";

    console.log("📌 TipoProducto corregido:", tipoProductoCorregido);

    const bodyData = { TipoProducto: tipoProductoCorregido };

    makeRequest(
        `${baseUrl}/stockminimo`,
        Method.POST,
        bodyData,
        ContentType.JSON,
        CallType.PUBLIC,
        (respuesta) => {
            console.log("✅ Productos recibidos:", respuesta);
            if (!respuesta || respuesta.length === 0) {
                mostrarMensaje("No hay productos en esta categoría.");
                return;
            }
            productosStockMinimo = respuesta;
            mostrarProductos(productosStockMinimo);
        },
        errorObtenerProductos
    );
}

function filtrarProductos() {
    let filtroSelect = document.getElementById("filtroSelect").value.trim();
    let filtroCorregido = TIPOS_PRODUCTO[filtroSelect.toUpperCase().replace(/\s+/g, "_")] || "";
    console.log("Aplicando filtro corregido:", filtroCorregido);
    obtenerProductosStockMinimo(filtroCorregido);
}

function quitarFiltro() {
    console.log("Quitando filtros...");
    document.getElementById("filtroSelect").value = "";
    obtenerProductosStockMinimo(TIPOS_PRODUCTO.SIN_FILTRO);
}

function mostrarMensaje(mensaje) {
    const mensajeContainer = document.getElementById("mensaje-container");
    const tabla = document.getElementById("elementosTable");

    if (!mensajeContainer) {
        console.error("Elemento 'mensaje-container' no encontrado en el DOM.");
        return;
    }

    mensajeContainer.innerHTML = `<p>${mensaje}</p>`;
    mensajeContainer.style.display = "block";
    if (tabla) tabla.style.display = "none";
}

function mostrarProductos(lista) {
    console.log("Mostrando productos:", lista);
    const elementosTable = document.getElementById("elementosTable");
    const tbody = elementosTable ? elementosTable.querySelector("tbody") : null;
    const mensajeContainer = document.getElementById("mensaje-container");

    if (!tbody) {
        console.error("No se encontró el tbody de la tabla.");
        return;
    }

    tbody.innerHTML = "";

    if (!lista || lista.length === 0) {
        mostrarMensaje("No se encontraron productos con stock mínimo");
        return;
    }

    lista.forEach((elemento) => {
        if (elemento.cantidadEnStock <= elemento.stockMinimo) {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${elemento.codigoProducto || ""}</td>
                <td>${elemento.nombre || ""}</td>
                <td>${elemento.tipoProducto || ""}</td>
                <td>${elemento.precioUnitario || 0}</td>
                <td>${elemento.pesoUnitario || 0}</td>
                <td>${elemento.stockMinimo || 0}</td>
                <td>${elemento.cantidadEnStock || 0}</td>
            `;
            tbody.appendChild(row);
        }
    });

    if (tbody.children.length === 0) {
        mostrarMensaje("No se encontraron productos con stock mínimo");
    } else {
        if (elementosTable) elementosTable.style.display = "table";
        if (mensajeContainer) mensajeContainer.style.display = "none";
    }
}

function errorObtenerProductos(error) {
    console.error("Error al obtener productos:", error);

    if (error && typeof error === "string" && error.includes("access_token")) {
        alert("Error de autenticación: No tienes acceso. Verifica tu sesión.");
    } else {
        alert("Error en la solicitud al servidor.");
    }

    mostrarMensaje("Error al cargar los productos. Intenta nuevamente.");
}

function volverAProductos() {
    window.location.href = "/web/productos/index_producto.html";
}

function actualizarInterfazProductos(productos) {
    const tabla = document.getElementById("elementosTable");
    const mensaje = document.getElementById("mensaje-container");

    if (!tabla || !mensaje) {
        console.error("No se encontraron los elementos en el DOM.");
        return;
    }

    if (productos.length === 0) {
        tabla.style.display = "none";
        mensaje.style.display = "block";
    } else {
        tabla.style.display = "table";
        mensaje.style.display = "none";
    }
}