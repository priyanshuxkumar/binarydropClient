CREATE TABLE IF NOT EXISTS files (
    id TEXT PRIMARY KEY,                          
    name TEXT NOT NULL,
    modified_time INTEGER NOT NULL,               
    size INTEGER NOT NULL,                        
    path TEXT NOT NULL,                           
    hash TEXT,                                    
    sync_status TEXT DEFAULT 'PENDING',           
    created_at INTEGER DEFAULT (strftime('%s','now')),
    updated_at INTEGER DEFAULT (strftime('%s','now')),
    is_deleted INTEGER DEFAULT 0                  
);

-- Indexes for files
CREATE INDEX IF NOT EXISTS idx_path ON files(path);
CREATE INDEX IF NOT EXISTS idx_hash ON files(hash);
CREATE INDEX IF NOT EXISTS idx_sync_status ON files(sync_status);


CREATE TABLE  IF NOT EXISTS chunks (
    id TEXT PRIMARY KEY,                          
    file_id TEXT NOT NULL,                        
    chunk_index INTEGER NOT NULL,
    chunk_hash TEXT NOT NULL,
    chunk_size INTEGER NOT NULL,
    cloud_url TEXT,                               
    upload_status TEXT DEFAULT 'PENDING',         -- PENDING, UPLOADING, UPLOADED, ERROR
    created_time INTEGER DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL,

    FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE,
    UNIQUE(file_id, chunk_index)
);

-- Indexes for chunks
CREATE INDEX IF NOT EXISTS idx_chunk_hash ON chunks(chunk_hash);
CREATE INDEX IF NOT EXISTS idx_upload_status ON chunks(upload_status);
