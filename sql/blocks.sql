DROP TABLE IF EXISTS blocks;
CREATE TABLE blocks (
    number              INTEGER UNSIGNED PRIMARY KEY,
    hash_id             INTEGER UNSIGNED,
    parent_hash_id      INTEGER UNSIGNED,
    timestamp           INTEGER UNSIGNED
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX hash_id ON blocks (hash_id);
CREATE INDEX parent_hash_id ON blocks (parent_hash_id);

