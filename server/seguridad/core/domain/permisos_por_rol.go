package domain

type PermisosPorRol struct {
	Crear     bool
	Editar    bool
	Consultar bool
	Eliminar  bool
	Permiso   Permiso
}
