package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

type Page struct {
	Title    string
	Articles []string
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/location/:id", getLocation)
	router.GET("/episode/:id", GetEpisodeHandler)
	router.GET("/character/:id", getCharacter)
	router.GET("/locations", getLocations)
	router.GET("/episodes", getEpisodes)
	router.GET("/characters", getCharacters)

	router.Run("localhost:8080")

}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Création d'une page
	p := Page{"Titre de ma page", []string{"Item 1", "Item 2", "Item 3"}}

	// Création d'une nouvelle instance de template
	t := template.New("Label de ma template")

	// Déclaration des fichiers à parser
	t = template.Must(t.ParseFiles("tmpl/layout.tmpl", "tmpl/content.tmpl"))

	// Exécution de la fusion et injection dans le flux de sortie
	// La variable p sera réprésentée par le "." dans le layout
	// Exemple {{.}} == p
	err := t.ExecuteTemplate(w, "layout", p)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func getLocation(c *gin.Context) {
	var toto Page
	toto.Articles = append(toto.Articles, "data", "data2", "data3")
	toto.Title = "myfirsttitle"
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	value1, err := GetLocation(id)

	golog.Info(value1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		// c.JSON(http.StatusOK, value1)
		c.HTML(http.StatusOK, "layout.tmpl", gin.H{
			"Title":     toto.Title,
			"Name":      value1.Name,
			"Created":   value1.Created,
			"Dimension": value1.Dimension,
			"Residents": value1.Residents,
			"Type":      value1.Type,
			"URL":       value1.URL,
		})
	}

}
func GetEpisodeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	value1, err := GetEpisode(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.HTML(http.StatusOK, "layout.tmpl", gin.H{
			"Title":     "title of page",
			"Name":      value1.Name,
			"Created":   value1.Created,
			"Dimension": value1.Characters,
			"Residents": value1.Episode,
			"Type":      value1.Air_Date,
			"URL":       value1.URL,
		})
	}

}
func getCharacter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	value1, err := GetCharacter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.HTML(http.StatusOK, "toto", value1)
	} else {
		c.HTML(http.StatusOK, "layout.tmpl", gin.H{
			"Title":     "title of page",
			"Name":      value1.Name,
			"Created":   value1.Created,
			"Dimension": value1.Episode,
			"Residents": value1.Gender,
			"Type":      value1.Image,
			"URL":       value1.URL,
		})
	}

}

func getLocations(c *gin.Context) {
	ids, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	golog.Info(ids)
	toto := []int{12, 13, 14}
	value1, err := GetLocationsArray(toto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, value1)
	}

}

func getEpisodes(c *gin.Context) {
	ids, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	golog.Info(ids)
	toto := []int{12, 13, 14}
	value1, err := GetEpisodesArray(toto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, value1)
	}

}
func getCharacters(c *gin.Context) {
	ids, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	golog.Info(ids)
	toto := []int{12, 13, 14}
	value1, err := GetCharactersArray(toto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, value1)
	}

}
