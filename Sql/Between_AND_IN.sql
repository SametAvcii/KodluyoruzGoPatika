SELECT * FROM FILM WHERE replacement_cost BETWEEN 12.99 AND 16.98; 

//replacement_cost değeri 12.99 ve 16,98 arasında olan verileri getirir.12.99 ve 16.98 dahil.

SELECT first_name, last_name FROM actor WHERE first_name IN('Penelope','Nick','Ed');

//first_name değerleri Penelope,Nick,Ed olan verileri getirir.

SELECT * FROM film WHERE rental_rate IN(0.99, 2.99, 4.99)AND replacement_cost IN(12.99, 15,99 ,28.99);

// rental_rate değeri 0.99, 2.99, 4.99 olan ve replacement_cost değerleri 12.99, 15,99 ,28.99 olan verileri getirir.