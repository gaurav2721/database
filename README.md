## Build and Run Commands

The binary is stored in the `build/` folder. Use the following Makefile commands to build, test, and run the application:

### Build and Test
```bash
make build
make test
```

### Run the Application
```bash
make run
```

### Docker Commands

#### Build Docker Image
```bash
make docker-build
```
This command builds a Docker image for the application.

#### Run Docker Container
```bash
make docker-run
```
This command runs the application in a Docker container.

#### Run Docker Container in Detached Mode
```bash
make docker-run-detached
```
This command runs the application in a Docker container in the background.

#### Build and Run Docker Container
```bash
make docker-build-run
```
This command builds the Docker image and runs the container in one step.

#### Stop Docker Container
```bash
make docker-stop
```
This command stops and removes the Docker container.

#### Clean Docker Resources
```bash
make docker-clean
```
This command stops the container, removes it, and deletes the Docker image.

#### View Docker Logs
```bash
make docker-logs
```
This command shows the logs from the running Docker container.

#### Check Docker Container Status
```bash
make docker-status
```
This command shows the status of the Docker container.

## Design Problem Solutions 

### Database Requirements

#### Functional Requirements (FR)

1. Index is associated with image files and videos  
2. Supports point lookups (exact key match, e.g., file ID)
3. Supports range lookups (e.g., fetch all files with size between X and Y or by timestamp)
4. Optimizes for additional attributes like size, genre, etc.

#### Non-Functional Requirements (Non-FR)

- Scales when the index grows.

### Assumptions/Design Decisions

1. Files (images/videos) are stored in blob/object storage (e.g., S3, filesystem) and we store only metadata and references (paths) in the database.

2. Each file has:
   - `file_id` (primary key, unique string/UUID)
   - `path` (location in storage)
   - `size` (bytes)
   - `genre` (string, e.g., "sports", "music")
   - `created_at` (timestamp)

3. Index should support:
   - Fast point lookup by file_id
   - Range lookup by attributes (e.g., size, created_at)


### Data Layout

Store metadata in a main table (sorted by file_id for quick access).

type FileMeta struct {
        ID        string    // Primary index for range lookups
        Path      string
        Size      int       // Secondary index for range queries 
        Genre     string
        CreatedAt int64     // Secondary index for range queries 
}

### Solution to Design Questions

**B+ Tree Implementation:**
- B+ trees can efficiently handle point lookups by traversing from root to leaf
- We'll use a B+Tree-based index for range queries (instead of a hash table which doesn't support range queries efficiently)
- For other attributes (genre, size), we use secondary indexes (inverted index or additional B+Trees)

**Index Structure:**
- Maintain B+Tree index for:
  - `file_id` (primary index for point lookups)
  - `size` and `created_at` (secondary indexes for range queries)
- For categorical attributes like genre, use a hash-based inverted index mapping genre → list of file_ids

### Scaling Considerations

**How solution will change if index size has to scale. What are the benefits, drawbacks and tradeoffs of your implementation?**

// I have implemented BPlusTreeIndex - Simplified B+Tree (simulated with sorted slice for demo)
// Note: This is in-memory and not scalable beyond moderate sizes.
// For scaling to millions+ of records, a disk-backed or distributed B+Tree/LSM-tree is required.
// Benefits: Range queries are efficient, multiple attributes can be indexed.
// Drawbacks: More complex I/O and caching, higher write amplification (LSM) vs. memory costs (B+Tree).
// Tradeoffs: LSM-trees give faster writes but slower range queries; B+Trees provide faster range reads, hence we have gone with B+ Trees 







