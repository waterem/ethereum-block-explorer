DROP TABLE IF EXISTS tokens;
CREATE TABLE tokens (
    id                  INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name                VARCHAR(129) NOT NULL,
    symbol              VARCHAR(129) NOT NULL,
    total_suply         INTEGER UNSIGNED,
    decimals            INTEGER UNSIGNED,
    address_id          INTEGER UNSIGNED
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
