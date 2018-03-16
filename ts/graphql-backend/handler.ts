interface CORSHeaders {
        'Access-Control-Allow-Origin': string,
        [propName: string]: string,
}

interface Response {
        Headers: CORSHeaders,
        StatusCode: number,
        Message: string,
}

const corsHeaders: CORSHeaders = {
        'Access-Control-Allow-Origin': '*',
}

const responses: { [propName: string]: Response } = {
        OK: {
                Headers: corsHeaders,
                StatusCode: 200,
                Message: "Everything OK!",
        },
        InternalError: {
                Headers: corsHeaders,
                StatusCode: 500,
                Message: "We messed up!",
        }
}

export function handler(event: AWSLambda.APIGatewayEvent, context: AWSLambda.Context, callback: AWSLambda.Callback) {
        callback(null, responses.OK)
}
