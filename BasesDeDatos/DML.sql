# INSERTS DE CUSTOMMER
INSERT INTO empresa_internet.customer VALUES (12345, 'Marcela', 'Lopez', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023910, 'Carlos', 'Pastrana', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023911, 'Maria', 'Lopez', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023912, 'Andres', 'Gonzales', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023913, 'Daniel', 'Perez', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023914, 'Jose', 'Uribe', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023915, 'Manuel', 'Obregon', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023916, 'Elizabeth', 'Porras', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023917, 'Sandra', 'Arero', '2000-01-10', 'Engativa', 'Bogota');
INSERT INTO empresa_internet.customer VALUES (1023918, 'Luigi', 'Moreno', '2000-01-10', 'Engativa', 'Bogota');

# INSERTS DE PLAN
INSERT INTO empresa_internet.plan VALUES (12345, 100, 200.000, 15.500);
INSERT INTO empresa_internet.plan VALUES (101, 50, 50.000, 5.000);
INSERT INTO empresa_internet.plan VALUES (102, 120, 210.000, 15.500);
INSERT INTO empresa_internet.plan VALUES (103, 150, 240.000, 16.500);
INSERT INTO empresa_internet.plan VALUES (104, 160, 250.000, 17.500);
INSERT INTO empresa_internet.plan VALUES (105, 180, 260.000, 18.500);
INSERT INTO empresa_internet.plan VALUES (106, 200, 280.000, 18.500);
INSERT INTO empresa_internet.plan VALUES (107, 250, 300.000, 18.500);
INSERT INTO empresa_internet.plan VALUES (108, 300, 320.000, 20.000);
INSERT INTO empresa_internet.plan VALUES (109, 500, 350.000, 30.000);

# INSERTS DE CONTRACT
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (12345, 12345);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (101, 1023910);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (101, 1023911);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (102, 1023912);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (102, 1023913);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (103, 1023914);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (104, 1023914);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (105, 1023915);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (106, 1023916);
INSERT INTO empresa_internet.contract (idPlan, dni) VALUES (107, 1023917);

	