
/* create sandbox schema if does not already exists */
CREATE SCHEMA IF NOT EXISTS wallet_sandbox;

/* create accounts table */
CREATE TABLE IF NOT EXISTS wallet_sandbox.accounts (
    id serial primary key,
    name text,
    created_at timestamp default CURRENT_TIMESTAMP
);

/* create wallets table */
CREATE TABLE IF NOT EXISTS wallet_sandbox.wallets (
    id serial primary key, 
    balance numeric default 0 CHECK (balance >= 0), 
    account_id integer references wallet_sandbox.accounts(id),
    created_at timestamp default CURRENT_TIMESTAMP
);

/* create transanctions table */
CREATE TABLE IF NOT EXISTS wallet_sandbox.transactions (
    id serial primary key, 
    type text not null,
    from_account_id text not null, 
    to_account_id text not null,
    amount text not null,
    status text not null default 'PENDING',
    processed_at timestamp default CURRENT_TIMESTAMP,
    created_at  timestamp default CURRENT_TIMESTAMP
);
