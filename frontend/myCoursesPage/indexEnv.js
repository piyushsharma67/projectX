/* eslint @typescript-eslint/no-var-requires: "off" */

const basePath = process.cwd();
const envType = process.argv[2].substring(4);
const fs = require('fs');


const devPath = `${basePath}/env/dev.env`;
const stagePath = `${basePath}/env/stage.env`;
const prodPath = `${basePath}/env/prod.env`;
const ngrokPath = `${basePath}/env/ngrok.env`;

const jsPath = `${basePath}/src/apiLink/apiLink.ts`;

// comman path created for all env based variable
const JsCommonPath = `${basePath}/env/index.js`;

// To fetch enums which decide current nature of enviorement
const enums = require('./env/enum');

// To fetch webview url as per enviorement
const webViewEnvUrl = require(`${basePath}/env/webViewUrl.js`);
const dsImageCheckerUrl = require(`${basePath}/env/dsApiEndpoint/dsImageCheckerUrl.js`)
const dsDocumentExtractorUrl = require(`${basePath}/env/dsApiEndpoint/dsDocumentExtractorUrl.js`)
const ticketingUrlFilePath = require('./env/ticketingUrl')

let link;
let webViewUrl;
let imageCheckerUrl;
let documentExtractorUrl;

switch (envType) {
    case 'dev':
        link = fs.readFileSync(devPath, { encoding: 'utf8', flag: 'r' });
        webViewUrl = webViewEnvUrl[enums.dev];
        imageCheckerUrl = dsImageCheckerUrl[enums.dev]
        documentExtractorUrl = dsDocumentExtractorUrl[enums.dev]
        ticketingUrl = ticketingUrlFilePath[enums.dev]
        break;
    case 'stage':
        link = fs.readFileSync(stagePath, { encoding: 'utf8', flag: 'r' });
        webViewUrl = webViewEnvUrl[enums.stage];
        imageCheckerUrl = dsImageCheckerUrl[enums.stage]
        documentExtractorUrl = dsDocumentExtractorUrl[enums.stage]
        ticketingUrl = ticketingUrlFilePath[enums.stage]
        break;
    case 'prod':
        link = fs.readFileSync(prodPath, { encoding: 'utf8', flag: 'r' });
        webViewUrl = webViewEnvUrl[enums.prod];
        imageCheckerUrl = dsImageCheckerUrl[enums.prod]
        documentExtractorUrl = dsDocumentExtractorUrl[enums.prod]
        ticketingUrl = ticketingUrlFilePath[enums.prod]
        break;
    case 'ngrok':
        link = fs.readFileSync(ngrokPath, { encoding: 'utf8', flag: 'r' });
        webViewUrl = webViewEnvUrl[enums.ngrok];
        imageCheckerUrl = dsImageCheckerUrl[enums.ngrok]
        documentExtractorUrl = dsDocumentExtractorUrl[enums.ngrok]
        ticketingUrl = ticketingUrlFilePath[enums.dev]
        break;
    default:
        link = fs.readFileSync(stagePath, { encoding: 'utf8', flag: 'r' });
        webViewUrl = webViewEnvUrl[enums.stage];
        imageCheckerUrl = dsImageCheckerUrl[enums.stage]
        documentExtractorUrl = dsDocumentExtractorUrl[enums.stage]
        ticketingUrl = ticketingUrlFilePath[enums.dev]

}
fs.writeFileSync(jsPath, `export const ApiLink = ${link}\nexport const env: 'prod' | 'stage' | 'dev' = '${envType}' \nexport const imageCheckerApi = '${imageCheckerUrl}'\nexport const documentExtractorApi = '${documentExtractorUrl}'\nexport const ticketUrl = '${ticketingUrl}'`);

fs.writeFileSync(
    JsCommonPath,
    `module.exports = { BaseUrl: ${link}, env: '${envType}', webViewUrl: '${webViewUrl}' };`,
);
