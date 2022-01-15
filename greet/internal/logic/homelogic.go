package logic

import (
	"context"
	"go-zero-demo/greet/internal/svc"
	"log"
	"net/http"
	"text/template"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeLogic {
	return HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home(w http.ResponseWriter) error {
	// todo: add your logic here and delete this line
	FileDemo(w)
	return nil
}

func FileDemo(w http.ResponseWriter) {
	templates := loadTemplates()

	// fileName := r.URL.Path[1:]
	t := templates.Lookup("home.html")
	// t := templates.Lookup(fileName)
	if t != nil {
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))
	// http.Handle("/", http.StripPrefix("/home", http.FileServer(http.Dir("wwwroot/css/"))))
	// http.Handle("/", http.StripPrefix("/home", http.FileServer(http.Dir("wwwroot/img/"))))

}

func loadTemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob("templates/*.html"))

	return result
}
