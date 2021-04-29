/****** Object:  Database [BaseLine] ******/
CREATE DATABASE BaseLine;

/****** Object:  Table [Product] ******/
CREATE TABLE Product (
	id SERIAL PRIMARY KEY,
	name varchar(200) NOT NULL,
	price decimal NOT NULL,
	id_category int NOT NULL,
	id_market uuid NOT NULL
);

/****** Insert ******/
insert into Product (id, name, price, id_category, id_market) values (1, 'Wine - Prem Select Charddonany', 2146, 1, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (2, 'Syrup - Monin, Irish Cream', 6269, 2), '90a92d0f-6565-4311-b0b0-77f9da9bffba';
insert into Product (id, name, price, id_category) values (3, 'Alize Red Passion', 16760, 3, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (4, 'Ginger - Crystalized', 20733, 4, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (5, 'Basil - Seedlings Cookstown', 20075, 1, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (6, 'Cranberries - Dry', 26798, 2, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (7, 'Flower - Carnations', 11210, 2, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (8, 'Sauce - Marinara', 25893, 1, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (9, 'Chick Peas - Canned', 17475, 3, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (10, 'Tea - Lemon Green Tea', 4618, 4, '90a92d0f-6565-4311-b0b0-77f9da9bffba');
insert into Product (id, name, price, id_category) values (11, 'Bread - Pita', 27608, 2, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (12, 'Pasta - Orzo, Dry', 25405, 1, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (13, 'Rhubarb', 13841, 2, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (14, 'Wine - Rhine Riesling Wolf Blass', 4461, 4, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (15, 'Mortadella', 17254, 1, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (16, 'Coffee - Almond Amaretto', 9282, 2, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (17, 'Sausage - Andouille', 7867, 4, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (18, 'Saskatoon Berries - Frozen', 21473, 3, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (19, 'Grapefruit - White', 2271, 4, 'd7070826-371d-47f0-a18a-3bb257b92527');
insert into Product (id, name, price, id_category) values (20, 'Wine - Cotes Du Rhone Parallele', 13089, 3, 'd7070826-371d-47f0-a18a-3bb257b92527');