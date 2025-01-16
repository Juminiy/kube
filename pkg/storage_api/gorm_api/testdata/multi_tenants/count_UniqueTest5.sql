SELECT count(*)
FROM `tbl_unique_test5`
WHERE
(
  1!=1
  OR (((1=1 AND `name` = "RR")
        AND `desc` = "ff-60")
        AND `perm` = "06-21")
)
AND `tbl_unique_test5`.`deleted_at` IS NULL
