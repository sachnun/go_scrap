# Web Scraping with Go

This is a simple Go application that performs web scraping on a given website URL and returns the website details as JSON.

## Prerequisites

To run this application, you need to have Go installed on your system. You can download and install Go from the official Go website: https://golang.org/

## Getting Started

1. Clone the repository or download the source code files.

2. Open a terminal or command prompt and navigate to the project's directory.

3. Install the project dependencies by running the following command:

   ```shell
   go get -u github.com/PuerkitoBio/goquery
   go get -u github.com/microcosm-cc/bluemonday
   ```

4. Build the application by running the following command:

   ```shell
   go build
   ```

5. Run the application:

   ```shell
   ./web-scraping-go
   ```

   By default, the application will start a server listening on port 8080.

## Endpoints

The application exposes the following endpoints:

- `GET /website-detail?url={url}`: Retrieves the details of the specified website. Replace `{url}` with the actual URL you want to scrape.

- `GET /`: Returns a simple message indicating the purpose of the application.

- `GET /ping`: Returns a simple message to check if the server is running.

## Example Usage

To retrieve the details of a website, send a GET request to the `/website-detail` endpoint with the `url` query parameter.

Example:

```
GET /website-detail?url=https://example.com
```

Response:

```json
{
  "url": "https://example.com",
  "results": "Example Domain Example Domain This domain is for use in illustrative examples in documents. You may use this domain in literature without prior coordination or asking for permission. More information...",
  "status_code": 200,
  "total_chars": 275,
  "length": 41
}
```

## Error Handling

If any error occurs during the scraping process, an error response with an appropriate status code will be returned. The response will contain a JSON object with an `"error"` field indicating the error message.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
