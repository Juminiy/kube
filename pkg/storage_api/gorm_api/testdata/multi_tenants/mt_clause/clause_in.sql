SELECT *
FROM `tbl_product`
WHERE
`tbl_product`.`tenant_id` IN (1,2,3,4,5,114514)
AND `tbl_product`.`deleted_at` IS NULL
