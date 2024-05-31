package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)



func (hnd *Handler) Blank(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("blank","blank.tmpl")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"blank",nil)
}


func (hnd *Handler) Notfound(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.LoadATemplate("404","404.html")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"404",nil)
}
