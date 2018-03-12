# XII Simposio de Contabilidad del Extremo Sur
> Infrastructure

This repo contains lambda functions and CloudFormation template to deploy backend for the app required by event *XII Simposio de Contabilidad del Extremo Sur*.

This project is done using [Serverless framework](https://serverless.com/).

## Deploy Cognito

Run this command: `AWS_DEFAULT_REGION="us-east-1" aws cloudformation create-stack --stack-name xii-simposio --template-body file://cognito-template.yml`

## License

All code is under MIT license, [read license](LICENSE). Distribution of logos, pictures, and other propietary imagery is forbidden unless under explicit authorization, you could contact [ponencias.simposio@gmail.com](mailto:ponencias.simposio@gmail.com).
