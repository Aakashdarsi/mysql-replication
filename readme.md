## MySQL Replication Setup 

### Prerequisites
- Docker
- Docker Compose

### Important Notes
- The `docker-compose.yml` file is configured to use MySQL 8.0.36.


### Setup
```bash
docker-compose up -d
```

### Verify 
```bash
docker-compose ps
```


### Connect to MySQl
```bash
docker exec -it mysql-primary mysql -u root -p rootpass
```
create replication user on primary database

```sql
CREATE USER 'repl'@'%' IDENTIFIED WITH mysql_native_password BY 'replpass';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
FLUSH PRIVILEGES;
```

Get Primary Binlog Position 
Still on Primary 
```sql
SHOW MASTER STATUS;
```

Configure Replica to Follow Primary

```bash
docker exec -it mysql-replica mysql -uroot -prootpass
```

```sql
STOP REPLICA;

CHANGE REPLICATION SOURCE TO
  SOURCE_HOST='mysql-primary',
  SOURCE_USER='repl',
  SOURCE_PASSWORD='replpass',
  SOURCE_LOG_FILE='mysql-bin.000003',
  SOURCE_LOG_POS=157;

START REPLICA;

```

Verify Replication
```sql
SHOW SLAVE STATUS\G
```

Test It (Write on Primary, Read on Replica)
```sql
USE appdb;
CREATE TABLE users(id INT PRIMARY KEY, name VARCHAR(50));
INSERT INTO users VALUES (1, 'Aakash');
```

#### Important Notes 

- To Disable Write on Replica 

    ```sql 
    STOP SLAVE 
    SET GLOBAL read_only = ON; 
    SET GLOBAL super_read_only = ON; 
    ```

- Then

    ```sql
    START REPLICA SQL_THREAD;
    ```

- If there are any changes you can down your docker-compose and remove the volumes and then start again. 
    ```bash
    docker-compose down -v
    docker-compose up -d
    ```
