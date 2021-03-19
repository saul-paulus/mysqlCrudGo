CREATE DATABASE db_mhsApi;
CREATE USER `saul`@`localhost` IDENTIFIED BY '123456789';
USE db_mhsApi;
GRANT ALL PRIVILEGES ON db_mhsApi.* TO `saul`@`localhost`;
FLUSH PRIVILEGES;


CREATE TABLE mahasiswa (
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `nim` INT NOT NULL,
    `nama` VARCHAR (60) NOT NULL,
    `semester` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);
