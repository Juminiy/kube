SELECT count(*)
FROM `tbl_product`
WHERE
(((
    (1!=1
    OR ((1!=1 OR ((1=1 AND `name` = "Beer") AND `desc` = "Local lager beer")) OR (1=1 AND `code` = 100007)))
    OR ((1!=1 OR ((1=1 AND `name` = "Noodles") AND `desc` = "Instant noodles")) OR (1=1 AND `code` = 100008)))
    OR ((1!=1 OR ((1=1 AND `name` = "Shampoo") AND `desc` = "Herbal shampoo")) OR (1=1 AND `code` = 100009)))
    OR ((1!=1 OR ((1=1 AND `name` = "Toothpaste") AND `desc` = "Mint toothpaste")) OR (1=1 AND `code` = 100010)))
AND `tbl_product`.`tenant_id` = 114514
AND `tbl_product`.`deleted_at` IS NULL
