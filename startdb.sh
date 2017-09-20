cockroach start --insecure --store=oshop --host=localhost &&
cockroach sql --insecure -e 'CREATE DATABASE oshop'