package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)


func (hnd *Handler) Dash(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("dash","dash.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"dash",nil)
}



func (hnd *Handler) Threatintel(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("dash","threat_intel.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"dash",nil)
}
