DROP TABLE IF EXISTS index_addresses;
CREATE TABLE index_addresses (
  id   INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  address VARCHAR(129) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX address on index_addresses (address(10));
