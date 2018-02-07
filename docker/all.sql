DROP TABLE IF EXISTS blocks;
CREATE TABLE blocks (
    number              INTEGER UNSIGNED PRIMARY KEY,
    hash_id             INTEGER UNSIGNED,
    parent_hash_id      INTEGER UNSIGNED,
    timestamp           INTEGER UNSIGNED
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX hash_id ON blocks (hash_id);
CREATE INDEX parent_hash_id ON blocks (parent_hash_id);

DROP TABLE IF EXISTS index_blocks;
CREATE TABLE index_blocks (
  id   INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  hash VARCHAR(129) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX hash on index_blocks (hash(10));

DROP TABLE IF EXISTS index_addresses;
CREATE TABLE index_addresses (
  id   INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  address VARCHAR(129) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX address on index_addresses (address(10));

