package helper

import (
	"log"

	"gorm.io/gorm"
)

func UpdateTransactionTrigger(db *gorm.DB) {
	err := db.Exec(`
	DROP TRIGGER IF EXISTS CheckCustomerTransactionNo on transactions;
	CREATE TRIGGER CheckCustomerTransactionNo
   BEFORE  insert 
   ON transactions
   FOR EACH ROW 
       EXECUTE PROCEDURE checkcustomer();`).Error
	if err != nil {
		panic(err)
	}
}
func TransactionTrigger(db *gorm.DB) {
	err := db.Exec(`
	DROP TRIGGER IF EXISTS CheckCustomerCountAfterUpdate on transactions;
	CREATE TRIGGER CheckCustomerCountAfterUpdate
   BEFORE  update of status
   ON transactions
   FOR EACH ROW 
       EXECUTE PROCEDURE CheckCustomerCountAfterUpdate();`).Error
	if err != nil {
		panic(err)
	}
}

func CartTrigger(db *gorm.DB) {
	err := db.Exec(`
	DROP TRIGGER IF EXISTS UpdateTotalPrice on carts;

	CREATE TRIGGER UpdateTotalPrice
   AFTER  insert OR update of sub_price
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
	getID integer;
    
BEGIN
  SELECT customer_transaction_no,count(customer_transaction_no) INTO cnt,getID
  FROM transactions
  WHERE customer_id =  NEW.customer_id and (transactions.status='pending' or transactions.status='completed' ) group by transactions.id, customer_transaction_no order by customer_transaction_no desc limit 1;
 
	if (cnt = 0 and 
	(new.customer_transaction_no is null or new.customer_transaction_no = 0)
	and (new.status='pending' or new.status='completed')) then
	new.customer_transaction_no=+1;
	elseif (cnt > 0 and (new.status = 'pending' or new.status = 'completed'))
	then 
	new.customer_transaction_no = cnt +1;
	elseif (new.status = 'draft') then
	new.customer_transaction_no = 0;
	else
	new.customer_transaction_no=1;

end if;
  RETURN NEW;
END;
$BODY$;`).Error
	if err != nil {
		panic(err)
	}
}

func CheckCustomerCountAfterUpdate(db *gorm.DB) {
	err := db.Exec(`CREATE OR REPLACE FUNCTION public.checkcustomercountafterupdate()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    IMMUTABLE
    COST 100
    SET event_triggers=false
AS $BODY$
DECLARE cnt INTEGER;

   
BEGIN
  SELECT customer_transaction_no INTO cnt
  FROM transactions
  WHERE customer_id =  NEW.customer_id and (transactions.status='pending' or transactions.status='completed') group by transactions.id, customer_transaction_no order by customer_transaction_no desc limit 1;

	if (cnt = 1) then
	NEW.customer_transaction_no = new.customer_transaction_no;
	elseif ((new.status = 'pending' or new.status = 'completed') and old.customer_transaction_no = 0 ) then
  NEW.customer_transaction_no = cnt + 1;
end if;
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
