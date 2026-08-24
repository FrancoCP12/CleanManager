package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Producto struct {
	ID          int
	Nombre      string
	SKU         string
	Precio      float64
	StockActual int
}

func main() {
	// 1. Conexión a MySQL
	dsn := "root:secreto@tcp(127.0.0.1:3306)/inventario_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	fmt.Println("¡Conexión exitosa a MySQL! 🚀")

	// --- RUTA 1: Listar productos (Página principal) ---
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nombre, sku, precio, stock_actual FROM productos")
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var productos []Producto
		for rows.Next() {
			var p Producto
			err := rows.Scan(&p.ID, &p.Nombre, &p.SKU, &p.Precio, &p.StockActual)
			if err != nil {
				log.Println("Error leyendo fila:", err)
				continue
			}
			productos = append(productos, p)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, "Error al procesar los productos", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error al cargar el template HTML", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, productos)
	})

	// --- RUTA 2: Guardar un nuevo producto enviado desde el formulario ---
	http.HandleFunc("/agregar", func(w http.ResponseWriter, r *http.Request) {
		// Validamos que la petición sea de tipo POST (que venga del formulario)
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Parseamos los datos que vienen del formulario HTML
		r.ParseForm()
		nombre := r.FormValue("nombre")
		sku := r.FormValue("sku")

		// Convertimos precio (string a float64) y stock (string a int)
		precio, err := strconv.ParseFloat(r.FormValue("precio"), 64)
		if err != nil {
			http.Error(w, "Precio inválido", http.StatusBadRequest)
			return
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			http.Error(w, "Stock inválido", http.StatusBadRequest)
			return
		}

		// Ejecutamos la consulta de inserción en MySQL de forma segura (usando ? para evitar inyección SQL)
		query := "INSERT INTO productos (nombre, sku, precio, stock_actual) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(query, nombre, sku, precio, stock)
		if err != nil {
			log.Println("Error al insertar producto:", err)
			http.Error(w, "Error al guardar el producto en la base de datos", http.StatusInternalServerError)
			return
		}

		// Una vez guardado con éxito, redirigimos al usuario de vuelta a la página principal
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// 3. Arrancamos el servidor
	fmt.Println("Servidor web corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// 3. Arrancamos el servidor
	fmt.Println("Servidor web corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
