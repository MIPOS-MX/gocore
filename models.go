package aumpi_core

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Configuration is
type Configuration struct {
	Tables      []interface{}
	Routes      []Routes
	BeforeSetup func(db *gorm.DB)
	AfterSetup  func(db *gorm.DB)
}

//Routes define las rtas en base a controladores
type Routes struct {
	Category    string          // Categoria de la ruta de separado por >
	Description string          // Descripcion corta de lo que hace la ruta
	Self        bool            // Poner en true si el controlador tiene acceso unicamente a informacion del usuario que solicito la peticion por ejemplo sus permisos, sus leads, su actividad
	NoAuth      bool            // Desactivar las validaciones de autenticacion para este controlador usado en inicio de sesion, registro y similares
	Path        string          // Ruta en la que respondera el controlador
	Method      string          // Metodo en el que respondera el controlador
	Function    gin.HandlerFunc // Controlador asociado a la ruta
}

//Permissions is
type Permissions struct {
	Pid         uuid.UUID `gorm:"primaryKey;type:uuid"`
	Category    string    // Categoria de la ruta de separado por >
	Description string    // Descripcion corta de lo que hace la ruta
	Self        bool      // Poner en true si el controlador tiene acceso unicamente a informacion del usuario que solicito la peticion por ejemplo sus permisos, sus leads, su actividad
	NoAuth      bool      // Desactivar las validaciones de autenticacion para este controlador usado en inicio de sesion, registro y similares
	Path        string    // Ruta en la que respondera el controlador
	Method      string    // Metodo en el que respondera el controlador
}

//Roles is
type Roles struct {
	Rid            uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name           string    `gorm:"type:varchar(25)"`
	Description    string    `gorm:"type:varchar(70)"`
	Permissions    string    `gorm:"type:text"`
	PermissionsWeb string    `gorm:"type:text"`
	Editable       bool
}

type User struct {
	Id          int       `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Rid         uuid.UUID `gorm:"type:uuid" json:"rid"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_ad"`
}
