CREATE EXTENSION btree_gist;
CREATE EXTENSION bdr;
SELECT bdr.bdr_group_create(
    local_node_name := 'node1',
    node_external_dsn := 'dbname=svip_db host=10.8.110.21'
);

SELECT bdr.bdr_node_join_wait_for_ready();

SELECT bdr.bdr_group_join(
    local_node_name := 'node2',
    node_external_dsn := 'dbname=svip_db host=10.8.110.12',
    join_using_dsn := 'dbname= host={{ ansible_host }}'
);

