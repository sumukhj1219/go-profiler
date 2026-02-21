package repository

const (
	createUserQuery = `
	INSERT INTO users (
		username,
		email,
		password,
		first_name,
		last_name,
		phone,
		profile_image,
		bio,
		date_of_birth,
		is_active,
		is_verified,
		created_at,
		updated_at
	) VALUES (
		:username,
		:email,
		:password,
		:first_name,
		:last_name,
		:phone,
		:profile_image,
		:bio,
		:date_of_birth,
		:is_active,
		:is_verified,
		:created_at,
		:updated_at
	)
	RETURNING id;
	`

	getUserByEmailQuery = `
	SELECT
	id,
	username,
	email,
	password,
	first_name,
	last_name,
	phone,
	profile_image,
	bio,
	date_of_birth,
	is_active,
	is_verified,
	created_at,
	updated_at,
	deleted_at
	FROM users
	WHERE email = $1
	AND deleted_at IS NULL;
	`

	getUserByIDQuery = `
	SELECT
	id,
	username,
	email,
	password,
	first_name,
	last_name,
	phone,
	profile_image,
	bio,
	date_of_birth,
	is_active,
	is_verified,
	created_at,
	updated_at,
	deleted_at
	FROM users
	WHERE id = $1
	AND deleted_at IS NULL;
	`

	getUsersQuery = `
	SELECT
		id,
		username,
		email,
		first_name,
		last_name,
		phone,
		profile_image,
		bio,
		date_of_birth,
		is_active,
		is_verified,
		created_at,
		updated_at,
		deleted_at
	FROM users
	WHERE deleted_at IS NULL
	ORDER BY created_at DESC;
	`
)
