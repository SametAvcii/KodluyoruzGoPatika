SELECT * FROM country WHERE country LIKE 'A%a';

//country değeri 'A' ile başlayıp 'a' ile biten verileri getirir.

SELECT * FROM country WHERE LENGTH (country) > 6 AND country LIKE '%n';

//Harf sayısı 6-dan büyük olan ne 'n' ile biten verileri getirir.

SELECT * FROM film WHERE title ILIKE '%T%T%T%T';

//İçinde en az 4 tane t harfi bulunan (büyük veya küçük) verileri getirir.

SELECT * FROM film WHERE title LIKE 'C%' AND length >90 AND rental_rate ='2.99';

//'C' ile başlayan length değeri 90dan büyük ve rental_rate değeri 2.99 olan verileri getirir.



