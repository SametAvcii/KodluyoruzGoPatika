SELECT AVG(rental_rate) FROM film;

//Bu sorgu bize rental rate sütunundaki değerlerin ortalamasını verir.

SELECT COUNT(*) FROM film WHERE title LIKE 'C%';

//Bu sorgu bize film tablosundaki ismi 'C' ile başlayan verilerin sayısını verir.

SELECT MAX(length) FROM film WHERE rental_rate=0.99;

//Bu sorgu bize rental_rate değeri 0.99 olan en uzun filmi verir.

SELECT COUNT(DISTINCT (replacement_cost)) FROM film WHERE length >150;

//Bu sorgu bize uzunluğu 150den büyük olan filmlerin kaç farklı replacement_cost değeri olduğunu gösterir.

