package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

type Url struct {
	Id  int64
	Url string
}

func handleCreateShortUrl(c *gin.Context) {
	originalUrl := c.Query("url")
	if originalUrl == "" {
		c.String(http.StatusBadRequest, "No url provided")
		return
	}
	log.Println(originalUrl)
	// try to find an existing url
	findUrlQuery := `SELECT * FROM Url WHERE url = ?`
	var foundUrl Url

	err := db.QueryRow(findUrlQuery, originalUrl).Scan(&foundUrl.Id, &foundUrl.Url)
	if err != nil {
		log.Println("Url: ", originalUrl, " not found, will create a db row")
		createQuery := `INSERT INTO Url (url) VALUES (?)`
		res, err := db.Exec(createQuery, originalUrl)
		if err != nil {
			log.Fatal("Failed to create a db row")
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal("Failed to get inserted id")
		}
		convertedId := idToBase62(id)
		c.String(http.StatusOK, "http://localhost:5000/"+string(convertedId))
		return
	}
	foundId := foundUrl.Id
	convertedId := idToBase62(foundId)

	c.String(http.StatusOK, "http://localhost:5000/"+string(convertedId))

	//if err != nil {
	//	c.String(http.StatusBadRequest, "Error in request")
	//}
	//
	//if foundUrl != "" {
	//	c.String(http.StatusOK, foundUrl)
	//	return
	//}
}

func handleRedirectById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.Redirect(http.StatusOK, "http://localhost:3000/")
	}
	convertedId := base62ToId(id)
	findQuery := `SELECT * FROM Url WHERE id = ?`
	var foundUrl Url
	err := db.QueryRow(findQuery, convertedId).Scan(&foundUrl.Id, &foundUrl.Url)

	if err != nil {
		c.Redirect(http.StatusFound, "http://localhost:3000/no-url")
		return
	}
	c.Redirect(http.StatusFound, foundUrl.Url)
}

func initDb() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
}

func closeDb() {
	err := db.Close()
	if err != nil {
		log.Fatal("error closing db: ", err)
	}
	log.Println("DB connection closed")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}
	initDb()
	//defer closeDb()
	// Print a log message to indicate that the server is starting
	log.Println("Starting server...")
	// Create a new router using the Gin framework
	router := gin.Default()

	// Use the CORS middleware to allow cross-origin requests
	router.Use(cors.Default())
	// Register the POST route for creating short URLs
	router.POST("/", handleCreateShortUrl)
	router.GET("/:id", handleRedirectById)
	// Start the server on port 5000
	if err := router.Run(":5000"); err != nil {
		// Log an error and exit if the server fails to start
		log.Fatal("Failed to start the server")
	}
}
