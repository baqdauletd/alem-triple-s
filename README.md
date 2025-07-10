# triple-s

## Overview
`triple-s` (Simple Storage Service) is a lightweight and simplified implementation of a cloud storage service, inspired by Amazon S3. This project is designed to provide a straightforward `RESTful API` that allows users to create and manage storage buckets and objects. The main goal of triple-s is to offer a practical and easy-to-understand foundation for key concepts related to cloud storage systems, as well as to demonstrate essential principles of RESTful API design, basic networking, and data management.

## Features

- **Bucket Management**: Create, list, and delete storage buckets.
- **Object Management**: Upload, retrieve, and delete objects within buckets.
- **Metadata Handling**: Store and manage object metadata in CSV files.
- **REST API**: Interact with the storage system using HTTP methods.
- **XML Responses**: All API responses conform to the Amazon S3 XML format.

## API Endpoints

### Bucket Management

#### Create a Bucket
- **HTTP Method**: `PUT`
- **Endpoint**: `/{BucketName}`
- **Request Body**: Empty
- **Behavior**:
  - Validate bucket name.
  - Ensure bucket name is unique.
  - Create a new bucket.
  - Respond with `200 OK` or an appropriate error message.

#### List Buckets
- **HTTP Method**: `GET`
- **Endpoint**: `/`
- **Behavior**:
  - List all existing buckets.
  - Respond with `200 OK` and bucket details.

#### Delete a Bucket
- **HTTP Method**: `DELETE`
- **Endpoint**: `/{BucketName}`
- **Behavior**:
  - Validate bucket existence.
  - Delete the bucket.
  - Respond with `204 No Content` or an appropriate error message.

### Object Management

#### Upload a New Object
- **HTTP Method**: `PUT`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Request Body**: Binary data of the object
- **Headers**:
  - `Content-Type`: The object's data type.
  - `Content-Length`: The length of the content in bytes.
- **Behavior**:
  - Validate bucket and object key.
  - Save the object content.
  - Store object metadata.
  - Respond with `200 OK` or an appropriate error message.

#### Retrieve an Object
- **HTTP Method**: `GET`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Behavior**:
  - Validate bucket and object existence.
  - Return the object data or an error.

#### Delete an Object
- **HTTP Method**: `DELETE`
- **Endpoint**: `/{BucketName}/{ObjectKey}`
- **Behavior**:
  - Validate bucket and object existence.
  - Delete the object and update metadata.
  - Respond with `204 No Content` or an appropriate error message.

## CSV File Structure for Object Metadata

Each bucket has its own object metadata CSV file (e.g., `data/{bucket-name}/objects.csv`).

- **Columns**:
  - `ObjectKey`: The unique key of the object.
  - `Size`: The size of the object in bytes.
  - `ContentType`: The MIME type of the object.
  - `LastModified`: The timestamp of the last modification.

## Running the Project

Follow these steps:

1. **Clone the Repository**:
    ```sh
    git clone git@github.com:baqdauletd/alem-triple-s.git
    cd triple-s
    ```

2. **Building the Project (already done - or - after each change)**
    ```sh
    go build -o triple-s . #creates executable triple-s (but it is already done)
    ```

3. **Run the Project**:
    ```sh
    ./triple-s # it will run on default port: 8080 and directory: ./data

    # to specify subcommands
    ./triple-s --port=8080 --dir="./data"
    
    # or 
    ./triple-s --help
    ```
