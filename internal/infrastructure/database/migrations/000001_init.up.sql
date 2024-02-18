create table accounts (
  id varchar(36) not null primary key,
  balance decimal(10, 2) not null,
  createdAt timestamp not null default current_timestamp,
  updatedAt timestamp not null default current_timestamp,
  deletedAt timestamp
);