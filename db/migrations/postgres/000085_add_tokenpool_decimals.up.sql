BEGIN;
ALTER TABLE tokenpool ADD COLUMN decimals INTEGER DEFAULT 0;
COMMIT;