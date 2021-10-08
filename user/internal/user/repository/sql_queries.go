package repository

const (
	createUserQuery = `INSERT INTO auth.users (id, first_name, last_name, email, password, avatar, role) 
	VALUES ($1,$2,$3,$4,$5,$6,$7) 
	RETURNING id, first_name, last_name, email, avatar, role, updated_at, created_at`

	getUserByIDQuery = `SELECT id, first_name, last_name, email, avatar, role, updated_at, created_at FROM auth.users WHERE id = $1`

	getUserByEmail = `SELECT id, first_name, last_name, email, password, avatar, role, updated_at, created_at 
	FROM auth.users WHERE email = $1`

	updateUserQuery = `UPDATE auth.users 
		SET first_name = COALESCE(NULLIF($1, ''), first_name), 
	    last_name = COALESCE(NULLIF($2, ''), last_name), 
	    email = COALESCE(NULLIF($3, ''), email), 
	    role = COALESCE(NULLIF($4, '')::role, role)
		WHERE id = $5
	    RETURNING id, first_name, last_name, email, role, avatar, updated_at, created_at`

	updateAvatarQuery = `UPDATE auth.users SET avatar = $1 WHERE id = $2 
	RETURNING id, first_name, last_name, email, role, avatar, updated_at, created_at`
)
