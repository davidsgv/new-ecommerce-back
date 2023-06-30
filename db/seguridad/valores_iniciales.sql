USE seguridad;

INSERT INTO tipo_identificacion (nombre, codigo) VALUES 
('Cédula de ciudadanía', 'CC'),
('Número de identificación tributaria', 'NIT');

INSERT INTO empresa (limite_usuario, razon_social, identificacion, tipo_identificacion_id) VALUES
(10, 'Automatsoft', '1112234567', 2);

INSERT INTO servidor (dominio, direccion_ip, nombre_bd, usuario, password, empresa_id) VALUES
('127.0.0.1:8081', 'a', 'a', 'a', 'a', 1);

INSERT INTO rol (nombre, descripcion) VALUES
('SuperAdmin', 'Administrador con control total sobre el sistema');

INSERT INTO permiso (modulo, recurso) VALUES
('seguridad', 'empresas'),
('seguridad', 'servidores'),
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

INSERT INTO usuario_por_empresa (usuario_id, empresa_id, activo) VALUES
(1, 1, 1)