## Sign Up and Login System

A secure authentication system built with Golang, featuring sign-up and login functionalities. User data is stored in a Postgres database and passwords are persisted using bcrypt.

### Features
    
- [x] Sign Up
- [x] Login
- [x] Password Hashing
- [x] Postgres Database

### Installation

1. Clone the repository
2. Install dependencies
    ```bash
    go get -u github.com/gorilla/mux
    go get -u github.com/lib/pq
    go get -u github.com/joho/godotenv
    ```
3. Run the application
    ```bash
    go run main.go
    ```
   
### Flow Diagram

[![](https://mermaid.ink/img/pako:eNp9kcFqwzAMhl_FGLJcuhfwoZc1x8FY2U65aPGfTpA4raxslJB3n72UrqO0Phjz6ZMlock2g4d1tigmDqzOTKV-okfpynYQRC3nuSjqEHEYERpsmHZCfR1MOnsS5Yb3FNS8Rcg13R6j4mRn43G9XpAzVVCI2UCJu2gezHb86FkXdXEu5FfsOGb_r8zZycil9KZBjOY5XbTDnZIZBuqRar5QjN-D-BtV36ljTwrzJPAIytTFRaVOL6EhgfnK9hK-3d8SRxdxlc_h_g-VyCD_50PwdmV7SE_s0xanjGv7u8HauvT0aGnstLZ1mJNKow7bY2isUxmxsuM-z3daqnVt6uVMK886yAnOP8-1u50?type=png)](https://mermaid.live/edit#pako:eNp9kcFqwzAMhl_FGLJcuhfwoZc1x8FY2U65aPGfTpA4raxslJB3n72UrqO0Phjz6ZMlock2g4d1tigmDqzOTKV-okfpynYQRC3nuSjqEHEYERpsmHZCfR1MOnsS5Yb3FNS8Rcg13R6j4mRn43G9XpAzVVCI2UCJu2gezHb86FkXdXEu5FfsOGb_r8zZycil9KZBjOY5XbTDnZIZBuqRar5QjN-D-BtV36ljTwrzJPAIytTFRaVOL6EhgfnK9hK-3d8SRxdxlc_h_g-VyCD_50PwdmV7SE_s0xanjGv7u8HauvT0aGnstLZ1mJNKow7bY2isUxmxsuM-z3daqnVt6uVMK886yAnOP8-1u50)
