#Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad 
FROM empleado e
INNER JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE e.puesto = 'Vendedor';
#Visualizar los departamentos con más de cinco empleados.
SELECT d.nombre_depto 
FROM departamento d
INNER JOIN empleado e ON e.depto_nro = d.depto_nro
GROUP BY d.nombre_depto
HAVING COUNT(e.nombre)>5;
#Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
SELECT e.nombre, e.apellido, e.salario, d.nombre_depto, e.puesto
FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
WHERE e.puesto = (SELECT puesto FROM empleado WHERE nombre LIKE 'mito' AND apellido LIKE 'Barchuk');
#Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT * 
FROM empleado e
inner join departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidad';
#Mostrar el nombre del empleado que tiene el salario más bajo.
select nombre 
from empleado 
where salario = (SELECT MIN(e.salario)
FROM empleado e);
#Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.nombre, e.salario, e.depto_nro, d.nombre_depto  FROM empleado e
INNER JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Ventas'
order by e.salario desc
limit 1;



