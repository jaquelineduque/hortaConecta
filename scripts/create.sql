CREATE DATABASE IF NOT EXISTS hortaConecta;
USE hortaConecta;

/*inserir dados de acesso.
--To store user data*/
CREATE TABLE IF NOT EXISTS userAccount(
	id INT AUTO_INCREMENT NOT NULL, 
    CONSTRAINT PK_userAccount_id PRIMARY KEY (id),
    nickname VARCHAR(255) NOT NULL,
    fullName VARCHAR(255) NOT NULL,
    creationDate DATETIME(3),
    lastAccess DATETIME(3)
);
    
/*To store kind of product. E.g.: Fruit, Vegetable*/
CREATE TABLE IF NOT EXISTS productType(
	id INT AUTO_INCREMENT NOT NULL, 
    CONSTRAINT PK_productType_id PRIMARY KEY (id),
    typeDescription VARCHAR(255) NOT NULL
);

/*To store product. E.g.: Strawberries, bananas, passion fruit*/
CREATE TABLE IF NOT EXISTS product(
	id INT AUTO_INCREMENT NOT NULL,
	CONSTRAINT PK_product_id PRIMARY KEY (id),
	idProductType INT,
	CONSTRAINT FK_product_idProductType FOREIGN KEY (idProductType)
    	REFERENCES productType(id),
	productName VARCHAR(255) NOT NULL
);

/*To store the advertisements state. E.g.: Active, Paused, Closed*/
CREATE TABLE IF NOT EXISTS advertisementState(
	id INT NOT NULL,
	CONSTRAINT PK_advertisementState_id PRIMARY KEY (id),
	stateDescription VARCHAR(255) NOT NULL
);

/*To store the advertisements*/
CREATE TABLE IF NOT EXISTS advertisement(
	id BIGINT AUTO_INCREMENT NOT NULL,
	CONSTRAINT PK_advertisement_id PRIMARY KEY (id),
	title VARCHAR(255) NOT NULL,
	description VARCHAR(1000) NOT NULL,
	idAdvertiser INT NOT NULL,
	CONSTRAINT FK_advertisement_idAdvertiser FOREIGN KEY (idAdvertiser)
    	REFERENCES userAccount(id),
    advertisementState INT NOT NULL,
    CONSTRAINT FK_advertisement_advertisementState FOREIGN KEY (advertisementState)
    	REFERENCES advertisementState(id)
);

/*To store the shop cart, before it turns order*/
CREATE TABLE IF NOT EXISTS shopCart(
	id BIGINT AUTO_INCREMENT NOT NULL,
	CONSTRAINT PK_shopCart_id PRIMARY KEY (id),
	idAdvertisement BIGINT  NOT NULL,
	CONSTRAINT FK_shopCart_idAdvertisement FOREIGN KEY (idAdvertisement)
		REFERENCES advertisement(id),
	quantity INT NOT NULL,
	creationDate DATETIME(3),
	lastInteration DATETIME(3)
);

/*To store orders*/
CREATE TABLE IF NOT EXISTS buyOrder(
	id BIGINT AUTO_INCREMENT NOT NULL,
	CONSTRAINT PK_buyOrder_id PRIMARY KEY (id),
	idConsumer INT NOT NULL,
	CONSTRAINT FK_buyOrder_idConsumer FOREIGN KEY (idConsumer)
		REFERENCES userAccount(id),
	finalizationDate DATETIME(3)	
);

/*To store de order state*/
CREATE TABLE IF NOT EXISTS orderState(
	id INT NOT NULL,
	CONSTRAINT PK_orderState_id PRIMARY KEY (id),
	stateDescription VARCHAR(255)
);


/*To store order's products*/
CREATE TABLE IF NOT EXISTS productOrder(
	id BIGINT AUTO_INCREMENT NOT NULL,
	CONSTRAINT PK_productOrder_id PRIMARY KEY (id),
	idOrder BIGINT NOT NULL,
	CONSTRAINT FK_productOrder_idOrder FOREIGN KEY (idOrder)
		REFERENCES buyOrder(id),
	idAdvertisement BIGINT  NOT NULL,
	CONSTRAINT FK_productOrder_idAdvertisement FOREIGN KEY (idAdvertisement)
		REFERENCES advertisement(id),
	idOrderState INT NOT NULL,
	CONSTRAINT FK_productOrder_idOrderState FOREIGN KEY (idOrderState)
		REFERENCES orderState(id),
	quantity INT NOT NULL, 
	shippingFee DECIMAL(15,2),
	unitaryPrice DECIMAL(15,2) NOT NULL,
	currencyName VARCHAR(3)
)

