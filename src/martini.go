package src

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"fmt"
	"encoding/base64"
)

func StartMartini()  {
	m := martini.Classic()
	m.Use(martini.Static("public"))
	m.Use(render.Renderer(render.Options{
		Extensions: []string{".tmpl", ".html"},
	}))
	m.Get("/", func(r render.Render) {
		var test struct{
			Title,Name,Age,Data string
		}
		query := Pay(1, "123456", 1.00)
		data,err := RsaEncrypt([]byte(query));
		if(err != nil){
			fmt.Printf("%s", err);
		}
		test.Name = "tingfeng"
		test.Title = "hello world"
		test.Age = "18"
		test.Data = "http://192.168.3.5:8080/api/merchOrder?"+query+"&sign="+base64.StdEncoding.EncodeToString(data);
		r.HTML(200, "index", test)
	})
	m.Run()
}
