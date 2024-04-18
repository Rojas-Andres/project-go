package users

import (
	"fmt"
	"time"

	"github.com/Rojas-Andres/project-go/modelos"
)

func AltaUsuario() {
	// Variables
	u:= new(modelos.User)
	// Operaciones
	u.AddUser(1, "Andres", time.Now(), true)
	fmt.Println(u)
}