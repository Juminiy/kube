SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-1"))
OR ((1=1 AND `loc_id` = 2110824664365475272) AND `app_id` = 49))
OR (((1=1 AND `app_id` = 49) AND `app_me` = "Cesar Trantow") AND `app_yr` = "Meditation"))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`updated_at`,`deleted_at`,`tenant_id`,`user_id`,`name`,`desc`,`pumping`,`elephant`,`loc_id`,`app_id`,`app_me`,`app_yr`)
VALUES ("2025-03-04 19:10:11.173","2025-03-04 19:10:11.173",NULL,
1919810,114514,"sandbox-1","",5.5,3.3,2110824664365475272,49,"Cesar Trantow","Meditation") RETURNING `id`
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-2"))
OR ((1=1 AND `loc_id` = 16021692958696634708) AND `app_id` = 27))
OR (((1=1 AND `app_id` = 27) AND `app_me` = "Cornell Nitzsche") AND `app_yr` = "Karate"))
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL
INSERT INTO `tbl_calico_weave`
(`app_id`,`app_me`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES (27,"Cornell Nitzsche","Karate","2025-03-04 19:10:11.174",3.3,16021692958696634708,"sandbox-2",5.5,1919810,"2025-03-04 19:10:11.174",114514) RETURNING `id`
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR ((1=1 AND `loc_id` = 13014341391514917410) AND `app_id` = 54))
OR (((1=1 AND `app_id` = 54) AND `app_me` = "Kelley Kihn") AND `app_yr` = "Hunting"))
OR (1=1 AND `name` = "sandbox-3"))
AND `tbl_calico_weave`.`tenant_id` = 1919810 AND `tbl_calico_weave`.`deleted_at` IS NULL
INSERT INTO `tbl_calico_weave`
(`app_id`,`app_me`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES (54,"Kelley Kihn","Hunting","2025-03-04 19:10:11.175",3.3,13014341391514917410,"sandbox-3",5.5,1919810,"2025-03-04 19:10:11.175",114514) RETURNING `id`
---