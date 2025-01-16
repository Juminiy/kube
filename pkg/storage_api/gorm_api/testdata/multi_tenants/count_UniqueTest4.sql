SELECT count(*)
FROM `tbl_unique_test4`
WHERE
((1!=1
  OR (1=1 AND `name` = "RR"))
  OR (1=1 AND `r_id` = 666))
AND `tbl_unique_test4`.`deleted_at` IS NULL
