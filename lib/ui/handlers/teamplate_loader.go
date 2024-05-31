package handlers

import (
	"fmt"
	"strings"
	"io/ioutil"
	"html/template"
)

/*
var PAGES []{string}{}

func (hnd *Handler) LoadPages(){
	// read the whole Directory
	pages,err := ioutil.ReadDir(hnd.TemplatesDir + "pages/")
	if err != nil {
		return fmt.Errorf("Error reading body Directory: %q",err)
	}
	// For each, read the pages and index it with the name.
	for _, page := range pages {
		name,data
	}
}
*/
// Creates a whole new base template for serving to clients
func (hnd *Handler) LoadBase() error {
	baseHTML, err := ioutil.ReadFile(hnd.TemplatesDir +"base.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading base template: %q",err)
	}
	clientBase, err := ioutil.ReadFile(hnd.TemplatesDir +"client_dash.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading base template: %q",err)
	}
	navigation_bar, err := ioutil.ReadFile(hnd.TemplatesDir + "navigation_bar.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading navigation bar template: %q",err)
	}
	sidebar, err := ioutil.ReadFile(hnd.TemplatesDir + "sidebar.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading sidebar template: %q",err)
	}
	notifications, err := ioutil.ReadFile(hnd.TemplatesDir + "notifications.tmpl")
	if err != nil {
		return  fmt.Errorf("Error loading notifications template: %q",err)
	}
	shortcuts, err := ioutil.ReadFile(hnd.TemplatesDir + "shortcuts.tmpl")
	if err != nil {
		return  fmt.Errorf("Error loading footer template: %q",err)
	}

	// Replace placeholders in the base HTML with actual content
	combinedHTML := strings.ReplaceAll(string(baseHTML), "{{.NAVBAR}}", string(navigation_bar))
	combinedHTML = strings.ReplaceAll(combinedHTML, "{{.SIDEBAR}}", string(sidebar))
	combinedHTML = strings.ReplaceAll(combinedHTML, "{{.NOTIFICATIONS}}", string(notifications))
	combinedHTML = strings.ReplaceAll(combinedHTML, "{{.SHORTCUTS}}", string(shortcuts))
	hnd.Base = combinedHTML

	combinedDash := strings.ReplaceAll(string(clientBase), "{{.NAVBAR}}", string(navigation_bar))
	combinedDash = strings.ReplaceAll(combinedDash, "{{.SIDEBAR}}", string(sidebar))
	combinedDash = strings.ReplaceAll(combinedDash, "{{.NOTIFICATIONS}}", string(notifications))
	combinedDash = strings.ReplaceAll(combinedDash, "{{.SHORTCUTS}}", string(shortcuts))
	hnd.ClientDash = combinedDash

	/*/var tpl = new(template.Template)
	var tpl = template.New("base")
	tpl,err = tpl.Parse(string(combinedHTML))
	if err != nil {
		return fmt.Errorf("Error parsing combined html to a template: %q",err)
	}
	hnd.Tpl = tpl*/
	return  nil
}

func (hnd *Handler) GetATemplate(name,templFile string) (*template.Template,error) {
	body,err := ioutil.ReadFile(hnd.TemplatesDir + "pages/" + templFile)
	if err != nil {
		return nil,fmt.Errorf("Error getting template %s: %q",templFile,err)
	}
	sTpl := strings.ReplaceAll(hnd.Base,"{{.BODY}}",string(body))
	var tpl = template.New(name)
	tpl,err = tpl.Parse(string(sTpl))
	if err != nil{
		return nil,fmt.Errorf("Error parsing body to template: %q",err)
	}
	return tpl,nil
}

// get a dash in question and fill it into the admin dash template
func (hnd *Handler) GetClientDash(name,templFile string) (*template.Template,error) {
	body,err := ioutil.ReadFile(hnd.TemplatesDir + "pages/" + templFile)
	if err != nil {
		return nil,fmt.Errorf("Error getting template %s: %q",templFile,err)
	}
	sTpl := strings.ReplaceAll(hnd.ClientDash,"{{.BODY}}",string(body))
	var tpl = template.New(name)
	tpl,err = tpl.Parse(string(sTpl))
	if err != nil{
		return nil,fmt.Errorf("Error parsing body to template: %q",err)
	}
	return tpl,nil
}

/* admin has different dash's, will do that later
func (hnd *Handler) GetAdminDash(name,templFile string) (*template.Template,error) {
	body,err := ioutil.ReadFile(hnd.TemplatesDir + "pages/" + templFile)
	if err != nil {
		return nil,fmt.Errorf("Error getting template %s: %q",templFile,err)
	}
	sTpl := strings.ReplaceAll(hnd.Dash,"{{.BODY}}",string(body))
	var tpl = template.New(name)
	tpl,err = tpl.Parse(string(sTpl))
	if err != nil{
		return nil,fmt.Errorf("Error parsing body to template: %q",err)
	}
	return tpl,nil
}
*/

func (hnd *Handler) CombineToBase(data string) string {
	newBase := strings.ReplaceAll(hnd.Base,"{{.BODY}}",data)
	return newBase
}

func (hnd *Handler) CombineAdminToBase(data string) string {
	newBase := strings.ReplaceAll(hnd.AdminBase,"{{.BODY}}",data)
	return newBase
}


func (hnd *Handler) LoadATemplate(name,templFile string) (*template.Template,error) {
	body,err := ioutil.ReadFile(hnd.TemplatesDir + templFile)
	if err != nil {
		return nil,fmt.Errorf("Error getting template %s: %q",templFile,err)
	}
	//sTpl := strings.ReplaceAll(hnd.Base,"{{.BODY}}",string(body))
	var tpl = template.New(name)
	tpl,err = tpl.Parse(string(body))
	if err != nil{
		return nil,fmt.Errorf("Error parsing body to template: %q",err)
	}
	return tpl,nil
}
