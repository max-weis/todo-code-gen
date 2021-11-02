CREATE TABLE todos
(
    id          BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title       text    NOT NULL,
    description text,
    status      integer NOT NULL
);