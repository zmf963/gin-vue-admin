openssl req \
-newkey rsa:2048 \
-x509 \
-nodes \
-keyout server.key \
-new \
-out server.crt \
-subj /CN=Hostname \
-reqexts SAN \
-extensions SAN \
-config <(cat /etc/ssl/openssl.cnf \
    <(printf '[SAN]\nsubjectAltName=DNS:hostname,IP:127.0.0.1,IP:0.0.0.0')) \
-sha256 \
-days 3650