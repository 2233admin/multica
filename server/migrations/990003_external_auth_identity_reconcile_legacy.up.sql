-- Converge databases that applied the original constraint-bearing version of
-- 433_external_auth_identity (foreign key to "user", two UNIQUE constraints and
-- a user_id index). Those databases skip the now-edited 433 by version, so the
-- legacy objects still exist there; this drops them so every database ends up
-- with the FK-free table plus the 990001 / 990002 unique indexes.
--
-- 990001 and 990002 sort before this file, so their unique indexes on
-- (provider, issuer, subject) and (user_id, provider, issuer) already exist when
-- the equivalent constraints are dropped and uniqueness is never lost. The
-- 990002 index also leads with user_id, which replaces the dropped
-- external_auth_identity_user_id_idx. Every statement uses IF EXISTS, so this
-- is a no-op on a fresh install.
ALTER TABLE external_auth_identity
    DROP CONSTRAINT IF EXISTS external_auth_identity_user_id_fkey;
ALTER TABLE external_auth_identity
    DROP CONSTRAINT IF EXISTS external_auth_identity_provider_issuer_subject_key;
ALTER TABLE external_auth_identity
    DROP CONSTRAINT IF EXISTS external_auth_identity_user_id_provider_issuer_key;
DROP INDEX IF EXISTS external_auth_identity_user_id_idx;
