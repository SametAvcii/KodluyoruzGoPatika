SELECT * FROM film  WHERE title LIKE '%n' ORDER BY length DESC LIMIT 5;

//Bu sorgu bize film tablosundaki ismi n ile biten en uzun 5 filmi getirir.

SELECT * FROM film  WHERE title LIKE '%n' ORDER BY length ASC OFFSET 5 LIMIT 5;

//Bu sorgu bize film tablosundaki ismi n ile biten en kısa ikince 5 filmi (6,7,8,9,10.) getirir.

SELECT * FROM customer  WHERE store_id=1 ORDER BY last_name DESC LIMIT 4;

//Bu sorgu bize customer tablosundaki verileri last_name sütünuna göre azalan sırada sıralayıp ilk 4 veriyi getirir.
