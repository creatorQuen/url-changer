CREATE TABLE IF EXISTS ulrs
(
    id       serial not null unique,
    longUrl  varchar(255),
    shortUrl varchar(255)
);