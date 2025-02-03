SELECT `id`,`name`,`desc`,`code`,`price`
FROM `tbl_product`
WHERE
(   id = 1
    AND name LIKE ""
    OR id = 1
    OR id = 2
    AND NOT id = 3
    AND NOT id = 4
)
AND `tbl_product`.`tenant_id` = 114514
AND `tbl_product`.`deleted_at` IS NULL
ORDER BY
    id desc,
    id asc,
    id DESC,
    id ASC
LIMIT 10
