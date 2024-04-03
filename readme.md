# Simple Bank Project

* API (RESTful): Use Gin server, a lightweight web framework for Go, to create RESTful API endpoints for banking operations.
* gRPC & gRPC Gateway: Implement gRPC services for inter-service communication and use gRPC Gateway to expose these services as RESTful endpoints. Define your API contracts using Protocol Buffers (protobuf).
* Database (PostgreSQL): Utilize PostgreSQL as your database engine. Use SQL Builder SQLC to generate type-safe Go code for interacting with the database.
* Migration: Manage database schema changes and versioning using a migration tool with Migrate.
* Authentication: Implement authentication using JWT or Paseto tokens. And applied role base access control.
* Email Verification: Integrate a mailer service to send email verification links to users during account registration. Implemented email verification using Redis as a task distributor for utilized Redis as a task queue to enqueue email verification tasks containing user information. Developed worker(s) to process tasks concurrently, ensuring timely email delivery. Enhanced scalability by adjusting the number of worker instances based on workload.
* [GitHub repo](https://github.com/InkZ-VJ/goBanking)
