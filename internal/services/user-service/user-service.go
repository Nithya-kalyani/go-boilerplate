package services

import (
	"encoding/json"
	"net/http"

	"github.com/Nithya-kalyani/go-boilerplate/internal/models"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/utils"
	"gorm.io/gorm"
)

// @Summary List Users
// @Description Get a list of all users
// @Tags User
// @Produce json
// @Success 200 {array} map[string]string
// @Router /user/list [get]
func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []map[string]string{
		{"id": "1", "name": "John"},
		{"id": "2", "name": "Peter"},
	}
	utils.RenderJSON(w, http.StatusOK, users)
}

// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /user/create [post]
var DB *gorm.DB

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RenderJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	// Save user to the database
	if err := DB.Create(&user).Error; err != nil {
		utils.RenderJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}

	utils.RenderJSON(w, http.StatusCreated, user)
}
