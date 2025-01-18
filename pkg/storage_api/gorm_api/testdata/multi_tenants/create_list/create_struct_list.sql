SELECT count(*)
FROM `tbl_product`
WHERE
(((((1!=1 OR ((1!=1 OR ((1=1 AND `name` = "Milk") AND `desc` = "Fresh milk")) OR (1=1 AND `code` = 100001))) OR ((1!=1 OR ((1=1 AND `name` = "Bread") AND `desc` = "Whole wheat bread")) OR (1=1 AND `code` = 100002))) OR ((1!=1 OR ((1=1 AND `name` = "Rice") AND `desc` = "Long grain rice")) OR (1=1 AND `code` = 100003))) OR ((1!=1 OR ((1=1 AND `name` = "Eggs") AND `desc` = "Free-range eggs")) OR (1=1 AND `code` = 100004))) OR ((1!=1 OR ((1=1 AND `name` = "Chicken") AND `desc` = "Fresh chicken breast")) OR (1=1 AND `code` = 100006)))
AND `tbl_product`.`tenant_id` = 114514
AND `tbl_product`.`deleted_at` IS NULL
