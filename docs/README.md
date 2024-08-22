---

### 1. Build Script (`build.sh`)

**Example Run**:
```bash
./scripts/build.sh
```
**Description**: This command compiles the load balancer application. Make sure you are in the root directory of the project before running this.

---

### 2. Dependency Installation Script (`install_deps.sh`)

**Example Run**:
```bash
./scripts/install_deps.sh
```
**Description**: This command installs and updates all Go dependencies across the project. Ensure you are in the root directory before executing it.

---

### 3. Load Testing Script (`load_test.sh`)

**Example Run**:
```bash
./scripts/load_test.sh -n 1000 -v 1 -u http://localhost:8080 -w 500
```
**Description**:
- `-n 1000` sets the number of requests to 1000.
- `-v 1` enables verbose output, showing HTTP status codes.
- `-u http://localhost:8080` specifies the URL of the load balancer.
- `-w 500` waits 500 milliseconds after every 1000 requests.

---

### 4. Test Script (`run_tests.sh`)

**Example Run**:
```bash
./scripts/run_tests.sh
```
**Description**: This command runs all unit tests located in the `tests` directory. Ensure you are in the root directory before executing it.

---

### 5. Start Servers and Load Balancer (`run.sh`)

**Example Run**:
```bash
./scripts/start_services.sh
```
**Description**: This command starts the example servers and load balancer, then waits for all processes to finish. Execute this from the root directory to ensure correct paths.

---

### 6. Update Dependencies Script (`update_deps.sh`)

**Example Run**:
```bash
./scripts/update_deps.sh
```
**Description**: This command updates all Go dependencies to their latest versions, tidies module files, and verifies dependencies. Make sure you are in the root directory of the project when running this.

---
