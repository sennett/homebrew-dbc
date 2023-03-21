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
2023/03/21 17:27:16 DBConnect
2023/03/21 17:27:16 Using bastion: i-0a884a55c892751e4
2023/03/21 17:27:16 Opening connection
2023/03/21 17:27:16 Connection Open at localhost:5432

...

```

You may connect to the database using `localhost:5432` as if `localhost` was the DB DNS.

`psql -h localhost -p 5432 -U ${user} -d ${db} --password`

## IAM Authentication

If your databases are super cool and secure, IAM Authentication will be enabled.

To use `dbc` with an IAM Enabled Database, you can use the `--iam` flag to generate a token while opening your connection!

```

➜ dbc connect -H ${db_host} --iam
2023/03/21 17:28:30 DBConnect IAM
2023/03/21 17:28:30 Token: ...
2023/03/21 17:28:30 Using bastion: i-0a884a55c892751e4
2023/03/21 17:28:30 Opening connection
2023/03/21 17:28:30 Connection Open at localhost:5432

```

Then when connecting to your DB...

`psql -h localhost -p 5432 -U ${user} -d ${db} --password`

Paste the token!

Or... If you're very fancy:

`export PGPASSWORD=${token} && psql -h localhost -p 5432 -U ${user} -d ${db}`

## exit status `255`

If you're experiencing trouble opening a session, and you're recieving a `255` error, it's likely due to a missing AWS SSM Plugin installation.

Run this handy script!

`wget -O - https://raw.githubusercontent.com/birdiecare/homebrew-dbc/main/install_ssm_plugin.sh | sh`

