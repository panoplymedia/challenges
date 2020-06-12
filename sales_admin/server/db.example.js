const { Pool } = require("pg");

const user = 'POSTGRESUSERNAME';
const password = 'POSTGRESSPASSWORD';
const postgresport = 5432;
const db = 'achs';
const connection = `postgresql://${user}:${password}@localhost:${postgresport}/${db}`;

const pool = new Pool({
    connectionString: connection,
    ssl: false
})


module.exports = pool;