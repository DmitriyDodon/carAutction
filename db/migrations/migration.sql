CREATE TABLE IF NOT EXISTS cars (
    id U			UID PRIMARY KEY,
    color        	VARCHAR(255) NOT NULL,
	price_in_cents 	BIGINT NOT NULL,
	max_speed_mph  	INT,
	max_speed_kmp 	INT NOT NULL,
	vendor_name   	VARCHAR(255) NOT NULL,
	model_name   	VARCHAR(255) NOT NULL,
    created_at 		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);