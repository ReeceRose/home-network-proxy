// Lambda function code

const { DynamoDBClient, GetItemCommand } = require("@aws-sdk/client-dynamodb");

const { NodeSSH } = require("node-ssh");

exports.handler = async (event) => {
  const userId = event.requestContext.authorizer.claims.sub;

  const { TABLE_NAME } = process.env;
  const headers = {
    "Content-Type": "application/json",
  };

  const body = JSON.parse(event.body);
  if (
    body === null ||
    body.id === null ||
    body.username === null ||
    body.password === null ||
    body.command === null
  ) {
    return {
      statusCode: 400,
      headers: headers,
      body: JSON.stringify({
        data: "invalid request body",
      }),
    };
  }

  const ssh = new NodeSSH();

  const client = new DynamoDBClient({ region: "us-east-1" });
  let command = new GetItemCommand({
    TableName: TABLE_NAME,
    Key: {
      Id: { S: body.id },
    },
  });

  try {
    const result = await client.send(command);

    if (result.Item.length > 0 || result.Item.UserId.S == userId) {
      const connection = await ssh.connect({
        host: result.Item.ExternalIP.S,
        username: body.username,
        password: body.password,
        port: body.port || "22",
      });
      const result = await connection.execCommand(body.command, {
        stream: "stdout",
      });
      await ssh.dispose();
      return {
        statusCode: 200,
        headers: headers,
        body: JSON.stringify({
          data: result,
        }),
      };
    } else {
      return {
        statusCode: 403,
        headers: headers,
        body: JSON.stringify({
          error: "invalid request",
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
