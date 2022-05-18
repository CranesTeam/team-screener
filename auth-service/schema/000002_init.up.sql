DROP TABLE IF EXISTS "public"."oauth2_clients";

CREATE TABLE "public"."oauth2_clients" (
    "id" text NOT NULL,
    "secret" text NOT NULL,
    "domain" text NOT NULL,
    "data" jsonb NOT NULL,
    CONSTRAINT "oauth2_clients_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "public"."oauth2_tokens";
DROP SEQUENCE IF EXISTS oauth2_tokens_id_seq;
CREATE SEQUENCE oauth2_tokens_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."oauth2_tokens" (
    "id" bigint DEFAULT nextval('oauth2_tokens_id_seq') NOT NULL,
    "created_at" timestamptz NOT NULL,
    "expires_at" timestamptz NOT NULL,
    "code" text NOT NULL,
    "access" text NOT NULL,
    "refresh" text NOT NULL,
    "data" jsonb NOT NULL,
    CONSTRAINT "oauth2_tokens_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_oauth2_tokens_access" ON "public"."oauth2_tokens" USING btree ("access");
CREATE INDEX "idx_oauth2_tokens_code" ON "public"."oauth2_tokens" USING btree ("code");
CREATE INDEX "idx_oauth2_tokens_expires_at" ON "public"."oauth2_tokens" USING btree ("expires_at");
CREATE INDEX "idx_oauth2_tokens_refresh" ON "public"."oauth2_tokens" USING btree ("refresh");