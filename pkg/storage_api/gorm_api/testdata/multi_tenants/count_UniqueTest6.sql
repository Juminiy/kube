SELECT count(*)
FROM `tbl_unique_test6`
WHERE
(((1!=1
    OR (((1=1 AND `char1` = 5) AND `char2` = 2) AND `char3` = 1))
    OR ((1=1 AND `name` = "RR") AND `desc` = "ff-06"))
    OR ((1=1 AND `height` = 183) AND `weight` = 73)
)
AND `tbl_unique_test6`.`deleted_at` IS NULL
