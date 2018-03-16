interface CORSHeaders {
        'Access-Control-Allow-Origin': string,
        [propName: string]: string | boolean | number,
}

const corsHeaders: CORSHeaders = {
        'Access-Control-Allow-Origin': '*',
}

const responses: { [propName: string]: AWSLambda.APIGatewayProxyResult } = {
        OK: {
                headers: corsHeaders,
                statusCode: 200,
                body: JSON.stringify({
                        Message: "Everything OK!",
                }),
        },
        InternalError: {
                headers: corsHeaders,
                statusCode: 500,
                body: JSON.stringify({
                        Message: "We messed up!",
                }),
        }
}

export function handler(event: AWSLambda.APIGatewayEvent, context: AWSLambda.Context, callback: AWSLambda.Callback) {
        console.log("Executing ...")
        callback(null, responses.OK)
}
