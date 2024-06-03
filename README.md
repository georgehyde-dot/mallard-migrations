# Mallard-Migrations
This is a work in progress to transition a simple rigid tool
into a more extensible project that has broader appeal

**Mallard-Migrations** is a Golang-based tool designed to integrate seamlessly
 into CI/CD pipelines to manage database migrations. The tool leverages git 
 for version control and uses message brokers to ensure smooth, automated 
 execution of SQL scripts on your databases. The goal is to support multiple 
 message brokers, databases, and versioncontrol systems.

## Table of Contents
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features
- **Version Control**: Uses git to manage and track database migration scripts. This section needs some thought to figure out how best to integrate into existing SQL version control systems. My initial idea is to use branch names as the determining factor for which actions are taken. Then something like git diff HEAD^1 to find newly added files to execute.
- **Message Broker Integration**: Supports RabbitMQ with plans to expand to others (SQS is first on my list) to send migration scripts to the execution environment.
- **Migration Tracking**: Logs execution details, including file names, execution time, and success status, into a SQLite database. This can be set up to use any persistent store.
- **CI/CD Integration**: TODO

## License
This project is licensed under the APACHE 2 License - see the [LICENSE](LICENSE) file for details.

## Contact
For any questions or feedback, please open an issue on GitHub

---

Happy migrating with Mallard-Migrations!