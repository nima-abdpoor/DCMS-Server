package controllers

import (
	"DCMS/api/helpers"
	"DCMS/db/postgresql/sqlc"
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(util.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func LoginPostHandler(store *db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(util.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}
		customer, err := store.GetCustomerTx(context.Background(), db.GetCustomerTxParams{
			Username: username,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				c.HTML(http.StatusNotFound, "login.html", gin.H{"content": err})
				return
			}
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": err})
			return
		}
		err = util.CheckPassword(password, customer.Customer.Password)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(util.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(util.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(util.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/logout")
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(util.Userkey)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(util.Userkey)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
