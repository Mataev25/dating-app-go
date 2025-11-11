CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	name VARCHAR(100) NOT NULL,
	age INTEGER NOT NULL CHECK (age >= 18 AND age <= 100),
	gender VARCHAR(10) NOT NULL CHECK (gender IN ('male', 'female')),
	description TEXT,
	looking_for VARCHAR(10) NOT NULL CHECK (looking_for IN ('male', 'female', 'both')),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_age_gender ON users(age, gender);
CREATE INDEX idx_users_looking_for ON users(looking_for);
