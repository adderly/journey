package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dimfeld/httptreemux/v5"

	"github.com/kabukky/journey/database"
	"github.com/kabukky/journey/filenames"
	"github.com/kabukky/journey/structure/methods"
	"github.com/kabukky/journey/templates"
)

func indexHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	number := params["number"]
	if number == "" {
		// Render index template (first page)
		err := templates.ShowIndexTemplate(w, r, 1)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	}
	page, err := strconv.Atoi(number)
	if err != nil || page <= 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// Render index template
	err = templates.ShowIndexTemplate(w, r, page)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			e404 := templates.ShowPostTemplate(w, r, "404")
			if e404 != nil {
				http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
				log.Println("404:", r.URL)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
		return
	}
}

func authorHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	slug := params["slug"]
	function := params["function"]
	number := params["number"]
	if function == "" {
		// Render author template (first page)
		err := templates.ShowAuthorTemplate(w, r, slug, 1)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	} else if function == "rss" {
		// Render author rss feed
		err := templates.ShowAuthorRss(w, slug)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	}
	page, err := strconv.Atoi(number)
	if err != nil || page <= 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// Render author template
	err = templates.ShowAuthorTemplate(w, r, slug, page)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			e404 := templates.ShowPostTemplate(w, r, "404")
			if e404 != nil {
				http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
				log.Println("404:", r.URL)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
		return
	}
}

func tagHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	slug := params["slug"]
	function := params["function"]
	number := params["number"]
	if function == "" {
		// Render tag template (first page)
		err := templates.ShowTagTemplate(w, r, slug, 1)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	} else if function == "rss" {
		// Render tag rss feed
		err := templates.ShowTagRss(w, slug)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	}
	page, err := strconv.Atoi(number)
	if err != nil || page <= 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// Render tag template
	err = templates.ShowTagTemplate(w, r, slug, page)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			e404 := templates.ShowPostTemplate(w, r, "404")
			if e404 != nil {
				http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
				log.Println("404:", r.URL)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var err error

	slug := params["slug"]
	if slug == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else if slug == "rss" {
		// Render index rss feed
		err := templates.ShowIndexRss(w)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.WriteHeader(http.StatusNotFound)
				e404 := templates.ShowPostTemplate(w, r, "404")
				if e404 != nil {
					http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
					log.Println("404:", r.URL)
				}
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("503:", r.URL)
			return
		}
		return
	}

	// Render post template
	if slug == "404" {
		err = fmt.Errorf("sql: no rows in result set")
	} else {
		err = templates.ShowPostTemplate(w, r, slug)
	}
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			e404 := templates.ShowPostTemplate(w, r, "404")
			if e404 != nil {
				http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
				log.Println("404:", r.URL)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
		return
	}
}

func postEditHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	slug := params["slug"]

	if slug == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	// Redirect to edit
	post, err := database.RetrievePostBySlug(slug)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
			e404 := templates.ShowPostTemplate(w, r, "404")
			if e404 != nil {
				http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
				log.Println("404:", r.URL)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
		return
	}

	url := fmt.Sprintf("/admin/#/edit/%d", post.Id)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func assetsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// Read lock global blog
	methods.Blog.RLock()
	defer methods.Blog.RUnlock()
	http.ServeFile(w, r, filepath.Join(filenames.ThemesFilepath, methods.Blog.ActiveTheme, "assets", params["filepath"]))
}

func imagesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	http.ServeFile(w, r, filepath.Join(filenames.ImagesFilepath, params["filepath"]))
}

func publicHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	http.ServeFile(w, r, filepath.Join(filenames.PublicFilepath, params["filepath"]))
}

func staticHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	filePath := filepath.Join(filenames.StaticFilepath, r.URL.Path)
	if _, err := os.Stat(filePath); err == nil {
		http.ServeFile(w, r, filePath)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("503:", r.URL)
	}
}

// InitializeBlog serving main blog
func InitializeBlog(router *httptreemux.TreeMux) {
	// For index
	router.GET("/", indexHandler)
	router.GET("/:slug/edit", postEditHandler)
	router.GET("/:slug/", postHandler)
	router.GET("/page/:number/", indexHandler)
	// For author
	router.GET("/author/:slug/", authorHandler)
	router.GET("/author/:slug/:function/", authorHandler)
	router.GET("/author/:slug/:function/:number/", authorHandler)
	// For tag
	router.GET("/tag/:slug/", tagHandler)
	router.GET("/tag/:slug/:function/", tagHandler)
	router.GET("/tag/:slug/:function/:number/", tagHandler)
	// For serving asset files
	router.GET("/assets/*filepath", assetsHandler)
	router.GET("/images/*filepath", imagesHandler)
	router.GET("/content/images/*filepath", imagesHandler) // This is here to keep compatibility with Ghost
	router.GET("/public/*filepath", publicHandler)
	// For static files
	router.GET("/favicon.ico", staticHandler)
	router.GET("/robots.txt", staticHandler)

	// 404
	router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		e404 := templates.ShowPostTemplate(w, r, "404")
		if e404 != nil {
			http.Error(w, "Nobody here but us chickens!", http.StatusNotFound)
			log.Println("404:", r.URL)
		}
	}
}
