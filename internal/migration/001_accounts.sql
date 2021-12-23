CREATE TABLE accounts (
    "id" SERIAL PRIMARY KEY,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL
);

INSERT INTO accounts (
    "email",
    "password"
) VALUES (
    'test@test.com',
    crypt('password', gen_salt('bf'))
);
