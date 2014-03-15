package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Server struct {
	storage    *Storage
	name, addr string
}

func NewServer(storage *Storage, name, addr string) *Server {
	return &Server{storage, name, addr}
}

// Root handler of bin. Validates and routes requests
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("http:", r.RemoteAddr, r.Method, r.URL.RequestURI())

	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			// "main" page
			s.HandleMainPage(w, r)

		} else {
			// paste page
			s.HandleGetPaste(w, r)
		}
	case "POST":
		// handle upload
		s.HandleUpload(w, r)
	default:
		http.Error(w, "", 405)
	}
}

func (s *Server) HandleMainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func (s *Server) HandleGetPaste(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if len(id) != ID_LENGTH {
		http.NotFound(w, r)
	}

	file, err := s.storage.Open(id)
	if err != nil {
		log.Print("http: failed to open paste: ", err)
		http.Error(w, "internal error\n", 500)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		log.Print("http: failed to deliver paste: ", err)
	}
}

func (s *Server) HandleUpload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, file, err := s.storage.Create()
	if err != nil {
		log.Print("upload: Failed to create paste: ", err)
		http.Error(w, "internal error", 500)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Print("upload: Failed to fill paste: ", err)
		http.Error(w, "internal error", 500)
		return
	}

	var host string
	if s.name == "" {
		host = r.Header.Get("Host")
		if host == "" {
			if s.addr[0] == ':' {
				host = "localhost" + s.addr
			} else {
				host = s.addr
			}
		}
	} else {
		host = s.name
	}

	pasteURL := &url.URL{Scheme: "http", Host: host, Path: "/" + id}

	_, err = w.Write([]byte(pasteURL.String() + "\n"))
	if err != nil {
		log.Print("upload: client does not received paste url: ", err)
	}
}
