DROP TABLE IF EXISTS blocks;
CREATE TABLE blocks (
    block_index           INTEGER UNSIGNED PRIMARY KEY,
    timestamp             INTEGER UNSIGNED,
    block_hash_id         INTEGER UNSIGNED,
    parent_block_hash_id  INTEGER UNSIGNED
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX block_hash_id ON blocks (block_hash_id);
CREATE INDEX parent_block_hash_id ON blocks (parent_block_hash_id);

DROP TABLE IF EXISTS index_blocks;
CREATE TABLE index_blocks (
  id   INTEGER UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  hash VARCHAR(129) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE INDEX hash on index_blocks (hash(10));
