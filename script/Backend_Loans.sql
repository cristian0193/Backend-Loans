-- public."Clients" definition

-- Drop table

-- DROP TABLE public."Clients";

CREATE TABLE public."Clients" (
	identification int8 NOT NULL,
	fullname varchar(200) NOT NULL,
	address varchar(100) NOT NULL,
	mobile varchar(100) NOT NULL,
	CONSTRAINT "Clients_pkey" PRIMARY KEY (identification)
);


-- public."Status" definition

-- Drop table

-- DROP TABLE public."Status";

CREATE TABLE public."Status" (
	id serial NOT NULL,
	description varchar(10) NOT NULL,
	CONSTRAINT "Status_pkey" PRIMARY KEY (id)
);


-- public."Types" definition

-- Drop table

-- DROP TABLE public."Types";

CREATE TABLE public."Types" (
	id serial NOT NULL,
	"name" varchar(50) NOT NULL,
	CONSTRAINT "Types_pkey" PRIMARY KEY (id)
);


-- public."Users" definition

-- Drop table

-- DROP TABLE public."Users";

CREATE TABLE public."Users" (
	id serial NOT NULL,
	"name" varchar(200) NOT NULL,
	username varchar(10) NOT NULL,
	"password" varchar(200) NOT NULL,
	creation_date date NOT NULL,
	CONSTRAINT "Users_pkey" PRIMARY KEY (id)
);


-- public."Loans" definition

-- Drop table

-- DROP TABLE public."Loans";

CREATE TABLE public."Loans" (
	id serial NOT NULL,
	identification_client int4 NOT NULL,
	borrowed_value numeric NOT NULL,
	interest_percentage numeric NOT NULL,
	paid_value numeric NOT NULL,
	pending_value numeric NOT NULL,
	interest_paid numeric NOT NULL,
	id_status int4 NOT NULL,
	creation_date date NOT NULL,
	CONSTRAINT "Loans_pkey" PRIMARY KEY (id),
	CONSTRAINT fk_loand_clients FOREIGN KEY (identification_client) REFERENCES "Clients"(identification),
	CONSTRAINT fk_loand_status FOREIGN KEY (id_status) REFERENCES "Status"(id)
);


-- public."Payments" definition

-- Drop table

-- DROP TABLE public."Payments";

CREATE TABLE public."Payments" (
	id serial NOT NULL,
	id_loan int4 NOT NULL,
	capital numeric NOT NULL,
	interest numeric NOT NULL,
	balance numeric NOT NULL,
	payment_date date NOT NULL,
	id_type int4 NOT NULL,
	CONSTRAINT "Payments_pkey" PRIMARY KEY (id),
	CONSTRAINT fk_payments_loans FOREIGN KEY (id_loan) REFERENCES "Loans"(id),
	CONSTRAINT fk_payments_types FOREIGN KEY (id_type) REFERENCES "Types"(id)
);


-- public."Interests" definition

-- Drop table

-- DROP TABLE public."Interests";

CREATE TABLE public."Interests" (
	id serial NOT NULL,
	id_loan int4 NOT NULL,
	"share" numeric NOT NULL,
	status varchar(3) NOT NULL,
	CONSTRAINT "Interests_pkey" PRIMARY KEY (id),
	CONSTRAINT fk_interests_loans FOREIGN KEY (id_loan) REFERENCES "Loans"(id)
);