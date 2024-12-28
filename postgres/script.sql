drop table if exists account_transaction;
drop table if exists daily_earning;
drop table if exists investment;
drop table if exists subscription;
drop table if exists transfer;
drop table if exists wallet;
drop table if exists weekly_payout;
drop table if exists withdrawal;
drop table if exists deposit;
drop table if exists account;

drop table if exists user_setting;
drop table if exists login_info;
drop table if exists security_code;
drop table if exists notification;
drop table if exists transaction;
drop table if exists payment_link;
drop table if exists beneficiary;


CREATE TABLE IF NOT EXISTS account 
(
    id character varying(64) NOT NULL PRIMARY KEY,
    referral_code character varying(256) not null UNIQUE,
    username character varying(256) NOT NULL UNIQUE,
    password character varying(256) NOT NULL,
    email character varying(256) not null,
    phone_number character varying(32) not null,
    created_at bigint NOT NULL,
    first_name character varying(256) NOT NULL default '',
    last_name character varying(256) NOT NULL default '',
    referral_id character varying(256) default '',
    referral_id_2 character varying(256) default '',
    referral_id_3 character varying(256) default '',
    withdrawal_addresss character varying(256) NOT NULL default '',
    balance bigint NOT NULL default 0,
    principal bigint NOT NULL default 0,
    matured_principal bigint not null default 0
);

CREATE TABLE IF NOT EXISTS wallet 
(
    id character varying(64) NOT NULL PRIMARY KEY,
    address character varying(64) NOT NULL UNIQUE,
    private_key character varying(124) NOT NULL UNIQUE,
    coin_symbol character varying(32) NOT NULL,
    account_id character varying(64) NOT NULL REFERENCES account(id)
);

CREATE TABLE IF NOT EXISTS package 
(
    id character varying(64) NOT NULL PRIMARY KEY,
    name character varying(64) NOT NULL UNIQUE,
    price BIGINT NOT NULL,
    min_return_per_month INT NOT NULL,
    max_return_per_month INT NOT NULL,
    trades_per_day INT NOT NULL,
    accuracy INT NOT NULL
);

CREATE TABLE IF NOT EXISTS subscription
(
    id character varying(64) not null primary key,
    account_id character varying(64) not null references account(id),
    package_id character varying(64) not null references package(id),
    start_date bigint not null,
    end_date bigint not null
);

create table if not exists daily_earning
(
    id serial not null primary key,
    account_id character varying(64) not null references account(id),
    date bigint not null,
    percentage int not null,
    principal bigint not null,
    unique(account_id, date)
);

create table if not exists deposit
(
    id character varying(64) not null primary key,
    amount bigint not null,
    account_id character varying(64) not null references account(id),
    date bigint not null,
    ref character varying(256) not null unique
);

create table if not exists transfer
(
    id character varying(64) not null primary key,
    amount bigint not null,
    sender_id character varying(64) not null references account(id),
    receiver_id character varying(64) not null references account(id),
    date bigint not null
);

create table if not exists withdrawal 
(
    id character varying(64) not null primary key,
    account_id character varying(64) not null references account(id),
    amount bigint not null,
    date bigint not null,
    destination character varying(256) not null,
    ref character varying(256) not null,
    status character varying(32) not null
);

CREATE TABLE IF NOT EXISTS account_transaction (
	id serial,
	account_id character varying(64) NOT NULL REFERENCES account(id),
	amount INT8 NOT NULL,
	tx_type VARCHAR(32) NOT NULL,
	opening_balance INT8 NOT NULL,
	closing_balance INT8 NOT NULL,
	date INT8 NOT NULL,
	description VARCHAR(256) NOT NULL,
	PRIMARY KEY(id)
);

create table if not exists investment (
    id character varying(64) not null primary key,
    account_id character varying(64) not null references account(id),
    amount bigint not null,
    date bigint not null,
    activation_date bigint not null,
    status int not null default 0
);

create table if not exists weekly_payout (
    id character varying(63) not null primary key,
    date bigint not null,
    amount bigint not null
);

