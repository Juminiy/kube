SELECT count(*)
FROM `tbl_unique_test1`
WHERE
(
  (1!=1
  OR ((1=1 AND `name` = "Galaxy") AND `number_id` = "0019527"))
  OR ((1=1 AND `number_id` = "0019527") AND `birth` = 1919730)
)
AND `tbl_unique_test1`.`deleted_at` IS NULL
