SELECT *
FROM `tbl_product`
WHERE
(
    (code BETWEEN 100007 AND 100010) OR code = 300179
)
AND `tbl_product`.`tenant_id` IN (1,2,3,4,5,114514)
AND `tbl_product`.`deleted_at` IS NULL
