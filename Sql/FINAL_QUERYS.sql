SELECT replacement_cost,title FROM film 
WHERE title LIKE 'K%'
GROUP BY replacement_cost,title
ORDER BY replacement_cost ASC
LIMIT 4;

//film tablosundan 'K' karakteri ile başlayan en uzun ve replacenet_cost u en düşük 4 filmin sıralamasını verir.

SELECT COUNT(rating) FROM film
GROUP BY rating
ORDER BY rating DESC
LIMIT 1;

//film tablosunda içerisinden en fazla sayıda film bulunduran rating kategorisini verir.

SELECT COUNT(customer.customer_id),first_name,last_name FROM payment
JOIN customer ON customer.customer_id =payment.customer_id
GROUP BY customer.customer_id,first_name,last_name
ORDER BY COUNT(customer.customer_id) DESC
LIMIT 1;

//customer tablosunda en çok alışveriş yapan müşterinin adıNU verir.




