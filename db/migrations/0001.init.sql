BEGIN;

-- Accounts table
CREATE TABLE accounts (
   id SERIAL PRIMARY KEY,
   domain VARCHAR(255) NOT NULL UNIQUE,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL
);

-- Plans table
CREATE TABLE plans (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   cost DECIMAL(10,2) NOT NULL,
   billing_period_days INTEGER
);

-- Subscriptions table
CREATE TABLE subscriptions (
   id SERIAL PRIMARY KEY,
   account_id INTEGER REFERENCES accounts(id) ON DELETE CASCADE,
   plan_id INTEGER REFERENCES plans(id) ON DELETE RESTRICT,
   start_date TIMESTAMPTZ,
   end_date TIMESTAMPTZ,
   trial_start_date TIMESTAMPTZ,
   trial_end_date TIMESTAMPTZ,
   status SMALLINT,
   cost DECIMAL(10,2) NOT NULL
);

END;
