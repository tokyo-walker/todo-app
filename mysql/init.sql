USE todo;

CREATE TABLE todos
(
  id           INT(10) NOT NULL AUTO_INCREMENT,
  title        VARCHAR(40),
  text      VARCHAR(40),
  favorite     BOOLEAN NOT NULL,
PRIMARY KEY (id)
);

