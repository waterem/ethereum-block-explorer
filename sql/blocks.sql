DROP TABLE IF EXISTS blocks;
CREATE TABLE blocks (
    block_index     INTEGER UNSIGNED PRIMARY KEY,
    timestamp       INTEGER UNSIGNED,
    hash            INTEGER UNSIGNED,
    parent_hash     INTEGER UNSIGNED
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX block_hash_id ON blocks (block_hash_id);

