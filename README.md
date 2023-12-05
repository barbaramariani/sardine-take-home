# Risk Evaluator Service

This service evaluates financial transaction risks based on predefined rules.

## Run

### Prerequisites

- Go (at least version 1.14)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/barbaramariani/sardine-take-home.git
    cd sardine-take-home
    ```

2. Run the service:

    ```bash
    go run .
    ```

The service should now be running at http://localhost:8090.

## Usage

The service exposes a single endpoint for risk evaluation:

- Endpoint: `http://localhost:8090/risk`
- Method: `POST`
- Request Body: JSON payload containing an array of financial transactions.

### Example Request

```bash
curl -X POST \
  http://localhost:8090/risk \
  -H 'Content-Type: application/json' \
  -d '{
    "transactions": [
      {"id": 1, "user_id": 1, "amount_us_cents": 200000, "card_id": 1},
      {"id": 2, "user_id": 1, "amount_us_cents": 600000, "card_id": 1},
      {"id": 3, "user_id": 1, "amount_us_cents": 1100000, "card_id": 1},
      {"id": 4, "user_id": 2, "amount_us_cents": 100000, "card_id": 2},
      {"id": 5, "user_id": 2, "amount_us_cents": 100000, "card_id": 3},
      {"id": 6, "user_id": 2, "amount_us_cents": 100000, "card_id": 4}
    ]
  }'
