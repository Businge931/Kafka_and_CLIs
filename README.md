
```markdown
# Command-line Driven Program for Kafka Message Exchange

This application allows you to send or receive messages from a Kafka cluster using a command-line interface (CLI). Built with Go and powered by the Cobra library, this program makes it easy to implement Kafka producers and consumers directly from the terminal.

## Features

- **Send** messages to a Kafka topic as a producer.
- **Receive** messages from a Kafka topic as a consumer.
- Easy configuration of channels (topics), server addresses, and consumer groups using command-line flags.

## Prerequisites

Before running the application, ensure you have the following installed on your machine:

- **Apache Kafka** and **Zookeeper**.
- **Go** programming language.

## Getting Started

### Clone the Repository

To get started, clone the repository:

```bash
git clone https://github.com/Businge931/Kafka_and_CLIs.git
cd Kafka_and_CLIs
```

### Start Zookeeper and Kafka

To send and receive messages, you'll need to start Zookeeper and Kafka servers in separate terminal instances. Follow the steps below:

1. **Start Zookeeper**:

   Open a terminal and run:

   ```bash
   bin/zookeeper-server-start.sh config/zookeeper.properties
   ```

2. **Start Kafka Server**:

   In a second terminal window, start Kafka:

   ```bash
   bin/kafka-server-start.sh config/server.properties
   ```

### Sending and Receiving Messages

The program allows you to either send messages (producer) or receive messages (consumer) by running different commands.

#### Send Messages (Producer)

To send messages to a Kafka topic, use the following command:

```bash
go run main.go send --channel <channel_name> --server <server:port> --group <group_name>
```

Replace the placeholders `<channel_name>`, `<server:port>`, and `<group_name>` with actual values.

Example:

```bash
go run main.go send --channel myTopic --server localhost:9092 --group myGroup
```

This will start a producer that sends messages to the specified channel (topic) on the Kafka server.

#### Receive Messages (Consumer)

To receive messages from a Kafka topic, run the following command:

```bash
go run main.go receive --channel <channel_name> --from start --server <server:port> --group <group_name>
```

Replace the placeholders `<channel_name>`, `<server:port>`, and `<group_name>` with actual values.

Example:

```bash
go run main.go receive --channel myTopic --from start --server localhost:9092 --group myGroup
```

This will start a consumer that listens to messages from the specified Kafka topic.

### Flags Description

The application uses several command-line flags to configure the Kafka producer and consumer:

- `--channel`: The Kafka topic you want to subscribe to.
- `--server`: The Kafka server address in the format `host:port` (e.g., `localhost:9092`).
- `--group`: The consumer group ID to join.
- `--from`: Specifies the starting point for consuming messages (e.g., `start` to consume from the beginning).

### Running Each Instance

Ensure that each instance (producer or consumer) is run in a separate terminal window. 

For example:
- In one terminal, run the **send** command to start the producer.
- In another terminal, run the **receive** command to start the consumer.

### Example Workflow

Here’s how you can run both producer and consumer in separate terminal windows:

1. **Producer**:

   ```bash
   go run main.go send --channel myTopic --server localhost:9092 --group myGroup
   ```

2. **Consumer**:

   ```bash
   go run main.go receive --channel myTopic --from start --server localhost:9092 --group myGroup
   ```

### Notes

- Ensure both Zookeeper and Kafka servers are running before executing any commands.
- Run the producer and consumer in separate terminals to send and receive messages concurrently.


## Contact

For any inquiries or issues, feel free to open an issue in the [repository](https://github.com/Businge931/Kafka_and_CLIs/issues).
```

This version includes all the commands and instructions necessary to run the application, including specific flags and example commands for both sending and receiving messages.