package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "./mock/layout"
var bootstrap *template.Template

type Data struct {
	Title string
	Image string
	Path  string
}

func main() {
	var err error
	bootstrap, err = template.ParseGlob("./mock/layout/*.html")
	if err != nil {
		panic(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./mock/static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/", dashboard)
	http.HandleFunc("/users", users)
	http.HandleFunc("/usersadd", usersadd)
	http.HandleFunc("/userprofile", userprofile)
	http.HandleFunc("/customers", customers)
	http.HandleFunc("/customeradd", customeradd)
	http.HandleFunc("/customerprofile", customerprofile)
	http.HandleFunc("/products", products)
	http.HandleFunc("/productprofile", productprofile)
	http.HandleFunc("/addproduct", addproduct)
	http.HandleFunc("/suppliers", supplier)
	http.HandleFunc("/addsupplier", addsupplier)
	http.HandleFunc("/analytics", analytics)
	http.HandleFunc("/invoices", invoices)
	http.HandleFunc("/addinvoice", addinvoice)
	http.HandleFunc("/modifyinvoice", modifyinvoice)
	http.HandleFunc("/quotations", quotations)
	http.HandleFunc("/addquotation", addquotation)
	http.HandleFunc("/modifyquotation", modifyquotation)
	http.HandleFunc("/settings/general", settings)
	http.HandleFunc("/settings/notification", notificationSettings)
	http.HandleFunc("/settings/authentication", authenticationSettings)
	http.HandleFunc("/documentation", documentation)
	http.ListenAndServe(":8085", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./mock/layout/header.html", "./mock/login.html"))
	err := t.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/dashboard.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Dashboard",
		Path:  "Home",
		Image: "/static/images/dashboard.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/users.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Users",
		Path:  "Home",
		Image: "/static/images/icon_user.jpeg",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func usersadd(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/users_add.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add User",
		Path:  "Home",
		Image: "/static/images/icon_new_user.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userprofile(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/user.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "User Profile",
		Path:  "Home",
		Image: "/static/images/icon_user.jpeg",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func customers(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/customers.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Customers",
		Path:  "Home",
		Image: "/static/images/icon_users.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func customeradd(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/customer_add.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add Customer",
		Path:  "Home",
		Image: "/static/images/icon_new_user.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func customerprofile(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/customer.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Customer Profile",
		Path:  "Home",
		Image: "/static/images/icon_profile.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func products(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/products.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Products",
		Path:  "Home",
		Image: "/static/images/icon_product.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func productprofile(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/product.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Product",
		Path:  "Home",
		Image: "/static/images/icon_product.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addproduct(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/addproduct.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add Product",
		Path:  "Home",
		Image: "/static/images/icon_product.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func supplier(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/supplier.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Suppliers",
		Path:  "Home",
		Image: "/static/images/icon_supplier.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addsupplier(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/addsupplier.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add Supplier",
		Path:  "Home",
		Image: "/static/images/icon_supplier.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func analytics(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/analytics.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Analytics",
		Path:  "Home",
		Image: "/static/images/icon_analytics.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func invoices(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/invoices.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Invoices",
		Path:  "Home",
		Image: "/static/images/icon_invoices.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addinvoice(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/addinvoice.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add Invoice",
		Path:  "Home",
		Image: "/static/images/icon_invoices.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func modifyinvoice(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/modify_invoice.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Modify Invoice",
		Path:  "Home",
		Image: "/static/images/icon_invoices.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func quotations(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/quotations.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Quotations",
		Path:  "Home",
		Image: "/static/images/icon_quotation.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addquotation(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/addquotation.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Add Quotation",
		Path:  "Home",
		Image: "/static/images/icon_quotation.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func modifyquotation(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/modifyquotation.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Modify Quotation",
		Path:  "Home",
		Image: "/static/images/icon_quotation.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/settings.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "General Settings",
		Path:  "Home",
		Image: "/static/images/icon_general_settings.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notificationSettings(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/notificationSettings.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Notification Settings",
		Path:  "Home",
		Image: "/static/images/icon_notification.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func authenticationSettings(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/authenticationSettings.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Authentication Settings",
		Path:  "Home",
		Image: "/static/images/icon_authentication.jpeg",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func documentation(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/documentation.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", Data{
		Title: "Documentation",
		Path:  "Home",
		Image: "/static/images/icon_documentation.png",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.html")
	if err != nil {
		panic(err)
	}
	return files
}
