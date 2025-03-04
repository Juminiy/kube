SELECT count(*) FROM `tbl_calico_weave`
WHERE ((1!=1
OR (1!=1 OR (1=1 AND `name` = "handsome-1")))
OR (1!=1 OR (1=1 AND `name` = "handsome-2")))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`updated_at`,`deleted_at`,`tenant_id`,`user_id`,`name`,`desc`,`pumping`,`elephant`,`loc_id`,`app_id`,`app_me`,`app_yr`,`app_secret`)
VALUES ("2025-03-04 20:45:51.134","2025-03-04 20:45:51.134",NULL,1919810,114514,"handsome-1","",5.5,3.3,0,0,"","",""),
("2025-03-04 20:45:51.134","2025-03-04 20:45:51.134",NULL,1919810,114514,"handsome-2","",5.5,3.3,0,0,"","","") RETURNING `id`
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE ((1!=1
OR (1!=1 OR (1=1 AND `name` = "handsome-3")))
OR (1!=1 OR (1=1 AND `name` = "handsome-4")))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`elephant`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES ("2025-03-04 20:48:40.492",3.3,"handsome-3",5.5,1919810,"2025-03-04 20:48:40.492",114514),
("2025-03-04 20:48:40.492",3.3,"handsome-4",5.5,1919810,"2025-03-04 20:48:40.492",114514) RETURNING `id`