create table if not exists referral_payout (
    id character varying(64) not null primary key,
    account_id character varying(64) not null references account(id),
    from_account_id character varying(64) not null references account(id),
    subscription_id character varying(64) not null references subscription(id),
    generation int not null,
    amount bigint not null,
    date bigint not null,
    payment_method int not null,
    payment_status int not null,
    payment_ref text not null,
    unique(subscription_id, generation)
);

create table if not exists notification (
    id character varying(64) not null primary key,
    account_id character varying(64) not null references account(id),
    status int not null,
    title character varying(128) not null,
    content character varying(500) not null,
    type int not null default 0,
    action_link character varying(64) not null default 0,
    action_text character varying(32) not null default 0,
    date bigint not null
);

create table if not exists trade_schedule (
    id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) not null references account(id),
    trade_no int not null,
    total_trades int not null,
    date bigint not null,
    target_profit_percentage int not null default 0,
    start_date int not null default 0,
    unique(account_id, date, trade_no)
);

create table if not exists trade (
    id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) not null references account(id), 
    trade_no int not null,
    date bigint not null,
    start_date bigint not null,
    end_date bigint not null,
    amount bigint not null,
    profit bigint not null,
    unique(account_id, date, trade_no)
);

create table if not exists user_setting (
    id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) not null references account(id),
    config_key character varying(64) not null,
    config_value character varying(500) not null,
    unique(account_id, config_key)
);

-- insert into user_setting(account_id, config_key, config_value) 
-- select id, 'EXCHANGE_RATE', 735 from account where email = 'ademuanthony@gmail.com';

create table if not exists login_info (
    id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) not null references account(id),
    date bigint not null,
    ip character varying(128) not null,
    platform character varying(128) not null
);

CREATE TABLE IF NOT EXISTS security_code (
	id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) not null references account(id),
	code VARCHAR(3200) NOT NULL,
	date bigint NOT NULL
);

create table if not exists transaction (
    id uuid not null default gen_random_uuid() primary key,
    bank_name character varying(64) not null,
    account_number character varying(64) not null,
    account_name character varying(64) not null,
    amount int8 not null,
    token_amount character varying(64) not null,
    amount_paid character varying(64) not null,
    email character varying(64) not null,
    network character varying(64) not null,
    currency character varying(64) not null,
    wallet_address character varying(64) not null,
    private_key character varying(64) not null,
    payment_link character varying(64) not null,
    type character varying(28) not null,
    status character varying(28) not null
);

create table if not exists payment_link ( 
    permalink character(64) not null primary key,
    account_id character varying(64) references account(id),
    email  character varying(64) not null,
    accountName character varying(64) not null,
    accountNumber character varying(64) not null,
    bankName character varying(64) not null,
    fixAmount int8 not null,
    title character varying(100) not null,
    description character varying(500) not null
);

create table if not exists beneficiary (
    id uuid not null default gen_random_uuid() primary key,
    account_id character varying(64) references account(id),
    bank character varying(64) not null,
    account_number character varying(64) not null,
    account_name character varying(64) not null,
    country character varying(64) not null,
    beneficial_email character varying(64) not null
);

create table if not exists agent (
    id serial not null primary key,
    slack_username character varying(64) not null unique,
    name character varying(64) not null,
    balance int8 not null,
    status int not null
);

create table if not exists transaction_assignment(
    id serial not null primary key,
    agent_id int not null references agent(id),
    transaction_id uuid not null references transaction(id),
    amount int8 not null,
    date int8 not null,
    status int not null
);

create table if not exists dfc_holders (
    wallet_address character varying(32) not null primary key,
    balance character varying(32) not null
);

create table if not exists dfc_wallet_transaction (
    id uuid not null default gen_random_uuid() primary key,
    wallet_address character varying(64) not null,
    amount float8 not null
);

create table if not exists cgold_holders (
    wallet_address character varying(32) not null primary key,
    balance character varying(32) not null
);

create table if not exists stable_naira_transaction (
    id uuid not null primary key default gen_random_uuid(),
    account_id character varying(64) not null references account(id),
    amount int8 not null,
    amount_receivable int8 not null,
    agent_id int not null references agent(id),
    date int8 not null,
    payment_date int8 not null,
    completion_date int8 not null 
);

