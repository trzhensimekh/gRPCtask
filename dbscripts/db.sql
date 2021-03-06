CREATE TABLE users (
    id INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
    firstname  varchar(255),
    lastname  varchar(255),
    email varchar(255)
);
CREATE TABLE messages (
     id INT GENERATED BY DEFAULT AS IDENTITY UNIQUE NOT NULL,
    message varchar(255),
    user_id INT REFERENCES users(id) on delete cascade on update cascade
);
INSERT INTO users(firstname,lastname,email) values('Tolia','Picus','tarakan@net.net');
INSERT INTO messages(message, user_id) values ('message one for Tolia',(SELECT id FROM users WHERE lastname='Picus'));
INSERT INTO messages(message, user_id) values ('message two for Tolia',(SELECT id FROM users WHERE lastname='Picus'));
INSERT INTO users(firstname,lastname,email) values('Olia','Vicus','kotik@net.net');
INSERT INTO messages(message, user_id) values ('message one for Olia',(SELECT id FROM users WHERE lastname='Vicus'));
INSERT INTO messages(message, user_id) values ('message two for Olia',(SELECT id FROM users WHERE lastname='Vicus'));
SELECT * FROM users;
SELECT * FROM messages;
