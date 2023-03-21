# `dbc`
**D**ata**b**ase **C**onnect

Connect to databases securely, utilising AWS SSM, EC2 & RDS IAM Authentication.

## Install `dbc`

`brew tap birdiecare/dbc`

`brew install dbc`

`dbc -h`

## Use `dbc`

The connect command opens a WebSocket port-forwarding session through an available Bastion host.
A connection to the specified Database host `-H` will then be available at `localhost:localport` (localport: `-lp`).

## Password Authentication

If the database you're connecting to doesn't have AWS IAM Authentication enabled, or doesn't have Users with the `rds_iam` role, you'll need to use a password to authenticate with the DB once your connection is open.

The following command will be enough to open a connection with a specified DB host with Birdie-specific defaults.

`dbc connect -H ${db_host}`

Once the connection is open:

```

➜ dbc connect -H ${db_host}
2023/03/21 10:50:44 DBConnect
2023/03/21 10:50:44 Using bastion: ${bastion_id}
2023/03/21 10:50:44 Opening connection @ localhost:5432
2023/03/21 10:50:44 listening on [::]:5432

...

```

You may connect to the database using `localhost:5432` as if `localhost` was the DB DNS.

`psql -h localhost -p 5432 -U ${user} -d ${db} --password`

## IAM Authentication

TODO
