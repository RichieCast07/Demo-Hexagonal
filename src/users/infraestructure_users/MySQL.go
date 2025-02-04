package infraestructure

import (
	"log"

	"demo/src/core"
	domain "demo/src/users/domain_users"
)

type UserRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewUserRepoMySQL() *UserRepoMySQL {
	conn := core.MySQLConection()
	if conn.Err != "" {
		log.Fatal("Error al configurar la pool de conexiones", conn.Err)
	}
	return &UserRepoMySQL{Connection: *conn}
}

func (r *UserRepoMySQL) SaveUserFunction(user domain.User) error {
	query := "INSERT INTO Users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Fatalf("Error al registrar Usuarios:", err)
	}
	return err
}

func (r *UserRepoMySQL) GetUserFunction(id int32) ([]domain.User, error) {
	query := "SELECT id, name, email FROM Users WHERE id = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var users []domain.User
	if err != nil {
		log.Fatalf("Error al obtener Usuarios:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int32
		var name string
		var email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Println("Error al escanear la fila:", err)
		}
		user := domain.User{ID: id, Name: name, Email: email, Password: ""}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepoMySQL) EditUserFunction(user domain.User) error {
	query := "UPDATE Users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		log.Fatalf("Error al editar info. dl usuario:", err)
	}
	return err
}

func (r *UserRepoMySQL) DeleteUserFunction(id int32) error {
	_, err := r.Connection.ExecPreparedQuerys("DELETE FROM Users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Error al eliminar usuario:", err)
	}
	return err
}
