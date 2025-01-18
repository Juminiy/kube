SELECT count(*)
FROM `tbl_unique_test2`
WHERE
(
  (((
    (1!=1
    OR (1=1 AND `mac_addr` = "02:1A:2B:3C:4D:5E"))
    OR ((1=1 AND `region_code` = 666) AND `ip_addr` = "10.101.22.10"))
    OR ((1=1 AND `region_code` = 666) AND `hostname` = "BJ01-HPC-0008"))
    OR ((1=1 AND `region_code` = 666) AND `node_id` = "pja-0x8090621"))
    OR ((1=1 AND `hostname` = "BJ01-HPC-0008") AND `node_id` = "pja-0x8090621")
)
AND `tbl_unique_test2`.`deleted_at` IS NULL
