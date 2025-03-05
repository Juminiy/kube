SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

UPDATE `tbl_calico_weave`
SET `updated_at`="2025-03-04 21:20:28.065",`name`="MyName"
WHERE `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `id` = 2
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

UPDATE `tbl_calico_weave`
SET `name`="MyName",`updated_at`="2025-03-04 21:20:28.066"
WHERE `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`id` = 2
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `id` = 2
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR ((1=1 AND `loc_id` = 33) AND `app_id` = 156))
OR (((1=1 AND `app_id` = 156) AND `app_me` = "which-is-my-handsome") AND `app_yr` = "my-bingo-done"))
OR (1=1 AND `name` = "Li-Hua"))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND NOT `tbl_calico_weave`.`id` = 3

UPDATE `tbl_calico_weave`
SET `app_id`=156,`app_me`="which-is-my-handsome",`app_yr`="my-bingo-done",
`loc_id`=33,`name`="Li-Hua",`updated_at`="2025-03-05 19:37:44.102"
WHERE `tbl_calico_weave`.`id` = 3
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
