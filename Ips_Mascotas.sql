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

ALTER TABLE "Atencion_Al_Cliente" ADD CONSTRAINT "Usuarios" FOREIGN KEY ("ID") REFERENCES "Usuarios" ("ID");

ALTER TABLE "Administradores" ADD CONSTRAINT "Usuarios" FOREIGN KEY ("ID") REFERENCES "Usuarios" ("ID");

ALTER TABLE "Clientes" ADD CONSTRAINT "Usuarios" FOREIGN KEY ("ID") REFERENCES "Usuarios" ("ID");

ALTER TABLE "telefonos" ADD CONSTRAINT "Usuarios" FOREIGN KEY ("ID_Usuario") REFERENCES "Usuarios" ("ID");

ALTER TABLE "Categorias_de_servicios" ADD CONSTRAINT "Administradores" FOREIGN KEY ("ID_Administrador") REFERENCES "Administradores" ("ID");

ALTER TABLE "Atencion_PQRS" ADD CONSTRAINT "Atencion_al_cliente" FOREIGN KEY ("ID_Atencion_al_cliente") REFERENCES "Atencion_Al_Cliente" ("ID");

ALTER TABLE "telefonos" ADD CONSTRAINT "tipos_de_telefono" FOREIGN KEY ("ID_Tipo_Telefono") REFERENCES "tipos_de_telefono" ("ID");

ALTER TABLE "Direcciones" ADD CONSTRAINT "Clientes" FOREIGN KEY ("ID_Clientes") REFERENCES "Clientes" ("ID");

ALTER TABLE "Duenos" ADD CONSTRAINT "Clientes" FOREIGN KEY ("ID") REFERENCES "Clientes" ("ID");

ALTER TABLE "Prestador_de_servicio" ADD CONSTRAINT "Clientes" FOREIGN KEY ("ID") REFERENCES "Clientes" ("ID");

ALTER TABLE "PQRS" ADD CONSTRAINT "Clientes" FOREIGN KEY ("ID_Cliente") REFERENCES "Clientes" ("ID");

ALTER TABLE "Mascotas" ADD CONSTRAINT "Dueno" FOREIGN KEY ("ID_Dueno") REFERENCES "Duenos" ("ID");

ALTER TABLE "Servicios" ADD CONSTRAINT "Prestador_de_servicio" FOREIGN KEY ("ID_Prestador") REFERENCES "Prestador_de_servicio" ("ID");

ALTER TABLE "Prestador_de_servicio" ADD CONSTRAINT "Empresas" FOREIGN KEY ("ID") REFERENCES "Empresas" ("ID");

ALTER TABLE "Servicios" ADD CONSTRAINT "Empresas" FOREIGN KEY ("ID_Prestador") REFERENCES "Empresas" ("ID");

ALTER TABLE "Mascotas" ADD CONSTRAINT "Tipo_Mascota" FOREIGN KEY ("ID_Tipo") REFERENCES "Tipo_Mascota" ("ID");

ALTER TABLE "Servicios_por_mascota" ADD CONSTRAINT "Mascotas" FOREIGN KEY ("ID_Mascota") REFERENCES "Mascotas" ("ID");

ALTER TABLE "Servicios" ADD CONSTRAINT "Categorias_de_servicios" FOREIGN KEY ("ID_Categoria") REFERENCES "Categorias_de_servicios" ("ID");

ALTER TABLE "Ofertas" ADD CONSTRAINT "Servicios" FOREIGN KEY ("ID") REFERENCES "Servicios" ("ID");

ALTER TABLE "Servicios_por_mascota" ADD CONSTRAINT "Ofertas" FOREIGN KEY ("ID_Oferta") REFERENCES "Ofertas" ("ID");

ALTER TABLE "Servicios_por_mascota" ADD CONSTRAINT "Pagos" FOREIGN KEY ("ID_Pago") REFERENCES "Pagos" ("ID");

ALTER TABLE "Resena" ADD CONSTRAINT "Servicios_por_mascota" FOREIGN KEY ("ID_Servicio_por_mascota") REFERENCES "Servicios_por_mascota" ("ID");

ALTER TABLE "Atencion_PQRS" ADD CONSTRAINT "PQRS" FOREIGN KEY ("ID_PQRS") REFERENCES "PQRS" ("ID");
