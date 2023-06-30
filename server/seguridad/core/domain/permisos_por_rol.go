package domain

type PermisosPorRol struct {
	Rol
	Crear     bool
	Editar    bool
	Consultar bool
	Eliminar  bool
	Permiso   Permiso
}
