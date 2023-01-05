-- name: CreateStrain :one
-- description: Create a strain
INSERT INTO strains (name,
                     type,
                     yield_average,
                     terp_average_total,
                     terp_1,
                     terp_1_value,
                     terp_2,
                     terp_2_value,
                     terp_3,
                     terp_3_value,
                     terp_4,
                     terp_4_value,
                     terp_5,
                     terp_5_value,
                     thc_average,
                     total_cannabinoid_average,
                     light_dep_2022,
                     fall_harvest_2022,
                     quantity_available)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING *;

-- name: GetStrain :one
-- description: Get a strain by ID
SELECT *
FROM strains
WHERE id = $1
LIMIT 1;

-- name: ListStrains :many
-- description: List all strains
SELECT *
FROM strains
ORDER BY name;

-- name: UpdateStrain :one
-- description: Update a strain
UPDATE strains
SET name                      = COALESCE(sqlc.narg(name), name),
    type                      = COALESCE(sqlc.narg(type), type),
    yield_average             = COALESCE(sqlc.narg(yield_average), yield_average),
    terp_average_total        = COALESCE(sqlc.narg(terp_average_total), terp_average_total),
    terp_1                    = COALESCE(sqlc.narg(terp_1), terp_1),
    terp_1_value              = COALESCE(sqlc.narg(terp_1_value), terp_1_value),
    terp_2                    = COALESCE(sqlc.narg(terp_2), terp_2),
    terp_2_value              = COALESCE(sqlc.narg(terp_2_value), terp_2_value),
    terp_3                    = COALESCE(sqlc.narg(terp_3), terp_3),
    terp_3_value              = COALESCE(sqlc.narg(terp_3_value), terp_3_value),
    terp_4                    = COALESCE(sqlc.narg(terp_4), terp_4),
    terp_4_value              = COALESCE(sqlc.narg(terp_4_value), terp_4_value),
    terp_5                    = COALESCE(sqlc.narg(terp_5), terp_5),
    terp_5_value              = COALESCE(sqlc.narg(terp_5_value), terp_5_value),
    thc_average               = COALESCE(sqlc.narg(thc_average), thc_average),
    total_cannabinoid_average = COALESCE(sqlc.narg(total_cannabinoid_average), total_cannabinoid_average),
    light_dep_2022            = COALESCE(sqlc.narg(light_dep_2022), light_dep_2022),
    fall_harvest_2022         = COALESCE(sqlc.narg(fall_harvest_2022), fall_harvest_2022),
    quantity_available        = COALESCE(sqlc.narg(quantity_available), quantity_available)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteStrain :exec
-- description: Delete a strain by ID
DELETE
FROM strains
WHERE id = $1;
