SELECT *
FROM `tbl_product`
WHERE
`tbl_product`.`tenant_id` = "9527"
AND `tbl_product`.`deleted_at` IS NULL
