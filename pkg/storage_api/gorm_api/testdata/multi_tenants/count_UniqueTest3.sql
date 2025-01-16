SELECT count(*)
FROM `tbl_unique_test3`
WHERE
(1!=1 OR (1=1 AND `name` = "RR"))
AND `tbl_unique_test3`.`deleted_at` IS NULL
