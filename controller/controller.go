package controller

import (
	"net/http"
	"simple-template/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetUserController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []model.UserResponse
		query := `
			SELECT id, name, address FROM users
		`
		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		for rows.Next() {
			var user model.UserResponse
			err = rows.Scan(
				&user.Id,
				&user.Name,
				&user.Address,
			)
			if err != nil {
				return err
			}
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    users,
		})
	}
}

func GetUserControllerById(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users model.UserResponse
		Id := c.Param("id")
		query := `
			SELECT id, name, address FROM users WHERE id = $1
		`
		rows, err := db.Query(query, Id)
		if err != nil {
			return err
		}
		if rows.Next() {
			err = rows.Scan(
				&users.Id,
				&users.Name,
				&users.Address,
			)
			if err != nil {
				return err
			}
		} else {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "user not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "get data user by id",
			"data":    users,
		})
	}
}

func CreateUserController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users model.UserResponse
		var req model.UserRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		query := `
			INSERT INTO users (name, address) VALUES ($1, $2) RETURNING name, address
		`
		rows := db.QueryRow(query, req.Name, req.Address)
		err = rows.Scan(
			&users.Name,
			&users.Address,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success add data",
			"data":    users,
		})
	}
}

func UpdateUserController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.UserRequest
		var user model.UserResponse
		Id := c.Param("id")
		err := c.Bind(&req)

		if err != nil {
			return err
		}
		query := `UPDATE users SET name = $1 WHERE id = $2 RETURNING name, address`

		rows := db.QueryRow(query, req.Name, Id)
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Address,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sucessfully updated",
			"data":    user,
		})
	}
}

func DeleteUserController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.UserRequest
		Id := c.Param("id")
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		query := `DELETE FROM users WHERE id = $1`

		_, err = db.Exec(query, Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "user not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sucessfully deleted",
		})
	}
}

func BulkDeleteUserController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.BulkDelete
		err := c.Bind(&req)

		if err != nil {
			return err
		}

		for _, Id := range req.Id {
			query := `DELETE FROM users WHERE id = $1`
			_, err = db.Exec(query, Id)
			if err != nil {
				return err
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Multiple item deleted successfully",
		})
	}
}
