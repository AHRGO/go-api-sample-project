create table produto (
	id_produto BIGINT not null generated always as identity,
	nome_produto VARCHAR(50) not null,
	preco NUMERIC(10,2) not null,
	
	constraint pk_produto primary key (id_produto)
);

select * from  produto;