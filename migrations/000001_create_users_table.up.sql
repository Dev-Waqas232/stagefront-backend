CREATE TYPE user_role AS ENUM ('organizer','attendee');

CREATE TABLE users
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role user_role NOT NULL
);

