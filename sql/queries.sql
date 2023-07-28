INSERT INTO `directorio`.`contacto`
(
`name`,
`phone_number`,
`e_mail`,
`last_update`,
`enabled`)
VALUES
(
"Dante",
"5512345679",
"dante@correo.com",
now(),
1);

UPDATE `directorio`.`contacto`
SET
`enabled` = 1,
`last_update` =now()
WHERE `idcontacto` = 5;



SELECT * FROM directorio.contacto where name = "Dante" and enabled = 1;
SELECT name as nombre,phone_number as phone, e_mail as mail FROM directorio.contacto order by name desc ;
-- delete FROM directorio.contacto where name = "Dante" and enabled = 1;