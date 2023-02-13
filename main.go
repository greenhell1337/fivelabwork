package main

import (
	"net/http"

	"github.com/greenhell1337/shopForDB/cmd/web"
	cfs "github.com/greenhell1337/shopForDB/pkg/closeFileSystem"
	ce "github.com/greenhell1337/shopForDB/pkg/checkerr"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(cfs.CloseFileSystem{Fs: http.Dir("./ui/static/")})

	mux.HandleFunc("/main", web.Main)
	mux.HandleFunc("/tea_list", web.TeaList)
	mux.HandleFunc("/about", web.About)
	mux.HandleFunc("/sign_up", web.SignUp)
	mux.HandleFunc("/log_in", web.LogIn)
	mux.HandleFunc("/basket", web.Basket)

	mux.HandleFunc("/tea_list/china_tea", web.ChinaList)
	mux.HandleFunc("/tea_list/bags_tea", web.BagsList)
	mux.HandleFunc("/tea_list/pyramids_tea", web.PyramidsList)
	mux.HandleFunc("/tea_list/soothing_tea", web.SoothingList)

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8080", mux)
	ce.CheckErr(err)
}