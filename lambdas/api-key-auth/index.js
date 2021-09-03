// Lambda function code

const { DynamoDBClient, ScanCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  let response = {
    isAuthorized: false,
    context: {},
  };

  const token = event.headers.authorization;
  const userId = event.headers.UserId || event.headers.userid;
  if (token === null || userId === null) {
    return response;
  }

  const { TABLE_NAME } = process.env;
  const client = new DynamoDBClient({ region: "us-east-1" });

  const command = new ScanCommand({
    TableName: TABLE_NAME,
    FilterExpression: "Id = :id and UserId = :userId",
    ExpressionAttributeValues: {
      ":id": { S: token },
      ":userId": { S: userId },
    },
  });

  try {
    let result = await client.send(command);
    if (result.Items.length >= 1) {
      return {
        isAuthorized: true,
        context: {},
        statusCode: 200,
      };
    }
  } catch (err) {
    console.log(err);
  }
  return response;
};
