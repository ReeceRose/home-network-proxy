// Lambda function code

const {
  DynamoDBClient,
  DeleteItemCommand,
} = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  const body = JSON.parse(event.body);
  // Not too much data validation going on, if any further data validation errors it will just return an error
  if (body === null || body.id === null) {
    return {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        data: "invalid request body",
      }),
    };
  }

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new DeleteItemCommand({
    TableName: TABLE_NAME,
    Key: {
      Id: { S: body.id },
    },
    ReturnValues: "ALL_OLD",
  });
  try {
    const result = await client.send(command);
    return {
      statusCode: result.$metadata.httpStatusCode,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        data: result.Attributes,
      }),
    };
  } catch (err) {
    return {
      statusCode: 500,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        error: err.message,
      }),
    };
  }
};
