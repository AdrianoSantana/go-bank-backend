CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "owner" varchar,
  "balance" bigint,
  "currency" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" serial PRIMARY KEY,
  "account_id" int NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" serial PRIMARY KEY,
  "from_account_id" int NOT NULL,
  "to_account_id" int NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");