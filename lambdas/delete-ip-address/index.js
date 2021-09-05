// Lambda function code

const {
  DynamoDBClient,
  GetItemCommand,
  DeleteItemCommand,
} = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  const userId = event.requestContext.authorizer.claims.sub;

  const headers = {
    "Content-Type": "application/json",
  };

  const body = JSON.parse(event.body);
  if (body === null || body.id === null) {
    return {
      statusCode: 400,
      headers: headers,
      body: JSON.stringify({
        data: "invalid request body",
      }),
    };
  }

  const { TABLE_NAME } = process.env;
  const client = new DynamoDBClient({ region: "us-east-1" });

  let command = new GetItemCommand({
    TableName: TABLE_NAME,
    Key: {
      Id: { S: body.id },
    },
  });
  try {
    let result = await client.send(command);
    if (result.Item) {
      if (result.Item.UserId.S == userId) {
        const command = new DeleteItemCommand({
          TableName: TABLE_NAME,
          Key: {
            Id: { S: body.id },
          },
          ReturnValues: "ALL_OLD",
        });
        const result = await client.send(command);
        return {
          statusCode: result.$metadata.httpStatusCode,
          headers: headers,
          body: JSON.stringify({
            data: result.Attributes,
          }),
        };
      } else {
        return {
          statusCode: 403,
          headers: headers,
        };
      }
    } else {
      return {
        statusCode: 404,
        headers: headers,
        body: JSON.stringify({
          error: "invalid id",
        }),
      };
    }
  } catch (err) {
    return {
      statusCode: 500,
      headers: headers,
      body: JSON.stringify({
        error: err.message,
      }),
    };
  }
};
