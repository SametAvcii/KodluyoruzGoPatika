(
	SELECT first_name FROM customer
)
UNION
(
	SELECT first_name FROM actor
)

// Bu sorgu actor ve customer tablolarındaki first_name sütunlarını beraber getirir.

(
	SELECT first_name FROM customer
)
INTERSECT
(
	SELECT first_name FROM actor
)

// Bu sorgu actor ve customer tablolarındaki first_name sütunlarında kesişen verileri getirir.

(
	SELECT first_name FROM customer
)
EXCEPT
(
	SELECT first_name FROM actor
)

// Bu sorgu actor ve customer tablolarındaki first_name sütunlarında customer tablosunda bulunan actor tablosunda bulunmayan verileri getirir.

