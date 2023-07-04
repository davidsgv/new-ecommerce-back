USE seguridad;

INSERT INTO rol (nombre, descripcion) VALUES
('SuperAdmin', 'Administrador con control total sobre el sistema');

INSERT INTO permiso (modulo, recurso) VALUES
('seguridad', 'usuarios'),
('seguridad', 'conexiones'),
('seguridad', 'bloqueos'),
('seguridad', 'roles'),
('seguridad', 'permisos');

INSERT INTO permiso_por_rol (permiso_id, rol_id, crear, editar, eliminar, consultar) VALUES
(1,1,1,1,1,1),
(2,1,1,1,1,1),
(3,1,1,1,1,1),
(4,1,1,1,1,1),
(5,1,1,1,1,1),
(6,1,1,1,1,1),
(7,1,1,1,1,1);

#contra: admin123
INSERT INTO usuario (password, correo, telefono, celular) VALUES
('9150a266c71a4cf0cbd01a60608f395ec1f8f7082f4041e49195cad98f6ee0cb08efe60cb8d148d2e40520b33922bf40', 'admin@gmail.com', '0601234567', '3135068997'/*, '1asd5a6sd1a6s51d6asd1651asd'*/);

INSERT INTO usuario_por_rol (usuario_id, rol_id) VALUES
(1,1);