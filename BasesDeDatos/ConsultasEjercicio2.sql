#1
SELECT * FROM customer;
#2
SELECT * FROM plan;
#3
SELECT c.firstName, c.lastName, p.speedMG, p.price 
FROM customer c, plan p, contract ct 
WHERE ct.idPlan = p.idPlan and c.dni = ct.dni;
#4
select COUNT(*) as totalClientes from customer;
#5
select COUNT(*) as totalPlanes from plan;
#6
SELECT * FROM contract;
#7
SELECT ct.idContract, c.firstName, c.lastName, p.speedMG, p.price 
FROM customer c
inner join contract ct on ct.dni=c.dni
inner join plan p on p.idPlan=ct.idPlan;
#8
SELECT ct.idContract, c.firstName, c.lastName, p.speedMG, p.price 
FROM customer c
inner join contract ct on ct.dni=c.dni
inner join plan p on p.idPlan=ct.idPlan
where p.idPlan=102;
#9
SELECT ct.idContract, c.firstName, c.lastName, p.speedMG, p.price 
FROM customer c
inner join contract ct on ct.dni=c.dni
inner join plan p on p.idPlan=ct.idPlan
order by c.firstName;
#10
select count(firstName) as ClientesDeEngativa from customer where province like 'Engativa';



