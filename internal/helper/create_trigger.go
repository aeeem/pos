package helper

import (
	"log"

	"gorm.io/gorm"
)

func TransactionTrigger(db *gorm.DB) {
	err := db.Exec(`
	DROP TRIGGER IF EXISTS CheckCustomerTransactionNo on transactions;
	CREATE TRIGGER CheckCustomerTransactionNo
   BEFORE  insert or update
   ON transactions
   FOR EACH ROW 
       EXECUTE PROCEDURE checkcustomer();`).Error
	if err != nil {
		panic(err)
	}
}

func CartTrigger(db *gorm.DB) {
	err := db.Exec(`
	DROP TRIGGER IF EXISTS UpdateTotalPrice on carts;

	CREATE TRIGGER UpdateTotalPrice
   BEFORE  insert OR update
   ON carts
   FOR EACH ROW 
       EXECUTE PROCEDURE update_total_price();`).Error
	if err != nil {
		panic(err)
	}
}
func UpdateTotalPrice(db *gorm.DB) {
	err := db.Exec(`CREATE OR REPLACE FUNCTION public.update_total_price()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    VOLATILE
    COST 100
AS $BODY$
declare total_price_cart int;
BEGIN
	select sum(carts.sub_price) from carts where transaction_id=NEW.transaction_id AND deleted_at is null into total_price_cart;
	RAISE NOTICE '% ' ,total_price_cart;
	update transactions SET "total_price" = total_price_cart where id=NEW.transaction_id;
	return NEW;
END;
$BODY$;`).Error
	if err != nil {
		panic(err)
	}
}
func CheckCustomer(db *gorm.DB) {
	err := db.Exec(`CREATE OR REPLACE FUNCTION public.checkcustomer()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    IMMUTABLE
    COST 100
    SET event_triggers=false
AS $BODY$
DECLARE cnt INTEGER;
  
BEGIN
  SELECT COUNT(*) INTO cnt
  FROM transactions
  WHERE customer_id = NEW.customer_id and (transaction.status='pending' or transaction.status='draft');

  NEW.customer_transaction_no = cnt + 1;

  RETURN NEW;
END;
$BODY$;`).Error
	if err != nil {
		panic(err)
	}
}

func CreateStatusEnum(db *gorm.DB) {
	err := db.Exec(`DROP TYPE IF EXISTS status;
CREATE TYPE status AS ENUM ('pending', 'completed', 'canceled','draft');`).Error
	if err != nil {
		log.Print(err)
	}
}