create table if not exists stable_naira_transaction_assignment(
    id serial not null primary key,
    agent_id int not null references agent(id),
    transaction_id uuid not null references stable_naira_transaction(id),
    amount int8 not null,
    date int8 not null,
    status int not null
);

create table if not exists deposit_wallet (
    address character varying(256) not null primary key,
    private_key character varying(256) not null,
    account_id character varying(64) not null references account(id),
    blockchain character varying(64) not null
);

create table if not exists wallet_balanace (
    id uuid not null primary key default gen_random_uuid(),
    wallet_address character varying(256) not null references deposit_wallet(address),
    balance character varying(64) not null,
    currency character varying(32) not null,
    network character varying(32) not null
);

create table if not exists crypto_deposit (
    transaction_hash character varying(256) not null primary key,
    wallet_address character varying(256) not null references deposit_wallet(address),
    token_amount character varying(256) not null,
    block_number int8 not null,
    currency character varying(64) not null,
    network character varying(64) not null,
    account_id character varying(64) not null references account(id),
    date int8 not null,
    status character varying(32) not null
);

create table if not exists crypto_asset (
    id uuid not null PRIMARY key default gen_random_uuid(),
    account_id character varying(64) not null references account(id),
    currency character varying(64) not null,
    balance int8 not null
);

create table if not exists trade_history (
    id uuid not null primary key default gen_random_uuid(),
    account_id character varying(64) not null,
    base_currency character varying(64) not null,
    qoute_currency character varying(64) not null,
    base_amount int8 not null,
    qoute_amount int8 not null,
    date int8 not null
);

create table if not exists kyc_form (
    id uuid not null primary key default gen_random_uuid(),
    account_id character varying(64) not null,
    first_name character varying(64) not null,
    last_name character varying(64) not null,
    middle_name character varying(64) not null,
    application_date int8 not null,
    date_of_birth int8 not null,
    identity_type character varying(64) not null,
    identity_isseu_date int8 not null,
    identity_expiry_date int8 not null,
    identity_country character varying(64) not null,
    country character varying(64) not null,
    state character varying(64) not null,
    address character varying(128) not null,
    identity_file character varying(256) not null,
    photo_file character varying(256) not null,
    verification_status int not null default 0
);

alter table account add role int not null default 0;

alter table transaction drop naira_amount;
alter table transaction add naira_amount int8 not null default 0;

alter table transaction add phone_number character varying(32) not null default '';
alter table transaction add mobile_network character varying(32) not null default '';
alter table transaction add smart_card_number character varying(32) not null default '';
alter table transaction add tv_provider character varying(32) not null default '';
alter table transaction add meter_number character varying(32) not null default '';
alter table transaction add power_provider character varying(200) not null default '';
alter table transaction add merchant_ref character varying(256) not null default '';
alter table transaction add success_url character varying(256) not null default '';
alter table transaction add failure_url character varying(256) not null default '';
alter table transaction add payment_method character varying(256) not null default '';


alter table beneficiary add phone_number character varying(32) not null default '';
alter table beneficiary add smart_card_number character varying(32) not null default '';
alter table beneficiary add meter_number character varying(32) not null default '';
alter table beneficiary add type int not null default 0;
alter table beneficiary add date int8 not null default 0;

alter table transaction add date int8 not null default 0;

alter table payment_link add date int8 not null default 0;
alter table payment_link add callback_url character varying(256) not null default '';

alter table transaction add request_id character varying(256) not null default '';
alter table transaction add package_id character varying(256) not null default '';
alter table transaction add electric_token character varying(256) not null default '';

alter table account add points int8 not null default 0;

alter table dfc_holders add sent int not null default 0;
alter table dfc_holders add hash character varying(124) not null default '';
alter table account add referral_earning int8 not null default 0;
alter table account add stable_naira int8 not null default 0;

alter table agent add account_number character varying(28) not null default '';
alter table agent add account_name character varying(128) not null default '';
alter table agent add bank_name character varying(128) not null default '';
alter table agent add stable_naira_balance int8 not null default '';

alter table transaction add billers_code character varying(64) not null default '';
alter table transaction add variation_code character varying(64) not null default '';
alter table transaction add operator_id character varying(64) not null default '';
alter table transaction add country_code character varying(64) not null default '';
alter table transaction add product_type_id character varying(64) not null default '';
