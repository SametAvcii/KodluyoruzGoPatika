SELECT COUNT(*) FROM film WHERE
length >
(
	SELECT AVG(length) FROM film
);

//Bu sorguda film tablosundaki filmlerin kaç tanesinin uzunluğu ortalama uzunluktan büyük olduğunun sayısını verir.

SELECT COUNT(*) FROM film WHERE rental_rate =
(
	SELECT MAX(rental_rate) FROM film
);

//Bu sorguda film tablosundaki filmlerin kaç tanesinin rental_rate değeri en büyük rental_rate değerine eşit olduğunun verisini verir.  

SELECT * FROM film WHERE rental_rate =
(
	SELECT MIN(rental_rate) FROM film
)
AND
replacement_cost =(
SELECT MIN(replacement_cost)FROM FILM
)

//film tablosunda en düşük rental_rate ve en düşün replacement_cost değerlerine sahip filmlerin sıralar.

SELECT customer_id, COUNT(payment_id) payment_count
FROM payment
GROUP BY customer_id
HAVING COUNT(payment_id) = 

(SELECT COUNT(payment_id) 
FROM payment 
GROUP BY customer_id 
ORDER BY COUNT(payment_id) DESC 
LIMIT 1);
//Payment tablosunda en çok alışveriş yapan müşterilerin sıralanmasıdır.