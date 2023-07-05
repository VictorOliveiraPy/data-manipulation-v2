-- Table "client_data"
CREATE TABLE raw_client_data (
                                 id VARCHAR(100) NOT NULL PRIMARY KEY,
                                 document VARCHAR(30) NOT NULL,
                                 is_private VARCHAR(14),
                                 is_incomplete VARCHAR(14),
                                 last_purchase_date VARCHAR(255),
                                 average_ticket VARCHAR(255),
                                 last_purchase_ticket VARCHAR(255),
                                 most_frequent_store VARCHAR(255),
                                 last_purchase_store VARCHAR(255),
                                 status VARCHAR(50),
                                 created_at TIMESTAMP,
                                 updated_at TIMESTAMP

);


CREATE TABLE client_data (
                             id UUID NOT NULL PRIMARY KEY,
                             document VARCHAR(30) NOT NULL,
                             document_type VARCHAR(30),
                             is_private BOOLEAN,
                             is_incomplete BOOLEAN,
                             last_purchase_date DATE,
                             average_ticket FLOAT,
                             last_purchase_ticket FLOAT,
                             most_frequent_store VARCHAR(30),
                             last_purchase_store VARCHAR(30),
                             status VARCHAR(30),
                             created_at TIMESTAMP
);


-- Criação do novo esquema 'dataloader_test'
CREATE SCHEMA IF NOT EXISTS dataloader_test;

-- Table "client_data" no esquema 'dataloader_test'
CREATE TABLE IF NOT EXISTS dataloader_test.raw_client_data (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    document VARCHAR(30) NOT NULL,
    is_private VARCHAR(14),
    is_incomplete VARCHAR(14),
    last_purchase_date VARCHAR(255),
    average_ticket VARCHAR(255),
    last_purchase_ticket VARCHAR(255),
    most_frequent_store VARCHAR(255),
    last_purchase_store VARCHAR(255),
    status VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    );


CREATE TABLE dataloader_test.client_data (
        id UUID NOT NULL PRIMARY KEY,
        document VARCHAR(30) NOT NULL,
        document_type VARCHAR(30),
        is_private BOOLEAN,
        is_incomplete BOOLEAN,
        last_purchase_date DATE,
        average_ticket FLOAT,
        last_purchase_ticket FLOAT,
        most_frequent_store VARCHAR(30),
        last_purchase_store VARCHAR(30),
        status VARCHAR(30),
        created_at TIMESTAMP
);
