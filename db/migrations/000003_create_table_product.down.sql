DROP INDEX IF EXISTS "idx_products_name";
DROP INDEX IF EXISTS "idx_products_price";
DROP INDEX IF EXISTS "idx_products_condition";
DROP INDEX IF EXISTS "idx_products_is_purchasable";

DROP TABLE IF EXISTS "public"."products";

DROP TYPE IF EXISTS condition;