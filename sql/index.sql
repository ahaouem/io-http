CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL
);

INSERT INTO Users (name, lastname) VALUES ('John', 'Doe');
INSERT INTO Users (name, lastname) VALUES ('Jane', 'Smith');
INSERT INTO Users (name, lastname) VALUES ('Alice', 'Johnson');
