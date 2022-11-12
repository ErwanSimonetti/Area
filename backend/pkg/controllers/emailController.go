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
	"io"

	"AREA/pkg/models"
	"github.com/tidwall/gjson"
)

// @endcond

/** @brief on a request, store the given email and password to the database.
 * @param w http.ResponseWriter, r *http.Request
 */
func AuthEmail(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
    email := gjson.GetBytes(b, "email")
    password := gjson.GetBytes(b, "password")
	requestUser, _ := GetUser(w, r)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email", email.String())
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email_password", password.String())
	w.WriteHeader(http.StatusOK)
}