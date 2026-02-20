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

	updateUserQuery = `
	UPDATE users SET
		username = :username,
		email = :email,
		first_name = :first_name,
		last_name = :last_name,
		phone = :phone,
		profile_image = :profile_image,
		bio = :bio,
		date_of_birth = :date_of_birth,
		is_active = :is_active,
		is_verified = :is_verified,
		updated_at = :updated_at
	WHERE id = :id
	AND deleted_at IS NULL;
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
	WHERE email = :email
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
	WHERE id = :id
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
