#Listar los datos de los autores.
SELECT * FROM autor;
#Listar nombre y edad de los estudiantes
SELECT e.Nombre, e.Edad
FROM estudiante e;
#¿Qué estudiantes pertenecen a la carrera informática?
SELECT e.Nombre
FROM estudiante e
WHERE e.Carrera = 'informática';
#¿Qué autores son de nacionalidad francesa o italiana?
SELECT a.Nombre
FROM autor a
WHERE a.Nacionalidad = 'francesa' OR a.Nacionalidad = 'italiana';
#¿Qué libros no son del área de internet?
SELECT l.Titulo
FROM libro l
WHERE l.Area != 'internet';
#Listar los libros de la editorial Salamandra.
SELECT l.Titulo
FROM libro l
WHERE l.Editorial = 'Salamandra';
#Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT *
FROM estudiante e
WHERE e.Edad > (SELECT AVG(e.Edad));
#Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT e.Nombre
FROM estudiante e
WHERE e.Apellido LIKE 'G%';
#Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT a.Nombre
FROM autor a, libro l
INNER JOIN libroautor la ON la.idLibro = l.idLibro
WHERE l.Titulo = 'El Universo: Guía de viaje';
#¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT l.Titulo
FROM libro l, estudiante e
INNER JOIN prestamo p ON p.IdLector = e.idLector
WHERE e.Nombre LIKE 'Filippo' AND e.Apellido LIKE 'Galli';
#Listar el nombre del estudiante de menor edad.
SELECT e.Nombre
FROM estudiante e
ORDER BY e.Edad
LIMIT 1;
#Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.

#Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.Titulo
FROM libro l, autor a
INNER JOIN libroautor la ON la.IdAutor = a.idAutor
WHERE a.Nombre = 'J.K. Rowling';
#Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT l.Titulo
FROM libro l
INNER JOIN prestamo p ON p.idLibro = l.idLibro
WHERE p.FechaDevolucion = '16/07/2021';


