package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)

func (hnd *Handler) Motherships(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("mother_ships","mother_ships.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"mother_ships",nil)
}
