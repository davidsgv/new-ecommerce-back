DROP DATABASE IF EXISTS seguridad;
CREATE DATABASE seguridad;

USE seguridad;

CREATE TABLE usuario(
    id INT AUTO_INCREMENT PRIMARY KEY,
    password VARCHAR(500) NOT NULL,
    correo VARCHAR(500) NOT NULL,
    #secreto VARCHAR(500) NOT NULL, #para firmar el token
    telefono VARCHAR(10),
    celular VARCHAR(15),
    direccion VARCHAR(200),
    CONSTRAINT unique_usuario_correo UNIQUE (correo),
    CONSTRAINT unique_usuario_telefono UNIQUE (telefono),
    CONSTRAINT unique_usuario_celular UNIQUE (celular)
);

CREATE TABLE conexion(
    id INT AUTO_INCREMENT PRIMARY KEY,
    token VARCHAR(1000) NOT NULL,
    ingreso DATETIME NOT NULL,
    equipo VARCHAR(500),
    vencimiento_token DATETIME NOT NULL,
    usuario_id INT NOT NULL,
    CONSTRAINT fk_conexion_usuario FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);

CREATE TABLE bloqueo(
    id INT AUTO_INCREMENT PRIMARY KEY,
    hora DATETIME NOT NULL DEFAULT NOW(),
    usuario_id INT NOT NULL,
    CONSTRAINT fk_bloqueo_usuario FOREIGN KEY (usuario_id) REFERENCES usuario (id)
);

CREATE TABLE rol(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(200) NOT NULL,
    descripcion VARCHAR(2000)  NOT NULL,
    CONSTRAINT unique_rol_nombre UNIQUE (correo),
);

CREATE TABLE usuario_por_rol(
    usuario_id INT NOT NULL,
    rol_id INT NOT NULL,
    PRIMARY KEY (usuario_id, rol_id),
    CONSTRAINT fk_usuario_por_rol_usuario FOREIGN KEY (usuario_id) REFERENCES usuario (id),
    CONSTRAINT fk_usuario_por_rol_rol FOREIGN KEY (rol_id) REFERENCES rol (id)
);

CREATE TABLE permiso(
    id INT AUTO_INCREMENT PRIMARY KEY,
    modulo VARCHAR(100) NOT NULL,
    recurso VARCHAR(100) NOT NULL,
    CONSTRAINT unique_permiso UNIQUE (modulo, recurso)
);

CREATE TABLE permiso_por_rol(
    permiso_id INT NOT NULL,
    rol_id INT NOT NULL,
    crear BOOLEAN NOT NULL DEFAULT 0,
    editar BOOLEAN NOT NULL DEFAULT 0,
    eliminar BOOLEAN NOT NULL DEFAULT 0,
    consultar BOOLEAN NOT NULL DEFAULT 0,
    PRIMARY KEY (permiso_id, rol_id),
    CONSTRAINT fk_permiso_por_rol_rol FOREIGN KEY (rol_id) REFERENCES rol (id),
    CONSTRAINT fk_permiso_por_rol_permiso FOREIGN KEY (permiso_id) REFERENCES permiso (id)
);
