## Messaging App

### requirements

    golang version >= 1.8
    docker

### How execute
    docker pull mysql:8.0
    docker compose up -d

### manual execution of migration

    cd repository\mysql
    sql-migrate up -config="dbconfig.yml" -env="development" 
    sql-migrate status -config="dbconfig.yml" -env="development" //to see status
    --if you want rollback scripts
    sql-migrate down -config="dbconfig.yml" -env="development"