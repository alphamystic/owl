package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)


func (hnd *Handler) Crm_Issues(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("issues","issues.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"issues",nil)
}


func (hnd *Handler) Client_Appointments(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("appointments","appointments.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"appointments",nil)
}

// set to appointment-detail and should take in parameters
func (hnd *Handler) Client_Appointments_Detail(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetClientDash("appointment_detail","appointment_detail.tmpl")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"appointment_detail",nil)
}
