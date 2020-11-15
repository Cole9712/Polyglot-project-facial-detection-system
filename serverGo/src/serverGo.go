package serverGo

import(
	"net/http"
    "log"
	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	img := parseImageJSON(r,w)

}


func main() {
	router := httprouter.New()
	router.GET("/",indexHandler)


	log.Fatal(http.ListenAndServe(":8082",router))
}
