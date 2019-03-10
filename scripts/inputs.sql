USE hortaConecta;

INSERT INTO userAccount(
	nickname,
	fullName,
	creationDate
) VALUES(
	'Jaqueline',
	'Jaqueline da Silva',
	now(3)
);

INSERT INTO productType(
	typeDescription
) VALUES(
	'Fruta'
),(
	'Legume'
),(
	'Hortaliça'
);

INSERT INTO product(
	idProductType,
	productName
) VALUES (
	1,
	'Morango'
),(
	2,
	'Abóbora'
);

INSERT INTO advertisementState(
	id,
	stateDescription
) VALUES (
	1,
	'Ativo'
),(
	2,
	'Pausado'
),(
	3,
	'Fechado'
);

INSERT INTO advertisement(
	title,
	description,
	idAdvertiser,
	advertisementState
) VALUES(
	'Morangos selecionados',
	'Os melhores morangos da nossa produção para você.#10#13 Excelente opção para doces ou consumo in natura.#10#13 Produção familiar e sem agrotóxicos',
	1,
	1	
);

INSERT INTO shopCart(
	idAdvertisement,
	quantity,
	creationDate
) VALUES(
	1,
	2,
	now(3)
);

INSERT INTO buyOrder(
	idConsumer,
	finalizationDate
) VALUES
(
	1,
	now(3)
);

INSERT INTO orderState(
	id,
	stateDescription
) VALUES(
	1,
	'Aguardando envio'
),(
	2,
	'Entregue'
);

INSERT INTO productOrder(
	idOrder,
	idAdvertisement,
	idOrderState,
	quantity,
	shippingFee,
	unitaryPrice,
	currencyName
) VALUES(
	1,
	1,
	1,
	2,
	7.00,
	5.00,
	'BRL'
);