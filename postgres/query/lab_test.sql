-- name: CreateLabTest :one
-- description: Create a lab test
INSERT INTO lab_tests (test_name,
                       batch_code,
                       test_id_code,
                       lab_facility_name,
                       test_performed_date_time,
                       overall_passed,
                       test_type_name,
                       test_passed,
                       test_comment,
                       thc_total_percent,
                       thc_total_value,
                       cbd_percent,
                       cbd_value,
                       terpene_total_percent,
                       terpene_total_value,
                       thc_a_percent,
                       thc_a_value,
                       delta9_thc_percent,
                       delta9_thc_value,
                       delta8_thc_percent,
                       delta8_thc_value,
                       thc_v_percent,
                       thc_v_value,
                       cbd_a_percent,
                       cbd_a_value,
                       cbn_percent,
                       cbn_value,
                       cbg_a_percent,
                       cbg_a_value,
                       cbg_percent,
                       cbg_value,
                       cbc_percent,
                       cbc_value,
                       total_cannabinoid_percent,
                       total_cannabinoid_value)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24,
        $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35)
RETURNING *;

-- name: GetLabTest :one
-- description: Get a lab test by ID
SELECT *
FROM lab_tests
WHERE id = $1
LIMIT 1;

-- name: ListLabTests :many
-- description: List all lab tests
SELECT *
FROM lab_tests
ORDER BY created_at;

-- name: UpdateLabTest :one
-- description: Update a lab test
UPDATE lab_tests
SET test_name                 = COALESCE(sqlc.narg(test_name), test_name),
    batch_code                = COALESCE(sqlc.narg(batch_code), batch_code),
    test_id_code              = COALESCE(sqlc.narg(test_id_code), test_id_code),
    lab_facility_name         = COALESCE(sqlc.narg(lab_facility_name), lab_facility_name),
    test_performed_date_time  = COALESCE(sqlc.narg(test_performed_date_time), test_performed_date_time),
    overall_passed            = COALESCE(sqlc.narg(overall_passed), overall_passed),
    test_type_name            = COALESCE(sqlc.narg(test_type_name), test_type_name),
    test_passed               = COALESCE(sqlc.narg(test_passed), test_passed),
    test_comment              = COALESCE(sqlc.narg(test_comment), test_comment),
    thc_total_percent         = COALESCE(sqlc.narg(thc_total_percent), thc_total_percent),
    thc_total_value           = COALESCE(sqlc.narg(thc_total_value), thc_total_value),
    cbd_percent               = COALESCE(sqlc.narg(cbd_percent), cbd_percent),
    cbd_value                 = COALESCE(sqlc.narg(cbd_value), cbd_value),
    terpene_total_percent     = COALESCE(sqlc.narg(terpene_total_percent), terpene_total_percent),
    terpene_total_value       = COALESCE(sqlc.narg(terpene_total_value), terpene_total_value),
    thc_a_percent             = COALESCE(sqlc.narg(thc_a_percent), thc_a_percent),
    thc_a_value               = COALESCE(sqlc.narg(thc_a_value), thc_a_value),
    delta9_thc_percent        = COALESCE(sqlc.narg(delta9_thc_percent), delta9_thc_percent),
    delta9_thc_value          = COALESCE(sqlc.narg(delta9_thc_value), delta9_thc_value),
    delta8_thc_percent        = COALESCE(sqlc.narg(delta8_thc_percent), delta8_thc_percent),
    delta8_thc_value          = COALESCE(sqlc.narg(delta8_thc_value), delta8_thc_value),
    thc_v_percent             = COALESCE(sqlc.narg(thc_v_percent), thc_v_percent),
    thc_v_value               = COALESCE(sqlc.narg(thc_v_value), thc_v_value),
    cbd_a_percent             = COALESCE(sqlc.narg(cbd_a_percent), cbd_a_percent),
    cbd_a_value               = COALESCE(sqlc.narg(cbd_a_value), cbd_a_value),
    cbn_percent               = COALESCE(sqlc.narg(cbn_percent), cbn_percent),
    cbn_value                 = COALESCE(sqlc.narg(cbn_value), cbn_value),
    cbg_a_percent             = COALESCE(sqlc.narg(cbg_a_percent), cbg_a_percent),
    cbg_a_value               = COALESCE(sqlc.narg(cbg_a_value), cbg_a_value),
    cbg_percent               = COALESCE(sqlc.narg(cbg_percent), cbg_percent),
    cbg_value                 = COALESCE(sqlc.narg(cbg_value), cbg_value),
    cbc_percent               = COALESCE(sqlc.narg(cbc_percent), cbc_percent),
    cbc_value                 = COALESCE(sqlc.narg(cbc_value), cbc_value),
    total_cannabinoid_percent = COALESCE(sqlc.narg(total_cannabinoid_percent), total_cannabinoid_percent),
    total_cannabinoid_value   = COALESCE(sqlc.narg(total_cannabinoid_value), total_cannabinoid_value)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteLabTest :exec
-- description: Delete a lab test by ID
DELETE
FROM lab_tests
WHERE id = $1;

-- name: AssignLabTestToPackage :exec
-- description: Assign a lab test to a package via junction table
INSERT INTO lab_tests_packages (lab_test_id, package_id)
VALUES ($1, $2);
