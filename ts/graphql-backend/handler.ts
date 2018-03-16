interface CORSHeaders {
        'Access-Control-Allow-Origin': string,
        [propName: string]: string,
}

interface Response {
        headers: CORSHeaders,
        statusCode: number,
        body: string,
}

const corsHeaders: CORSHeaders = {
        'Access-Control-Allow-Origin': '*',
}

const responses: { [propName: string]: Response } = {
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
        callback(null, JSON.stringify(responses.OK))
}
