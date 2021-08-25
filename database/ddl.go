package database

const JuniorDDL = `create table if not exists Juniors
(
  id serial primary key,
  name varchar(30) not null
)`

const ClientsDDL = `create table if not exists Clients
(
  id serial primary key,
  Name varchar(20) not null,
  surname varchar(30) not null)`

const AccountsDDL = `create table if not exists Accounts
(
  id serial primary key,
  AccountNumber text not null,
  ClientId int reference Clients(id) on delete cascade,
  MerchantId int reference Merchants(id) on delete cascade,
  Balance int default 0,
  SWIFT text null,
  CreatedAt date not null,
  UpdateAt date not null,
  Locked boolean,
  Remove boolean
)`

const ATMsDDL = `create table if not exists ATMs
(
  id serial primary key,
  Address text not null,
  Balance int default 0,
  Lock boolean
)`

const MerchantsDDL = `create table if not exists Merchants
(
  id serial primary key,
  CompanyName text not null,
  Address text not null,
  CompanyDescription text not null,
  OwnerName varchar(20) not null,
  OwnerSurname varchar(20) not null,
  Login text not null unique,
  Password text not null,
  Locked boolean not null,
  CreatedAt date not null,
  UpdatedAt date,
  Remove boolean
)`

const ServicesDDL = `create table if not exists Services
(
  id serial primary key,
  Name text not null,
  MerchantID text not null references Merchants(id) on delete cascade,
  Price int default 200
)`

const TransactionsMerchantsDDL = `create table if not exists Transactions Merchants
(
  id serial primary key,
  SenderAcc text not null references Merchants(id) on delete cascade,
  ReceiverAcc text not null references Merchants(id) on delete cascade,
  Amount int default 0,
  CreatedAt date not null
)`

const TransactionsClientsDDL = `create table if not exists Transactions Clients
(
  id serial primary key,
  SenderAcc text not null references Clients(id) on delete cascade,
  ReceiverAcc text not null references Clients(id) on delete cascade,
  Amount int default 0,
  CreatedAt date not null
)`

//TODO: clientsDDL / accountsDDL / atmsDDL / merchantsDDL / servicesDDL / transactionsDDL