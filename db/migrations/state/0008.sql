-- +migrate Up
ALTER TABLE state.batch
    ADD COLUMN batch_hash BYTEA,
    ADD COLUMN da_block_number BIGINT,
    ADD COLUMN da_proof TEXT[],
    ADD COLUMN da_width SMALLINT,
    ADD COLUMN da_index SMALLINT;

-- +migrate Down
ALTER TABLE state.batch
    DROP COLUMN batch_hash,
    DROP COLUMN da_block_number,
    DROP COLUMN da_proof,
    DROP COLUMN da_width,
    DROP COLUMN da_index;
