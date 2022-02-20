SELECT country, city FROM city
LEFT JOIN country ON country.country_id=city.country_id;

////City tablosundaki şehirlerin  country tablosundaki ülkelerini beraber getiren LEFT JOIN sorgusudur.

SELECT payment_id ,first_name, last_name FROM payment
RIGHT JOIN customer ON payment.customer_id = customer.customer_id;

//Customer tablosundaki first_name ve last_name ile payment tablosundaki payment_id sütunlarını RIGHT JOIN yapısına göre beraber getiren sorgudur.

SELECT rental_id ,first_name, last_name FROM customer
FULL JOIN rental ON rental.customer_id = customer.customer_id;

//Customer tablosundaki first_name ve last_name ile rental tablosundaki rental_id sütunlarını FULL JOIN yapısına göre  beraber getiren sorgudur.
