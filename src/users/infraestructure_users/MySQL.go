package infraestructure_users

import (
	"log"
	"time"

	"demo/src/core"
	domain "demo/src/users/domain_users"
)

type UserRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewUserRepoMySQL() *UserRepoMySQL {
	conn := core.MySQLConection()
	if conn.Err != "" {
		log.Fatal("Error al configurar la pool de conexiones:", conn.Err)
	}
	return &UserRepoMySQL{Connection: *conn}
}

func (r *UserRepoMySQL) Create(user *domain.User) error {
	return r.SaveUserFunction(*user)
}

func (r *UserRepoMySQL) SaveUserFunction(user domain.User) error {
	query := "INSERT INTO Users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Fatalf("Error al registrar Usuarios: %v", err)
	}
	return err
}

func (r *UserRepoMySQL) GetUserById(id int32) (*domain.User, error) {
	users, err := r.GetUserFunction(id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}

func (r *UserRepoMySQL) GetUserFunction(id int32) ([]domain.User, error) {
	query := "SELECT id, name, email FROM Users WHERE id = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var users []domain.User
	if err != nil {
		log.Fatalf("Error al obtener Usuarios: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var uid int32
		var name, email string
		if err := rows.Scan(&uid, &name, &email); err != nil {
			log.Println("Error al escanear la fila:", err)
		}
		user := domain.User{ID: uid, Name: name, Email: email, Password: ""}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepoMySQL) GetAllUsers() ([]*domain.User, error) {
	query := "SELECT id, name, email, UNIX_TIMESTAMP(updated_at) FROM users"
	rows, err := r.Connection.FetchRows(query)
	if err != nil {
		log.Printf("Error al obtener todos los usuarios: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		var id int32
		var name, email string
		var updatedAt int64
		if err := rows.Scan(&id, &name, &email, &updatedAt); err != nil {
			log.Println("Error al escanear la fila:", err)
			continue
		}
		user := &domain.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  "",
			UpdatedAt: updatedAt,
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepoMySQL) GetUsers() ([]*domain.User, error) {
	return r.GetAllUsers()
}

func (r *UserRepoMySQL) UpdateUser(user *domain.User) error {
	return r.EditUserFunction(*user)
}

func (r *UserRepoMySQL) EditUserFunction(user domain.User) error {
	newUpdatedAt := time.Now().Unix()
	query := "UPDATE Users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Email, user.Password, newUpdatedAt, user.ID)
	if err != nil {
		log.Fatalf("Error al editar info. del usuario: %v", err)
	}
	return err
}

func (r *UserRepoMySQL) DeleteUser(id int32) error {
	return r.DeleteUserFunction(id)
}

func (r *UserRepoMySQL) DeleteUserFunction(id int32) error {
	_, err := r.Connection.ExecPreparedQuerys("DELETE FROM Users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Error al eliminar usuario: %v", err)
	}
	return err
}
