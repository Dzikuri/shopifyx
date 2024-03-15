CREATE TYPE condition AS ENUM ('new','second');

CREATE TABLE IF NOT EXISTS "public"."products" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "name" varchar(60) NOT NULL,
    "price" integer NOT NULL,
    "image_url" varchar(255) NOT NULL,
    "stock" integer NOT NULL,
    "condition" condition NOT NULL,
    "tags" text[] NOT NULL, 
    "is_purchasable" boolean NOT NULL DEFAULT TRUE,
    "created_at" timestamptz(6),
    "updated_at" timestamptz(6)

     CONSTRAINT "fk_user"
      FOREIGN KEY("user_id") 
	  REFERENCES users(id)
	  ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_products_name" ON "public"."products"("name");
CREATE INDEX IF NOT EXISTS "idx_products_price" ON "public"."products"("price");
CREATE INDEX IF NOT EXISTS "idx_products_condition" ON "public"."products"("condition");
CREATE INDEX IF NOT EXISTS "idx_products_is_purchasable" ON "public"."products"("is_purchasable");