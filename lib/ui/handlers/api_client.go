package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)


func (hnd *Handler) ClientListapikeys(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("client_list_api_key","client_list_api_key.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"client_list_api_key",nil)
}

func (hnd *Handler) CreateApikeys(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("client_create_api_key","client_create_api_key.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"client_create_api_key",nil)
}
