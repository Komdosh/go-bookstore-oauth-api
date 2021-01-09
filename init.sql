CREATE KEYSPACE oauth WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE oauth.access_token(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);