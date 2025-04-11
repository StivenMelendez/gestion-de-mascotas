CREATE TABLE "Usuarios" (
  "ID" integer PRIMARY KEY,
  "Nombre" varchar,
  "Apellido" varchar,
  "Correo" varchar NOT NULL,
  "Contrasena" varchar NOT NULL,
  "confirmacion" boolean NOT NULL DEFAULT false,
  "NID" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "tipos_de_telefono" (
  "ID" integer PRIMARY KEY,
  "nombre_tipo" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "telefonos" (
  "ID" integer PRIMARY KEY,
  "numero" varchar NOT NULL,
  "ID_Usuario" integer,
  "ID_Tipo_Telefono" Integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Atencion_Al_Cliente" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Administradores" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Clientes" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Direcciones" (
  "ID" integer PRIMARY KEY,
  "localizacion" varchar,
  "ID_Clientes" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Duenos" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Prestador_de_servicio" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Empresas" (
  "ID" integer PRIMARY KEY,
  "nombre_Empresa" varchar,
  "NIT" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Tipo_Mascota" (
  "ID" integer PRIMARY KEY,
  "Nombre_tipo" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Mascotas" (
  "ID" integer PRIMARY KEY,
  "Raza" varchar,
  "Peso" float,
  "ID_Dueno" integer,
  "ID_Tipo" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Categorias_de_servicios" (
  "ID" integer PRIMARY KEY,
  "Descripcion" varchar,
  "ID_Administrador" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Servicios" (
  "ID" integer PRIMARY KEY,
  "Descripcion" varchar,
  "ID_Prestador" integer,
  "ID_Categoria" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Ofertas" (
  "ID" integer PRIMARY KEY,
  "Costo" float,
  "Fecha_Inicio" timestamp,
  "Fecha_Fin" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Pagos" (
  "ID" integer PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Servicios_por_mascota" (
  "ID" integer PRIMARY KEY,
  "ID_Oferta" integer,
  "ID_Pago" integer,
  "ID_Mascota" integer,
  "Fecha_inicio" timestamp,
  "Fecha_final" timestamp,
  "Atencion" boolean NOT NULL DEFAULT false,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Resena" (
  "ID" integer PRIMARY KEY,
  "Calificacion" float,
  "Comentario" varchar,
  "ID_Servicio_por_mascota" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Tipo_PQRS" (
  "ID" integer PRIMARY KEY,
  "Nombre_tipo" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "PQRS" (
  "ID" integer PRIMARY KEY,
  "ID_Cliente" integer,
  "Asunto" varchar,
  "Descripcion" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "Atencion_PQRS" (
  "ID_PQRS" integer,
  "ID_Atencion_al_cliente" integer,
  "Estado" boolean NOT NULL DEFAULT false,
  "Fecha_inicio" timestamp,
  "Fecha_final" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  PRIMARY KEY ("ID_PQRS", "ID_Atencion_al_cliente")
);