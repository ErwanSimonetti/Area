/** @file emailController.go
 * @brief This file contain a functions for the email API
 * @author Juliette Destang
 * @version
 */

// @cond

package controllers

import (
	"net/http"
	"strconv"

	"AREA/pkg/models"
)

// @endcond

/** @brief on a request, store the given email and password to the database.
 * @param w http.ResponseWriter, r *http.Request
 */
func AuthEmail(w http.ResponseWriter, r *http.Request) {
  email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	requestUser, _ := GetUser(w, r)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email", email)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email_password", password)
}