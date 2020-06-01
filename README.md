# auth service

## Init database

#### Init postgres
```shell script
docker pull postgres
docker run --name auth-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
psql -h 172.17.0.2 -p 5432 -U postgres
```

NOTE: Container ip address you can find in output ```docker inspect``` command.

#### Create user and database
```sql
create database auth;
create user user_auth with encrypted password 'secretpassword';
grant all privileges on database auth to user_auth;
```

#### Init tables
```shell script
go build
./auth --vv -c ./examples/auth.ini adm initdb
```

## Run service

```shell script
go build
./auth --vv -c ./examples/auth.ini srv
```