package webserver

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	fileRoot string
}

func (s *Server) ListenAndServe() error {
	fmt.Println("Start to Serve.")
	s.fileRoot = "/../pkg/webserver/docs"
	fmt.Println("fileRoot = ", s.fileRoot)
	http.HandleFunc("/", s.HandleRequest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) HandleRequest(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if is200(req, s.fileRoot) {
		s.Write200(res, req)
	} else {
		s.Write404(res)
	}
}

func is200(req *http.Request, fileRoot string) bool {
	path := fileRoot + req.URL.Path
	if path[len(path)-1] == '/' {
		path += "index.html"
	}
	path = simplifyPath(path)
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsRegular()
}

func simplifyPath(path string) string {
	stack := make([]string, 0)
	ret := strings.Split(path, "/")
	for _, v := range ret {

		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "" && v != "." {
			stack = append(stack, v)
		}
	}
	rst := ""
	if len(stack) == 0 {
		rst = "/"
	} else {
		for _, v := range stack {
			if len(rst) != 0 {
				rst += "/"
			}
			rst += v
		}
	}
	return rst
}

func (s *Server) Write200(res http.ResponseWriter, req *http.Request) {
	path := s.fileRoot + req.URL.Path
	if path[len(path)-1] == '/' {
		path += "index.html"
	}
	path = simplifyPath(path)
	f, _ := os.ReadFile(path)
	res.WriteHeader(200)
	res.Write(f)
}

func (s *Server) Write404(res http.ResponseWriter) {
	res.WriteHeader(404)
}
