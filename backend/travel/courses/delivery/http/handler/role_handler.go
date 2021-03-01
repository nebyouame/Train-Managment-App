package handler

import (
	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"

	"html/template"

	"net/http"

	"strconv"
)

type AdminRoleHandler struct {
	tmp1 *template.Template
	roleSrv schedule.RoleService
}

func NewAdminRoleHandler(ss schedule.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleSrv:ss}
}

func (arh *AdminRoleHandler) AdminRoles(w http.ResponseWriter, r *http.Request){
	roles, errs := arh.roleSrv.Roles()
	if errs != nil {
		panic(errs)
	}
	arh.tmp1.ExecuteTemplate(w, "admin.roles.layout", roles)
}

func (arh *AdminRoleHandler) AdminRolesNew(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost {
		sch := &model.Role{}
		sch.Name = r.FormValue("name")



		_, errs := arh.roleSrv.StoreRole(sch)

		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/roles", http.StatusSeeOther)
	} else {
		arh.tmp1.ExecuteTemplate(w, "admin.roles.new.layout", nil)
	}
}

func (arh *AdminRoleHandler) AdminRolesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method ==http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		sat, errs := arh.roleSrv.Role(uint(id))

		if len(errs) > 0 {
			panic(errs)
		}
		arh.tmp1.ExecuteTemplate(w, "admin.roles.update.layout", sat)
	}  else if r.Method == http.MethodPost {
		sch := &model.Role{}
		id,_ := strconv.Atoi(r.FormValue("id"))
		sch.ID= uint(id)
		sch.Name= r.FormValue("name")


		_, errs := arh.roleSrv.UpdateRole(sch)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/roles", http.StatusSeeOther)
	} else {
		http.Redirect(w,r, "/admin/roles", http.StatusSeeOther)
	}
}

func (arh *AdminRoleHandler) AdminRolesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := arh.roleSrv.DeleteRole(uint(id))
		if len(errs) > 0 {
			panic(err)
		}
	}
	http.Redirect(w,r, "/admin/roles", http.StatusSeeOther)
}
