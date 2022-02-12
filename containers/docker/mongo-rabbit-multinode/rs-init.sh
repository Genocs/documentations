#!/bin/bash

mongosh <<EOF
var config = {
    "_id": "rs0",
    "version": 1,
    "members": [
        {
            "_id": 1,
            "host": "mongodb_node_1:27017",
            "priority": 3
        },
        {
            "_id": 2,
            "host": "mongodb_node_2:27017",
            "priority": 2
        },
        {
            "_id": 3,
            "host": "mongodb_node_3:27017",
            "priority": 1
        }
    ]
};
rs.initiate(config, { force: true });
rs.status();
EOF
