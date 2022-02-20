SELECT city, country FROM city
INNER JOIN country ON city.country_id = country.country_id;

//City tablosundaki şehirlerin  country tablosundaki ülkelerini beraber getiren INNER JOIN sorgusudur.

SELECT payment_id ,first_name, last_name FROM customer
INNER JOIN payment ON payment.customer_id = customer.customer_id;

//Customer tablosundaki first_name ve last_name ile payment tablosundaki payment_id sütunlarını beraber getiren sorgudur.

SELECT rental_id ,first_name, last_name FROM customer
INNER JOIN rental ON rental.customer_id = customer.customer_id;

//Customer tablosundaki first_name ve last_name ile rental tablosundaki rental_id sütunlarını beraber getiren sorgudur.
