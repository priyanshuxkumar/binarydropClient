use std::fs;

use rusqlite::{Connection, Result};

pub fn init_db(path: &str) -> Result<Connection> {
    let conn = Connection::open(path)?;

    let schema = fs::read_to_string("schema.sql").expect("Failed to read schema.sql");

    conn.execute_batch(&schema)?;

    Ok(conn)
}
