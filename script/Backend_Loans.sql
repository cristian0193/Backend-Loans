
/****** Object:  Table [User] ******/
CREATE TABLE "Users" (
	id SERIAL PRIMARY KEY,
	name varchar(200) NOT NULL,
	username varchar(10) NOT NULL,
	password varchar(200) NOT NULL,
	creation_date date NOT NULL
);

/****** Object:  Table [Type] ******/
CREATE TABLE "Types" (
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL
);

/****** Object:  Table [Status] ******/
CREATE TABLE "Status" (
	id SERIAL PRIMARY KEY,
	description varchar(10) NOT NULL
);

/****** Object:  Table [Loans] ******/
CREATE TABLE "Loans" (
	id SERIAL PRIMARY KEY,
	identification_client int4 NOT NULL,
	borrowed_value numeric NOT NULL,
	interest_percentage numeric NOT NULL,
	paid_value numeric NOT NULL,
	pending_value numeric NOT NULL,
	interest_paid numeric NOT NULL,
	months_arrears int4 NOT NULL,
	id_status int4 NOT NULL,
	creation_date date NOT NULL
);

/****** Object:  Table [Clients] ******/
CREATE TABLE "Clients" (
	identification int8 PRIMARY KEY,
	fullname varchar(200) NOT NULL,
	address varchar(10) NOT NULL,
	mobile varchar(10) NOT NULL
);

/****** Object:  Table [Historial] ******/
CREATE TABLE "Historial" (
	id SERIAL PRIMARY KEY,
	id_loan int4 NOT NULL,
	capital numeric NOT NULL,
	interest numeric NOT NULL,
	balance numeric NOT NULL,
	payment_date date NOT NULL,
	id_type int4 NOT NULL
);

/****** Object:  Table [Interests] ******/
CREATE TABLE "Interests" (
	id SERIAL PRIMARY KEY,
	identification_client int4 NOT NULL,
	share numeric NOT NULL,
	status varchar(3) NOT NULL
);

alter table "Loans" add constraint FK_loand_status foreign key (id_status) references "Status"(id);
alter table "Loans" add constraint FK_loand_clients foreign key (identification_client) references "Clients"(identification);
alter table "Historial" add constraint FK_historial_loans foreign key (id_loan) references "Loans"(id);
alter table "Historial" add constraint FK_historial_types foreign key (id_type) references "Types"(id);
alter table "Interests" add constraint FK_interests_client foreign key (identification_client) references "Clients"(identification);