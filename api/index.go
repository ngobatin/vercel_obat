package api

import (
	"net/http"

	"github.com/mytodolist1/todolist_be/config"
	h "github.com/mytodolist1/todolist_be/handler"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	ngobat "github.com/Febriand1/api_obat"
)

var (
	user     ngobat.User
	obat     ngobat.Obat
	penyakit ngobat.Penyakit
	rs       ngobat.RumahSakit
)

var mconn = config.MongoConnect("MONGOSTRING", "psikofarmaka")

func Handler(w http.ResponseWriter, r *http.Request) {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("_id")

		switch r.URL.Path {
		case "/":
			if r.Method == "GET" {
				h.StatusOK(w, "Welcome to Ngobat API")
				return
			}

		case "/login":
			if r.Method == "POST" {
				if r.ContentLength == 0 {
					h.StatusMethodNotAllowed(w, "Request body is empty")
					return
				}
				err := h.JDecoder(w, r, &user)
				if err != nil {
					h.StatusBadRequest(w, "error parsing application/json: "+err.Error())
					return
				}
				users, _, err := ngobat.Login(mconn, "user", user)
				if err != nil {
					h.StatusBadRequest(w, err.Error())
					return
				}
				h.StatusOK(w, "Selamat Datang "+users.Username)
				return
			}

		case "/register":
			if r.Method == "POST" {
				if r.ContentLength == 0 {
					h.StatusMethodNotAllowed(w, "Request body is empty")
					return
				}
				err := h.JDecoder(w, r, &user)
				if err != nil {
					h.StatusBadRequest(w, "error parsing application/json: "+err.Error())
					return
				}
				err = ngobat.Register(mconn, "user", user)
				if err != nil {
					h.StatusBadRequest(w, err.Error())
					return
				}
				h.StatusCreated(w, "Registrasi Berhasil")
				return
			}

		case "/user":
			if r.Method == "GET" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					user.ID = ID
					user, err := ngobat.GetUserByID(mconn, "user", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get User By ID Success", "data", user)
					return

				} else {
					user, err := ngobat.GetAllUser(mconn, "user")
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get All User Success", "data", user)
					return
				}
			}

		case "/obat":
			if r.Method == "DELETE" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					obat.ID = ID
					_, err = ngobat.DeleteObat(mconn, "obat", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Delete Obat Success")
					return
				}
			}

			if r.Method == "POST" {
				if r.ContentLength == 0 {
					h.StatusMethodNotAllowed(w, "Request body is empty")
					return
				}
				data, err := ngobat.InsertObat(mconn, "obat", r)
				if err != nil {
					h.StatusBadRequest(w, err.Error())
					return
				}
				h.StatusOK(w, "Insert Obat Success", "data", data)
				return
			}

			if r.Method == "PUT" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					obat.ID = ID
					obat, err = ngobat.UpdateObat(mconn, "obat", ID, r)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Update Obat Success", "data", obat)
					return
				}
			}

			if r.Method == "GET" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					obat.ID = ID
					obat, err := ngobat.GetObatByID(mconn, "obat", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get Obat By ID Success", "data", obat)
					return

				} else {
					obat, err := ngobat.GetAllObat(mconn, "obat")
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get All Obat Success", "data", obat)
					return
				}
			}

		case "/penyakit":
			if r.Method == "DELETE" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					penyakit.ID = ID
					_, err = ngobat.DeletePenyakit(mconn, "penyakit", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Delete Penyakit Success")
					return
				}
			}

			if r.Method == "POST" {
				if r.ContentLength == 0 {
					h.StatusMethodNotAllowed(w, "Request body is empty")
					return
				}
				data, err := ngobat.InsertPenyakit(mconn, "penyakit", r)
				if err != nil {
					h.StatusBadRequest(w, err.Error())
					return
				}
				h.StatusOK(w, "Insert Penyakit Success", "data", data)
				return
			}

			if r.Method == "PUT" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					penyakit.ID = ID
					penyakit, err = ngobat.UpdatePenyakit(mconn, "penyakit", ID, r)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Update Penyakit Success", "data", penyakit)
					return
				}
			}

			if r.Method == "GET" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					penyakit.ID = ID
					penyakit, err := ngobat.GetPenyakitByID(mconn, "penyakit", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get Penyakit By ID Success", "data", penyakit)
					return

				} else {
					penyakit, err := ngobat.GetAllPenyakit(mconn, "penyakit")
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get All Penyakit Success", "data", penyakit)
					return
				}
			}

		case "/rs":
			if r.Method == "DELETE" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					rs.ID = ID
					_, err = ngobat.DeleteRS(mconn, "rs", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Delete RS Success")
					return
				}
			}

			if r.Method == "POST" {
				if r.ContentLength == 0 {
					h.StatusMethodNotAllowed(w, "Request body is empty")
					return
				}
				data, err := ngobat.InsertRS(mconn, "rs", r)
				if err != nil {
					h.StatusBadRequest(w, err.Error())
					return
				}
				h.StatusOK(w, "Insert RS Success", "data", data)
				return
			}

			if r.Method == "PUT" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					rs.ID = ID
					rs, err = ngobat.UpdateRS(mconn, "rs", ID, r)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Update RS Success", "data", rs)
					return
				}
			}

			if r.Method == "GET" {
				if id != "" {
					ID, err := primitive.ObjectIDFromHex(id)
					if err != nil {
						h.StatusBadRequest(w, "Invalid '_id' parameter in the URL")
						return
					}
					rs.ID = ID
					rs, err := ngobat.GetRSByID(mconn, "rs", ID)
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get RS By ID Success", "data", rs)
					return

				} else {
					rs, err := ngobat.GetAllRS(mconn, "rs")
					if err != nil {
						h.StatusBadRequest(w, err.Error())
						return
					}
					h.StatusOK(w, "Get All RS Success", "data", rs)
					return
				}
			}
		}
	}))

	handler.ServeHTTP(w, r)
}
