CREATE DATABASE BaseDB
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci;

CREATE TABLE BaseDB.Users (
  userId int unsigned AUTO_INCREMENT COMMENT 'ID autoincrementarl del User',
  email char(250) DEFAULT NULL COMMENT 'Email registrado por el user',
  firstName varchar(100) DEFAULT NULL COMMENT 'Nombres del user',
  lastName varchar(100) DEFAULT NULL COMMENT 'Apellidos del user',
  icCode char(10) DEFAULT NULL COMMENT 'International Calling Code: +51 -> Peru',
  phone char(20) DEFAULT NULL COMMENT 'Telefono movil',
  ip varbinary(16) DEFAULT '1' COMMENT 'IP usada al momento del registro',
  birthDate date DEFAULT NULL COMMENT 'Fecha de nacimiento del user',
  birthYear smallint unsigned DEFAULT NULL COMMENT 'Año de nacimiento del user',
  birthMonth tinyint unsigned DEFAULT NULL COMMENT 'Mes de nacimiento del user',
  countryId char(3) DEFAULT NULL COMMENT 'ID que referencia a la tabla countries',
  address text COMMENT 'Direccion del user',
  docIdType tinyint unsigned DEFAULT '1' COMMENT '1 => DNI, 2 => X, 3 => Y',
  docId char(50) DEFAULT NULL COMMENT 'Nro documento de identidad',
  completed tinyint DEFAULT '0' COMMENT '1 -> true | 0 -> false, indica que un registro ya no puede ser actualizado por el propio usuario, solo se puede actualizar desde el administrador',
  active tinyint DEFAULT '1' COMMENT '1 -> Activo | 0 -> Inactivo',
  tester tinyint DEFAULT '0' COMMENT '1 -> true | 0 -> false',
  regUserId int unsigned DEFAULT NULL COMMENT 'ID del usuario que inserto',
  regDate date DEFAULT NULL COMMENT 'Fecha registro',
  regDatetime datetime DEFAULT NULL COMMENT 'Fecha Hora registro',
  regTimestamp bigint DEFAULT NULL COMMENT 'Epoch registro',
  PRIMARY KEY (userId)
)
ENGINE = INNODB,
CHARACTER SET utf8mb4,
COLLATE utf8mb4_general_ci,
COMMENT='Tabla de Usuarios';

ALTER TABLE BaseDB.Users ADD UNIQUE unqDocs (docIdType, docId);
ALTER TABLE BaseDB.Users ADD INDEX idxEmail (email);
ALTER TABLE BaseDB.Users ADD INDEX idxIcCode (icCode);
ALTER TABLE BaseDB.Users ADD INDEX idxPhone (phone);
ALTER TABLE BaseDB.Users ADD INDEX idxIcCodePhone (icCode, phone);
ALTER TABLE BaseDB.Users ADD INDEX idxIp (ip);
ALTER TABLE BaseDB.Users ADD INDEX idxBirthDate (birthDate);
ALTER TABLE BaseDB.Users ADD INDEX idxBirthYear (birthYear);
ALTER TABLE BaseDB.Users ADD INDEX idxBirthMonth (birthMonth);
ALTER TABLE BaseDB.Users ADD INDEX idxCountryId (countryId);
ALTER TABLE BaseDB.Users ADD INDEX idxCompleted (completed);
ALTER TABLE BaseDB.Users ADD INDEX idxActive (active);
ALTER TABLE BaseDB.Users ADD INDEX idxTester (tester);
ALTER TABLE BaseDB.Users ADD INDEX idxRegUserId (regUserId);
ALTER TABLE BaseDB.Users ADD INDEX idxRegDate (regDate);
ALTER TABLE BaseDB.Users ADD INDEX idxInsRegTimestamp (regTimestamp);

INSERT INTO BaseDB.Users(firstName,lastName,docId) VALUES("Pedro","Picapiedra","12345043");
INSERT INTO BaseDB.Users(firstName,lastName,docId) VALUES("Pablo","Marmol","1234523");
INSERT INTO BaseDB.Users(firstName,lastName,docId) VALUES("Vilma","Picapiedra","12345879");
